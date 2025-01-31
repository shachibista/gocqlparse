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

type DropTableStatement struct {
	IfExists     bool
	ColumnFamily *ObjectRef
}
