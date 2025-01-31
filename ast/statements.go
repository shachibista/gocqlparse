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

type (
	ListUsersStatement      struct{}
	ListSuperUsersStatement struct{}
)

type DropKeyspaceStatement struct {
	IfExists bool
	Name     Identifier
}

type DropIdentityStatement struct {
	IfExists bool
	Identity Identifier
}

type AddIdentityStatement struct {
	IfNotExists bool
	Identity    Identifier
	Role        Identifier
}

type DropUserStatement struct {
	IfExists bool
	Username Identifier
}

type DropIndexStatement struct {
	IfExists bool
	Name     *ObjectRef
}

type DropMaterializedViewStatement struct {
	IfExists bool
	Name     *ObjectRef
}

type GrantRoleStatement struct {
	Role    Identifier
	Grantee Identifier
}

type RevokeRoleStatement struct {
	Role    Identifier
	Revokee Identifier
}

type ListRolesStatement struct {
	Role      Identifier
	Recursive bool
}

type DropRoleStatement struct {
	IfExists bool
	Role     Identifier
}
