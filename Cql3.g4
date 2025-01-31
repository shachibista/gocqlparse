/*
 * SPDX-License-Identifier: Apache-2.0
 */

grammar Cql3;

options {
    language = Go;
}

/** STATEMENTS **/

cqlStatements
    : st=cqlStatement
      (';' st=cqlStatement)*
      (';')* EOF
    ;

cqlStatement
//     @after{ if (stmt != null) stmt.setBindVariables(bindVariables); }
//     : st1= selectStatement                 { $stmt = st1; }
//     | st2= insertStatement                 { $stmt = st2; }
//     | st3= updateStatement                 { $stmt = st3; }
//     | st4= batchStatement                  { $stmt = st4; }
//     | st5= deleteStatement                 { $stmt = st5; }
       : st6= useStatement
       | st7= truncateStatement
       | st8= createKeyspaceStatement
       | st9= createTableStatement
//     | st10=createIndexStatement            { $stmt = st10; }
       | st11=dropKeyspaceStatement
       | st12=dropTableStatement
//     | st13=dropIndexStatement              { $stmt = st13; }
//     | st14=alterTableStatement             { $stmt = st14; }
//     | st15=alterKeyspaceStatement          { $stmt = st15; }
//     | st16=grantPermissionsStatement       { $stmt = st16; }
//     | st17=revokePermissionsStatement      { $stmt = st17; }
//     | st18=listPermissionsStatement        { $stmt = st18; }
//     | st19=createUserStatement             { $stmt = st19; }
//     | st20=alterUserStatement              { $stmt = st20; }
//     | st21=dropUserStatement               { $stmt = st21; }
       | st22=listUsersStatement
//     | st23=createTriggerStatement          { $stmt = st23; }
//     | st24=dropTriggerStatement            { $stmt = st24; }
       | st25=createTypeStatement
//     | st26=alterTypeStatement              { $stmt = st26; }
       | st27=dropTypeStatement
//     | st28=createFunctionStatement         { $stmt = st28; }
//     | st29=dropFunctionStatement           { $stmt = st29; }
//     | st30=createAggregateStatement        { $stmt = st30; }
//     | st31=dropAggregateStatement          { $stmt = st31; }
//     | st32=createRoleStatement             { $stmt = st32; }
//     | st33=alterRoleStatement              { $stmt = st33; }
//     | st34=dropRoleStatement               { $stmt = st34; }
//     | st35=listRolesStatement              { $stmt = st35; }
//     | st36=grantRoleStatement              { $stmt = st36; }
//     | st37=revokeRoleStatement             { $stmt = st37; }
//     | st38=createMaterializedViewStatement { $stmt = st38; }
//     | st39=dropMaterializedViewStatement   { $stmt = st39; }
//     | st40=alterMaterializedViewStatement  { $stmt = st40; }
//     | st41=describeStatement               { $stmt = st41; }
//     | st42=addIdentityStatement            { $stmt = st42; }
//     | st43=dropIdentityStatement           { $stmt = st43; }
//     | st44=listSuperUsersStatement         { $stmt = st44; }
    ;

// /*
//  * USE <KEYSPACE>;
//  */
useStatement
    : K_USE ks=keyspaceName
    ;

// /**
//  * SELECT <expression>
//  * FROM <CF>
//  * WHERE KEY = "key1" AND COL > 1 AND COL < 100
//  * LIMIT <NUMBER>;
//  */
// selectStatement
//     @init {
//         Term.Raw limit = null;
//         Term.Raw perPartitionLimit = null;
//         List<Ordering.Raw> orderings = new ArrayList<>();
//         List<Selectable.Raw> groups = new ArrayList<>();
//         boolean allowFiltering = false;
//         boolean isJson = false;
//     }
//     : K_SELECT
//         // json is a valid column name. By consequence, we need to resolve the ambiguity for "json - json"
//       ( (K_JSON selectClause)=> K_JSON { isJson = true; } )? sclause=selectClause
//       K_FROM cf=columnFamilyName
//       ( K_WHERE wclause=whereClause )?
//       ( K_GROUP K_BY groupByClause[groups] ( ',' groupByClause[groups] )* )?
//       ( K_ORDER K_BY orderByClause[orderings] ( ',' orderByClause[orderings] )* )?
//       ( K_PER K_PARTITION K_LIMIT rows=intValue { perPartitionLimit = rows; } )?
//       ( K_LIMIT rows=intValue { limit = rows; } )?
//       ( K_ALLOW K_FILTERING  { allowFiltering = true; } )?
//       {
//           SelectStatement.Parameters params = new SelectStatement.Parameters(orderings,
//                                                                              groups,
//                                                                              $sclause.isDistinct,
//                                                                              allowFiltering,
//                                                                              isJson);
//           WhereClause where = wclause == null ? WhereClause.empty() : wclause.build();
//           $expr = new SelectStatement.RawStatement(cf, params, $sclause.selectors, where, limit, perPartitionLimit);
//       }
//     ;

// selectClause
//     @init{ $isDistinct = false; }
//     // distinct is a valid column name. By consequence, we need to resolve the ambiguity for "distinct - distinct"
//     : (K_DISTINCT selectors)=> K_DISTINCT s=selectors { $isDistinct = true; $selectors = s; }
//     | s=selectors { $selectors = s; }
//     ;

// selectors
//     : t1=selector { $expr = new ArrayList<RawSelector>(); $expr.add(t1); } (',' tN=selector { $expr.add(tN); })*
//     | '\*' { $expr = Collections.<RawSelector>emptyList();}
//     ;

// selector
//     @init{ ColumnIdentifier alias = null; }
//     : us=unaliasedSelector (K_AS c=noncol_ident { alias = c; })? { $s = new RawSelector(us, alias); }
//     ;

// unaliasedSelector
//     : a=selectionAddition {$s = a;}
//     ;

// selectionAddition
//     :   l=selectionMultiplication   {$s = l;}
//         ( '+' r=selectionMultiplication {$s = Selectable.WithFunction.Raw.newOperation('+', $s, r);}
//         | '-' r=selectionMultiplication {$s = Selectable.WithFunction.Raw.newOperation('-', $s, r);}
//         )*
//     ;

// selectionMultiplication
//     :   l=selectionGroup   {$s = l;}
//         ( '\*' r=selectionGroup {$s = Selectable.WithFunction.Raw.newOperation('*', $s, r);}
//         | '/' r=selectionGroup {$s = Selectable.WithFunction.Raw.newOperation('/', $s, r);}
//         | '%' r=selectionGroup {$s = Selectable.WithFunction.Raw.newOperation('\%', $s, r);}
//         )*
//     ;

// selectionGroup
//     : (selectionGroupWithField)=>  f=selectionGroupWithField { $s=f; }
//     | g=selectionGroupWithoutField { $s=g; }
//     | '-' g=selectionGroup {$s = Selectable.WithFunction.Raw.newNegation(g);}
//     ;

// selectionGroupWithField
//     : g=selectionGroupWithoutField m=selectorModifier[g] {$s = m;}
//     ;

// selectorModifier[Selectable.Raw receiver]
//     : f=fieldSelectorModifier[receiver] m=selectorModifier[f] { $s = m; }
//     | '[' ss=collectionSubSelection[receiver] ']' m=selectorModifier[ss] { $s = m; }
//     | { $s = receiver; }
//     ;

// fieldSelectorModifier[Selectable.Raw receiver]
//     : '.' fi=fident { $s = new Selectable.WithFieldSelection.Raw(receiver, fi); }
//     ;

// collectionSubSelection [Selectable.Raw receiver]
//     @init { boolean isSlice=false; }
//     : ( t1=term ( { isSlice=true; } RANGE (t2=term)? )?
//       | RANGE { isSlice=true; } t2=term
//       ) {
//           $s = isSlice
//              ? new Selectable.WithSliceSelection.Raw(receiver, t1, t2)
//              : new Selectable.WithElementSelection.Raw(receiver, t1);
//       }
//      ;

// selectionGroupWithoutField
//     @init { Selectable.Raw tmp = null; }
//     @after { $s = tmp; }
//     : sn=simpleUnaliasedSelector  { tmp=sn; }
//     | (selectionTypeHint)=> h=selectionTypeHint { tmp=h; }
//     | t=selectionTupleOrNestedSelector { tmp=t; }
//     | l=selectionList { tmp=l; }
//     | m=selectionMapOrSet { tmp=m; }
//     // UDTs are equivalent to maps from the syntax point of view, so the final decision will be done in Selectable.WithMapOrUdt
//     ;

// selectionTypeHint
//     : '(' ct=comparatorType ')' a=selectionGroupWithoutField { $s = new Selectable.WithTypeHint.Raw(ct, a); }
//     ;

// selectionList
//     @init { List<Selectable.Raw> l = new ArrayList<>(); }
//     @after { $s = new Selectable.WithArrayLiteral.Raw(l); }
//     : '[' ( t1=unaliasedSelector { l.add(t1); } ( ',' tn=unaliasedSelector { l.add(tn); } )* )? ']'
//     ;

// selectionMapOrSet
//     : '{' t1=unaliasedSelector ( m=selectionMap[t1] { $s = m; } | st=selectionSet[t1] { $s = st; }) '}'
//     | '{' '}' { $s = new Selectable.WithSet.Raw(Collections.emptyList());}
//     ;

// selectionMap [Selectable.Raw k1]
//     @init { List<Pair<Selectable.Raw, Selectable.Raw>> m = new ArrayList<>(); }
//     @after { $s = new Selectable.WithMapOrUdt.Raw(m); }
//       : ':' v1=unaliasedSelector   { m.add(Pair.create(k1, v1)); } ( ',' kn=unaliasedSelector ':' vn=unaliasedSelector { m.add(Pair.create(kn, vn)); } )*
//       ;

// selectionSet [Selectable.Raw t1]
//     @init { List<Selectable.Raw> l = new ArrayList<>(); l.add(t1); }
//     @after { $s = new Selectable.WithSet.Raw(l); }
//       : ( ',' tn=unaliasedSelector { l.add(tn); } )*
//       ;

// selectionTupleOrNestedSelector
//     @init { List<Selectable.Raw> l = new ArrayList<>(); }
//     @after { $s = new Selectable.BetweenParenthesesOrWithTuple.Raw(l); }
//     : '(' t1=unaliasedSelector { l.add(t1); } (',' tn=unaliasedSelector { l.add(tn); } )* ')'
//     ;

// /*
//  * A single selection. The core of it is selecting a column, but we also allow any term and function, as well as
//  * sub-element selection for UDT.
//  */
// simpleUnaliasedSelector
//     : c=sident                                   { $s = c; }
//     | l=selectionLiteral                         { $s = new Selectable.WithTerm.Raw(l); }
//     | f=selectionFunction                        { $s = f; }
//     ;

