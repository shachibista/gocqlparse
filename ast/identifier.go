package ast

import (
	"fmt"
	"strings"
)

type Identifier interface {
	fmt.Stringer
}
type (
	QuotedIdentifier   string
	UnquotedIdentifier string
)

func (qi QuotedIdentifier) String() string {
	return string(qi)
}

func (ui UnquotedIdentifier) String() string {
	return strings.ToLower(string(ui))
}
