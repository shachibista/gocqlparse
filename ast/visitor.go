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

	return &Property{
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

	return elements
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
		return unwrapStr(ctx.GetText())
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

func (v *Visitor) VisitProperties(ctx *parser.PropertiesContext) any {
	var properties []*Property

	for _, p := range ctx.AllProperty() {
		properties = append(properties, v.Visit(p).(*Property))
	}

	return properties
}

func (v *Visitor) VisitIdentity(ctx *parser.IdentityContext) any {
	if ctx.IDENT() != nil {
		return visitIdent(ctx)
	}

	if ctx.STRING_LITERAL() != nil {
		return visitQuotedIdent(ctx)
	}

	v.Err(ctx, "quoted strings are not supported for identity")
	return nil
}

func (v *Visitor) VisitUsername(ctx *parser.UsernameContext) any {
	if ctx.IDENT() != nil {
		return visitIdent(ctx)
	}

	if ctx.STRING_LITERAL() != nil {
		return visitQuotedIdent(ctx)
	}

	v.Err(ctx, "quoted strings are not supported for username")
	return nil
}

func (v *Visitor) VisitIndexName(ctx *parser.IndexNameContext) any {
	var keyspace Identifier
	if ctx.KsName() != nil {
		keyspace = v.Visit(ctx.KsName()).(Identifier)
	}

	ident := v.Visit(ctx.IdxName()).(Identifier)

	return &ObjectRef{
		Keyspace: keyspace,
		Name:     ident,
	}
}

func (v *Visitor) VisitIdxName(ctx *parser.IdxNameContext) any {
	if ctx.IDENT() != nil {
		return visitIdent(ctx)
	}

	if ctx.QUOTED_NAME() != nil {
		return visitQuotedIdent(ctx)
	}

	if ctx.Unreserved_keyword() != nil {
		return visitIdent(ctx)
	}

	v.Err(ctx, "bind variables cannot be used for index names")
	return nil
}

func (v *Visitor) VisitRoleName(ctx *parser.RoleNameContext) any {
	if ctx.IDENT() != nil {
		return visitIdent(ctx)
	}

	if ctx.STRING_LITERAL() != nil {
		return visitQuotedIdent(ctx)
	}

	if ctx.QUOTED_NAME() != nil {
		return visitQuotedIdent(ctx)
	}

	if ctx.Unreserved_keyword() != nil {
		return visitIdent(ctx)
	}

	v.Err(ctx, "bind variables cannot be used for role names")
	return nil
}

func (v *Visitor) VisitUserOrRoleName(ctx *parser.UserOrRoleNameContext) any {
	return v.Visit(ctx.RoleName())
}

func (v *Visitor) VisitUserPassword(ctx *parser.UserPasswordContext) any {
	if ctx.K_GENERATED() != nil {
		return GeneratedPassword
	}

	if ctx.K_HASHED() != nil {
		return HashedPassword(unwrapStr(ctx.GetV().GetText()))
	}

	if ctx.K_PASSWORD() != nil {
		return PlainPassword(unwrapStr(ctx.GetV().GetText()))
	}

	v.Err(ctx, "unrecognized password setting")
	return nil
}

func (v *Visitor) VisitSident(ctx *parser.SidentContext) any {
	if ctx.IDENT() != nil {
		return visitIdent(ctx)
	}

	if ctx.QUOTED_NAME() != nil {
		return visitQuotedIdent(ctx)
	}

	return visitIdent(ctx)
}

func (v *Visitor) VisitCident(ctx *parser.CidentContext) any {
	if ctx.EMPTY_QUOTED_NAME() != nil {
		return UnquotedIdentifier("")
	}

	return v.Visit(ctx.Ident())
}

func (v *Visitor) VisitSelectionLiteral(ctx *parser.SelectionLiteralContext) any {
	switch {
	case ctx.GetC() != nil:
		return v.Visit(ctx.GetC())
	case ctx.K_NULL() != nil:
		return NullLiteral
	case ctx.GetM() != nil:
		return v.Visit(ctx.GetM())
	}

	v.Err(ctx, "unrecognized literal")
	return nil
}

func (v *Visitor) VisitSimpleUnaliasedSelector(ctx *parser.SimpleUnaliasedSelectorContext) any {
	switch {
	case ctx.GetC() != nil:
		return &ColumnSelector{Name: v.Visit(ctx.GetC()).(Identifier)}
	case ctx.GetL() != nil:
		return v.Visit(ctx.GetL())
	case ctx.GetF() != nil:
		return v.Visit(ctx.GetF())
	}

	v.Err(ctx, "unrecognized selector")
	return nil
}

func (v *Visitor) VisitSelectionFunction(ctx *parser.SelectionFunctionContext) any {
	if ctx.K_COUNT() != nil {
		return &Function{
			Name: &ObjectRef{Name: UnquotedIdentifier("count")},
			Arguments: []Term{
				"*",
			},
		}
	}

	if ctx.K_MAXWRITETIME() != nil {
		var modifiers []SelectorModifier

		if len(ctx.AllSelectorModifier()) > 0 {
			for _, m := range ctx.AllSelectorModifier() {
				modifiers = append(modifiers, v.Visit(m))
			}
		}

		return &Function{
			Name: &ObjectRef{Name: UnquotedIdentifier("maxwritetime")},
			Arguments: []Term{
				&ColumnSelector{
					Name:      v.Visit(ctx.GetC()).(Identifier),
					Modifiers: modifiers,
				},
			},
		}
	}

	if ctx.K_WRITETIME() != nil {
		var modifiers []SelectorModifier

		if len(ctx.AllSelectorModifier()) > 0 {
			for _, m := range ctx.AllSelectorModifier() {
				modifiers = append(modifiers, v.Visit(m))
			}
		}

		return &Function{
			Name: &ObjectRef{Name: UnquotedIdentifier("writetime")},
			Arguments: []Term{
				&ColumnSelector{
					Name:      v.Visit(ctx.GetC()).(Identifier),
					Modifiers: modifiers,
				},
			},
		}
	}

	if ctx.K_TTL() != nil {
		var modifiers []SelectorModifier

		if len(ctx.AllSelectorModifier()) > 0 {
			for _, m := range ctx.AllSelectorModifier() {
				modifiers = append(modifiers, v.Visit(m))
			}
		}

		return &Function{
			Name: &ObjectRef{Name: UnquotedIdentifier("ttl")},
			Arguments: []Term{
				&ColumnSelector{
					Name:      v.Visit(ctx.GetC()).(Identifier),
					Modifiers: modifiers,
				},
			},
		}
	}

	if ctx.K_CAST() != nil {
		return &TermBinaryOperation{
			Operator: CastOperator,
			Left:     v.Visit(ctx.GetSn()),
			Right:    v.Visit(ctx.GetT()),
		}
	}

	return &Function{
		Name:      v.Visit(ctx.GetF()).(*ObjectRef),
		Arguments: v.Visit(ctx.GetArgs()).([]Term),
	}
}

func (v *Visitor) VisitSelectionFunctionArgs(ctx *parser.SelectionFunctionArgsContext) any {
	var args []Term

	for _, s := range ctx.AllUnaliasedSelector() {
		args = append(args, v.Visit(s))
	}

	return args
}

func (v *Visitor) VisitSelectionTypeHint(ctx *parser.SelectionTypeHintContext) any {
	return &TermBinaryOperation{
		Operator: TypeHint,
		Left:     v.Visit(ctx.GetCt()),
		Right:    v.Visit(ctx.GetA()),
	}
}

func (v *Visitor) VisitSelectorModifier(ctx *parser.SelectorModifierContext) any {
	if ctx.GetFi() != nil {
		return &FieldSelector{
			Name: v.Visit(ctx.GetFi()).(Identifier),
		}
	}

	if ctx.GetSs() != nil {
		return v.Visit(ctx.GetSs())
	}

	v.Err(ctx, "unrecognized selector modifier")
	return nil
}

func (v *Visitor) VisitSelectTupleOrNestedSelector(ctx *parser.SelectionTupleOrNestedSelectorContext) any {
	var elements TupleLiteral

	for _, e := range ctx.AllUnaliasedSelector() {
		elements = append(elements, v.Visit(e))
	}

	return elements
}

func (v *Visitor) VisitSelectionList(ctx *parser.SelectionListContext) any {
	var elements ListLiteral

	for _, e := range ctx.AllUnaliasedSelector() {
		elements = append(elements, v.Visit(e))
	}

	return elements
}

func (v *Visitor) VisitSelectionMapOrSet(ctx *parser.SelectionMapOrSetContext) any {
	firstTerm := v.Visit(ctx.GetT1())

	if ctx.SelectionSet() != nil {
		elements := SetLiteral{firstTerm}

		for _, e := range ctx.SelectionSet().AllUnaliasedSelector() {
			elements = append(elements, v.Visit(e))
		}

		return elements
	}

	if ctx.SelectionMap() != nil {
		var terms []any
		for _, t := range ctx.SelectionMap().AllUnaliasedSelector() {
			terms = append(terms, v.Visit(t))
		}

		elements := MapLiteral{
			firstTerm: terms[0],
		}

		if len(terms[1:])%2 != 0 {
			v.Err(ctx, "unpaired key-value pair")
			return nil
		}

		for i := 1; i < len(terms)-1; i += 2 {
			elements[terms[i]] = terms[i+1]
		}

		return elements
	}

	return SetLiteral{}
}

func (v *Visitor) VisitCollectionSubSelection(ctx *parser.CollectionSubSelectionContext) any {
	if ctx.RANGE() == nil {
		return &CollectionSelector{
			Index: v.Visit(ctx.GetT1()),
		}
	}

	var left, right Term

	if ctx.GetT1() != nil {
		// a.. | a..b

		left = v.Visit(ctx.GetT1())

		if ctx.GetT2() != nil {
			right = v.Visit(ctx.GetT2())
		}
	} else {
		// ..b
		right = v.Visit(ctx.GetT2())
	}

	return &CollectionSelector{
		Range: &ListSliceRange{
			Left:  left,
			Right: right,
		},
	}
}

func (v *Visitor) VisitUnaliasedSelector(ctx *parser.UnaliasedSelectorContext) any {
	return v.Visit(ctx.SelectionAddition())
}

func (v *Visitor) VisitSelectionAddition(ctx *parser.SelectionAdditionContext) any {
	lterm := v.Visit(ctx.GetL())

	if ctx.GetR() == nil {
		return lterm
	}

	rterm := v.Visit(ctx.GetR())

	otype := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if otype != "+" && otype != "-" {
		v.Err(ctx, "unrecognized operator: "+otype)
		return nil
	}

	return &TermBinaryOperation{
		Operator: Operator(otype),
		Left:     lterm,
		Right:    rterm,
	}
}

func (v *Visitor) VisitSelectionMultiplication(ctx *parser.SelectionMultiplicationContext) any {
	lterm := v.Visit(ctx.GetL())

	if ctx.GetR() == nil {
		return lterm
	}

	rterm := v.Visit(ctx.GetR())

	otype := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if otype != "*" && otype != "/" && otype != "%" {
		v.Err(ctx, "unrecognized operation: "+otype)
		return nil
	}

	return &TermBinaryOperation{
		Operator: Operator(otype),
		Left:     lterm,
		Right:    rterm,
	}
}

func (v *Visitor) VisitSelectionGroup(ctx *parser.SelectionGroupContext) any {
	if ctx.GetF() != nil {
		return v.Visit(ctx.GetF())
	}

	if ctx.GetG() != nil {
		return v.Visit(ctx.GetG())
	}

	return &TermUnaryOperation{
		Operator: Operator("-"),
		Operand:  v.Visit(ctx.GetH()),
	}
}

func (v *Visitor) VisitSelectionTupleOrNestedSelector(ctx *parser.SelectionTupleOrNestedSelectorContext) any {
	elements := ctx.AllUnaliasedSelector()

	// If there is a single element, this is a grouping and we don't need to wrap it.
	if len(elements) == 1 {
		return v.Visit(elements[0])
	}

	var group TupleLiteral

	for _, s := range ctx.AllUnaliasedSelector() {
		group = append(group, v.Visit(s))
	}

	return group
}

func (v *Visitor) VisitCustomExpression(ctx *parser.CustomIndexExpressionContext) any {
	idxName := v.Visit(ctx.IdxName())
	arguments := v.Visit(ctx.Term())

	return &Function{
		Name: &ObjectRef{Name: UnquotedIdentifier("expr")},
		Arguments: []Term{
			idxName,
			arguments,
		},
	}
}

func (v *Visitor) VisitRelationRecursive(ctx *parser.RelationRecursiveContext) any {
	return v.Visit(ctx.Relation())
}

func (v *Visitor) VisitTupleOfIdentifiers(ctx *parser.TupleOfIdentifiersContext) any {
	var identifiers TupleLiteral

	for _, id := range ctx.AllCident() {
		identifiers = append(identifiers, v.Visit(id))
	}

	return identifiers
}

func (v *Visitor) VisitSingleColumnBetweenValues(ctx *parser.SingleColumnBetweenValuesContext) any {
	var terms []Term

	for _, t := range ctx.AllTerm() {
		terms = append(terms, v.Visit(t))
	}

	return terms
}

func (v *Visitor) VisitSingleColumnInValues(ctx *parser.SingleColumnInValuesContext) any {
	if ctx.GetT() != nil {
		var tup TupleLiteral
		for _, t := range ctx.GetT().AllTerm() {
			tup = append(tup, v.Visit(t))
		}

		return &tup
	}

	if ctx.GetM() != nil {
		return v.Visit(ctx.GetM())
	}

	v.Err(ctx, "unrecognized expression")
	return nil
}

func (v *Visitor) VisitRelationColumn(ctx *parser.RelationColumnContext) any {
	lterm := &ColumnSelector{Name: v.Visit(ctx.GetName()).(Identifier)}

	switch {
	case ctx.GetType_() != nil:
		rel, invert := getRelationType(ctx.GetType_().GetText())
		rterm := v.Visit(ctx.GetT())

		if !validRelationTypes[rel] {
			v.Err(ctx, "unrecognized relation: "+ctx.GetType_().GetText())

			return nil
		}

		return &Relation{
			Type:   rel,
			Invert: invert,
			Left:   lterm,
			Right:  rterm,
		}
	case ctx.K_BETWEEN() != nil:
		rterm := v.Visit(ctx.GetBetweenValues())

		return &Relation{
			Type:  BetweenRelationType,
			Left:  lterm,
			Right: rterm,
		}
	case ctx.K_LIKE() != nil:
		rterm := v.Visit(ctx.GetT())

		return &Relation{
			Type:  LikeRelationType,
			Left:  lterm,
			Right: rterm,
		}
	case ctx.K_NULL() != nil:
		return &Relation{
			Type:   IsRelationType,
			Invert: true,
			Left:   lterm,
			Right:  NullLiteral,
		}
	case ctx.InOperator() != nil:
		rterm := v.Visit(ctx.GetInValue())

		return &Relation{
			Type:   InRelationType,
			Invert: ctx.InOperator().K_NOT() != nil,
			Left:   lterm,
			Right:  rterm,
		}
	case ctx.ContainsOperator() != nil:
		rterm := v.Visit(ctx.GetT())

		var rel RelationType
		if ctx.ContainsOperator().K_KEY() != nil {
			rel = ContainsKeyRelationType
		} else {
			rel = ContainsRelationType
		}

		return &Relation{
			Type:  rel,
			Left:  lterm,
			Right: rterm,
		}
	}

	v.Err(ctx, "unrecognized column relation")
	return nil
}

func (v *Visitor) VisitRelationCollection(ctx *parser.RelationCollectionContext) any {
	lterm := &ColumnSelector{
		Name: v.Visit(ctx.GetName()).(Identifier),
		Modifiers: []SelectorModifier{
			&CollectionSelector{
				Index: v.Visit(ctx.GetKey()),
			},
		},
	}

	rel, invert := getRelationType(ctx.GetType_().GetText())

	if !validRelationTypes[rel] {
		v.Err(ctx, "unrecognized relation: "+ctx.GetType_().GetText())

		return nil
	}

	rterm := v.Visit(ctx.GetT())

	return &Relation{
		Type:   rel,
		Invert: invert,
		Left:   lterm,
		Right:  rterm,
	}
}

func (v *Visitor) VisitRelationToken(ctx *parser.RelationTokenContext) any {
	arguments := v.Visit(ctx.TupleOfIdentifiers()).(TupleLiteral)

	var t []Term
	for _, a := range arguments {
		t = append(t, a)
	}

	lterm := &Function{
		Name:      &ObjectRef{Name: UnquotedIdentifier("token")},
		Arguments: t,
	}

	switch {
	case ctx.GetType_() != nil:
		rel, invert := getRelationType(ctx.GetType_().GetText())
		rterm := v.Visit(ctx.GetT())

		if !validRelationTypes[rel] {
			v.Err(ctx, "unrecognized relation: "+ctx.GetType_().GetText())

			return nil
		}

		return &Relation{
			Type:   rel,
			Invert: invert,
			Left:   lterm,
			Right:  rterm,
		}
	case ctx.K_BETWEEN() != nil:
		rterm := v.Visit(ctx.GetBetweenValues())

		return &Relation{
			Type:  BetweenRelationType,
			Left:  lterm,
			Right: rterm,
		}
	}

	v.Err(ctx, "unrecognized token relation")
	return nil
}

func (v *Visitor) VisitRelationTuple(ctx *parser.RelationTupleContext) any {
	ids := v.Visit(ctx.GetIds()).([]Identifier)

	lterm := TupleSelector{}
	for _, id := range ids {
		lterm.Elements = append(lterm.Elements, id)
	}

	switch {
	case ctx.InOperator() != nil:
		rterm := v.Visit(ctx.GetMinValue())

		return &Relation{
			Type:   InRelationType,
			Invert: ctx.InOperator().K_NOT() != nil,
			Left:   lterm,
			Right:  rterm,
		}
	case ctx.RelationType() != nil:
		rel, invert := getRelationType(ctx.GetType_().GetText())
		rterm := v.Visit(ctx.GetV())

		if !validRelationTypes[rel] {
			v.Err(ctx, "unrecognized relation: "+ctx.GetType_().GetText())

			return nil
		}

		return &Relation{
			Type:   rel,
			Invert: invert,
			Left:   lterm,
			Right:  rterm,
		}
	case ctx.K_BETWEEN() != nil:
		lterm := v.Visit(ctx.GetT1())
		rterm := v.Visit(ctx.GetT2())

		return &Relation{
			Type:  BetweenRelationType,
			Left:  lterm,
			Right: rterm,
		}
	}

	return nil
}

func (v *Visitor) VisitMultiColumnValue(ctx *parser.MultiColumnValueContext) any {
	if ctx.GetL() != nil {
		return v.Visit(ctx.GetL())
	}

	if ctx.GetM() != nil {
		return v.Visit(ctx.GetM())
	}

	v.Err(ctx, "unrecognized multi column value")
	return nil
}

func (v *Visitor) VisitMultiColumnInValues(ctx *parser.MultiColumnInValuesContext) any {
	if ctx.GetM() != nil {
		return v.Visit(ctx.GetM())
	}

	tup := TupleLiteral{}

	if ctx.GetTl() != nil {
		return v.Visit(ctx.GetTl())
	}

	if ctx.GetTm() != nil {
		for _, m := range ctx.GetTm().AllMarker() {
			tup = append(tup, v.Visit(m))
		}
	}

	return &tup
}

func (v *Visitor) VisitTupleOfTupleLiterals(ctx *parser.TupleOfTupleLiteralsContext) any {
	tup := TupleLiteral{}

	for _, tt := range ctx.AllTupleLiteral() {
		tup = append(tup, v.Visit(tt))
	}

	return tup
}

func (v *Visitor) VisitTupleOfMarkersForTuples(ctx *parser.TupleOfMarkersForTuplesContext) any {
	tup := TupleLiteral{}

	for _, tm := range ctx.AllMarker() {
		tup = append(tup, v.Visit(tm))
	}

	return tup
}

func (v *Visitor) VisitRelationOrExpression(ctx *parser.RelationOrExpressionContext) any {
	if ctx.Relation() != nil {
		return v.Visit(ctx.Relation())
	}

	if ctx.CustomIndexExpression() != nil {
		return v.Visit(ctx.CustomIndexExpression())
	}

	v.Err(ctx, "unrecognized relation or expression")
	return nil
}

func (v *Visitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) any {
	col := v.Visit(ctx.GetC()).(Identifier)
	asc := ctx.K_DESC() == nil

	var ann Term
	if ctx.K_ANN() != nil {
		ann = v.Visit(ctx.GetT())
	}

	return &OrderByClause{
		Column:                       col,
		Ascending:                    asc,
		ApproximateNearestNeighborOf: ann,
	}
}

func (v *Visitor) VisitIntValue(ctx *parser.IntValueContext) any {
	if ctx.GetM() != nil {
		return v.Visit(ctx.GetM())
	}

	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		v.Err(ctx, "bad int value")
		return nil
	}

	return i
}