// selectionFunction
//     : K_COUNT        '(' '\*' ')'                                    { $s = Selectable.WithFunction.Raw.newCountRowsFunction(); }
//     | K_MAXWRITETIME '(' c=sident m=selectorModifier[c] ')'          { $s = new Selectable.WritetimeOrTTL.Raw(c, m, Selectable.WritetimeOrTTL.Kind.MAX_WRITE_TIME); }
//     | K_WRITETIME    '(' c=sident m=selectorModifier[c] ')'          { $s = new Selectable.WritetimeOrTTL.Raw(c, m, Selectable.WritetimeOrTTL.Kind.WRITE_TIME); }
//     | K_TTL          '(' c=sident m=selectorModifier[c] ')'          { $s = new Selectable.WritetimeOrTTL.Raw(c, m, Selectable.WritetimeOrTTL.Kind.TTL); }
//     | K_CAST         '(' sn=unaliasedSelector K_AS t=native_type ')' { $s = new Selectable.WithCast.Raw(sn, t);}
//     | f=functionName args=selectionFunctionArgs                      { $s = new Selectable.WithFunction.Raw(f, args); }
//     ;

// selectionLiteral
//     : c=constant  { $value = c; }
//     | K_NULL      { $value = Constants.NULL_LITERAL; }
//     | m=marker    { $value = m; }
//     ;

marker
    : ':' id=noncol_ident
    | QMARK
    ;

// selectionFunctionArgs
//     @init{ $a = new ArrayList<>(); }
//     : '(' (s1=unaliasedSelector { $a.add(s1); }
//           ( ',' sn=unaliasedSelector { $a.add(sn); } )*)?
//       ')'
//     ;

// sident
//     : t=IDENT              { $id = Selectable.RawIdentifier.forUnquoted($t.text); }
//     | t=QUOTED_NAME        { $id = Selectable.RawIdentifier.forQuoted($t.text); }
//     | k=unreserved_keyword { $id = Selectable.RawIdentifier.forUnquoted(k); }
//     ;

// whereClause
//     @init{ $clause = new WhereClause.Builder(); }
//     : relationOrExpression[$clause] (K_AND relationOrExpression[$clause])*
//     ;

// relationOrExpression [WhereClause.Builder clause]
//     : relation[$clause]
//     | customIndexExpression[$clause]
//     ;

// customIndexExpression [WhereClause.Builder clause]
//     @init{QualifiedName name = new QualifiedName();}
//     : 'expr(' idxName[name] ',' t=term ')' { clause.add(new CustomIndexExpression(name, t));}
//     ;

// orderByClause[List<Ordering.Raw> orderings]
//     @init{
//         Ordering.Direction direction = Ordering.Direction.ASC;
//     }
//     : c=cident (K_ANN K_OF t=term)? (K_ASC | K_DESC { direction = Ordering.Direction.DESC; })?
//     {
//         Ordering.Raw.Expression expr = (t == null)
//             ? new Ordering.Raw.SingleColumn(c)
//             : new Ordering.Raw.Ann(c, t);
//         orderings.add(new Ordering.Raw(expr, direction));
//     }
//     ;

// groupByClause[List<Selectable.Raw> groups]
//     : s=unaliasedSelector { groups.add(s); }
//     ;

// /**
//  * INSERT INTO <CF> (<column>, <column>, <column>, ...)
//  * VALUES (<value>, <value>, <value>, ...)
//  * USING TIMESTAMP <long>;
//  *
//  */
// insertStatement
//     : K_INSERT K_INTO cf=columnFamilyName
//         ( st1=normalInsertStatement[cf] { $expr = st1; }
//         | K_JSON st2=jsonInsertStatement[cf] { $expr = st2; })
//     ;

// normalInsertStatement [QualifiedName qn]
//     @init {
//         Attributes.Raw attrs = new Attributes.Raw();
//         List<ColumnIdentifier> columnNames  = new ArrayList<>();
//         List<Term.Raw> values = new ArrayList<>();
//         boolean ifNotExists = false;
//     }
//     : '(' c1=cident { columnNames.add(c1); }  ( ',' cn=cident { columnNames.add(cn); } )* ')'
//       K_VALUES
//       '(' v1=term { values.add(v1); } ( ',' vn=term { values.add(vn); } )* ')'
//       ( K_IF K_NOT K_EXISTS { ifNotExists = true; } )?
//       ( usingClause[attrs] )?
//       {
//           $expr = new UpdateStatement.ParsedInsert(qn, attrs, columnNames, values, ifNotExists);
//       }
//     ;

// jsonInsertStatement [QualifiedName qn]
//     @init {
//         Attributes.Raw attrs = new Attributes.Raw();
//         boolean ifNotExists = false;
//         boolean defaultUnset = false;
//     }
//     : val=jsonValue
//       ( K_DEFAULT ( K_NULL | ( { defaultUnset = true; } K_UNSET) ) )?
//       ( K_IF K_NOT K_EXISTS { ifNotExists = true; } )?
//       ( usingClause[attrs] )?
//       {
//           $expr = new UpdateStatement.ParsedInsertJson(qn, attrs, val, defaultUnset, ifNotExists);
//       }
//     ;

// jsonValue
//     : s=STRING_LITERAL { $value = new Json.Literal($s.text); }
//     | ':' id=noncol_ident     { $value = newJsonBindVariables(id); }
//     | QMARK            { $value = newJsonBindVariables(null); }
//     ;

// usingClause[Attributes.Raw attrs]
//     : K_USING usingClauseObjective[attrs] ( K_AND usingClauseObjective[attrs] )*
//     ;

// usingClauseObjective[Attributes.Raw attrs]
//     : K_TIMESTAMP ts=intValue { attrs.timestamp = ts; }
//     | K_TTL t=intValue { attrs.timeToLive = t; }
//     ;

// /**
//  * UPDATE <CF>
//  * USING TIMESTAMP <long>
//  * SET name1 = value1, name2 = value2
//  * WHERE key = value;
//  * [IF (EXISTS | name = value, ...)];
//  */
// updateStatement
//     @init {
//         Attributes.Raw attrs = new Attributes.Raw();
//         List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations = new ArrayList<>();
//         boolean ifExists = false;
//     }
//     : K_UPDATE cf=columnFamilyName
//       ( usingClause[attrs] )?
//       K_SET columnOperation[operations] (',' columnOperation[operations])*
//       K_WHERE wclause=whereClause
//       ( K_IF ( K_EXISTS { ifExists = true; } | conditions=updateConditions ))?
//       {
//           $expr = new UpdateStatement.ParsedUpdate(cf,
//                                                    attrs,
//                                                    operations,
//                                                    wclause.build(),
//                                                    conditions == null ? Collections.<ColumnCondition.Raw>emptyList() : conditions,
//                                                    ifExists);
//      }
//     ;

// updateConditions
//     @init { conditions = new ArrayList<ColumnCondition.Raw>(); }
//     : c1=columnCondition { $conditions.add(c1);} ( K_AND cn=columnCondition { $conditions.add(cn); })*
//     ;

// /**
//  * DELETE name1, name2
//  * FROM <CF>
//  * USING TIMESTAMP <long>
//  * WHERE KEY = keyname
//    [IF (EXISTS | name = value, ...)];
//  */
// deleteStatement
//     @init {
//         Attributes.Raw attrs = new Attributes.Raw();
//         List<Operation.RawDeletion> columnDeletions = Collections.emptyList();
//         boolean ifExists = false;
//     }
//     : K_DELETE ( dels=deleteSelection { columnDeletions = dels; } )?
//       K_FROM cf=columnFamilyName
//       ( usingClauseDelete[attrs] )?
//       K_WHERE wclause=whereClause
//       ( K_IF ( K_EXISTS { ifExists = true; } | conditions=updateConditions ))?
//       {
//           $expr = new DeleteStatement.Parsed(cf,
//                                              attrs,
//                                              columnDeletions,
//                                              wclause.build(),
//                                              conditions == null ? Collections.<ColumnCondition.Raw>emptyList() : conditions,
//                                              ifExists);
//       }
//     ;

// deleteSelection
//     : { $operations = new ArrayList<Operation.RawDeletion>(); }
//           t1=deleteOp { $operations.add(t1); }
//           (',' tN=deleteOp { $operations.add(tN); })*
//     ;

// deleteOp
//     : c=cident                { $op = new Operation.ColumnDeletion(c); }
//     | c=cident '[' t=term ']' { $op = new Operation.ElementDeletion(c, t); }
//     | c=cident '.' field=fident { $op = new Operation.FieldDeletion(c, field); }
//     ;

// usingClauseDelete[Attributes.Raw attrs]
//     : K_USING K_TIMESTAMP ts=intValue { attrs.timestamp = ts; }
//     ;

// /**
//  * BEGIN BATCH
//  *   UPDATE <CF> SET name1 = value1 WHERE KEY = keyname1;
//  *   UPDATE <CF> SET name2 = value2 WHERE KEY = keyname2;
//  *   UPDATE <CF> SET name3 = value3 WHERE KEY = keyname3;
//  *   ...
//  * APPLY BATCH
//  *
//  * OR
//  *
//  * BEGIN BATCH
//  *   INSERT INTO <CF> (KEY, <name>) VALUES ('<key>', '<value>');
//  *   INSERT INTO <CF> (KEY, <name>) VALUES ('<key>', '<value>');
//  *   ...
//  * APPLY BATCH
//  *
//  * OR
//  *
//  * BEGIN BATCH
//  *   DELETE name1, name2 FROM <CF> WHERE key = <key>
//  *   DELETE name3, name4 FROM <CF> WHERE key = <key>
//  *   ...
//  * APPLY BATCH
//  */
// batchStatement
//     @init {
//         BatchStatement.Type type = BatchStatement.Type.LOGGED;
//         List<ModificationStatement.Parsed> statements = new ArrayList<ModificationStatement.Parsed>();
//         Attributes.Raw attrs = new Attributes.Raw();
//     }
//     : K_BEGIN
//       ( K_UNLOGGED { type = BatchStatement.Type.UNLOGGED; } | K_COUNTER { type = BatchStatement.Type.COUNTER; } )?
//       K_BATCH ( usingClause[attrs] )?
//           ( s=batchStatementObjective ';'? { statements.add(s); } )*
//       K_APPLY K_BATCH
//       {
//           $expr = new BatchStatement.Parsed(type, attrs, statements);
//       }
//     ;

// batchStatementObjective
//     : i=insertStatement  { $statement = i; }
//     | u=updateStatement  { $statement = u; }
//     | d=deleteStatement  { $statement = d; }
//     ;

// createAggregateStatement
//     @init {
//         boolean orReplace = false;
//         boolean ifNotExists = false;

//         List<CQL3Type.Raw> argTypes = new ArrayList<>();
//     }
//     : K_CREATE (K_OR K_REPLACE { orReplace = true; })?
//       K_AGGREGATE
//       (K_IF K_NOT K_EXISTS { ifNotExists = true; })?
//       fn=functionName
//       '('
//         (
//           v=comparatorType { argTypes.add(v); }
//           ( ',' v=comparatorType { argTypes.add(v); } )*
//         )?
//       ')'
//       K_SFUNC sfunc = allowedFunctionName
//       K_STYPE stype = comparatorType
//       (
//         K_FINALFUNC ffunc = allowedFunctionName
//       )?
//       (
//         K_INITCOND ival = term
//       )?
//       { $stmt = new CreateAggregateStatement.Raw(fn, argTypes, stype, sfunc, ffunc, ival, orReplace, ifNotExists); }
//     ;

