package ast

import "github.com/shachibista/gocqlparse/internal/parser"

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
				p := v.Visit(tp.Property()).(*TableProperty)

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