func (v *Visitor) VisitSelectionGroupWithoutField(ctx *parser.SelectionGroupWithoutFieldContext) any {
	if ctx.GetSn() != nil {
		return v.Visit(ctx.GetSn())
	}

	if ctx.GetH() != nil {
		return v.Visit(ctx.GetH())
	}

	if ctx.GetT() != nil {
		return v.Visit(ctx.GetT())
	}

	if ctx.GetL() != nil {
		return v.Visit(ctx.GetL())
	}

	if ctx.GetM() != nil {
		return v.Visit(ctx.GetM())
	}

	v.Err(ctx, "unrecognized selection group")
	return nil
}

func (v *Visitor) VisitSelectionGroupWithField(ctx *parser.SelectionGroupWithFieldContext) any {
	var modifiers []SelectorModifier

	for _, m := range ctx.AllSelectorModifier() {
		modifiers = append(modifiers, v.Visit(m))
	}

	sel := v.Visit(ctx.GetG()).(*ColumnSelector)
	sel.Modifiers = modifiers

	return sel
}

func (v *Visitor) VisitSelectors(ctx *parser.SelectorsContext) any {
	var projections []*Projection

	selectors := ctx.AllSelector()

	if len(selectors) > 0 {
		for _, s := range selectors {
			projections = append(projections, v.Visit(s).(*Projection))
		}
	} else {
		projections = append(projections, &Projection{Expression: "*"})
	}

	return projections
}

func (v *Visitor) VisitSelector(ctx *parser.SelectorContext) any {
	var alias Identifier
	if ctx.GetC() != nil {
		alias = v.Visit(ctx.GetC()).(Identifier)
	}

	return &Projection{
		Expression: v.Visit(ctx.GetUs()),
		Alias:      alias,
	}
}

func visitIdent(ctx antlr.ParserRuleContext) Identifier {
	id := ctx.GetText()

	return UnquotedIdentifier(id)
}

func visitQuotedIdent(ctx antlr.ParserRuleContext) Identifier {
	return QuotedIdentifier(unwrapStr(ctx.GetText()))
}

func unwrapStr(src string) string {
	if src[0] == '$' {
		return src[2 : len(src)-2]
	} else {
		return src[1 : len(src)-1]
	}
}