// dropAggregateStatement
//     @init {
//         boolean ifExists = false;
//         List<CQL3Type.Raw> argTypes = new ArrayList<>();
//         boolean argsSpecified = false;
//     }
//     : K_DROP K_AGGREGATE
//       (K_IF K_EXISTS { ifExists = true; } )?
//       fn=functionName
//       (
//         '('
//           (
//             v=comparatorType { argTypes.add(v); }
//             ( ',' v=comparatorType { argTypes.add(v); } )*
//           )?
//         ')'
//         { argsSpecified = true; }
//       )?
//       { $stmt = new DropAggregateStatement.Raw(fn, argTypes, argsSpecified, ifExists); }
//     ;

// createFunctionStatement
//     @init {
//         boolean orReplace = false;
//         boolean ifNotExists = false;

//         List<ColumnIdentifier> argNames = new ArrayList<>();
//         List<CQL3Type.Raw> argTypes = new ArrayList<>();
//         boolean calledOnNullInput = false;
//     }
//     : K_CREATE (K_OR K_REPLACE { orReplace = true; })?
//       K_FUNCTION
//       (K_IF K_NOT K_EXISTS { ifNotExists = true; })?
//       fn=functionName
//       '('
//         (
//           k=noncol_ident v=comparatorType { argNames.add(k); argTypes.add(v); }
//           ( ',' k=noncol_ident v=comparatorType { argNames.add(k); argTypes.add(v); } )*
//         )?
//       ')'
//       ( (K_RETURNS K_NULL) | (K_CALLED { calledOnNullInput=true; })) K_ON K_NULL K_INPUT
//       K_RETURNS returnType = comparatorType
//       K_LANGUAGE language = IDENT
//       K_AS body = STRING_LITERAL
//       { $stmt = new CreateFunctionStatement.Raw(
//           fn, argNames, argTypes, returnType, calledOnNullInput, LocalizeString.toLowerCaseLocalized($language.text), $body.text, orReplace, ifNotExists);
//       }
//     ;

// dropFunctionStatement
//     @init {
//         boolean ifExists = false;
//         List<CQL3Type.Raw> argTypes = new ArrayList<>();
//         boolean argsSpecified = false;
//     }
//     : K_DROP K_FUNCTION
//       (K_IF K_EXISTS { ifExists = true; } )?
//       fn=functionName
//       (
//         '('
//           (
//             v=comparatorType { argTypes.add(v); }
//             ( ',' v=comparatorType { argTypes.add(v); } )*
//           )?
//         ')'
//         { argsSpecified = true; }
//       )?
//       { $stmt = new DropFunctionStatement.Raw(fn, argTypes, argsSpecified, ifExists); }
//     ;

/**
 * CREATE KEYSPACE [IF NOT EXISTS] <KEYSPACE> WITH attr1 = value1 AND attr2 = value2;
 */
createKeyspaceStatement
    : K_CREATE K_KEYSPACE ifNotExists? ks=keyspaceName
      K_WITH properties
    ;

// /**
//  * CREATE TABLE [IF NOT EXISTS] <CF> (
//  *     <name1> <type>,
//  *     <name2> <type>,
//  *     <name3> <type>
//  * ) WITH <property> = <value> AND ...;
//  */
createTableStatement
    : K_CREATE K_COLUMNFAMILY ifNotExists?
      cf=columnFamilyName
      tableDefinition
    ;

ifNotExists
    : K_IF K_NOT K_EXISTS
    ;

ifExists
    : K_IF K_EXISTS
    ;

tableDefinition
    : '(' tableColumns ( ',' tableColumns? )* ')'
      ( K_WITH tableProperty ( K_AND tableProperty )*)?
    ;

tableColumns
    : k=ident v=comparatorType (K_STATIC)? (mask=columnMask)? (K_PRIMARY K_KEY)? # TableColumnDeclaration
    | K_PRIMARY K_KEY '(' pk=tablePartitionKey (',' c=ident )* ')'               # TableKeyDeclaration
    ;

columnMask
    : K_MASKED K_WITH name=functionName columnMaskArguments
    | K_MASKED K_WITH K_DEFAULT
    ;

columnMaskArguments
    : '('  ')' | '(' c=term (',' cn=term)* ')'
    ;

tablePartitionKey
    : k1=ident
    | '(' k1=ident ( ',' kn=ident )* ')'
    ;

tableProperty
    : property
    | K_COMPACT K_STORAGE
    | K_CLUSTERING K_ORDER K_BY '(' tableClusteringOrder (',' tableClusteringOrder)* ')'
    ;

tableClusteringOrder
    : k=ident (K_ASC | K_DESC )
    ;

/**
 * CREATE TYPE foo (
 *    <name1> <type1>,
 *    <name2> <type2>,
 *    ....
 * )
 */
createTypeStatement
    : K_CREATE K_TYPE ifNotExists?
         tn=userTypeName
         '(' typeColumns ( ',' typeColumns? )* ')'
    ;

typeColumns
    : k=fident v=comparatorType
    ;

// /**
//  * CREATE INDEX [IF NOT EXISTS] [indexName] ON <columnFamily> (<columnName>);
//  * CREATE CUSTOM INDEX [IF NOT EXISTS] [indexName] ON <columnFamily> (<columnName>) USING <indexClass>;
//  */
// createIndexStatement
//     @init {
//         IndexAttributes props = new IndexAttributes();
//         boolean ifNotExists = false;
//         QualifiedName name = new QualifiedName();
//         List<IndexTarget.Raw> targets = new ArrayList<>();
//     }
//     : K_CREATE (K_CUSTOM { props.isCustom = true; })? K_INDEX (K_IF K_NOT K_EXISTS { ifNotExists = true; } )?
//         (idxName[name])? K_ON cf=columnFamilyName '(' (indexIdent[targets] (',' indexIdent[targets])*)? ')'
//         (K_USING cls=STRING_LITERAL { props.customClass = $cls.text; })?
//         (K_WITH properties[props])?
//       { $stmt = new CreateIndexStatement.Raw(cf, name, targets, props, ifNotExists); }
//     ;

// indexIdent [List<IndexTarget.Raw> targets]
//     : c=cident                   { $targets.add(IndexTarget.Raw.simpleIndexOn(c)); }
//     | K_VALUES '(' c=cident ')'  { $targets.add(IndexTarget.Raw.valuesOf(c)); }
//     | K_KEYS '(' c=cident ')'    { $targets.add(IndexTarget.Raw.keysOf(c)); }
//     | K_ENTRIES '(' c=cident ')' { $targets.add(IndexTarget.Raw.keysAndValuesOf(c)); }
//     | K_FULL '(' c=cident ')'    { $targets.add(IndexTarget.Raw.fullCollection(c)); }
//     ;

// /**
//  * CREATE MATERIALIZED VIEW <viewName> AS
//  *  SELECT <columns>
//  *  FROM <CF>
//  *  WHERE <pkColumns> IS NOT NULL
//  *  PRIMARY KEY (<pkColumns>)
//  *  WITH <property> = <value> AND ...;
//  */
// createMaterializedViewStatement
//     @init {
//         boolean ifNotExists = false;
//     }
//     : K_CREATE K_MATERIALIZED K_VIEW (K_IF K_NOT K_EXISTS { ifNotExists = true; })? cf=columnFamilyName K_AS
//         K_SELECT sclause=selectors K_FROM basecf=columnFamilyName
//         (K_WHERE wclause=whereClause)?
//         {
//              WhereClause where = wclause == null ? WhereClause.empty() : wclause.build();
//              $stmt = new CreateViewStatement.Raw(basecf, cf, sclause, where, ifNotExists);
//         }
//         viewPrimaryKey[stmt]
//         ( K_WITH viewProperty[stmt] ( K_AND viewProperty[stmt] )*)?
//     ;

// viewPrimaryKey[CreateViewStatement.Raw stmt]
//     : K_PRIMARY K_KEY '(' viewPartitionKey[stmt] (',' c=ident { $stmt.markClusteringColumn(c); } )* ')'
//     ;

// viewPartitionKey[CreateViewStatement.Raw stmt]
//     @init {List<ColumnIdentifier> l = new ArrayList<ColumnIdentifier>();}
//     @after{ $stmt.setPartitionKeyColumns(l); }
//     : k1=ident { l.add(k1);}
//     | '(' k1=ident { l.add(k1); } ( ',' kn=ident { l.add(kn); } )* ')'
//     ;

// viewProperty[CreateViewStatement.Raw stmt]
//     : property[stmt.attrs]
//     | K_COMPACT K_STORAGE { throw new SyntaxException("COMPACT STORAGE tables are not allowed starting with version 4.0"); }
//     | K_CLUSTERING K_ORDER K_BY '(' viewClusteringOrder[stmt] (',' viewClusteringOrder[stmt])* ')'
//     ;

// viewClusteringOrder[CreateViewStatement.Raw stmt]
//     @init{ boolean ascending = true; }
//     : k=ident (K_ASC | K_DESC { ascending = false; } ) { $stmt.extendClusteringOrder(k, ascending); }
//     ;

// /**
//  * CREATE TRIGGER triggerName ON columnFamily USING 'triggerClass';
//  */
// createTriggerStatement
//     @init {
//         boolean ifNotExists = false;
//     }
//     : K_CREATE K_TRIGGER (K_IF K_NOT K_EXISTS { ifNotExists = true; } )? (name=ident)
//         K_ON cf=columnFamilyName K_USING cls=STRING_LITERAL
//       { $stmt = new CreateTriggerStatement.Raw(cf, name.toString(), $cls.text, ifNotExists); }
//     ;

// /**
//  * DROP TRIGGER [IF EXISTS] triggerName ON columnFamily;
//  */
// dropTriggerStatement
//      @init { boolean ifExists = false; }
//     : K_DROP K_TRIGGER (K_IF K_EXISTS { ifExists = true; } )? (name=ident) K_ON cf=columnFamilyName
//       { $stmt = new DropTriggerStatement.Raw(cf, name.toString(), ifExists); }
//     ;

// /**
//  * ALTER KEYSPACE [IF EXISTS] <KS> WITH <property> = <value>;
//  */
// alterKeyspaceStatement
//     @init {
//      KeyspaceAttributes attrs = new KeyspaceAttributes();
//      boolean ifExists = false;
//     }
//     : K_ALTER K_KEYSPACE (K_IF K_EXISTS { ifExists = true; } )? ks=keyspaceName
//         K_WITH properties[attrs] { $stmt = new AlterKeyspaceStatement.Raw(ks, attrs, ifExists); }
//     ;

// /**
//  * ALTER TABLE <table> ALTER <column> TYPE <newtype>;
//  * ALTER TABLE [IF EXISTS] <table> ALTER [IF EXISTS] <column> MASKED WITH <maskFunction>);
//  * ALTER TABLE [IF EXISTS] <table> ALTER [IF EXISTS] <column> DROP MASKED;
//  * ALTER TABLE [IF EXISTS] <table> ADD [IF NOT EXISTS] <column> <newtype> <maskFunction>; | ALTER TABLE [IF EXISTS] <table> ADD [IF NOT EXISTS] (<column> <newtype> <maskFunction>, <column1> <newtype1>  <maskFunction1>..... <column n> <newtype n>  <maskFunction n>)
//  * ALTER TABLE [IF EXISTS] <table> DROP [IF EXISTS] <column>; | ALTER TABLE [IF EXISTS] <table> DROP [IF EXISTS] ( <column>,<column1>.....<column n>)
//  * ALTER TABLE [IF EXISTS] <table> RENAME [IF EXISTS] <column> TO <column>;
//  * ALTER TABLE [IF EXISTS] <table> WITH <property> = <value>;
//  */
// alterTableStatement
//     @init { boolean ifExists = false; }
//     : K_ALTER K_COLUMNFAMILY (K_IF K_EXISTS { ifExists = true; } )?
//       cf=columnFamilyName { $stmt = new AlterTableStatement.Raw(cf, ifExists); }
//       (
//         K_ALTER id=cident K_TYPE v=comparatorType { $stmt.alter(id, v); }

