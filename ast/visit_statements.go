package ast

import (
	"github.com/shachibista/gocqlparse/internal/parser"
)

func (v *Visitor) VisitCreateTableStatement(ctx *parser.CreateTableStatementContext) any {
	cts := &CreateTableStatement{}
	tableDef := ctx.TableDefinition()

	var implictlyDeclaredPK []*TableColumn
	var columns []*TableColumn
	keys := &TableKeys{}
	for _, tc := range tableDef.AllTableColumns() {
		c := v.Visit(tc)

		switch ct := c.(type) {
		case *TableColumn:
			columns = append(columns, ct)

			if ct.ImplicitPrimaryKey {
				implictlyDeclaredPK = append(implictlyDeclaredPK, ct)
			}
		case *TableKeys:
			keys = ct
		default:
			v.Err(ctx, "bad table definition")
			return nil
		}
	}

	if tableDef.K_WITH() != nil {
		// table properties
		for _, tp := range tableDef.AllTableProperty() {
			if tp.K_COMPACT() != nil {
				// compact storage
				cts.CompactStorage = true
			} else if tp.K_CLUSTERING() != nil {
				// table clustering
				for _, co := range tp.AllTableClusteringOrder() {
					cc := v.Visit(co).(*ColumnOrder)
					cts.ClusteringOrder = append(cts.ClusteringOrder, cc)
				}
			} else {
				// generic properties
				p := v.Visit(tp.Property()).(*Property)

				cts.Properties = append(cts.Properties, p)
			}
		}
	}

	// You can declare PRIMARY KEY inline or separately, but not both.
	if implictlyDeclaredPK != nil && keys != nil && len(keys.Partition) != 0 {
		v.Err(ctx, "multiple PRIMARY KEYs specified")
		return nil
	}

	if len(implictlyDeclaredPK) > 1 {
		// Only one column can be declared PRIMARY KEY implicitly.
		v.Err(ctx, "multiple PRIMARY KEYs specified")
		return nil
	}

	if len(implictlyDeclaredPK) > 0 {
		keys.Partition = []Identifier{
			implictlyDeclaredPK[0].Name,
		}
	}

	cts.IfNotExists = ctx.IfNotExists() != nil
	cts.ColumnFamily = v.Visit(ctx.ColumnFamilyName()).(*ObjectRef)
	cts.Columns = columns
	cts.PartitionKeys = keys.Partition
	cts.ClusteringKeys = keys.Clustering

	return cts
}

func (v *Visitor) VisitCreateTypeStatement(ctx *parser.CreateTypeStatementContext) any {
	cts := &CreateTypeStatement{}

	tn := v.Visit(ctx.GetTn()).(*ObjectRef)

	for _, tc := range ctx.AllTypeColumns() {
		f := v.Visit(tc).(*TypeField)
		cts.Fields = append(cts.Fields, f)

		if ut, ok := f.Type.(*ObjectRef); ok {
			if tn.Equals(ut) {
				v.Err(ctx, "circular type reference")

				return nil
			}
		}
	}

	cts.IfNotExists = ctx.IfNotExists() != nil
	cts.Name = tn

	return cts
}

func (v *Visitor) VisitUseStatement(ctx *parser.UseStatementContext) any {
	ks := v.Visit(ctx.GetKs()).(Identifier)

	return &UseStatement{
		Keyspace: ks,
	}
}

func (v *Visitor) VisitDropTableStatement(ctx *parser.DropTableStatementContext) any {
	cf := v.Visit(ctx.ColumnFamilyName()).(*ObjectRef)

	return &DropTableStatement{
		ColumnFamily: cf,
		IfExists:     ctx.IfExists() != nil,
	}
}

func (v *Visitor) VisitTruncateStatement(ctx *parser.TruncateStatementContext) any {
	cf := v.Visit(ctx.ColumnFamilyName()).(*ObjectRef)

	return &TruncateStatement{
		ColumnFamily: cf,
	}
}

func (v *Visitor) VisitCreateKeyspaceStatement(ctx *parser.CreateKeyspaceStatementContext) any {
	ks := v.Visit(ctx.KeyspaceName()).(Identifier)
	props := v.Visit(ctx.Properties()).([]*Property)

	return &CreateKeyspaceStatement{
		IfNotExists: ctx.IfNotExists() != nil,
		Name:        ks,
		Properties:  props,
	}
}

func (v *Visitor) VisitAlterKeyspaceStatement(ctx *parser.AlterKeyspaceStatementContext) any {
	ks := v.Visit(ctx.KeyspaceName()).(Identifier)
	props := v.Visit(ctx.Properties()).([]*Property)

	return &AlterKeyspaceStatement{
		IfExists:   ctx.IfExists() != nil,
		Keyspace:   ks,
		Properties: props,
	}
}

func (v *Visitor) VisitDropTypeStatement(ctx *parser.DropTypeStatementContext) any {
	name := v.Visit(ctx.UserTypeName()).(*ObjectRef)

	return &DropTypeStatement{
		Name:     name,
		IfExists: ctx.IfExists() != nil,
	}
}

func (v *Visitor) VisitListUsersStatement(ctx *parser.ListUsersStatementContext) any {
	return &ListUsersStatement{}
}

