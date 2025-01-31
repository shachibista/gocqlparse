package ast

type (
	Term            interface{}
	SetLiteral      []any
	MapLiteral      map[any]any
	UserTypeLiteral map[any]any
	nullLiteral     struct{}
)

var NullLiteral nullLiteral

type TermOperator string

const (
	LiteralCastTermOperator    TermOperator = "literal_cast"
	CastTermOperator           TermOperator = "cast"
	AdditionTermOperator       TermOperator = "+"
	SubtractionTermOperator    TermOperator = "-"
	MultiplicationTermOperator TermOperator = "*"
	DivisionTermOperator       TermOperator = "/"
	RemainderTermOperator      TermOperator = "%"
)

type TermOperation struct {
	Operator TermOperator
	Left     Term
	Right    Term
}

type SimpleTerm struct{}

type Function struct {
	Name      *ObjectRef
	Arguments []Term
}
