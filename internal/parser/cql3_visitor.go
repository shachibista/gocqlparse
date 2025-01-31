// Code generated from Cql3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql3

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Cql3Parser.
type Cql3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Cql3Parser#cqlStatements.
	VisitCqlStatements(ctx *CqlStatementsContext) interface{}

	// Visit a parse tree produced by Cql3Parser#cqlStatement.
	VisitCqlStatement(ctx *CqlStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#marker.
	VisitMarker(ctx *MarkerContext) interface{}

	// Visit a parse tree produced by Cql3Parser#createKeyspaceStatement.
	VisitCreateKeyspaceStatement(ctx *CreateKeyspaceStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by Cql3Parser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by Cql3Parser#tableDefinition.
	VisitTableDefinition(ctx *TableDefinitionContext) interface{}

	// Visit a parse tree produced by Cql3Parser#TableColumnDeclaration.
	VisitTableColumnDeclaration(ctx *TableColumnDeclarationContext) interface{}

	// Visit a parse tree produced by Cql3Parser#TableKeyDeclaration.
	VisitTableKeyDeclaration(ctx *TableKeyDeclarationContext) interface{}

	// Visit a parse tree produced by Cql3Parser#columnMask.
	VisitColumnMask(ctx *ColumnMaskContext) interface{}

	// Visit a parse tree produced by Cql3Parser#columnMaskArguments.
	VisitColumnMaskArguments(ctx *ColumnMaskArgumentsContext) interface{}

	// Visit a parse tree produced by Cql3Parser#tablePartitionKey.
	VisitTablePartitionKey(ctx *TablePartitionKeyContext) interface{}

	// Visit a parse tree produced by Cql3Parser#tableProperty.
	VisitTableProperty(ctx *TablePropertyContext) interface{}

	// Visit a parse tree produced by Cql3Parser#tableClusteringOrder.
	VisitTableClusteringOrder(ctx *TableClusteringOrderContext) interface{}

	// Visit a parse tree produced by Cql3Parser#createTypeStatement.
	VisitCreateTypeStatement(ctx *CreateTypeStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#typeColumns.
	VisitTypeColumns(ctx *TypeColumnsContext) interface{}

	// Visit a parse tree produced by Cql3Parser#dropKeyspaceStatement.
	VisitDropKeyspaceStatement(ctx *DropKeyspaceStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#dropTableStatement.
	VisitDropTableStatement(ctx *DropTableStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#dropTypeStatement.
	VisitDropTypeStatement(ctx *DropTypeStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#truncateStatement.
	VisitTruncateStatement(ctx *TruncateStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#listUsersStatement.
	VisitListUsersStatement(ctx *ListUsersStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#listSuperUsersStatement.
	VisitListSuperUsersStatement(ctx *ListSuperUsersStatementContext) interface{}

	// Visit a parse tree produced by Cql3Parser#IdentIdent.
	VisitIdentIdent(ctx *IdentIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#IdentQuotedIdent.
	VisitIdentQuotedIdent(ctx *IdentQuotedIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#FIdentIdent.
	VisitFIdentIdent(ctx *FIdentIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#FIdentQuotedIdent.
	VisitFIdentQuotedIdent(ctx *FIdentQuotedIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#NonColIdent.
	VisitNonColIdent(ctx *NonColIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#NonColQuoted.
	VisitNonColQuoted(ctx *NonColQuotedContext) interface{}

	// Visit a parse tree produced by Cql3Parser#keyspaceName.
	VisitKeyspaceName(ctx *KeyspaceNameContext) interface{}

	// Visit a parse tree produced by Cql3Parser#columnFamilyName.
	VisitColumnFamilyName(ctx *ColumnFamilyNameContext) interface{}

	// Visit a parse tree produced by Cql3Parser#userTypeName.
	VisitUserTypeName(ctx *UserTypeNameContext) interface{}

	// Visit a parse tree produced by Cql3Parser#KsNameIdent.
	VisitKsNameIdent(ctx *KsNameIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#KsNameQuotedIdent.
	VisitKsNameQuotedIdent(ctx *KsNameQuotedIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#KsNameInvalidBind.
	VisitKsNameInvalidBind(ctx *KsNameInvalidBindContext) interface{}

	// Visit a parse tree produced by Cql3Parser#CfNameIdent.
	VisitCfNameIdent(ctx *CfNameIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#CfNameQuotedIdent.
	VisitCfNameQuotedIdent(ctx *CfNameQuotedIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#CfNameInvalidBind.
	VisitCfNameInvalidBind(ctx *CfNameInvalidBindContext) interface{}

	// Visit a parse tree produced by Cql3Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by Cql3Parser#fullMapLiteral.
	VisitFullMapLiteral(ctx *FullMapLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#setOrMapLiteral.
	VisitSetOrMapLiteral(ctx *SetOrMapLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#setLiteral.
	VisitSetLiteral(ctx *SetLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#collectionLiteral.
	VisitCollectionLiteral(ctx *CollectionLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#usertypeLiteral.
	VisitUsertypeLiteral(ctx *UsertypeLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#tupleLiteral.
	VisitTupleLiteral(ctx *TupleLiteralContext) interface{}

	// Visit a parse tree produced by Cql3Parser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by Cql3Parser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by Cql3Parser#AllowedFunctionNameIdent.
	VisitAllowedFunctionNameIdent(ctx *AllowedFunctionNameIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#AllowedFunctionNameQuotedIdent.
	VisitAllowedFunctionNameQuotedIdent(ctx *AllowedFunctionNameQuotedIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by Cql3Parser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by Cql3Parser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by Cql3Parser#termAddition.
	VisitTermAddition(ctx *TermAdditionContext) interface{}

	// Visit a parse tree produced by Cql3Parser#termMultiplication.
	VisitTermMultiplication(ctx *TermMultiplicationContext) interface{}

	// Visit a parse tree produced by Cql3Parser#termGroup.
	VisitTermGroup(ctx *TermGroupContext) interface{}

	// Visit a parse tree produced by Cql3Parser#simpleTerm.
	VisitSimpleTerm(ctx *SimpleTermContext) interface{}

	// Visit a parse tree produced by Cql3Parser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by Cql3Parser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by Cql3Parser#propertyValue.
	VisitPropertyValue(ctx *PropertyValueContext) interface{}

	// Visit a parse tree produced by Cql3Parser#comparatorType.
	VisitComparatorType(ctx *ComparatorTypeContext) interface{}

	// Visit a parse tree produced by Cql3Parser#native_type.
	VisitNative_type(ctx *Native_typeContext) interface{}

	// Visit a parse tree produced by Cql3Parser#collection_type.
	VisitCollection_type(ctx *Collection_typeContext) interface{}

	// Visit a parse tree produced by Cql3Parser#tuple_type.
	VisitTuple_type(ctx *Tuple_typeContext) interface{}

	// Visit a parse tree produced by Cql3Parser#vector_type.
	VisitVector_type(ctx *Vector_typeContext) interface{}

	// Visit a parse tree produced by Cql3Parser#NonTypeIdent.
	VisitNonTypeIdent(ctx *NonTypeIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#NonTypeQuotedIdent.
	VisitNonTypeQuotedIdent(ctx *NonTypeQuotedIdentContext) interface{}

	// Visit a parse tree produced by Cql3Parser#unreserved_keyword.
	VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{}

	// Visit a parse tree produced by Cql3Parser#unreserved_function_keyword.
	VisitUnreserved_function_keyword(ctx *Unreserved_function_keywordContext) interface{}

	// Visit a parse tree produced by Cql3Parser#basic_unreserved_keyword.
	VisitBasic_unreserved_keyword(ctx *Basic_unreserved_keywordContext) interface{}
}
