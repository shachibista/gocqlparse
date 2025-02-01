package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/shachibista/gocqlparse/ast"
	"github.com/shachibista/gocqlparse/internal/parser"
)

type Parser struct {
	*parser.Cql3Parser
	lexer  *parser.Cql3Lexer
	tokens *antlr.CommonTokenStream
}

type syntaxError struct {
	line, col int
	msg       string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("Syntax error at line %d:%d - %s", e.line, e.col, e.msg)
}

type cql3ErrorListener struct {
	*antlr.DefaultErrorListener
	err error
}

var _ parser.Cql3Visitor = &ast.Visitor{}

func newCql3ErrorListener() *cql3ErrorListener {
	return &cql3ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *cql3ErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line, col int,
	msg string,
	ex antlr.RecognitionException,
) {
	panic(&syntaxError{line, col, msg})
}

func NewParser(input string) *Parser {
	lexer := parser.NewCql3Lexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newCql3ErrorListener())

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCql3Parser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(newCql3ErrorListener())

	return &Parser{
		p,
		lexer,
		tokens,
	}
}

func parseRule[T any](rule func() antlr.ParseTree) (st T, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(*syntaxError); ok {
				var z T
				st, err = z, rErr
			} else {
				panic(r)
			}
		}
	}()

	visitor := ast.NewVisitor()
	stu := visitor.Visit(rule())
	err = visitor.Errs()
	if err != nil {
		return st, err
	}

	return stu.(T), nil
}
