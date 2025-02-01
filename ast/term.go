package ast

type (
	Term            interface{}
	ListLiteral     []any
	TupleLiteral    []any
	SetLiteral      []any
	MapLiteral      map[any]any
	UserTypeLiteral map[any]any
	nullLiteral     struct{}
)

var NullLiteral nullLiteral

type Operator string

const (
	LiteralCastOperator    Operator = "literal_cast"
	CastOperator           Operator = "cast"
	AdditionOperator       Operator = "+"
	SubtractionOperator    Operator = "-"
	MultiplicationOperator Operator = "*"
	DivisionOperator       Operator = "/"
	RemainderOperator      Operator = "%"
)

type TermOperation struct {
	Operator Operator
	Left     Term
	Right    Term
}

type SimpleTerm struct{}

type Function struct {
	Name      *ObjectRef
	Arguments []Term
}
