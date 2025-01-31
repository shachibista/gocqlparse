package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	var id1, id2 Identifier

	id1 = QuotedIdentifier("foo")
	id2 = QuotedIdentifier("foo")

	assert.True(t, id1.String() == id2.String())

	id1 = QuotedIdentifier("FOO")
	id2 = QuotedIdentifier("FOO")

	assert.True(t, id1.String() == id2.String())

	id1 = QuotedIdentifier("FOO")
	id2 = QuotedIdentifier("foo")

	assert.False(t, id1.String() == id2.String())

	id1 = QuotedIdentifier("FOO")
	id2 = UnquotedIdentifier("FOO")

	assert.False(t, id1.String() == id2.String())
}
