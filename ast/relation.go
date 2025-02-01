package ast

type RelationType string

const (
	EqualityRelationType     RelationType = "="
	LesserRelationType       RelationType = "<"
	LesserEqualRelationType  RelationType = "<="
	GreaterRelationType      RelationType = ">"
	GreaterEqualRelationType RelationType = ">="
	BetweenRelationType      RelationType = "between"
	LikeRelationType         RelationType = "like"
	IsRelationType           RelationType = "is"
	InRelationType           RelationType = "in"
	ContainsRelationType     RelationType = "contains"
	ContainsKeyRelationType  RelationType = "contains_key"
)

var validRelationTypes = map[RelationType]bool{
	EqualityRelationType:     true,
	GreaterRelationType:      true,
	GreaterEqualRelationType: true,
	LesserRelationType:       true,
	LesserEqualRelationType:  true,
	BetweenRelationType:      true,
	LikeRelationType:         true,
	IsRelationType:           true,
	InRelationType:           true,
	ContainsRelationType:     true,
}

type Relation struct {
	Type   RelationType
	Invert bool
	Left   Term
	Right  Term
}

func getRelationType(rel string) (RelationType, bool) {
	if rel == "!=" {
		return EqualityRelationType, true
	}

	return RelationType(rel), false
}
