package ast

type (
	Password          interface{}
	PlainPassword     string
	HashedPassword    string
	generatedPassword struct{}
)

var GeneratedPassword generatedPassword