//       | K_ALTER ( K_IF K_EXISTS { $stmt.ifColumnExists(true); } )? id=cident
//               ( mask=columnMask { $stmt.mask(id, mask); }
//               | K_DROP K_MASKED { $stmt.mask(id, null); } )

//       | K_ADD ( K_IF K_NOT K_EXISTS { $stmt.ifColumnNotExists(true); } )?
//               (        id=ident  v=comparatorType  b=isStaticColumn (m=columnMask)? { $stmt.add(id,  v,  b, m);  }
//                | ('('  id1=ident v1=comparatorType b1=isStaticColumn (m1=columnMask)? { $stmt.add(id1, v1, b1, m1); }
//                  ( ',' idn=ident vn=comparatorType bn=isStaticColumn (mn=columnMask)? { $stmt.add(idn, vn, bn, mn); mn=null; } )* ')') )

//       | K_DROP ( K_IF K_EXISTS { $stmt.ifColumnExists(true); } )?
//                (       id=ident { $stmt.drop(id);  }
//                | ('('  id1=ident { $stmt.drop(id1); }
//                  ( ',' idn=ident { $stmt.drop(idn); } )* ')') )
//                ( K_USING K_TIMESTAMP t=INTEGER { $stmt.timestamp(Long.parseLong(Constants.Literal.integer($t.text).getText())); } )?

//       | K_RENAME ( K_IF K_EXISTS { $stmt.ifColumnExists(true); } )?
//                (        id1=ident K_TO toId1=ident { $stmt.rename(id1, toId1); }
//                 ( K_AND idn=ident K_TO toIdn=ident { $stmt.rename(idn, toIdn); } )* )

//       | K_DROP K_COMPACT K_STORAGE { $stmt.dropCompactStorage(); }

//       | K_WITH properties[$stmt.attrs] { $stmt.attrs(); }
//       )
//     ;

// isStaticColumn
//     @init { boolean isStatic = false; }
//     : (K_STATIC { isStatic=true; })? { $isStaticColumn = isStatic; }
//     ;

// alterMaterializedViewStatement
//     @init {
//         TableAttributes attrs = new TableAttributes();
//         boolean ifExists = false;
//     }
//     : K_ALTER K_MATERIALIZED K_VIEW (K_IF K_EXISTS { ifExists = true; } )? name=columnFamilyName
//           K_WITH properties[attrs]
//     {
//         $stmt = new AlterViewStatement.Raw(name, attrs, ifExists);
//     }
//     ;


// /**
//  * ALTER TYPE [IF EXISTS] <name> ALTER <field> TYPE <newtype>;
//  * ALTER TYPE [IF EXISTS] <name> ADD [IF NOT EXISTS]<field> <newtype>;
//  * ALTER TYPE [IF EXISTS] <name> RENAME [IF EXISTS] <field> TO <newtype> AND ...;
//  */
// alterTypeStatement
//     @init {
//         boolean ifExists = false;
//     }
//     : K_ALTER K_TYPE (K_IF K_EXISTS { ifExists = true; } )? name=userTypeName { $stmt = new AlterTypeStatement.Raw(name, ifExists); }
//       (
//         K_ALTER   f=fident K_TYPE v=comparatorType { $stmt.alter(f, v); }

//       | K_ADD (K_IF K_NOT K_EXISTS { $stmt.ifFieldNotExists(true); } )?     f=fident v=comparatorType        { $stmt.add(f, v); }

//       | K_RENAME (K_IF K_EXISTS { $stmt.ifFieldExists(true); } )? f1=fident K_TO toF1=fident        { $stmt.rename(f1, toF1); }
//          ( K_AND fn=fident K_TO toFn=fident        { $stmt.rename(fn, toFn); } )*
//       )
//     ;

/**
 * DROP KEYSPACE [IF EXISTS] <KSP>;
 */
dropKeyspaceStatement
    : K_DROP K_KEYSPACE ifExists? ks=keyspaceName
    ;

/**
 * DROP TABLE [IF EXISTS] <table>;
 */
dropTableStatement
    : K_DROP K_COLUMNFAMILY ifExists? name=columnFamilyName
    ;

/**
 * DROP TYPE <name>;
 */
dropTypeStatement
    : K_DROP K_TYPE ifExists? name=userTypeName
    ;

// /**
//  * DROP INDEX [IF EXISTS] <INDEX_NAME>
//  */
// dropIndexStatement
//     @init { boolean ifExists = false; }
//     : K_DROP K_INDEX (K_IF K_EXISTS { ifExists = true; } )? index=indexName
//       { $stmt = new DropIndexStatement.Raw(index, ifExists); }
//     ;

// /**
//  * DROP MATERIALIZED VIEW [IF EXISTS] <view_name>
//  */
// dropMaterializedViewStatement
//     @init { boolean ifExists = false; }
//     : K_DROP K_MATERIALIZED K_VIEW (K_IF K_EXISTS { ifExists = true; } )? cf=columnFamilyName
//       { $stmt = new DropViewStatement.Raw(cf, ifExists); }
//     ;

// /**
//   * TRUNCATE <CF>;
//   */
truncateStatement
    : K_TRUNCATE (K_COLUMNFAMILY)? cf=columnFamilyName
    ;

// /**
//  * GRANT <permission>[, <permission>]* | ALL ON <resource> TO <rolename>
//  */
// grantPermissionsStatement
//     : K_GRANT
//           permissionOrAll
//       K_ON
//           resource
//       K_TO
//           grantee=userOrRoleName
//       { $stmt = new GrantPermissionsStatement(filterPermissions($permissionOrAll.perms, $resource.res), $resource.res, grantee); }
//     ;

// /**
//  * REVOKE <permission>[, <permission>]* | ALL ON <resource> FROM <rolename>
//  */
// revokePermissionsStatement
//     : K_REVOKE
//           permissionOrAll
//       K_ON
//           resource
//       K_FROM
//           revokee=userOrRoleName
//       { $stmt = new RevokePermissionsStatement(filterPermissions($permissionOrAll.perms, $resource.res), $resource.res, revokee); }
//     ;

// /**
//  * GRANT ROLE <rolename> TO <grantee>
//  */
// grantRoleStatement
//     : K_GRANT
//           role=userOrRoleName
//       K_TO
//           grantee=userOrRoleName
//       { $stmt = new GrantRoleStatement(role, grantee); }
//     ;

// /**
//  * REVOKE ROLE <rolename> FROM <revokee>
//  */
// revokeRoleStatement
//     : K_REVOKE
//           role=userOrRoleName
//       K_FROM
//           revokee=userOrRoleName
//       { $stmt = new RevokeRoleStatement(role, revokee); }
//     ;

// listPermissionsStatement
//     @init {
//         IResource resource = null;
//         boolean recursive = true;
//         RoleName grantee = new RoleName();
//     }
//     : K_LIST
//           permissionOrAll
//       ( K_ON resource { resource = $resource.res; } )?
//       ( K_OF roleName[grantee] )?
//       ( K_NORECURSIVE { recursive = false; } )?
//       { $stmt = new ListPermissionsStatement($permissionOrAll.perms, resource, grantee, recursive); }
//     ;

// permission
//     : p=(K_CREATE | K_ALTER | K_DROP | K_SELECT | K_MODIFY | K_AUTHORIZE | K_DESCRIBE | K_EXECUTE | K_UNMASK | K_SELECT_MASKED)
//     { $perm = Permission.valueOf(LocalizeString.toUpperCaseLocalized($p.text)); }
//     ;

// permissionOrAll
//     : K_ALL ( K_PERMISSIONS )?       { $perms = Permission.ALL; }
//     | p=permission ( K_PERMISSION )? { $perms = EnumSet.of($p.perm); } ( ',' p=permission ( K_PERMISSION )? { $perms.add($p.perm); } )*
//     ;

// resource
//     : d=dataResource { $res = $d.res; }
//     | r=roleResource { $res = $r.res; }
//     | f=functionResource { $res = $f.res; }
//     | j=jmxResource { $res = $j.res; }
//     ;

// dataResource
//     : K_ALL K_KEYSPACES { $res = DataResource.root(); }
//     | K_KEYSPACE ks = keyspaceName { $res = DataResource.keyspace($ks.id); }
//     | ( K_COLUMNFAMILY )? cf = columnFamilyName { $res = DataResource.table($cf.name.getKeyspace(), $cf.name.getName()); }
//     | K_ALL K_TABLES K_IN K_KEYSPACE ks = keyspaceName { $res = DataResource.allTables($ks.id); }
//     ;

// jmxResource
//     : K_ALL K_MBEANS { $res = JMXResource.root(); }
//     // when a bean name (or pattern) is supplied, validate that it's a legal ObjectName
//     // also, just to be picky, if the "MBEANS" form is used, only allow a pattern style names
//     | K_MBEAN mbean { $res = JMXResource.mbean(canonicalizeObjectName($mbean.text, false)); }
//     | K_MBEANS mbean { $res = JMXResource.mbean(canonicalizeObjectName($mbean.text, true)); }
//     ;

// roleResource
//     : K_ALL K_ROLES { $res = RoleResource.root(); }
//     | K_ROLE role = userOrRoleName { $res = RoleResource.role($role.name.getName()); }
//     ;

// functionResource
//     @init {
//         List<CQL3Type.Raw> argsTypes = new ArrayList<>();
//     }
//     : K_ALL K_FUNCTIONS { $res = FunctionResource.root(); }
//     | K_ALL K_FUNCTIONS K_IN K_KEYSPACE ks = keyspaceName { $res = FunctionResource.keyspace($ks.id); }
//     // Arg types are mandatory for DCL statements on Functions
//     | K_FUNCTION fn=functionName
//       (
//         '('
//           (
//             v=comparatorType { argsTypes.add(v); }
//             ( ',' v=comparatorType { argsTypes.add(v); } )*
//           )?
//         ')'
//       )
//       { $res = FunctionResource.functionFromCql($fn.s.keyspace, $fn.s.name, argsTypes); }
//     ;

