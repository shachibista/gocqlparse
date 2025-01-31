// Code generated from Cql3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql3

import "github.com/antlr4-go/antlr/v4"

type BaseCql3Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCql3Visitor) VisitCqlStatements(ctx *CqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCqlStatement(ctx *CqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUseStatement(ctx *UseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitMarker(ctx *MarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCreateKeyspaceStatement(ctx *CreateKeyspaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTableDefinition(ctx *TableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTableColumnDeclaration(ctx *TableColumnDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTableKeyDeclaration(ctx *TableKeyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitColumnMask(ctx *ColumnMaskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitColumnMaskArguments(ctx *ColumnMaskArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTablePartitionKey(ctx *TablePartitionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTableProperty(ctx *TablePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTableClusteringOrder(ctx *TableClusteringOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCreateTypeStatement(ctx *CreateTypeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTypeColumns(ctx *TypeColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropKeyspaceStatement(ctx *DropKeyspaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropTypeStatement(ctx *DropTypeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTruncateStatement(ctx *TruncateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitGrantRoleStatement(ctx *GrantRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropUserStatement(ctx *DropUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitAddIdentityStatement(ctx *AddIdentityStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitDropIdentityStatement(ctx *DropIdentityStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitListUsersStatement(ctx *ListUsersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitListSuperUsersStatement(ctx *ListSuperUsersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIdentIdent(ctx *IdentIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIdentQuotedIdent(ctx *IdentQuotedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitFIdentIdent(ctx *FIdentIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitFIdentQuotedIdent(ctx *FIdentQuotedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitNonColIdent(ctx *NonColIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitNonColQuoted(ctx *NonColQuotedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitKeyspaceName(ctx *KeyspaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitColumnFamilyName(ctx *ColumnFamilyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUserTypeName(ctx *UserTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUserOrRoleName(ctx *UserOrRoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitKsNameIdent(ctx *KsNameIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitKsNameQuotedIdent(ctx *KsNameQuotedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitKsNameInvalidBind(ctx *KsNameInvalidBindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCfNameIdent(ctx *CfNameIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCfNameQuotedIdent(ctx *CfNameQuotedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCfNameInvalidBind(ctx *CfNameInvalidBindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIdxName(ctx *IdxNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitRoleName(ctx *RoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitFullMapLiteral(ctx *FullMapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitSetOrMapLiteral(ctx *SetOrMapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitSetLiteral(ctx *SetLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCollectionLiteral(ctx *CollectionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitListLiteral(ctx *ListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUsertypeLiteral(ctx *UsertypeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTupleLiteral(ctx *TupleLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitAllowedFunctionNameIdent(ctx *AllowedFunctionNameIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitAllowedFunctionNameQuotedIdent(ctx *AllowedFunctionNameQuotedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTermAddition(ctx *TermAdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTermMultiplication(ctx *TermMultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTermGroup(ctx *TermGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitSimpleTerm(ctx *SimpleTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitPropertyValue(ctx *PropertyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitComparatorType(ctx *ComparatorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitNative_type(ctx *Native_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitCollection_type(ctx *Collection_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitTuple_type(ctx *Tuple_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitVector_type(ctx *Vector_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUsername(ctx *UsernameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitIdentity(ctx *IdentityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitNonTypeIdent(ctx *NonTypeIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitNonTypeQuotedIdent(ctx *NonTypeQuotedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitUnreserved_function_keyword(ctx *Unreserved_function_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCql3Visitor) VisitBasic_unreserved_keyword(ctx *Basic_unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}
