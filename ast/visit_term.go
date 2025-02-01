package ast

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/shachibista/gocqlparse/internal/parser"
)

func (v *Visitor) VisitTerm(ctx *parser.TermContext) any {
	return v.Visit(ctx.GetT())
}

func (v *Visitor) VisitTerms(ctx *parser.TermsContext) any {
	var terms []Term

	for _, t := range ctx.AllTerm() {
		terms = append(terms, v.Visit(t))
	}

	return terms
}

func (v *Visitor) VisitTermAddition(ctx *parser.TermAdditionContext) any {
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

func (v *Visitor) VisitTermMultiplication(ctx *parser.TermMultiplicationContext) any {
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

func (v *Visitor) VisitTermGroup(ctx *parser.TermGroupContext) any {
	return v.Visit(ctx.GetT())
}

func (v *Visitor) VisitSimpleTerm(ctx *parser.SimpleTermContext) any {
	switch {
	case ctx.GetV() != nil:
		return v.Visit(ctx.GetV())
	case ctx.GetF() != nil:
		return v.Visit(ctx.GetF())
	case ctx.GetC() != nil:
		return &TermBinaryOperation{
			Operator: LiteralCastOperator,
			Left:     v.Visit(ctx.GetC()),
			Right:    v.Visit(ctx.GetT()),
		}
	case ctx.K_CAST() != nil:
		return &TermBinaryOperation{
			Operator: CastOperator,
			Left:     v.Visit(ctx.GetT()),
			Right:    v.Visit(ctx.GetN()),
		}
	default:
	}

	v.Err(ctx, "unrecognized term")

	return nil
}

func (v *Visitor) VisitValue(ctx *parser.ValueContext) any {
	switch {
	case ctx.GetC() != nil:
		return v.Visit(ctx.GetC())
	case ctx.GetL() != nil:
		return v.Visit(ctx.GetL())
	case ctx.GetU() != nil:
		return v.Visit(ctx.GetU())
	case ctx.GetT() != nil:
		return v.Visit(ctx.GetT())
	case ctx.K_NULL() != nil:
		return NullLiteral
	case ctx.GetM() != nil:
		return v.Visit(ctx.GetM())
	}
	return nil
}

func (v *Visitor) VisitCollectionLiteral(ctx *parser.CollectionLiteralContext) any {
	if ctx.GetL() != nil {
		var elements ListLiteral

		for _, e := range ctx.GetL().AllTerm() {
			elements = append(elements, v.Visit(e))
		}

		return elements
	}

	if ctx.GetV() != nil {
		val := ctx.GetV()
		m := val.GetM()
		s := val.GetS()

		if m != nil {
			firstK := v.Visit(ctx.GetT())

			var terms []any
			for _, t := range m.AllTerm() {
				terms = append(terms, v.Visit(t))
			}

			elements := MapLiteral{
				firstK: terms[0],
			}

			if len(terms[1:])%2 != 0 {
				v.Err(ctx, "unpaired key-value pair")
				return nil
			}

			for i := 1; i < len(terms)-1; i += 2 {
				elements[terms[i]] = terms[i+1]
			}

			return elements
		} else if s != nil {
			elements := SetLiteral{
				v.Visit(ctx.GetT()),
			}

			for _, t := range s.AllTerm() {
				elements = append(elements, v.Visit(t))
			}

			return elements
		} else {
			v.Err(ctx, "unrecognized literal")

			return nil
		}
	}

	return SetLiteral{}
}

func (v *Visitor) VisitTupleLiteral(ctx *parser.TupleLiteralContext) any {
	tup := TupleLiteral{}

	for _, t := range ctx.AllTerm() {
		tup = append(tup, v.Visit(t))
	}

	return &tup
}

func (v *Visitor) VisitUsertypeLiteral(ctx *parser.UsertypeLiteralContext) any {
	allK := ctx.AllFident()
	allV := ctx.AllTerm()

	elements := UserTypeLiteral{}

	if len(allK) != len(allV) {
		v.Err(ctx, "unpaired key-value pair")
		return nil
	}

	for i := range len(allK) {
		k := v.Visit(allK[i])
		v := v.Visit(allV[i])

		elements[k] = v
	}

	return &elements
}

func (v *Visitor) VisitFunction(ctx *parser.FunctionContext) any {
	fn := v.Visit(ctx.GetF()).(*ObjectRef)

	var args []Term

	if ctx.GetArgs() != nil {
		for _, arg := range ctx.GetArgs().AllTerm() {
			args = append(args, v.Visit(arg))
		}
	}

	return &Function{
		Name:      fn,
		Arguments: args,
	}
}