// /**
//  * CREATE USER [IF NOT EXISTS] <username> [WITH PASSWORD <password>] [SUPERUSER|NOSUPERUSER]
//  */
// createUserStatement
//     @init {
//         RoleOptions opts = new RoleOptions();
//         opts.setOption(IRoleManager.Option.LOGIN, true);
//         boolean superuser = false;
//         boolean ifNotExists = false;
//         RoleName name = new RoleName();
//     }
//     : K_CREATE K_USER (K_IF K_NOT K_EXISTS { ifNotExists = true; })? u=username { name.setName($u.text, true); }
//       ( K_WITH userPassword[opts] )?
//       ( K_SUPERUSER { superuser = true; } | K_NOSUPERUSER { superuser = false; } )?
//       { opts.setOption(IRoleManager.Option.SUPERUSER, superuser);
//         if (opts.getPassword().isPresent() && opts.getHashedPassword().isPresent())
//         {
//            throw new SyntaxException("Options 'password' and 'hashed password' are mutually exclusive");
//         }
//         if (opts.getPassword().isPresent() && opts.isGeneratedPassword())
//         {
//            throw new SyntaxException("Options 'password' and 'generated password' are mutually exclusive");
//         }
//         if (opts.getHashedPassword().isPresent() && opts.isGeneratedPassword())
//         {
//            throw new SyntaxException("Options 'hashed password' and 'generated password' are mutually exclusive");
//         }
//         $stmt = new CreateRoleStatement(name, opts, DCPermissions.all(), CIDRPermissions.all(), ifNotExists); }
//     ;

// /**
//  * ALTER USER [IF EXISTS] <username> [WITH PASSWORD <password>] [SUPERUSER|NOSUPERUSER]
//  */
// alterUserStatement
//     @init {
//         RoleOptions opts = new RoleOptions();
//         RoleName name = new RoleName();
//         boolean ifExists = false;
//     }
//     : K_ALTER K_USER (K_IF K_EXISTS { ifExists = true; })? u=username { name.setName($u.text, true); }
//       ( K_WITH userPassword[opts] )?
//       ( K_SUPERUSER { opts.setOption(IRoleManager.Option.SUPERUSER, true); }
//         | K_NOSUPERUSER { opts.setOption(IRoleManager.Option.SUPERUSER, false); } ) ?
//       {
//          if (opts.getPassword().isPresent() && opts.getHashedPassword().isPresent())
//          {
//             throw new SyntaxException("Options 'password' and 'hashed password' are mutually exclusive");
//          }
//          if (opts.getPassword().isPresent() && opts.isGeneratedPassword())
//          {
//             throw new SyntaxException("Options 'password' and 'generated password' are mutually exclusive");
//          }
//          if (opts.getHashedPassword().isPresent() && opts.isGeneratedPassword())
//          {
//             throw new SyntaxException("Options 'hashed password' and 'generated password' are mutually exclusive");
//          }
//          $stmt = new AlterRoleStatement(name, opts, null, null, ifExists);
//       }
//     ;

// /**
//  * DROP USER [IF EXISTS] <username>
//  */
// dropUserStatement
//     @init {
//         boolean ifExists = false;
//         RoleName name = new RoleName();
//     }
//     : K_DROP K_USER (K_IF K_EXISTS { ifExists = true; })? u=username { name.setName($u.text, true); $stmt = new DropRoleStatement(name, ifExists); }
//     ;
// /**
//  * ADD IDENTITY [IF NOT EXISTS] <identity> TO ROLE <role>
//  */
// addIdentityStatement
//     @init {
//         String identity = null;
//         String role = null;
//         boolean ifNotExists = false;
//     }
//     : K_ADD K_IDENTITY (K_IF K_NOT K_EXISTS { ifNotExists = true; })? u=identity { identity= $u.text; } K_TO K_ROLE r=identity { role=$r.text; $stmt = new AddIdentityStatement(identity, role, ifNotExists); }
//     ;

// /**
//  * DROP IDENTITY [IF EXISTS] <identity>
//  */
//  dropIdentityStatement
//       @init {
//           boolean ifExists = false;
//           String identity = null;
//       }
//       : K_DROP K_IDENTITY (K_IF K_EXISTS { ifExists = true; })? u=identity { identity= $u.text; $stmt = new DropIdentityStatement(identity, ifExists);}
//       ;

/**
 * LIST USERS
 */
listUsersStatement
    : K_LIST K_USERS
    ;

// /**
//  * CREATE ROLE [IF NOT EXISTS] <rolename> [ [WITH] option [ [AND] option ]* ]
//  *
//  * where option can be:
//  *  PASSWORD = '<password>'
//  *  SUPERUSER = (true|false)
//  *  LOGIN = (true|false)
//  *  OPTIONS = { 'k1':'v1', 'k2':'v2'}
//  *  ACCESS TO ALL DATACENTERS
//  *  ACCESS TO DATACENTERS { dcPermission (, dcPermission)* }
//  *  ACCESS FROM ALL CIDRS
//  *  ACCESS FROM CIDRS { cidrPermission (, cidrPermission)* }
//  */
// createRoleStatement
//     @init {
//         RoleOptions opts = new RoleOptions();
//         DCPermissions.Builder dcperms = DCPermissions.builder();
//         CIDRPermissions.Builder cidrperms = CIDRPermissions.builder();
//         boolean ifNotExists = false;
//     }
//     : K_CREATE K_ROLE (K_IF K_NOT K_EXISTS { ifNotExists = true; })? name=userOrRoleName
//       ( K_WITH roleOptions[opts, dcperms, cidrperms] )?
//       {
//         // set defaults if they weren't explictly supplied
//         if (!opts.getLogin().isPresent())
//         {
//             opts.setOption(IRoleManager.Option.LOGIN, false);
//         }
//         if (!opts.getSuperuser().isPresent())
//         {
//             opts.setOption(IRoleManager.Option.SUPERUSER, false);
//         }
//         if (opts.getPassword().isPresent() && opts.getHashedPassword().isPresent())
//         {
//             throw new SyntaxException("Options 'password' and 'hashed password' are mutually exclusive");
//         }
//         if (opts.getPassword().isPresent() && opts.isGeneratedPassword())
//         {
//             throw new SyntaxException("Options 'password' and 'generated password' are mutually exclusive");
//         }
//         if (opts.getHashedPassword().isPresent() && opts.isGeneratedPassword())
//         {
//            throw new SyntaxException("Options 'hashed password' and 'generated password' are mutually exclusive");
//         }
//         $stmt = new CreateRoleStatement(name, opts, dcperms.build(), cidrperms.build(), ifNotExists);
//       }
//     ;

// /**
//  * ALTER ROLE [IF EXISTS] <rolename> [ [WITH] option [ [AND] option ]* ]
//  *
//  * where option can be:
//  *  PASSWORD = '<password>'
//  *  SUPERUSER = (true|false)
//  *  LOGIN = (true|false)
//  *  OPTIONS = { 'k1':'v1', 'k2':'v2'}
//  *  ACCESS TO ALL DATACENTERS
//  *  ACCESS TO DATACENTERS { dcPermission (, dcPermission)* }
//  *  ACCESS FROM ALL CIDRS
//  *  ACCESS FROM CIDRS { cidrPermission (, cidrPermission)* }
//  */
// alterRoleStatement
//     @init {
//         RoleOptions opts = new RoleOptions();
//         DCPermissions.Builder dcperms = DCPermissions.builder();
//         CIDRPermissions.Builder cidrperms = CIDRPermissions.builder();
//         boolean ifExists = false;
//     }
//     : K_ALTER K_ROLE (K_IF K_EXISTS { ifExists = true; })? name=userOrRoleName
//       ( K_WITH roleOptions[opts, dcperms, cidrperms] )?
//       {
//          if (opts.getPassword().isPresent() && opts.getHashedPassword().isPresent())
//          {
//             throw new SyntaxException("Options 'password' and 'hashed password' are mutually exclusive");
//          }
//          if (opts.getPassword().isPresent() && opts.isGeneratedPassword())
//          {
//             throw new SyntaxException("Options 'password' and 'generated password' are mutually exclusive");
//          }
//          if (opts.getHashedPassword().isPresent() && opts.isGeneratedPassword())
//          {
//             throw new SyntaxException("Options 'hashed password' and 'generated password' are mutually exclusive");
//          }
//          $stmt = new AlterRoleStatement(name, opts, dcperms.isModified() ? dcperms.build() : null, cidrperms.isModified() ? cidrperms.build() : null, ifExists);
//       }
//     ;

// /**
//  * DROP ROLE [IF EXISTS] <rolename>
//  */
// dropRoleStatement
//     @init {
//         boolean ifExists = false;
//     }
//     : K_DROP K_ROLE (K_IF K_EXISTS { ifExists = true; })? name=userOrRoleName
//       { $stmt = new DropRoleStatement(name, ifExists); }
//     ;

// /**
//  * LIST ROLES [OF <rolename>] [NORECURSIVE]
//  */
// listRolesStatement
//     @init {
//         boolean recursive = true;
//         RoleName grantee = new RoleName();
//     }
//     : K_LIST K_ROLES
//       ( K_OF roleName[grantee])?
//       ( K_NORECURSIVE { recursive = false; } )?
//       { $stmt = new ListRolesStatement(grantee, recursive); }
//     ;

// /**
//  * LIST SUPERUSERS
//  */
// listSuperUsersStatement
//     @init {
//     }
//     : K_LIST K_SUPERUSERS { $stmt = new ListSuperUsersStatement(); }
//     ;

// roleOptions[RoleOptions opts, DCPermissions.Builder dcperms, CIDRPermissions.Builder cidrperms]
//     : roleOption[opts, dcperms, cidrperms] (K_AND roleOption[opts, dcperms, cidrperms])*
//     ;

// roleOption[RoleOptions opts, DCPermissions.Builder dcperms, CIDRPermissions.Builder cidrperms]
//     :  K_PASSWORD '=' v=STRING_LITERAL { opts.setOption(IRoleManager.Option.PASSWORD, $v.text); }
//     |  K_GENERATED K_PASSWORD { opts.setOption(IRoleManager.Option.GENERATED_PASSWORD, Boolean.TRUE); }
//     |  K_HASHED K_PASSWORD '=' v=STRING_LITERAL { opts.setOption(IRoleManager.Option.HASHED_PASSWORD, $v.text); }
//     |  K_OPTIONS '=' m=fullMapLiteral { opts.setOption(IRoleManager.Option.OPTIONS, convertPropertyMap(m)); }
//     |  K_SUPERUSER '=' b=BOOLEAN { opts.setOption(IRoleManager.Option.SUPERUSER, Boolean.valueOf($b.text)); }
//     |  K_LOGIN '=' b=BOOLEAN { opts.setOption(IRoleManager.Option.LOGIN, Boolean.valueOf($b.text)); }
//     |  K_ACCESS K_TO K_ALL K_DATACENTERS { dcperms.all(); }
//     |  K_ACCESS K_TO K_DATACENTERS '{' dcPermission[dcperms] (',' dcPermission[dcperms])* '}'
//     |  K_ACCESS K_FROM K_ALL K_CIDRS { cidrperms.all(); }
//     |  K_ACCESS K_FROM K_CIDRS '{' cidrPermission[cidrperms] (',' cidrPermission[cidrperms])* '}'
//     ;

// dcPermission[DCPermissions.Builder builder]
//     : dc=STRING_LITERAL { builder.add($dc.text); }
//     ;

// cidrPermission[CIDRPermissions.Builder builder]
//     : cidr=STRING_LITERAL { builder.add($cidr.text); }
//     ;

// // for backwards compatibility in CREATE/ALTER USER, this has no '='
// userPassword[RoleOptions opts]
//     :  K_PASSWORD v=STRING_LITERAL { opts.setOption(IRoleManager.Option.PASSWORD, $v.text); }
//     |  K_HASHED K_PASSWORD v=STRING_LITERAL { opts.setOption(IRoleManager.Option.HASHED_PASSWORD, $v.text); }
//     |  K_GENERATED K_PASSWORD { opts.setOption(IRoleManager.Option.GENERATED_PASSWORD, Boolean.TRUE); }
//     ;

