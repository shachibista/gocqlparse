package ast

import (
	"encoding/hex"
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/uuid"
	"github.com/shachibista/gocqlparse/internal/parser"
)

type Visitor struct {
	parser.BaseCql3Visitor

	errors []error
}

var ReservedTypeNames = map[string]bool{
	"byte":      true,
	"complex":   true,
	"enum":      true,
	"date":      true,
	"interval":  true,
	"macaddr":   true,
	"bitstring": true,
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) Errs() error {
	return errors.Join(v.errors...)
}

func (v *Visitor) Err(ctx antlr.ParserRuleContext, msg string) {
	start := ctx.GetStart()
	line := start.GetLine()
	col := start.GetColumn()

	v.errors = append(v.errors, &visitorError{
		line: line,
		col:  col,
		msg:  msg,
	})
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitComparatorType(ctx *parser.ComparatorTypeContext) any {
	if ctx.Native_type() != nil {
		return v.Visit(ctx.Native_type())
	}

	if ctx.Collection_type() != nil {
		return v.Visit(ctx.Collection_type())
	}

	if ctx.Tuple_type() != nil {
		return v.Visit(ctx.Tuple_type())
	}

	if ctx.Vector_type() != nil {
		return v.Visit(ctx.Vector_type())
	}

	if ctx.UserTypeName() != nil {
		return v.Visit(ctx.UserTypeName())
	}

	if ctx.K_FROZEN() != nil {
		return &FrozenType{
			Elem: v.Visit(ctx.ComparatorType()),
		}
	}

	return &CustomType{
		Name: v.Visit(ctx.STRING_LITERAL()).(string),
	}
}

func (v *Visitor) VisitNative_type(ctx *parser.Native_typeContext) any {
	return &NativeType{
		Name: strings.ToLower(ctx.GetText()),
	}
}

func (v *Visitor) VisitCollection_type(ctx *parser.Collection_typeContext) any {
	switch {
	case ctx.K_LIST() != nil:
		elementType := v.Visit(ctx.GetT())

		return &ListType{
			Elem: elementType,
		}
	case ctx.K_SET() != nil:
		elementType := v.Visit(ctx.GetT())

		return &SetType{
			Elem: elementType,
		}
	case ctx.K_MAP() != nil:
		keyType := v.Visit(ctx.GetT1())
		valueType := v.Visit(ctx.GetT2())

		return &MapType{
			Key:   keyType,
			Value: valueType,
		}
	}

	ctype := ctx.GetChild(0).(antlr.TerminalNode)

	v.Err(ctx, "unsupported collection type: "+ctype.GetText())

	return nil
}

func (v *Visitor) VisitTuple_type(ctx *parser.Tuple_typeContext) any {
	tup := &TupleType{}

	for _, t := range ctx.AllComparatorType() {
		tup.Elements = append(tup.Elements, v.Visit(t))
	}

	return tup
}

func (v *Visitor) VisitVector_type(ctx *parser.Vector_typeContext) any {
	elem := v.Visit(ctx.ComparatorType())
	dims, err := strconv.Atoi(ctx.GetD().GetText())
	if err != nil {
		v.Err(ctx, "non-integer dimension: "+ctx.GetD().String())
		return nil
	}

	return &VectorType{
		Elem:       elem,
		Dimensions: dims,
	}
}

func (v *Visitor) VisitNonTypeIdent(ctx *parser.NonTypeIdentContext) any {
	id := visitIdent(ctx)

	if ReservedTypeNames[id.String()] {
		v.Err(ctx, "invalid (reserved) user type name: "+ctx.GetText())

		return nil
	}

	return id
}

func (v *Visitor) VisitNonTypeQuotedIdent(ctx *parser.NonTypeQuotedIdentContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitNonColIdent(ctx *parser.NonColIdentContext) any {
	return visitIdent(ctx)
}

func (v *Visitor) VisitNonColQuoted(ctx *parser.NonColQuotedContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitUserTypeName(ctx *parser.UserTypeNameContext) any {
	var keyspace Identifier
	if ctx.GetKs() != nil {
		keyspace = v.Visit(ctx.GetKs()).(Identifier)
	}

	ident := v.Visit(ctx.GetUt()).(Identifier)

	return &ObjectRef{
		Keyspace: keyspace,
		Name:     ident,
	}
}

func (v *Visitor) VisitKeyspaceName(ctx *parser.KeyspaceNameContext) any {
	return v.Visit(ctx.KsName())
}

func (v *Visitor) VisitKsNameIdent(ctx *parser.KsNameIdentContext) any {
	return visitIdent(ctx)
}

func (v *Visitor) VisitKsNameQuotedIdent(ctx *parser.KsNameQuotedIdentContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitKsNameInvalidBind(ctx *parser.KsNameInvalidBindContext) any {
	v.Err(ctx, "bind variables cannot be used for keyspace names")

	return nil
}

func (v *Visitor) VisitCfNameIdent(ctx *parser.CfNameIdentContext) any {
	return visitIdent(ctx)
}

func (v *Visitor) VisitCfNameQuotedIdent(ctx *parser.CfNameQuotedIdentContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitIdentIdent(ctx *parser.IdentIdentContext) any {
	return visitIdent(ctx)
}

func (v *Visitor) VisitIdentQuotedIdent(ctx *parser.IdentQuotedIdentContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitFIdentIdent(ctx *parser.FIdentIdentContext) any {
	return visitIdent(ctx)
}

func (v *Visitor) VisitFIdentQuotedIdent(ctx *parser.FIdentQuotedIdentContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitCfNameInvalidBind(ctx *parser.CfNameInvalidBindContext) any {
	v.Err(ctx, "bind variables cannot be used for column family names")

	return nil
}

func (v *Visitor) VisitTableColumnDeclaration(ctx *parser.TableColumnDeclarationContext) any {
	key := v.Visit(ctx.GetK()).(Identifier)
	val := v.Visit(ctx.GetV()).(GenericType)

	var mask *Function
	if ctx.GetMask() != nil {
		mask = v.Visit(ctx.GetMask()).(*Function)
	}

	return &TableColumn{
		Name:               key,
		Type:               val,
		Static:             ctx.K_STATIC() != nil,
		ImplicitPrimaryKey: ctx.K_PRIMARY() != nil,
		Mask:               mask,
	}
}

func (v *Visitor) VisitTableKeyDeclaration(ctx *parser.TableKeyDeclarationContext) any {
	pks := v.Visit(ctx.GetPk()).([]Identifier)

	var cks []Identifier
	for _, ck := range ctx.AllIdent() {
		cks = append(cks, v.Visit(ck).(Identifier))
	}

	return &TableKeys{
		Partition:  pks,
		Clustering: cks,
	}
}

func (v *Visitor) VisitTablePartitionKey(ctx *parser.TablePartitionKeyContext) any {
	var keys []Identifier

	for _, k := range ctx.AllIdent() {
		keys = append(keys, v.Visit(k).(Identifier))
	}

	return keys
}

func (v *Visitor) VisitAllowedFunctionNameIdent(ctx *parser.AllowedFunctionNameIdentContext) any {
	return visitIdent(ctx)
}

func (v *Visitor) VisitAllowedFunctionNameQuotedIdent(ctx *parser.AllowedFunctionNameQuotedIdentContext) any {
	return visitQuotedIdent(ctx)
}

func (v *Visitor) VisitFunctionName(ctx *parser.FunctionNameContext) any {
	var ks Identifier

	if ctx.GetKs() != nil {
		ks = v.Visit(ctx.GetKs()).(Identifier)
	}

	f := v.Visit(ctx.GetF()).(Identifier)

	return &ObjectRef{
		Keyspace: ks,
		Name:     f,
	}
}

func (v *Visitor) VisitColumnFamilyName(ctx *parser.ColumnFamilyNameContext) any {
	var ks Identifier
	if ctx.KsName() != nil {
		ks = v.Visit(ctx.KsName()).(Identifier)
	}

	cf := v.Visit(ctx.CfName()).(Identifier)

	return &ObjectRef{
		Keyspace: ks,
		Name:     cf,
	}
}

func (v *Visitor) VisitProperty(ctx *parser.PropertyContext) any {
	k := v.Visit(ctx.GetK()).(Identifier)

	var value any
	if ctx.GetSimple() != nil {
		value = v.Visit(ctx.GetSimple())
	} else if ctx.GetMap_() != nil {
		value = v.Visit(ctx.GetMap_())
	}

	return &TableProperty{
		Key:   k,
		Value: value,
	}
}

func (v *Visitor) VisitFullMapLiteral(ctx *parser.FullMapLiteralContext) any {
	var terms []any
	for _, t := range ctx.AllTerm() {
		terms = append(terms, v.Visit(t))
	}

	if len(terms)%2 != 0 {
		v.Err(ctx, "unpaired key-value pair")
		return nil
	}

	elements := make(MapLiteral, len(terms)/2)
	for i := 0; i < len(terms); i += 2 {
		elements[terms[i]] = terms[i+1]
	}

	return &elements
}

func (v *Visitor) VisitPropertyValue(ctx *parser.PropertyValueContext) any {
	if ctx.GetC() != nil {
		return v.Visit(ctx.GetC())
	}

	if ctx.GetU() != nil {
		return v.Visit(ctx.GetU())
	}

	v.Err(ctx, "bad property value")
	return nil
}

func (v *Visitor) VisitTableClusteringOrder(ctx *parser.TableClusteringOrderContext) any {
	k := v.Visit(ctx.GetK()).(Identifier)
	asc := ctx.K_ASC() != nil

	return &ColumnOrder{
		Name:      k,
		Ascending: asc,
	}
}

func (v *Visitor) VisitTypeColumns(ctx *parser.TypeColumnsContext) any {
	key := v.Visit(ctx.GetK()).(Identifier)
	value := v.Visit(ctx.GetV()).(GenericType)

	return &TypeField{
		Name: key,
		Type: value,
	}
}

func (v *Visitor) VisitConstant(ctx *parser.ConstantContext) any {
	switch {
	case ctx.STRING_LITERAL() != nil:
		str := ctx.GetText()

		if str[0] == '$' {
			return str[2 : len(str)-2]
		} else {
			return str[1 : len(str)-1]
		}
	case ctx.INTEGER() != nil:
		str := ctx.GetText()

		val, err := strconv.Atoi(str)
		if err != nil {
			v.Err(ctx, "could not parse integer")
			return 0
		}

		return val
	case ctx.FLOAT() != nil:
		str := ctx.GetText()

		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			v.Err(ctx, "could not parse float")
			return nil
		}

		return val
	case ctx.BOOLEAN() != nil:
		return strings.ToLower(ctx.GetText()) == "true"
	case ctx.DURATION() != nil:
		// TODO: Implement duration parsing.
		v.Err(ctx, "duration literal not implemented")
		return nil
	case ctx.UUID() != nil:
		id, err := uuid.Parse(ctx.GetText())
		if err != nil {
			v.Err(ctx, "could not parse uuid")
			return nil
		}

		return id
	case ctx.HEXNUMBER() != nil:
		str := ctx.GetText()
		val, err := hex.DecodeString(str[2:])
		if err != nil {
			v.Err(ctx, "could not decode hex string")
			return nil
		}

		return val
	case ctx.K_POSITIVE_NAN() != nil:
		return math.NaN()
	case ctx.K_NEGATIVE_NAN() != nil:
		return -math.NaN()
	case ctx.K_POSITIVE_INFINITY() != nil:
		return math.Inf(1)
	case ctx.K_NEGATIVE_INFINITY() != nil:
		return math.Inf(-1)
	}
	return nil
}

func (v *Visitor) VisitMarker(ctx *parser.MarkerContext) any {
	if ctx.QMARK() != nil {
		return &Marker{}
	}

	id := v.Visit(ctx.GetId()).(Identifier)
	return &Marker{
		Name: id,
	}
}

func (v *Visitor) VisitColumnMask(ctx *parser.ColumnMaskContext) any {
	if ctx.K_DEFAULT() != nil {
		return &Function{
			Name: &ObjectRef{
				Name: UnquotedIdentifier("mask_default"),
			},
		}
	}

	name := v.Visit(ctx.GetName()).(*ObjectRef)
	args := v.Visit(ctx.ColumnMaskArguments()).([]Term)

	return &Function{
		Name:      name,
		Arguments: args,
	}
}

func (v *Visitor) VisitColumnMaskArguments(ctx *parser.ColumnMaskArgumentsContext) any {
	var terms []Term
	for _, t := range ctx.AllTerm() {
		terms = append(terms, v.Visit(t))
	}

	return terms
}

func visitIdent(ctx antlr.ParserRuleContext) Identifier {
	id := ctx.GetText()

	return UnquotedIdentifier(id)
}

func visitQuotedIdent(ctx antlr.ParserRuleContext) Identifier {
	qid := ctx.GetText()
	id := qid[1 : len(qid)-1]

	return QuotedIdentifier(id)
}
