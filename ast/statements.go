package ast

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

	Property struct {
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
	Properties      []*Property
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

type DropTableStatement struct {
	IfExists     bool
	ColumnFamily *ObjectRef
}

type TruncateStatement struct {
	ColumnFamily *ObjectRef
}

type CreateKeyspaceStatement struct {
	IfNotExists bool
	Name        Identifier
	Properties  []*Property
}

type DropTypeStatement struct {
	IfExists bool
	Name     *ObjectRef
}