// /**
//  * DESCRIBE statement(s)
//  *
//  * Must be in sync with the javadoc for org.apache.cassandra.cql3.statements.DescribeStatement and the
//  * cqlsh syntax definition in for cqlsh_describe_cmd_syntax_rules pylib/cqlshlib/cqlshhandling.py.
//  */
// describeStatement
//     @init {
//         boolean fullSchema = false;
//         boolean pending = false;
//         boolean config = false;
//         boolean only = false;
//         QualifiedName gen = new QualifiedName();
//     }
//     : ( K_DESCRIBE | K_DESC )
//     ( (K_CLUSTER)=> K_CLUSTER                     { $stmt = DescribeStatement.cluster(); }
//     | (K_FULL { fullSchema=true; })? K_SCHEMA     { $stmt = DescribeStatement.schema(fullSchema); }
//     | (K_KEYSPACES)=> K_KEYSPACES                 { $stmt = DescribeStatement.keyspaces(); }
//     | (K_ONLY { only=true; })? K_KEYSPACE ( ks=keyspaceName )?
//                                                   { $stmt = DescribeStatement.keyspace(ks, only); }
//     | (K_TABLES) => K_TABLES                      { $stmt = DescribeStatement.tables(); }
//     | K_COLUMNFAMILY cf=columnFamilyName          { $stmt = DescribeStatement.table(cf.getKeyspace(), cf.getName()); }
//     | K_INDEX idx=columnFamilyName                { $stmt = DescribeStatement.index(idx.getKeyspace(), idx.getName()); }
//     | K_MATERIALIZED K_VIEW view=columnFamilyName { $stmt = DescribeStatement.view(view.getKeyspace(), view.getName()); }
//     | (K_TYPES) => K_TYPES                        { $stmt = DescribeStatement.types(); }
//     | K_TYPE tn=userTypeName                      { $stmt = DescribeStatement.type(tn.getKeyspace(), tn.getStringTypeName()); }
//     | (K_FUNCTIONS) => K_FUNCTIONS                { $stmt = DescribeStatement.functions(); }
//     | K_FUNCTION fn=functionName                  { $stmt = DescribeStatement.function(fn.keyspace, fn.name); }
//     | (K_AGGREGATES) => K_AGGREGATES              { $stmt = DescribeStatement.aggregates(); }
//     | K_AGGREGATE ag=functionName                 { $stmt = DescribeStatement.aggregate(ag.keyspace, ag.name); }
//     | ( ( ksT=IDENT                       { gen.setKeyspace($ksT.text, false);}
//           | ksT=QUOTED_NAME                 { gen.setKeyspace($ksT.text, true);}
//           | ksK=unreserved_keyword          { gen.setKeyspace(ksK, false);} ) '.' )?
//         ( tT=IDENT                          { gen.setName($tT.text, false);}
//         | tT=QUOTED_NAME                    { gen.setName($tT.text, true);}
//         | tK=unreserved_keyword             { gen.setName(tK, false);} )
//                                                     { $stmt = DescribeStatement.generic(gen.getKeyspace(), gen.getName()); }
//     )
//     ( K_WITH K_INTERNALS { $stmt.withInternalDetails(); } )?
//     ;

// /** DEFINITIONS **/

// // Like ident, but for case where we take a column name that can be the legacy super column empty name. Importantly,
// // this should not be used in DDL statements, as we don't want to let users create such column.
// cident
//     : EMPTY_QUOTED_NAME    { $id = ColumnIdentifier.getInterned("", true); }
//     | t=ident              { $id = t; }
//     ;

ident
    : t=IDENT                   # IdentIdent
    | t=QUOTED_NAME             # IdentQuotedIdent
    | k=unreserved_keyword      # IdentIdent
    ;

fident
    : t=IDENT                   # FIdentIdent
    | t=QUOTED_NAME             # FIdentQuotedIdent
    | k=unreserved_keyword      # FIdentIdent
    ;

// Identifiers that do not refer to columns
noncol_ident
    : i=IDENT                   # NonColIdent
    | t=QUOTED_NAME             # NonColQuoted
    | k=unreserved_keyword      # NonColIdent
    ;

// Keyspace & Column family names
keyspaceName
    : ksName
    ;

// indexName
//     @init { $name = new QualifiedName(); }
//     : (ksName[name] '.')? idxName[name]
//     ;

columnFamilyName
    : (ksName '.')? cfName
    ;

userTypeName
    : (ks=noncol_ident '.')? ut=non_type_ident
    ;

// userOrRoleName
//     @init { RoleName role = new RoleName(); }
//     : roleName[role] {$name = role;}
//     ;

ksName
    : t=IDENT                   # KsNameIdent
    | t=QUOTED_NAME             # KsNameQuotedIdent
    | k=unreserved_keyword      # KsNameIdent
    | QMARK                     # KsNameInvalidBind
    ;

cfName
    : t=IDENT                   # CfNameIdent
    | t=QUOTED_NAME             # CfNameQuotedIdent
    | k=unreserved_keyword      # CfNameIdent
    | QMARK                     # CfNameInvalidBind
    ;

// idxName[QualifiedName name]
//     : t=IDENT              { $name.setName($t.text, false); }
//     | t=QUOTED_NAME        { $name.setName($t.text, true);}
//     | k=unreserved_keyword { $name.setName(k, false); }
//     | QMARK {addRecognitionError("Bind variables cannot be used for index names");}
//     ;

// roleName[RoleName name]
//     : t=IDENT              { $name.setName($t.text, false); }
//     | s=STRING_LITERAL     { $name.setName($s.text, true); }
//     | t=QUOTED_NAME        { $name.setName($t.text, true); }
//     | k=unreserved_keyword { $name.setName(k, false); }
//     | QMARK {addRecognitionError("Bind variables cannot be used for role names");}
//     ;

constant
    : t=STRING_LITERAL
    | t=INTEGER
    | t=FLOAT
    | t=BOOLEAN
    | t=DURATION
    | t=UUID
    | t=HEXNUMBER
    | ((K_POSITIVE_NAN | K_NEGATIVE_NAN)
        | K_POSITIVE_INFINITY
        | K_NEGATIVE_INFINITY )
    ;

fullMapLiteral
    : '{' ( k1=term ':' v1=term ( ',' kn=term ':' vn=term )* )? '}'
    ;

setOrMapLiteral
    : m=mapLiteral
    | s=setLiteral
    ;

setLiteral
    : ( ',' tn=term )*
    ;

mapLiteral
    : ':' v=term ( ',' kn=term ':' vn=term )*
    ;

collectionLiteral
    : l=listLiteral
    | '{' t=term v=setOrMapLiteral '}'
    // Note that we have an ambiguity between maps and set for "{}". So we force it to a set literal,
    // and deal with it later based on the type of the column (SetLiteral.java).
    | '{' '}'
    ;

listLiteral
    : '[' ( t1=term ( ',' tn=term )* )? ']'
    ;

usertypeLiteral
    // We don't allow empty literals because that conflicts with sets/maps and is currently useless since we don't allow empty user types
    : '{' k1=fident ':' v1=term ( ',' kn=fident ':' vn=term )* '}'
    ;

tupleLiteral
    : '(' t1=term ( ',' tn=term )* ')'
    ;

value
    : c=constant
    | l=collectionLiteral
    | u=usertypeLiteral
    | t=tupleLiteral
    | K_NULL
    | m=marker
    ;

// intValue
//     : t=INTEGER { $value = Constants.Literal.integer($t.text); }
//     | m=marker  { $value = m; }
//     ;

functionName
     // antlr might try to recover and give a null for f. It will still error out in the end, but FunctionName
     // wouldn't be happy with that so we should bypass this for now or we'll have a weird user-facing error
    : (ks=keyspaceName '.')? f=allowedFunctionName
    ;

allowedFunctionName
    : f=IDENT                           # AllowedFunctionNameIdent
    | f=QUOTED_NAME                     # AllowedFunctionNameQuotedIdent
    | u=unreserved_function_keyword     # AllowedFunctionNameIdent
    | K_TOKEN                           # AllowedFunctionNameIdent
    | K_COUNT                           # AllowedFunctionNameIdent
    ;

function
    : f=functionName '(' ')'
    | f=functionName '(' args=functionArgs ')'
    ;

functionArgs
    : t1=term ( ',' tn=term )*
    ;

term
    : t=termAddition
    ;

termAddition
    :   l=termMultiplication
        ( '+' r=termMultiplication
        | '-' r=termMultiplication
        )*
    ;

termMultiplication
    :   l=termGroup
        ( '*' r=termGroup
        | '/' r=termGroup
        | '%' r=termGroup
        )*
    ;

termGroup
    : t=simpleTerm
    | '-'  t=simpleTerm
    ;

simpleTerm
    : v=value
    | f=function
    | '(' c=comparatorType ')' t=simpleTerm
    | K_CAST '(' t=simpleTerm K_AS n=native_type ')'
    ;

// columnOperation[List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations]
//     : key=cident columnOperationDifferentiator[operations, key]
//     ;

// columnOperationDifferentiator[List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations, ColumnIdentifier key]
//     : '=' normalColumnOperation[operations, key]
//     | shorthandColumnOperation[operations, key]
//     | '[' k=term ']' collectionColumnOperation[operations, key, k]
//     | '.' field=fident udtColumnOperation[operations, key, field]
//     ;

// normalColumnOperation[List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations, ColumnIdentifier key]
//     : t=term ('+' c=cident )?
//       {
//           if (c == null)
//           {
//               addRawUpdate(operations, key, new Operation.SetValue(t));
//           }
//           else
//           {
//               if (!key.equals(c))
//                   addRecognitionError("Only expressions of the form X = <value> + X are supported.");
//               addRawUpdate(operations, key, new Operation.Prepend(t));
//           }
//       }
//     | c=cident sig=('+' | '-') t=term
//       {
//           if (!key.equals(c))
//               addRecognitionError("Only expressions of the form X = X " + $sig.text + "<value> are supported.");
//           addRawUpdate(operations, key, $sig.text.equals("+") ? new Operation.Addition(t) : new Operation.Substraction(t));
//       }
//     | c=cident i=INTEGER
//       {
//           // Note that this production *is* necessary because X = X - 3 will in fact be lexed as [ X, '=', X, INTEGER].
//           if (!key.equals(c))
//               // We don't yet allow a '+' in front of an integer, but we could in the future really, so let's be future-proof in our error message
//               addRecognitionError("Only expressions of the form X = X " + ($i.text.charAt(0) == '-' ? '-' : '+') + " <value> are supported.");
//           addRawUpdate(operations, key, new Operation.Addition(Constants.Literal.integer($i.text)));
//       }
//     ;

// shorthandColumnOperation[List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations, ColumnIdentifier key]
//     : sig=('+=' | '-=') t=term
//       {
//           addRawUpdate(operations, key, $sig.text.equals("+=") ? new Operation.Addition(t) : new Operation.Substraction(t));
//       }
//     ;

// collectionColumnOperation[List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations, ColumnIdentifier key, Term.Raw k]
//     : '=' t=term
//       {
//           addRawUpdate(operations, key, new Operation.SetElement(k, t));
//       }
//     ;

