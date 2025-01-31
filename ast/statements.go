package ast

import (
	"fmt"
	"strings"
)

type (
	QuotedIdentifier   string
	UnquotedIdentifier string
)

type Identifier interface {
	fmt.Stringer
}

type (
	TypeField struct {
		Name Identifier
		Type GenericType
	}

	TableColumn struct {
		Name               Identifier
		Type               GenericType
		Static             bool
		ImplicitPrimaryKey bool
		Mask               *Function
	}

	TableProperty struct {
		Key   Identifier
		Value any
	}
)

type TableKeys struct {
	Partition  []Identifier
	Clustering []Identifier
}

type ColumnOrder struct {
	Name      Identifier
	Ascending bool
}

type CreateTableStatement struct {
	IfNotExists     bool
	ColumnFamily    *ObjectRef
	Columns         []*TableColumn
	PartitionKeys   []Identifier
	ClusteringKeys  []Identifier
	Properties      []*TableProperty
	CompactStorage  bool
	ClusteringOrder []*ColumnOrder
}

type UseStatement struct {
	Keyspace Identifier
}

type CreateTypeStatement struct {
	IfNotExists bool
	Name        *ObjectRef
	Fields      []*TypeField
}

func (qi QuotedIdentifier) String() string {
	return string(qi)
}

func (ui UnquotedIdentifier) String() string {
	return strings.ToLower(string(ui))
}
