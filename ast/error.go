package ast

import (
	"fmt"
)

type visitorError struct {
	line, col int
	msg       string
}

func (e *visitorError) Line() int {
	return e.line
}

func (e *visitorError) Col() int {
	return e.col
}

func (e *visitorError) Error() string {
	return fmt.Sprintf("line %d, column %d: %s", e.line, e.col, e.msg)
}