// udtColumnOperation[List<Pair<ColumnIdentifier, Operation.RawUpdate>> operations, ColumnIdentifier key, FieldIdentifier field]
//     : '=' t=term
//       {
//           addRawUpdate(operations, key, new Operation.SetField(field, t));
//       }
//     ;

// columnCondition
//     // Note: we'll reject duplicates later
//     : column=cident
//         ( op=relationType t=term       { $condition = ColumnCondition.Raw.simpleCondition(column, op, Terms.Raw.of(t)); }
//         | K_CONTAINS (K_KEY)? t=term   { $condition = ColumnCondition.Raw.simpleCondition(column, $K_KEY != null ? Operator.CONTAINS_KEY : Operator.CONTAINS, Terms.Raw.of(t)); }
//         | K_IN v=singleColumnInValues  { $condition = ColumnCondition.Raw.simpleCondition(column, Operator.IN, v); }
//         | '[' element=term ']'
//             ( op=relationType t=term      { $condition = ColumnCondition.Raw.collectionElementCondition(column, element, op, Terms.Raw.of(t)); }
//             | K_IN v=singleColumnInValues { $condition = ColumnCondition.Raw.collectionElementCondition(column, element, Operator.IN, v); }
//             )
//         | '.' field=fident
//             ( op=relationType t=term      { $condition = ColumnCondition.Raw.udtFieldCondition(column, field, op, Terms.Raw.of(t)); }
//             | K_IN v=singleColumnInValues { $condition = ColumnCondition.Raw.udtFieldCondition(column, field, Operator.IN, v); }
//             )
//         )
//     ;

properties
    : property (K_AND property)*
    ;

property
    : k=noncol_ident '=' simple=propertyValue
    | k=noncol_ident '=' map=fullMapLiteral
    ;

propertyValue
    : c=constant
    | u=unreserved_keyword
    ;

// singleColumnBetweenValues
//     @init { List<Term.Raw> list = new ArrayList<>(); }
//     @after { $terms = Terms.Raw.of(list); }
//     : t1=term { list.add(t1); } K_AND t2=term { list.add(t2); }
//     ;

// relationType
//     : '='  { $op = Operator.EQ; }
//     | '<'  { $op = Operator.LT; }
//     | '<=' { $op = Operator.LTE; }
//     | '>'  { $op = Operator.GT; }
//     | '>=' { $op = Operator.GTE; }
//     | '!=' { $op = Operator.NEQ; }
//     ;

// relation[WhereClause.Builder clauses]
//     : name=cident
//            ( type=relationType t=term { $clauses.add(Relation.singleColumn(name, type, t)); }
//            | K_BETWEEN betweenValues=singleColumnBetweenValues { $clauses.add(Relation.singleColumn(name, Operator.BETWEEN, betweenValues)); }
//            | K_LIKE t=term { $clauses.add(Relation.singleColumn(name, Operator.LIKE, t)); }
//            | K_IS K_NOT K_NULL { $clauses.add(Relation.singleColumn(name, Operator.IS_NOT, Constants.NULL_LITERAL)); }
//            | rtInOperator=inOperator inValue=singleColumnInValues { $clauses.add(Relation.singleColumn(name, rtInOperator, inValue)); }
//            | rtContainsOperator=containsOperator t=term { $clauses.add(Relation.singleColumn(name, rtContainsOperator, t)); }
//            )
//     | K_TOKEN l=tupleOfIdentifiers
//         ( type=relationType t=term { $clauses.add(Relation.token(l, type, t)); }
//         | K_BETWEEN betweenValues=singleColumnBetweenValues { $clauses.add(Relation.token(l, Operator.BETWEEN, betweenValues)); }
//         )
//     | name=cident '[' key=term ']' type=relationType t=term { $clauses.add(Relation.mapElement(name, key, type, t)); }
//     | ids=tupleOfIdentifiers
//         ( rt=inOperator inValue=multiColumnInValues { $clauses.add(Relation.multiColumn(ids, rt, inValue)); }
//         | type=relationType v=multiColumnValue {$clauses.add(Relation.multiColumn(ids, type, v)); }
//         | K_BETWEEN t1=multiColumnValue K_AND t2=multiColumnValue { $clauses.add(Relation.multiColumn(ids, Operator.BETWEEN, Terms.Raw.of(List.of(t1, t2)))); }
//         )
//     | '(' relation[$clauses] ')'
//     ;

// containsOperator
//     : K_CONTAINS { o = Operator.CONTAINS; } (K_KEY { o = Operator.CONTAINS_KEY; })?
//     | K_NOT K_CONTAINS { o = Operator.NOT_CONTAINS; } (K_KEY { o = Operator.NOT_CONTAINS_KEY; })?
//     ;

// inOperator
//     : K_IN { o = Operator.IN; }
//     | K_NOT K_IN { o = Operator.NOT_IN; }
//     ;

// inMarker
//     : QMARK { $marker = newINBindVariables(null); }
//     | ':' name=noncol_ident { $marker = newINBindVariables(name); }
//     ;

// tupleOfIdentifiers
//     @init { $ids = new ArrayList<ColumnIdentifier>(); }
//     : '(' n1=cident { $ids.add(n1); } (',' ni=cident { $ids.add(ni); })* ')'
//     ;

// singleColumnInValues
//     : t=terms     { $terms = t;}
//     | m=inMarker  { $terms = m;}
//     ;

// terms
//     @init { List<Term.Raw> list = new ArrayList<>(); }
//     @after { $terms = Terms.Raw.of(list); }
//     : '(' ( t1 = term { list.add(t1); } (',' ti=term { list.add(ti); })* )? ')'
//     ;

// multiColumnValue
//     : l=tupleLiteral { $term = l; } /* (a, b, c) > (1, 2, 3) or (a, b, c) > (?, ?, ?) */
//     | m=marker       { $term = m; } /* (a, b, c) >= ? */
//     ;

// multiColumnInValues
//     : '(' ')'                    { $terms = Terms.Raw.of();}  /* (a, b, c) IN () */
//     | m=inMarker                 { $terms = m; }              /* (a, b, c) IN ? */
//     | tl=tupleOfTupleLiterals    { $terms = tl; }             /* (a, b, c) IN ((1, 2, 3), (4, 5, 6), ...) */
//     | tm=tupleOfMarkersForTuples { $terms = tm; }             /* (a, b, c) IN (?, ?, ...) */
//     ;

// tupleOfTupleLiterals
//     @init { List<Term.Raw> list = new ArrayList<>(); }
//     @after { $literals = Terms.Raw.of(list); }
//     : '(' t1=tupleLiteral { list.add(t1); } (',' ti=tupleLiteral { list.add(ti); })* ')'
//     ;

// tupleOfMarkersForTuples
//     @init { List<Term.Raw> list = new ArrayList<>(); }
//     @after { $markers = Terms.Raw.of(list); }
//     : '(' m1=marker { list.add(m1); } (',' mi=marker { list.add(mi); })* ')'
//     ;

comparatorType
    : n=native_type
    | c=collection_type
    | tt=tuple_type
    | vc=vector_type
    | id=userTypeName
    | K_FROZEN '<' f=comparatorType '>'
    | s=STRING_LITERAL
    ;

native_type
    : K_ASCII
    | K_BIGINT
    | K_BLOB
    | K_BOOLEAN
    | K_COUNTER
    | K_DECIMAL
    | K_DOUBLE
    | K_DURATION
    | K_FLOAT
    | K_INET
    | K_INT
    | K_SMALLINT
    | K_TEXT
    | K_TIMESTAMP
    | K_TINYINT
    | K_UUID
    | K_VARCHAR
    | K_VARINT
    | K_TIMEUUID
    | K_DATE
    | K_TIME
    ;

collection_type
    : K_MAP  '<' t1=comparatorType ',' t2=comparatorType '>'
    | K_LIST '<' t=comparatorType '>'
    | K_SET  '<' t=comparatorType '>'
    ;

tuple_type
    : K_TUPLE '<' t1=comparatorType (',' tn=comparatorType)* '>'
    ;

vector_type
    : K_VECTOR '<' t1=comparatorType ','  d=INTEGER '>'
    ;

// username
//     : IDENT
//     | STRING_LITERAL
//     | QUOTED_NAME { addRecognitionError("Quoted strings are are not supported for user names and USER is deprecated, please use ROLE");}
//     ;

// identity
//     : IDENT
//     | STRING_LITERAL
//     | QUOTED_NAME { addRecognitionError("Quoted strings are are not supported for identity");}
//     ;

// mbean
//     : STRING_LITERAL
//     ;

// Basically the same as cident, but we need to exlude existing CQL3 types
// (which for some reason are not reserved otherwise)
non_type_ident
    : i=(IDENT | K_DATE)            # NonTypeIdent
    | t=QUOTED_NAME                 # NonTypeQuotedIdent
    | k=basic_unreserved_keyword    # NonTypeIdent
    | kk=K_KEY                      # NonTypeIdent
    ;

unreserved_keyword
    : u=unreserved_function_keyword
    | k=(K_TTL | K_COUNT | K_WRITETIME | K_MAXWRITETIME | K_KEY | K_CAST | K_JSON | K_DISTINCT)
    ;

unreserved_function_keyword
    : u=basic_unreserved_keyword
    | t=native_type
    ;

basic_unreserved_keyword
    : k=( K_KEYS
        | K_AS
        | K_CLUSTER
        | K_CLUSTERING
        | K_COMPACT
        | K_STORAGE
        | K_TABLES
        | K_TYPE
        | K_TYPES
        | K_VALUES
        | K_MAP
        | K_LIST
        | K_FILTERING
        | K_PERMISSION
        | K_PERMISSIONS
        | K_KEYSPACES
        | K_ALL
        | K_USER
        | K_USERS
        | K_ROLE
        | K_ROLES
        | K_IDENTITY
        | K_SUPERUSER
        | K_SUPERUSERS
        | K_NOSUPERUSER
        | K_LOGIN
        | K_NOLOGIN
        | K_OPTIONS
        | K_PASSWORD
        | K_GENERATED
        | K_HASHED
        | K_EXISTS
        | K_CUSTOM
        | K_TRIGGER
        | K_CONTAINS
        | K_INTERNALS
        | K_ONLY
        | K_STATIC
        | K_FROZEN
        | K_TUPLE
        | K_FUNCTION
        | K_FUNCTIONS
        | K_AGGREGATE
        | K_AGGREGATES
        | K_SFUNC
        | K_STYPE
        | K_FINALFUNC
        | K_INITCOND
        | K_RETURNS
        | K_LANGUAGE
        | K_CALLED
        | K_INPUT
        | K_LIKE
        | K_PER
        | K_PARTITION
        | K_GROUP
        | K_DATACENTERS
        | K_CIDRS
        | K_ACCESS
        | K_DEFAULT
        | K_MBEAN
        | K_MBEANS
        | K_REPLACE
        | K_UNSET
        | K_MASKED
        | K_UNMASK
        | K_SELECT_MASKED
        | K_VECTOR
        | K_ANN
        | K_BETWEEN
        )
    ;