func (v *Visitor) VisitListSuperUsersStatement(ctx *parser.ListSuperUsersStatementContext) any {
	return &ListSuperUsersStatement{}
}

func (v *Visitor) VisitDropKeyspaceStatement(ctx *parser.DropKeyspaceStatementContext) any {
	ks := v.Visit(ctx.KeyspaceName()).(Identifier)

	return &DropKeyspaceStatement{
		Name:     ks,
		IfExists: ctx.IfExists() != nil,
	}
}

func (v *Visitor) VisitDropIdentityStatement(ctx *parser.DropIdentityStatementContext) any {
	return &DropIdentityStatement{
		Identity: v.Visit(ctx.Identity()).(Identifier),
		IfExists: ctx.IfExists() != nil,
	}
}

func (v *Visitor) VisitAddIdentityStatement(ctx *parser.AddIdentityStatementContext) any {
	return &AddIdentityStatement{
		IfNotExists: ctx.IfNotExists() != nil,
		Identity:    v.Visit(ctx.GetU()).(Identifier),
		Role:        v.Visit(ctx.GetR()).(Identifier),
	}
}

func (v *Visitor) VisitDropUserStatement(ctx *parser.DropUserStatementContext) any {
	return &DropUserStatement{
		IfExists: ctx.IfExists() != nil,
		Username: v.Visit(ctx.Username()).(Identifier),
	}
}

func (v *Visitor) VisitDropIndexStatement(ctx *parser.DropIndexStatementContext) any {
	return &DropIndexStatement{
		IfExists: ctx.IfExists() != nil,
		Name:     v.Visit(ctx.IndexName()).(*ObjectRef),
	}
}

func (v *Visitor) VisitDropMaterializedViewStatement(ctx *parser.DropMaterializedViewStatementContext) any {
	return &DropMaterializedViewStatement{
		IfExists: ctx.IfExists() != nil,
		Name:     v.Visit(ctx.ColumnFamilyName()).(*ObjectRef),
	}
}

func (v *Visitor) VisitGrantRoleStatement(ctx *parser.GrantRoleStatementContext) any {
	return &GrantRoleStatement{
		Role:    v.Visit(ctx.GetRole()).(Identifier),
		Grantee: v.Visit(ctx.GetGrantee()).(Identifier),
	}
}

func (v *Visitor) VisitRevokeRoleStatement(ctx *parser.RevokeRoleStatementContext) any {
	return &RevokeRoleStatement{
		Role:    v.Visit(ctx.GetRole()).(Identifier),
		Revokee: v.Visit(ctx.GetRevokee()).(Identifier),
	}
}

func (v *Visitor) VisitListRolesStatement(ctx *parser.ListRolesStatementContext) any {
	var role Identifier
	if ctx.RoleName() != nil {
		role = v.Visit(ctx.RoleName()).(Identifier)
	}

	return &ListRolesStatement{
		Role:      role,
		Recursive: ctx.K_NORECURSIVE() == nil,
	}
}

func (v *Visitor) VisitDropRoleStatement(ctx *parser.DropRoleStatementContext) any {
	return &DropRoleStatement{
		Role:     v.Visit(ctx.UserOrRoleName()).(Identifier),
		IfExists: ctx.IfExists() != nil,
	}
}

func (v *Visitor) VisitCreateTriggerStatement(ctx *parser.CreateTriggerStatementContext) any {
	return &CreateTriggerStatement{
		IfNotExists: ctx.IfNotExists() != nil,
		Name:        v.Visit(ctx.GetName()).(Identifier),
		On:          v.Visit(ctx.GetCf()).(*ObjectRef),
		Using:       unwrapStr(ctx.GetCls().GetText()),
	}
}

func (v *Visitor) VisitDropTriggerStatement(ctx *parser.DropTriggerStatementContext) any {
	return &DropTriggerStatement{
		IfExists: ctx.IfExists() != nil,
		Name:     v.Visit(ctx.GetName()).(Identifier),
		On:       v.Visit(ctx.GetCf()).(*ObjectRef),
	}
}

func (v *Visitor) VisitAlterMaterializedViewStatement(ctx *parser.AlterMaterializedViewStatementContext) any {
	return &AlterMaterializedViewStatement{
		IfExists:   ctx.IfExists() != nil,
		Name:       v.Visit(ctx.GetName()).(*ObjectRef),
		Properties: v.Visit(ctx.Properties()).([]*Property),
	}
}

func (v *Visitor) VisitCreateUserStatement(ctx *parser.CreateUserStatementContext) any {
	var password Password
	if ctx.K_WITH() != nil {
		password = v.Visit(ctx.UserPassword())
	}

	return &CreateUserStatement{
		IfNotExists: ctx.IfNotExists() != nil,
		Username:    v.Visit(ctx.Username()).(Identifier),
		Password:    password,
		Superuser:   ctx.K_SUPERUSER() != nil,
	}
}

func (v *Visitor) VisitAlterUserStatement(ctx *parser.AlterUserStatementContext) any {
	var password Password
	if ctx.K_WITH() != nil {
		password = v.Visit(ctx.UserPassword())
	}

	return &AlterUserStatement{
		IfExists:  ctx.IfExists() != nil,
		Username:  v.Visit(ctx.Username()).(Identifier),
		Password:  password,
		Superuser: ctx.K_SUPERUSER() != nil,
	}
}