// Case-insensitive keywords
// When adding a new reserved keyword, add entry to o.a.c.cql3.ReservedKeywords and
// pylib/cqlshlib/cqlhandling.py::cql_keywords_reserved.
// When adding a new unreserved keyword, add entry to unreserved keywords in Parser.g.
K_SELECT:      S E L E C T;
K_FROM:        F R O M;
K_AS:          A S;
K_WHERE:       W H E R E;
K_AND:         A N D;
K_KEY:         K E Y;
K_KEYS:        K E Y S;
K_ENTRIES:     E N T R I E S;
K_FULL:        F U L L;
K_INSERT:      I N S E R T;
K_UPDATE:      U P D A T E;
K_WITH:        W I T H;
K_LIMIT:       L I M I T;
K_PER:         P E R;
K_PARTITION:   P A R T I T I O N;
K_USING:       U S I N G;
K_USE:         U S E;
K_DISTINCT:    D I S T I N C T;
K_COUNT:       C O U N T;
K_SET:         S E T;
K_BEGIN:       B E G I N;
K_UNLOGGED:    U N L O G G E D;
K_BATCH:       B A T C H;
K_APPLY:       A P P L Y;
K_TRUNCATE:    T R U N C A T E;
K_DELETE:      D E L E T E;
K_IN:          I N;
K_CREATE:      C R E A T E;
K_SCHEMA:      S C H E M A;
K_KEYSPACE:    ( K E Y S P A C E
                 | K_SCHEMA );
K_KEYSPACES:   K E Y S P A C E S;
K_COLUMNFAMILY:( C O L U M N F A M I L Y
                 | T A B L E );
K_TABLES:      ( C O L U M N F A M I L I E S
                 | T A B L E S );
K_MATERIALIZED:M A T E R I A L I Z E D;
K_VIEW:        V I E W;
K_INDEX:       I N D E X;
K_CUSTOM:      C U S T O M;
K_ON:          O N;
K_TO:          T O;
K_DROP:        D R O P;
K_PRIMARY:     P R I M A R Y;
K_INTO:        I N T O;
K_VALUES:      V A L U E S;
K_TIMESTAMP:   T I M E S T A M P;
K_TTL:         T T L;
K_CAST:        C A S T;
K_ALTER:       A L T E R;
K_RENAME:      R E N A M E;
K_ADD:         A D D;
K_TYPE:        T Y P E;
K_TYPES:       T Y P E S;
K_COMPACT:     C O M P A C T;
K_STORAGE:     S T O R A G E;
K_ORDER:       O R D E R;
K_BY:          B Y;
K_ASC:         A S C;
K_DESC:        D E S C;
K_ALLOW:       A L L O W;
K_FILTERING:   F I L T E R I N G;
K_IF:          I F;
K_IS:          I S;
K_CONTAINS:    C O N T A I N S;
K_BETWEEN:     B E T W E E N;
K_GROUP:       G R O U P;
K_CLUSTER:     C L U S T E R;
K_INTERNALS:   I N T E R N A L S;
K_ONLY:        O N L Y;

K_GRANT:       G R A N T;
K_ALL:         A L L;
K_PERMISSION:  P E R M I S S I O N;
K_PERMISSIONS: P E R M I S S I O N S;
K_OF:          O F;
K_REVOKE:      R E V O K E;
K_MODIFY:      M O D I F Y;
K_AUTHORIZE:   A U T H O R I Z E;
K_DESCRIBE:    D E S C R I B E;
K_EXECUTE:     E X E C U T E;
K_NORECURSIVE: N O R E C U R S I V E;
K_MBEAN:       M B E A N;
K_MBEANS:      M B E A N S;

K_USER:        U S E R;
K_USERS:       U S E R S;
K_ROLE:        R O L E;
K_ROLES:       R O L E S;
K_SUPERUSERS:  S U P E R U S E R S;
K_SUPERUSER:   S U P E R U S E R;
K_NOSUPERUSER: N O S U P E R U S E R;
K_GENERATED:   G E N E R A T E D;
K_PASSWORD:    P A S S W O R D;
K_HASHED:      H A S H E D;
K_LOGIN:       L O G I N;
K_NOLOGIN:     N O L O G I N;
K_OPTIONS:     O P T I O N S;
K_ACCESS:      A C C E S S;
K_DATACENTERS: D A T A C E N T E R S;
K_CIDRS:       C I D R S;
K_IDENTITY:    I D E N T I T Y;

K_CLUSTERING:  C L U S T E R I N G;
K_ASCII:       A S C I I;
K_BIGINT:      B I G I N T;
K_BLOB:        B L O B;
K_BOOLEAN:     B O O L E A N;
K_COUNTER:     C O U N T E R;
K_DECIMAL:     D E C I M A L;
K_DOUBLE:      D O U B L E;
K_DURATION:    D U R A T I O N;
K_FLOAT:       F L O A T;
K_INET:        I N E T;
K_INT:         I N T;
K_SMALLINT:    S M A L L I N T;
K_TINYINT:     T I N Y I N T;
K_TEXT:        T E X T;
K_UUID:        U U I D;
K_VARCHAR:     V A R C H A R;
K_VARINT:      V A R I N T;
K_TIMEUUID:    T I M E U U I D;
K_TOKEN:       T O K E N;
K_WRITETIME:   W R I T E T I M E;
K_MAXWRITETIME:M A X W R I T E T I M E;
K_DATE:        D A T E;
K_TIME:        T I M E;

K_NULL:        N U L L;
K_NOT:         N O T;
K_EXISTS:      E X I S T S;

K_MAP:         M A P;
K_LIST:        L I S T;
K_POSITIVE_NAN: N A N;
K_NEGATIVE_NAN: '-' N A N;
K_POSITIVE_INFINITY:    I N F I N I T Y;
K_NEGATIVE_INFINITY: '-' I N F I N I T Y;
K_TUPLE:       T U P L E;

K_TRIGGER:     T R I G G E R;
K_STATIC:      S T A T I C;
K_FROZEN:      F R O Z E N;

K_FUNCTION:    F U N C T I O N;
K_FUNCTIONS:   F U N C T I O N S;
K_AGGREGATE:   A G G R E G A T E;
K_AGGREGATES:  A G G R E G A T E S;
K_SFUNC:       S F U N C;
K_STYPE:       S T Y P E;
K_FINALFUNC:   F I N A L F U N C;
K_INITCOND:    I N I T C O N D;
K_RETURNS:     R E T U R N S;
K_CALLED:      C A L L E D;
K_INPUT:       I N P U T;
K_LANGUAGE:    L A N G U A G E;
K_OR:          O R;
K_REPLACE:     R E P L A C E;

K_JSON:        J S O N;
K_DEFAULT:     D E F A U L T;
K_UNSET:       U N S E T;
K_LIKE:        L I K E;

K_MASKED:      M A S K E D;
K_UNMASK:      U N M A S K;
K_SELECT_MASKED: S E L E C T '_' M A S K E D;

K_VECTOR:      V E C T O R;
K_ANN:         A N N;

// Case-insensitive alpha characters
fragment A: ('a'|'A');
fragment B: ('b'|'B');
fragment C: ('c'|'C');
fragment D: ('d'|'D');
fragment E: ('e'|'E');
fragment F: ('f'|'F');
fragment G: ('g'|'G');
fragment H: ('h'|'H');
fragment I: ('i'|'I');
fragment J: ('j'|'J');
fragment K: ('k'|'K');
fragment L: ('l'|'L');
fragment M: ('m'|'M');
fragment N: ('n'|'N');
fragment O: ('o'|'O');
fragment P: ('p'|'P');
fragment Q: ('q'|'Q');
fragment R: ('r'|'R');
fragment S: ('s'|'S');
fragment T: ('t'|'T');
fragment U: ('u'|'U');
fragment V: ('v'|'V');
fragment W: ('w'|'W');
fragment X: ('x'|'X');
fragment Y: ('y'|'Y');
fragment Z: ('z'|'Z');

STRING_LITERAL
    : '$$' .*? '$$'
    | '\'' .*? '\''
    ;

QUOTED_NAME
    : '"' .*? '"'
    ;

EMPTY_QUOTED_NAME
    : '"' '"'
    ;

fragment DIGIT
    : '0'..'9'
    ;

fragment LETTER
    : ('A'..'Z' | 'a'..'z')
    ;

fragment HEX
    : ('A'..'F' | 'a'..'f' | '0'..'9')
    ;

fragment EXPONENT
    : E ('+' | '-')? DIGIT+
    ;

fragment DURATION_ISO_8601_PERIOD_DESIGNATORS
    : '-'? 'P' DIGIT+ 'Y' (DIGIT+ 'M')? (DIGIT+ 'D')?
    | '-'? 'P' DIGIT+ 'M' (DIGIT+ 'D')?
    | '-'? 'P' DIGIT+ 'D'
    ;

fragment DURATION_ISO_8601_TIME_DESIGNATORS
    : 'T' DIGIT+ 'H' (DIGIT+ 'M')? (DIGIT+ 'S')?
    | 'T' DIGIT+ 'M' (DIGIT+ 'S')?
    | 'T' DIGIT+ 'S'
    ;

fragment DURATION_ISO_8601_WEEK_PERIOD_DESIGNATOR
    : '-'? 'P' DIGIT+ 'W'
    ;

fragment DURATION_UNIT
    : Y
    | M O
    | W
    | D
    | H
    | M
    | S
    | M S
    | U S
    | '\u00B5' S
    | N S
    ;

INTEGER
    : '-'? DIGIT+
    ;

QMARK
    : '?'
    ;

RANGE
    : '..'
    ;

/*
 * Normally a lexer only emits one token at a time, but ours is tricked out
 * to support multiple (see @lexer::members near the top of the grammar).
 */
FLOAT
    : INTEGER EXPONENT
    | INTEGER '.' DIGIT* EXPONENT?
    ;

/*
 * This has to be before IDENT so it takes precendence over it.
 */
BOOLEAN
    : T R U E | F A L S E
    ;

DURATION
    : '-'? DIGIT+ DURATION_UNIT (DIGIT+ DURATION_UNIT)*
    | '-'? 'P' DIGIT DIGIT DIGIT DIGIT '-' DIGIT DIGIT '-' DIGIT DIGIT 'T' DIGIT DIGIT ':' DIGIT DIGIT ':' DIGIT DIGIT // ISO 8601 "alternative format"
    | '-'? 'P' DURATION_ISO_8601_TIME_DESIGNATORS
    | DURATION_ISO_8601_WEEK_PERIOD_DESIGNATOR
    | DURATION_ISO_8601_PERIOD_DESIGNATORS DURATION_ISO_8601_TIME_DESIGNATORS?
    ;

IDENT
    : LETTER (LETTER | DIGIT | '_')*
    ;

HEXNUMBER
    : '0' X HEX*
    ;

UUID
    : HEX HEX HEX HEX HEX HEX HEX HEX '-'
      HEX HEX HEX HEX '-'
      HEX HEX HEX HEX '-'
      HEX HEX HEX HEX '-'
      HEX HEX HEX HEX HEX HEX HEX HEX HEX HEX HEX HEX
    ;

WS
    : (' ' | '\t' | '\n' | '\r')+ -> channel(HIDDEN)
    ;

COMMENT
    : ('--' | '//') .*? ('\n'|'\r') -> channel(HIDDEN)
    ;

MULTILINE_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;
