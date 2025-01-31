// Code generated from Cql3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql3

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Cql3Parser struct {
	*antlr.BaseParser
}

var Cql3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cql3ParserInit() {
	staticData := &Cql3ParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "':'", "'('", "','", "')'", "'.'", "'{'", "'}'", "'['", "']'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'='", "'<'", "'>'", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "'?'", "'..'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "K_SELECT", "K_FROM", "K_AS", "K_WHERE", "K_AND", "K_KEY", "K_KEYS",
		"K_ENTRIES", "K_FULL", "K_INSERT", "K_UPDATE", "K_WITH", "K_LIMIT",
		"K_PER", "K_PARTITION", "K_USING", "K_USE", "K_DISTINCT", "K_COUNT",
		"K_SET", "K_BEGIN", "K_UNLOGGED", "K_BATCH", "K_APPLY", "K_TRUNCATE",
		"K_DELETE", "K_IN", "K_CREATE", "K_SCHEMA", "K_KEYSPACE", "K_KEYSPACES",
		"K_COLUMNFAMILY", "K_TABLES", "K_MATERIALIZED", "K_VIEW", "K_INDEX",
		"K_CUSTOM", "K_ON", "K_TO", "K_DROP", "K_PRIMARY", "K_INTO", "K_VALUES",
		"K_TIMESTAMP", "K_TTL", "K_CAST", "K_ALTER", "K_RENAME", "K_ADD", "K_TYPE",
		"K_TYPES", "K_COMPACT", "K_STORAGE", "K_ORDER", "K_BY", "K_ASC", "K_DESC",
		"K_ALLOW", "K_FILTERING", "K_IF", "K_IS", "K_CONTAINS", "K_BETWEEN",
		"K_GROUP", "K_CLUSTER", "K_INTERNALS", "K_ONLY", "K_GRANT", "K_ALL",
		"K_PERMISSION", "K_PERMISSIONS", "K_OF", "K_REVOKE", "K_MODIFY", "K_AUTHORIZE",
		"K_DESCRIBE", "K_EXECUTE", "K_NORECURSIVE", "K_MBEAN", "K_MBEANS", "K_USER",
		"K_USERS", "K_ROLE", "K_ROLES", "K_SUPERUSERS", "K_SUPERUSER", "K_NOSUPERUSER",
		"K_GENERATED", "K_PASSWORD", "K_HASHED", "K_LOGIN", "K_NOLOGIN", "K_OPTIONS",
		"K_ACCESS", "K_DATACENTERS", "K_CIDRS", "K_IDENTITY", "K_CLUSTERING",
		"K_ASCII", "K_BIGINT", "K_BLOB", "K_BOOLEAN", "K_COUNTER", "K_DECIMAL",
		"K_DOUBLE", "K_DURATION", "K_FLOAT", "K_INET", "K_INT", "K_SMALLINT",
		"K_TINYINT", "K_TEXT", "K_UUID", "K_VARCHAR", "K_VARINT", "K_TIMEUUID",
		"K_TOKEN", "K_WRITETIME", "K_MAXWRITETIME", "K_DATE", "K_TIME", "K_NULL",
		"K_NOT", "K_EXISTS", "K_MAP", "K_LIST", "K_POSITIVE_NAN", "K_NEGATIVE_NAN",
		"K_POSITIVE_INFINITY", "K_NEGATIVE_INFINITY", "K_TUPLE", "K_TRIGGER",
		"K_STATIC", "K_FROZEN", "K_FUNCTION", "K_FUNCTIONS", "K_AGGREGATE",
		"K_AGGREGATES", "K_SFUNC", "K_STYPE", "K_FINALFUNC", "K_INITCOND", "K_RETURNS",
		"K_CALLED", "K_INPUT", "K_LANGUAGE", "K_OR", "K_REPLACE", "K_JSON",
		"K_DEFAULT", "K_UNSET", "K_LIKE", "K_MASKED", "K_UNMASK", "K_SELECT_MASKED",
		"K_VECTOR", "K_ANN", "STRING_LITERAL", "QUOTED_NAME", "EMPTY_QUOTED_NAME",
		"INTEGER", "QMARK", "RANGE", "FLOAT", "BOOLEAN", "DURATION", "IDENT",
		"HEXNUMBER", "UUID", "WS", "COMMENT", "MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"cqlStatements", "cqlStatement", "useStatement", "marker", "createKeyspaceStatement",
		"createTableStatement", "ifNotExists", "ifExists", "tableDefinition",
		"tableColumns", "columnMask", "columnMaskArguments", "tablePartitionKey",
		"tableProperty", "tableClusteringOrder", "createTypeStatement", "typeColumns",
		"dropTableStatement", "truncateStatement", "ident", "fident", "noncol_ident",
		"keyspaceName", "columnFamilyName", "userTypeName", "ksName", "cfName",
		"constant", "fullMapLiteral", "setOrMapLiteral", "setLiteral", "mapLiteral",
		"collectionLiteral", "listLiteral", "usertypeLiteral", "tupleLiteral",
		"value", "functionName", "allowedFunctionName", "function", "functionArgs",
		"term", "termAddition", "termMultiplication", "termGroup", "simpleTerm",
		"properties", "property", "propertyValue", "comparatorType", "native_type",
		"collection_type", "tuple_type", "vector_type", "non_type_ident", "unreserved_keyword",
		"unreserved_function_keyword", "basic_unreserved_keyword",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 190, 644, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		1, 0, 1, 0, 1, 0, 5, 0, 120, 8, 0, 10, 0, 12, 0, 123, 9, 0, 1, 0, 5, 0,
		126, 8, 0, 10, 0, 12, 0, 129, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 139, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3,
		147, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 152, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 3, 5, 161, 8, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 177, 8, 8, 5, 8, 179,
		8, 8, 10, 8, 12, 8, 182, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 189,
		8, 8, 10, 8, 12, 8, 192, 9, 8, 3, 8, 194, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9,
		199, 8, 9, 1, 9, 3, 9, 202, 8, 9, 1, 9, 1, 9, 3, 9, 206, 8, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 214, 8, 9, 10, 9, 12, 9, 217, 9, 9, 1,
		9, 1, 9, 3, 9, 221, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 231, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 239, 8, 11, 10, 11, 12, 11, 242, 9, 11, 1, 11, 1, 11, 3, 11, 246,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 253, 8, 12, 10, 12, 12,
		12, 256, 9, 12, 1, 12, 1, 12, 3, 12, 260, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 272, 8, 13, 10, 13,
		12, 13, 275, 9, 13, 1, 13, 1, 13, 3, 13, 279, 8, 13, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 3, 15, 287, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 294, 8, 15, 5, 15, 296, 8, 15, 10, 15, 12, 15, 299, 9, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 309, 8, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 3, 18, 315, 8, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 3, 19, 322, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 327, 8, 20, 1,
		21, 1, 21, 1, 21, 3, 21, 332, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		3, 23, 339, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 346, 8, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 354, 8, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 3, 26, 360, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 372, 8, 27, 3, 27, 374, 8, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 385,
		8, 28, 10, 28, 12, 28, 388, 9, 28, 3, 28, 390, 8, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 3, 29, 396, 8, 29, 1, 30, 1, 30, 5, 30, 400, 8, 30, 10, 30,
		12, 30, 403, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5,
		31, 412, 8, 31, 10, 31, 12, 31, 415, 9, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 425, 8, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 5, 33, 431, 8, 33, 10, 33, 12, 33, 434, 9, 33, 3, 33, 436, 8, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		5, 34, 449, 8, 34, 10, 34, 12, 34, 452, 9, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 35, 5, 35, 460, 8, 35, 10, 35, 12, 35, 463, 9, 35, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 473, 8, 36, 1,
		37, 1, 37, 1, 37, 3, 37, 478, 8, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 3, 38, 487, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 3, 39, 498, 8, 39, 1, 40, 1, 40, 1, 40, 5, 40,
		503, 8, 40, 10, 40, 12, 40, 506, 9, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 5, 42, 515, 8, 42, 10, 42, 12, 42, 518, 9, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 527, 8, 43, 10, 43, 12,
		43, 530, 9, 43, 1, 44, 1, 44, 1, 44, 3, 44, 535, 8, 44, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 3, 45, 551, 8, 45, 1, 46, 1, 46, 1, 46, 5, 46, 556, 8, 46, 10, 46,
		12, 46, 559, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 3, 47, 569, 8, 47, 1, 48, 1, 48, 3, 48, 573, 8, 48, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 586,
		8, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51,
		607, 8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 614, 8, 52, 10, 52,
		12, 52, 617, 9, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 632, 8, 54, 1, 55, 1, 55,
		3, 55, 636, 8, 55, 1, 56, 1, 56, 3, 56, 640, 8, 56, 1, 57, 1, 57, 1, 57,
		0, 0, 58, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 110, 112, 114, 0, 6, 1, 0, 74, 75, 1, 0, 145, 146, 3, 0, 62,
		62, 117, 134, 138, 139, 2, 0, 138, 138, 185, 185, 5, 0, 24, 24, 36, 37,
		63, 64, 136, 137, 167, 167, 16, 0, 21, 21, 25, 25, 32, 33, 49, 49, 51,
		51, 55, 55, 61, 61, 68, 71, 77, 77, 80, 85, 87, 89, 97, 116, 142, 144,
		149, 164, 166, 166, 168, 175, 689, 0, 116, 1, 0, 0, 0, 2, 138, 1, 0, 0,
		0, 4, 140, 1, 0, 0, 0, 6, 146, 1, 0, 0, 0, 8, 148, 1, 0, 0, 0, 10, 157,
		1, 0, 0, 0, 12, 165, 1, 0, 0, 0, 14, 169, 1, 0, 0, 0, 16, 172, 1, 0, 0,
		0, 18, 220, 1, 0, 0, 0, 20, 230, 1, 0, 0, 0, 22, 245, 1, 0, 0, 0, 24, 259,
		1, 0, 0, 0, 26, 278, 1, 0, 0, 0, 28, 280, 1, 0, 0, 0, 30, 283, 1, 0, 0,
		0, 32, 302, 1, 0, 0, 0, 34, 305, 1, 0, 0, 0, 36, 312, 1, 0, 0, 0, 38, 321,
		1, 0, 0, 0, 40, 326, 1, 0, 0, 0, 42, 331, 1, 0, 0, 0, 44, 333, 1, 0, 0,
		0, 46, 338, 1, 0, 0, 0, 48, 345, 1, 0, 0, 0, 50, 353, 1, 0, 0, 0, 52, 359,
		1, 0, 0, 0, 54, 373, 1, 0, 0, 0, 56, 375, 1, 0, 0, 0, 58, 395, 1, 0, 0,
		0, 60, 401, 1, 0, 0, 0, 62, 404, 1, 0, 0, 0, 64, 424, 1, 0, 0, 0, 66, 426,
		1, 0, 0, 0, 68, 439, 1, 0, 0, 0, 70, 455, 1, 0, 0, 0, 72, 472, 1, 0, 0,
		0, 74, 477, 1, 0, 0, 0, 76, 486, 1, 0, 0, 0, 78, 497, 1, 0, 0, 0, 80, 499,
		1, 0, 0, 0, 82, 507, 1, 0, 0, 0, 84, 509, 1, 0, 0, 0, 86, 519, 1, 0, 0,
		0, 88, 534, 1, 0, 0, 0, 90, 550, 1, 0, 0, 0, 92, 552, 1, 0, 0, 0, 94, 568,
		1, 0, 0, 0, 96, 572, 1, 0, 0, 0, 98, 585, 1, 0, 0, 0, 100, 587, 1, 0, 0,
		0, 102, 606, 1, 0, 0, 0, 104, 608, 1, 0, 0, 0, 106, 620, 1, 0, 0, 0, 108,
		631, 1, 0, 0, 0, 110, 635, 1, 0, 0, 0, 112, 639, 1, 0, 0, 0, 114, 641,
		1, 0, 0, 0, 116, 121, 3, 2, 1, 0, 117, 118, 5, 1, 0, 0, 118, 120, 3, 2,
		1, 0, 119, 117, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 127, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124,
		126, 5, 1, 0, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125,
		1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130, 1, 0, 0, 0, 129, 127, 1, 0,
		0, 0, 130, 131, 5, 0, 0, 1, 131, 1, 1, 0, 0, 0, 132, 139, 3, 4, 2, 0, 133,
		139, 3, 36, 18, 0, 134, 139, 3, 8, 4, 0, 135, 139, 3, 10, 5, 0, 136, 139,
		3, 34, 17, 0, 137, 139, 3, 30, 15, 0, 138, 132, 1, 0, 0, 0, 138, 133, 1,
		0, 0, 0, 138, 134, 1, 0, 0, 0, 138, 135, 1, 0, 0, 0, 138, 136, 1, 0, 0,
		0, 138, 137, 1, 0, 0, 0, 139, 3, 1, 0, 0, 0, 140, 141, 5, 35, 0, 0, 141,
		142, 3, 44, 22, 0, 142, 5, 1, 0, 0, 0, 143, 144, 5, 2, 0, 0, 144, 147,
		3, 42, 21, 0, 145, 147, 5, 180, 0, 0, 146, 143, 1, 0, 0, 0, 146, 145, 1,
		0, 0, 0, 147, 7, 1, 0, 0, 0, 148, 149, 5, 46, 0, 0, 149, 151, 5, 48, 0,
		0, 150, 152, 3, 12, 6, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 154, 3, 44, 22, 0, 154, 155, 5, 30, 0, 0, 155, 156,
		3, 92, 46, 0, 156, 9, 1, 0, 0, 0, 157, 158, 5, 46, 0, 0, 158, 160, 5, 50,
		0, 0, 159, 161, 3, 12, 6, 0, 160, 159, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 163, 3, 46, 23, 0, 163, 164, 3, 16, 8, 0, 164,
		11, 1, 0, 0, 0, 165, 166, 5, 78, 0, 0, 166, 167, 5, 141, 0, 0, 167, 168,
		5, 142, 0, 0, 168, 13, 1, 0, 0, 0, 169, 170, 5, 78, 0, 0, 170, 171, 5,
		142, 0, 0, 171, 15, 1, 0, 0, 0, 172, 173, 5, 3, 0, 0, 173, 180, 3, 18,
		9, 0, 174, 176, 5, 4, 0, 0, 175, 177, 3, 18, 9, 0, 176, 175, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 174, 1, 0, 0, 0, 179,
		182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 193, 5, 5, 0, 0, 184, 185, 5, 30,
		0, 0, 185, 190, 3, 26, 13, 0, 186, 187, 5, 23, 0, 0, 187, 189, 3, 26, 13,
		0, 188, 186, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 184,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 17, 1, 0, 0, 0, 195, 196, 3, 38,
		19, 0, 196, 198, 3, 98, 49, 0, 197, 199, 5, 151, 0, 0, 198, 197, 1, 0,
		0, 0, 198, 199, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200, 202, 3, 20, 10,
		0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203,
		204, 5, 59, 0, 0, 204, 206, 5, 24, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206,
		1, 0, 0, 0, 206, 221, 1, 0, 0, 0, 207, 208, 5, 59, 0, 0, 208, 209, 5, 24,
		0, 0, 209, 210, 5, 3, 0, 0, 210, 215, 3, 24, 12, 0, 211, 212, 5, 4, 0,
		0, 212, 214, 3, 38, 19, 0, 213, 211, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0,
		215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217,
		215, 1, 0, 0, 0, 218, 219, 5, 5, 0, 0, 219, 221, 1, 0, 0, 0, 220, 195,
		1, 0, 0, 0, 220, 207, 1, 0, 0, 0, 221, 19, 1, 0, 0, 0, 222, 223, 5, 171,
		0, 0, 223, 224, 5, 30, 0, 0, 224, 225, 3, 74, 37, 0, 225, 226, 3, 22, 11,
		0, 226, 231, 1, 0, 0, 0, 227, 228, 5, 171, 0, 0, 228, 229, 5, 30, 0, 0,
		229, 231, 5, 168, 0, 0, 230, 222, 1, 0, 0, 0, 230, 227, 1, 0, 0, 0, 231,
		21, 1, 0, 0, 0, 232, 233, 5, 3, 0, 0, 233, 246, 5, 5, 0, 0, 234, 235, 5,
		3, 0, 0, 235, 240, 3, 82, 41, 0, 236, 237, 5, 4, 0, 0, 237, 239, 3, 82,
		41, 0, 238, 236, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0,
		240, 241, 1, 0, 0, 0, 241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243,
		244, 5, 5, 0, 0, 244, 246, 1, 0, 0, 0, 245, 232, 1, 0, 0, 0, 245, 234,
		1, 0, 0, 0, 246, 23, 1, 0, 0, 0, 247, 260, 3, 38, 19, 0, 248, 249, 5, 3,
		0, 0, 249, 254, 3, 38, 19, 0, 250, 251, 5, 4, 0, 0, 251, 253, 3, 38, 19,
		0, 252, 250, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254,
		255, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 258,
		5, 5, 0, 0, 258, 260, 1, 0, 0, 0, 259, 247, 1, 0, 0, 0, 259, 248, 1, 0,
		0, 0, 260, 25, 1, 0, 0, 0, 261, 279, 3, 94, 47, 0, 262, 263, 5, 70, 0,
		0, 263, 279, 5, 71, 0, 0, 264, 265, 5, 116, 0, 0, 265, 266, 5, 72, 0, 0,
		266, 267, 5, 73, 0, 0, 267, 268, 5, 3, 0, 0, 268, 273, 3, 28, 14, 0, 269,
		270, 5, 4, 0, 0, 270, 272, 3, 28, 14, 0, 271, 269, 1, 0, 0, 0, 272, 275,
		1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0,
		0, 0, 275, 273, 1, 0, 0, 0, 276, 277, 5, 5, 0, 0, 277, 279, 1, 0, 0, 0,
		278, 261, 1, 0, 0, 0, 278, 262, 1, 0, 0, 0, 278, 264, 1, 0, 0, 0, 279,
		27, 1, 0, 0, 0, 280, 281, 3, 38, 19, 0, 281, 282, 7, 0, 0, 0, 282, 29,
		1, 0, 0, 0, 283, 284, 5, 46, 0, 0, 284, 286, 5, 68, 0, 0, 285, 287, 3,
		12, 6, 0, 286, 285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 1, 0, 0,
		0, 288, 289, 3, 48, 24, 0, 289, 290, 5, 3, 0, 0, 290, 297, 3, 32, 16, 0,
		291, 293, 5, 4, 0, 0, 292, 294, 3, 32, 16, 0, 293, 292, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295, 291, 1, 0, 0, 0, 296, 299,
		1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0,
		0, 0, 299, 297, 1, 0, 0, 0, 300, 301, 5, 5, 0, 0, 301, 31, 1, 0, 0, 0,
		302, 303, 3, 40, 20, 0, 303, 304, 3, 98, 49, 0, 304, 33, 1, 0, 0, 0, 305,
		306, 5, 58, 0, 0, 306, 308, 5, 50, 0, 0, 307, 309, 3, 14, 7, 0, 308, 307,
		1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 311, 3, 46,
		23, 0, 311, 35, 1, 0, 0, 0, 312, 314, 5, 43, 0, 0, 313, 315, 5, 50, 0,
		0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316,
		317, 3, 46, 23, 0, 317, 37, 1, 0, 0, 0, 318, 322, 5, 185, 0, 0, 319, 322,
		5, 177, 0, 0, 320, 322, 3, 110, 55, 0, 321, 318, 1, 0, 0, 0, 321, 319,
		1, 0, 0, 0, 321, 320, 1, 0, 0, 0, 322, 39, 1, 0, 0, 0, 323, 327, 5, 185,
		0, 0, 324, 327, 5, 177, 0, 0, 325, 327, 3, 110, 55, 0, 326, 323, 1, 0,
		0, 0, 326, 324, 1, 0, 0, 0, 326, 325, 1, 0, 0, 0, 327, 41, 1, 0, 0, 0,
		328, 332, 5, 185, 0, 0, 329, 332, 5, 177, 0, 0, 330, 332, 3, 110, 55, 0,
		331, 328, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332,
		43, 1, 0, 0, 0, 333, 334, 3, 50, 25, 0, 334, 45, 1, 0, 0, 0, 335, 336,
		3, 50, 25, 0, 336, 337, 5, 6, 0, 0, 337, 339, 1, 0, 0, 0, 338, 335, 1,
		0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341, 3, 52, 26,
		0, 341, 47, 1, 0, 0, 0, 342, 343, 3, 42, 21, 0, 343, 344, 5, 6, 0, 0, 344,
		346, 1, 0, 0, 0, 345, 342, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347,
		1, 0, 0, 0, 347, 348, 3, 108, 54, 0, 348, 49, 1, 0, 0, 0, 349, 354, 5,
		185, 0, 0, 350, 354, 5, 177, 0, 0, 351, 354, 3, 110, 55, 0, 352, 354, 5,
		180, 0, 0, 353, 349, 1, 0, 0, 0, 353, 350, 1, 0, 0, 0, 353, 351, 1, 0,
		0, 0, 353, 352, 1, 0, 0, 0, 354, 51, 1, 0, 0, 0, 355, 360, 5, 185, 0, 0,
		356, 360, 5, 177, 0, 0, 357, 360, 3, 110, 55, 0, 358, 360, 5, 180, 0, 0,
		359, 355, 1, 0, 0, 0, 359, 356, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359,
		358, 1, 0, 0, 0, 360, 53, 1, 0, 0, 0, 361, 374, 5, 176, 0, 0, 362, 374,
		5, 179, 0, 0, 363, 374, 5, 182, 0, 0, 364, 374, 5, 183, 0, 0, 365, 374,
		5, 184, 0, 0, 366, 374, 5, 187, 0, 0, 367, 374, 5, 186, 0, 0, 368, 372,
		7, 1, 0, 0, 369, 372, 5, 147, 0, 0, 370, 372, 5, 148, 0, 0, 371, 368, 1,
		0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 370, 1, 0, 0, 0, 372, 374, 1, 0, 0,
		0, 373, 361, 1, 0, 0, 0, 373, 362, 1, 0, 0, 0, 373, 363, 1, 0, 0, 0, 373,
		364, 1, 0, 0, 0, 373, 365, 1, 0, 0, 0, 373, 366, 1, 0, 0, 0, 373, 367,
		1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 55, 1, 0, 0, 0, 375, 389, 5, 7,
		0, 0, 376, 377, 3, 82, 41, 0, 377, 378, 5, 2, 0, 0, 378, 386, 3, 82, 41,
		0, 379, 380, 5, 4, 0, 0, 380, 381, 3, 82, 41, 0, 381, 382, 5, 2, 0, 0,
		382, 383, 3, 82, 41, 0, 383, 385, 1, 0, 0, 0, 384, 379, 1, 0, 0, 0, 385,
		388, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 390,
		1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 389, 376, 1, 0, 0, 0, 389, 390, 1, 0,
		0, 0, 390, 391, 1, 0, 0, 0, 391, 392, 5, 8, 0, 0, 392, 57, 1, 0, 0, 0,
		393, 396, 3, 62, 31, 0, 394, 396, 3, 60, 30, 0, 395, 393, 1, 0, 0, 0, 395,
		394, 1, 0, 0, 0, 396, 59, 1, 0, 0, 0, 397, 398, 5, 4, 0, 0, 398, 400, 3,
		82, 41, 0, 399, 397, 1, 0, 0, 0, 400, 403, 1, 0, 0, 0, 401, 399, 1, 0,
		0, 0, 401, 402, 1, 0, 0, 0, 402, 61, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0,
		404, 405, 5, 2, 0, 0, 405, 413, 3, 82, 41, 0, 406, 407, 5, 4, 0, 0, 407,
		408, 3, 82, 41, 0, 408, 409, 5, 2, 0, 0, 409, 410, 3, 82, 41, 0, 410, 412,
		1, 0, 0, 0, 411, 406, 1, 0, 0, 0, 412, 415, 1, 0, 0, 0, 413, 411, 1, 0,
		0, 0, 413, 414, 1, 0, 0, 0, 414, 63, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0,
		416, 425, 3, 66, 33, 0, 417, 418, 5, 7, 0, 0, 418, 419, 3, 82, 41, 0, 419,
		420, 3, 58, 29, 0, 420, 421, 5, 8, 0, 0, 421, 425, 1, 0, 0, 0, 422, 423,
		5, 7, 0, 0, 423, 425, 5, 8, 0, 0, 424, 416, 1, 0, 0, 0, 424, 417, 1, 0,
		0, 0, 424, 422, 1, 0, 0, 0, 425, 65, 1, 0, 0, 0, 426, 435, 5, 9, 0, 0,
		427, 432, 3, 82, 41, 0, 428, 429, 5, 4, 0, 0, 429, 431, 3, 82, 41, 0, 430,
		428, 1, 0, 0, 0, 431, 434, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 432, 433,
		1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 435, 427, 1, 0,
		0, 0, 435, 436, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 438, 5, 10, 0, 0,
		438, 67, 1, 0, 0, 0, 439, 440, 5, 7, 0, 0, 440, 441, 3, 40, 20, 0, 441,
		442, 5, 2, 0, 0, 442, 450, 3, 82, 41, 0, 443, 444, 5, 4, 0, 0, 444, 445,
		3, 40, 20, 0, 445, 446, 5, 2, 0, 0, 446, 447, 3, 82, 41, 0, 447, 449, 1,
		0, 0, 0, 448, 443, 1, 0, 0, 0, 449, 452, 1, 0, 0, 0, 450, 448, 1, 0, 0,
		0, 450, 451, 1, 0, 0, 0, 451, 453, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 453,
		454, 5, 8, 0, 0, 454, 69, 1, 0, 0, 0, 455, 456, 5, 3, 0, 0, 456, 461, 3,
		82, 41, 0, 457, 458, 5, 4, 0, 0, 458, 460, 3, 82, 41, 0, 459, 457, 1, 0,
		0, 0, 460, 463, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0,
		462, 464, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 464, 465, 5, 5, 0, 0, 465,
		71, 1, 0, 0, 0, 466, 473, 3, 54, 27, 0, 467, 473, 3, 64, 32, 0, 468, 473,
		3, 68, 34, 0, 469, 473, 3, 70, 35, 0, 470, 473, 5, 140, 0, 0, 471, 473,
		3, 6, 3, 0, 472, 466, 1, 0, 0, 0, 472, 467, 1, 0, 0, 0, 472, 468, 1, 0,
		0, 0, 472, 469, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 471, 1, 0, 0, 0,
		473, 73, 1, 0, 0, 0, 474, 475, 3, 44, 22, 0, 475, 476, 5, 6, 0, 0, 476,
		478, 1, 0, 0, 0, 477, 474, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 479,
		1, 0, 0, 0, 479, 480, 3, 76, 38, 0, 480, 75, 1, 0, 0, 0, 481, 487, 5, 185,
		0, 0, 482, 487, 5, 177, 0, 0, 483, 487, 3, 112, 56, 0, 484, 487, 5, 135,
		0, 0, 485, 487, 5, 37, 0, 0, 486, 481, 1, 0, 0, 0, 486, 482, 1, 0, 0, 0,
		486, 483, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 486, 485, 1, 0, 0, 0, 487,
		77, 1, 0, 0, 0, 488, 489, 3, 74, 37, 0, 489, 490, 5, 3, 0, 0, 490, 491,
		5, 5, 0, 0, 491, 498, 1, 0, 0, 0, 492, 493, 3, 74, 37, 0, 493, 494, 5,
		3, 0, 0, 494, 495, 3, 80, 40, 0, 495, 496, 5, 5, 0, 0, 496, 498, 1, 0,
		0, 0, 497, 488, 1, 0, 0, 0, 497, 492, 1, 0, 0, 0, 498, 79, 1, 0, 0, 0,
		499, 504, 3, 82, 41, 0, 500, 501, 5, 4, 0, 0, 501, 503, 3, 82, 41, 0, 502,
		500, 1, 0, 0, 0, 503, 506, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 504, 505,
		1, 0, 0, 0, 505, 81, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 507, 508, 3, 84,
		42, 0, 508, 83, 1, 0, 0, 0, 509, 516, 3, 86, 43, 0, 510, 511, 5, 11, 0,
		0, 511, 515, 3, 86, 43, 0, 512, 513, 5, 12, 0, 0, 513, 515, 3, 86, 43,
		0, 514, 510, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 515, 518, 1, 0, 0, 0, 516,
		514, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 85, 1, 0, 0, 0, 518, 516, 1,
		0, 0, 0, 519, 528, 3, 88, 44, 0, 520, 521, 5, 13, 0, 0, 521, 527, 3, 88,
		44, 0, 522, 523, 5, 14, 0, 0, 523, 527, 3, 88, 44, 0, 524, 525, 5, 15,
		0, 0, 525, 527, 3, 88, 44, 0, 526, 520, 1, 0, 0, 0, 526, 522, 1, 0, 0,
		0, 526, 524, 1, 0, 0, 0, 527, 530, 1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 528,
		529, 1, 0, 0, 0, 529, 87, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 531, 535, 3,
		90, 45, 0, 532, 533, 5, 12, 0, 0, 533, 535, 3, 90, 45, 0, 534, 531, 1,
		0, 0, 0, 534, 532, 1, 0, 0, 0, 535, 89, 1, 0, 0, 0, 536, 551, 3, 72, 36,
		0, 537, 551, 3, 78, 39, 0, 538, 539, 5, 3, 0, 0, 539, 540, 3, 98, 49, 0,
		540, 541, 5, 5, 0, 0, 541, 542, 3, 90, 45, 0, 542, 551, 1, 0, 0, 0, 543,
		544, 5, 64, 0, 0, 544, 545, 5, 3, 0, 0, 545, 546, 3, 90, 45, 0, 546, 547,
		5, 21, 0, 0, 547, 548, 3, 100, 50, 0, 548, 549, 5, 5, 0, 0, 549, 551, 1,
		0, 0, 0, 550, 536, 1, 0, 0, 0, 550, 537, 1, 0, 0, 0, 550, 538, 1, 0, 0,
		0, 550, 543, 1, 0, 0, 0, 551, 91, 1, 0, 0, 0, 552, 557, 3, 94, 47, 0, 553,
		554, 5, 23, 0, 0, 554, 556, 3, 94, 47, 0, 555, 553, 1, 0, 0, 0, 556, 559,
		1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 93, 1, 0,
		0, 0, 559, 557, 1, 0, 0, 0, 560, 561, 3, 42, 21, 0, 561, 562, 5, 16, 0,
		0, 562, 563, 3, 96, 48, 0, 563, 569, 1, 0, 0, 0, 564, 565, 3, 42, 21, 0,
		565, 566, 5, 16, 0, 0, 566, 567, 3, 56, 28, 0, 567, 569, 1, 0, 0, 0, 568,
		560, 1, 0, 0, 0, 568, 564, 1, 0, 0, 0, 569, 95, 1, 0, 0, 0, 570, 573, 3,
		54, 27, 0, 571, 573, 3, 110, 55, 0, 572, 570, 1, 0, 0, 0, 572, 571, 1,
		0, 0, 0, 573, 97, 1, 0, 0, 0, 574, 586, 3, 100, 50, 0, 575, 586, 3, 102,
		51, 0, 576, 586, 3, 104, 52, 0, 577, 586, 3, 106, 53, 0, 578, 586, 3, 48,
		24, 0, 579, 580, 5, 152, 0, 0, 580, 581, 5, 17, 0, 0, 581, 582, 3, 98,
		49, 0, 582, 583, 5, 18, 0, 0, 583, 586, 1, 0, 0, 0, 584, 586, 5, 176, 0,
		0, 585, 574, 1, 0, 0, 0, 585, 575, 1, 0, 0, 0, 585, 576, 1, 0, 0, 0, 585,
		577, 1, 0, 0, 0, 585, 578, 1, 0, 0, 0, 585, 579, 1, 0, 0, 0, 585, 584,
		1, 0, 0, 0, 586, 99, 1, 0, 0, 0, 587, 588, 7, 2, 0, 0, 588, 101, 1, 0,
		0, 0, 589, 590, 5, 143, 0, 0, 590, 591, 5, 17, 0, 0, 591, 592, 3, 98, 49,
		0, 592, 593, 5, 4, 0, 0, 593, 594, 3, 98, 49, 0, 594, 595, 5, 18, 0, 0,
		595, 607, 1, 0, 0, 0, 596, 597, 5, 144, 0, 0, 597, 598, 5, 17, 0, 0, 598,
		599, 3, 98, 49, 0, 599, 600, 5, 18, 0, 0, 600, 607, 1, 0, 0, 0, 601, 602,
		5, 38, 0, 0, 602, 603, 5, 17, 0, 0, 603, 604, 3, 98, 49, 0, 604, 605, 5,
		18, 0, 0, 605, 607, 1, 0, 0, 0, 606, 589, 1, 0, 0, 0, 606, 596, 1, 0, 0,
		0, 606, 601, 1, 0, 0, 0, 607, 103, 1, 0, 0, 0, 608, 609, 5, 149, 0, 0,
		609, 610, 5, 17, 0, 0, 610, 615, 3, 98, 49, 0, 611, 612, 5, 4, 0, 0, 612,
		614, 3, 98, 49, 0, 613, 611, 1, 0, 0, 0, 614, 617, 1, 0, 0, 0, 615, 613,
		1, 0, 0, 0, 615, 616, 1, 0, 0, 0, 616, 618, 1, 0, 0, 0, 617, 615, 1, 0,
		0, 0, 618, 619, 5, 18, 0, 0, 619, 105, 1, 0, 0, 0, 620, 621, 5, 174, 0,
		0, 621, 622, 5, 17, 0, 0, 622, 623, 3, 98, 49, 0, 623, 624, 5, 4, 0, 0,
		624, 625, 5, 179, 0, 0, 625, 626, 5, 18, 0, 0, 626, 107, 1, 0, 0, 0, 627,
		632, 7, 3, 0, 0, 628, 632, 5, 177, 0, 0, 629, 632, 3, 114, 57, 0, 630,
		632, 5, 24, 0, 0, 631, 627, 1, 0, 0, 0, 631, 628, 1, 0, 0, 0, 631, 629,
		1, 0, 0, 0, 631, 630, 1, 0, 0, 0, 632, 109, 1, 0, 0, 0, 633, 636, 3, 112,
		56, 0, 634, 636, 7, 4, 0, 0, 635, 633, 1, 0, 0, 0, 635, 634, 1, 0, 0, 0,
		636, 111, 1, 0, 0, 0, 637, 640, 3, 114, 57, 0, 638, 640, 3, 100, 50, 0,
		639, 637, 1, 0, 0, 0, 639, 638, 1, 0, 0, 0, 640, 113, 1, 0, 0, 0, 641,
		642, 7, 5, 0, 0, 642, 115, 1, 0, 0, 0, 66, 121, 127, 138, 146, 151, 160,
		176, 180, 190, 193, 198, 201, 205, 215, 220, 230, 240, 245, 254, 259, 273,
		278, 286, 293, 297, 308, 314, 321, 326, 331, 338, 345, 353, 359, 371, 373,
		386, 389, 395, 401, 413, 424, 432, 435, 450, 461, 472, 477, 486, 497, 504,
		514, 516, 526, 528, 534, 550, 557, 568, 572, 585, 606, 615, 631, 635, 639,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// Cql3ParserInit initializes any static state used to implement Cql3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCql3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Cql3ParserInit() {
	staticData := &Cql3ParserStaticData
	staticData.once.Do(cql3ParserInit)
}

// NewCql3Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewCql3Parser(input antlr.TokenStream) *Cql3Parser {
	Cql3ParserInit()
	this := new(Cql3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Cql3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Cql3.g4"

	return this
}

// Cql3Parser tokens.
const (
	Cql3ParserEOF                 = antlr.TokenEOF
	Cql3ParserT__0                = 1
	Cql3ParserT__1                = 2
	Cql3ParserT__2                = 3
	Cql3ParserT__3                = 4
	Cql3ParserT__4                = 5
	Cql3ParserT__5                = 6
	Cql3ParserT__6                = 7
	Cql3ParserT__7                = 8
	Cql3ParserT__8                = 9
	Cql3ParserT__9                = 10
	Cql3ParserT__10               = 11
	Cql3ParserT__11               = 12
	Cql3ParserT__12               = 13
	Cql3ParserT__13               = 14
	Cql3ParserT__14               = 15
	Cql3ParserT__15               = 16
	Cql3ParserT__16               = 17
	Cql3ParserT__17               = 18
	Cql3ParserK_SELECT            = 19
	Cql3ParserK_FROM              = 20
	Cql3ParserK_AS                = 21
	Cql3ParserK_WHERE             = 22
	Cql3ParserK_AND               = 23
	Cql3ParserK_KEY               = 24
	Cql3ParserK_KEYS              = 25
	Cql3ParserK_ENTRIES           = 26
	Cql3ParserK_FULL              = 27
	Cql3ParserK_INSERT            = 28
	Cql3ParserK_UPDATE            = 29
	Cql3ParserK_WITH              = 30
	Cql3ParserK_LIMIT             = 31
	Cql3ParserK_PER               = 32
	Cql3ParserK_PARTITION         = 33
	Cql3ParserK_USING             = 34
	Cql3ParserK_USE               = 35
	Cql3ParserK_DISTINCT          = 36
	Cql3ParserK_COUNT             = 37
	Cql3ParserK_SET               = 38
	Cql3ParserK_BEGIN             = 39
	Cql3ParserK_UNLOGGED          = 40
	Cql3ParserK_BATCH             = 41
	Cql3ParserK_APPLY             = 42
	Cql3ParserK_TRUNCATE          = 43
	Cql3ParserK_DELETE            = 44
	Cql3ParserK_IN                = 45
	Cql3ParserK_CREATE            = 46
	Cql3ParserK_SCHEMA            = 47
	Cql3ParserK_KEYSPACE          = 48
	Cql3ParserK_KEYSPACES         = 49
	Cql3ParserK_COLUMNFAMILY      = 50
	Cql3ParserK_TABLES            = 51
	Cql3ParserK_MATERIALIZED      = 52
	Cql3ParserK_VIEW              = 53
	Cql3ParserK_INDEX             = 54
	Cql3ParserK_CUSTOM            = 55
	Cql3ParserK_ON                = 56
	Cql3ParserK_TO                = 57
	Cql3ParserK_DROP              = 58
	Cql3ParserK_PRIMARY           = 59
	Cql3ParserK_INTO              = 60
	Cql3ParserK_VALUES            = 61
	Cql3ParserK_TIMESTAMP         = 62
	Cql3ParserK_TTL               = 63
	Cql3ParserK_CAST              = 64
	Cql3ParserK_ALTER             = 65
	Cql3ParserK_RENAME            = 66
	Cql3ParserK_ADD               = 67
	Cql3ParserK_TYPE              = 68
	Cql3ParserK_TYPES             = 69
	Cql3ParserK_COMPACT           = 70
	Cql3ParserK_STORAGE           = 71
	Cql3ParserK_ORDER             = 72
	Cql3ParserK_BY                = 73
	Cql3ParserK_ASC               = 74
	Cql3ParserK_DESC              = 75
	Cql3ParserK_ALLOW             = 76
	Cql3ParserK_FILTERING         = 77
	Cql3ParserK_IF                = 78
	Cql3ParserK_IS                = 79
	Cql3ParserK_CONTAINS          = 80
	Cql3ParserK_BETWEEN           = 81
	Cql3ParserK_GROUP             = 82
	Cql3ParserK_CLUSTER           = 83
	Cql3ParserK_INTERNALS         = 84
	Cql3ParserK_ONLY              = 85
	Cql3ParserK_GRANT             = 86
	Cql3ParserK_ALL               = 87
	Cql3ParserK_PERMISSION        = 88
	Cql3ParserK_PERMISSIONS       = 89
	Cql3ParserK_OF                = 90
	Cql3ParserK_REVOKE            = 91
	Cql3ParserK_MODIFY            = 92
	Cql3ParserK_AUTHORIZE         = 93
	Cql3ParserK_DESCRIBE          = 94
	Cql3ParserK_EXECUTE           = 95
	Cql3ParserK_NORECURSIVE       = 96
	Cql3ParserK_MBEAN             = 97
	Cql3ParserK_MBEANS            = 98
	Cql3ParserK_USER              = 99
	Cql3ParserK_USERS             = 100
	Cql3ParserK_ROLE              = 101
	Cql3ParserK_ROLES             = 102
	Cql3ParserK_SUPERUSERS        = 103
	Cql3ParserK_SUPERUSER         = 104
	Cql3ParserK_NOSUPERUSER       = 105
	Cql3ParserK_GENERATED         = 106
	Cql3ParserK_PASSWORD          = 107
	Cql3ParserK_HASHED            = 108
	Cql3ParserK_LOGIN             = 109
	Cql3ParserK_NOLOGIN           = 110
	Cql3ParserK_OPTIONS           = 111
	Cql3ParserK_ACCESS            = 112
	Cql3ParserK_DATACENTERS       = 113
	Cql3ParserK_CIDRS             = 114
	Cql3ParserK_IDENTITY          = 115
	Cql3ParserK_CLUSTERING        = 116
	Cql3ParserK_ASCII             = 117
	Cql3ParserK_BIGINT            = 118
	Cql3ParserK_BLOB              = 119
	Cql3ParserK_BOOLEAN           = 120
	Cql3ParserK_COUNTER           = 121
	Cql3ParserK_DECIMAL           = 122
	Cql3ParserK_DOUBLE            = 123
	Cql3ParserK_DURATION          = 124
	Cql3ParserK_FLOAT             = 125
	Cql3ParserK_INET              = 126
	Cql3ParserK_INT               = 127
	Cql3ParserK_SMALLINT          = 128
	Cql3ParserK_TINYINT           = 129
	Cql3ParserK_TEXT              = 130
	Cql3ParserK_UUID              = 131
	Cql3ParserK_VARCHAR           = 132
	Cql3ParserK_VARINT            = 133
	Cql3ParserK_TIMEUUID          = 134
	Cql3ParserK_TOKEN             = 135
	Cql3ParserK_WRITETIME         = 136
	Cql3ParserK_MAXWRITETIME      = 137
	Cql3ParserK_DATE              = 138
	Cql3ParserK_TIME              = 139
	Cql3ParserK_NULL              = 140
	Cql3ParserK_NOT               = 141
	Cql3ParserK_EXISTS            = 142
	Cql3ParserK_MAP               = 143
	Cql3ParserK_LIST              = 144
	Cql3ParserK_POSITIVE_NAN      = 145
	Cql3ParserK_NEGATIVE_NAN      = 146
	Cql3ParserK_POSITIVE_INFINITY = 147
	Cql3ParserK_NEGATIVE_INFINITY = 148
	Cql3ParserK_TUPLE             = 149
	Cql3ParserK_TRIGGER           = 150
	Cql3ParserK_STATIC            = 151
	Cql3ParserK_FROZEN            = 152
	Cql3ParserK_FUNCTION          = 153
	Cql3ParserK_FUNCTIONS         = 154
	Cql3ParserK_AGGREGATE         = 155
	Cql3ParserK_AGGREGATES        = 156
	Cql3ParserK_SFUNC             = 157
	Cql3ParserK_STYPE             = 158
	Cql3ParserK_FINALFUNC         = 159
	Cql3ParserK_INITCOND          = 160
	Cql3ParserK_RETURNS           = 161
	Cql3ParserK_CALLED            = 162
	Cql3ParserK_INPUT             = 163
	Cql3ParserK_LANGUAGE          = 164
	Cql3ParserK_OR                = 165
	Cql3ParserK_REPLACE           = 166
	Cql3ParserK_JSON              = 167
	Cql3ParserK_DEFAULT           = 168
	Cql3ParserK_UNSET             = 169
	Cql3ParserK_LIKE              = 170
	Cql3ParserK_MASKED            = 171
	Cql3ParserK_UNMASK            = 172
	Cql3ParserK_SELECT_MASKED     = 173
	Cql3ParserK_VECTOR            = 174
	Cql3ParserK_ANN               = 175
	Cql3ParserSTRING_LITERAL      = 176
	Cql3ParserQUOTED_NAME         = 177
	Cql3ParserEMPTY_QUOTED_NAME   = 178
	Cql3ParserINTEGER             = 179
	Cql3ParserQMARK               = 180
	Cql3ParserRANGE               = 181
	Cql3ParserFLOAT               = 182
	Cql3ParserBOOLEAN             = 183
	Cql3ParserDURATION            = 184
	Cql3ParserIDENT               = 185
	Cql3ParserHEXNUMBER           = 186
	Cql3ParserUUID                = 187
	Cql3ParserWS                  = 188
	Cql3ParserCOMMENT             = 189
	Cql3ParserMULTILINE_COMMENT   = 190
)

// Cql3Parser rules.
const (
	Cql3ParserRULE_cqlStatements               = 0
	Cql3ParserRULE_cqlStatement                = 1
	Cql3ParserRULE_useStatement                = 2
	Cql3ParserRULE_marker                      = 3
	Cql3ParserRULE_createKeyspaceStatement     = 4
	Cql3ParserRULE_createTableStatement        = 5
	Cql3ParserRULE_ifNotExists                 = 6
	Cql3ParserRULE_ifExists                    = 7
	Cql3ParserRULE_tableDefinition             = 8
	Cql3ParserRULE_tableColumns                = 9
	Cql3ParserRULE_columnMask                  = 10
	Cql3ParserRULE_columnMaskArguments         = 11
	Cql3ParserRULE_tablePartitionKey           = 12
	Cql3ParserRULE_tableProperty               = 13
	Cql3ParserRULE_tableClusteringOrder        = 14
	Cql3ParserRULE_createTypeStatement         = 15
	Cql3ParserRULE_typeColumns                 = 16
	Cql3ParserRULE_dropTableStatement          = 17
	Cql3ParserRULE_truncateStatement           = 18
	Cql3ParserRULE_ident                       = 19
	Cql3ParserRULE_fident                      = 20
	Cql3ParserRULE_noncol_ident                = 21
	Cql3ParserRULE_keyspaceName                = 22
	Cql3ParserRULE_columnFamilyName            = 23
	Cql3ParserRULE_userTypeName                = 24
	Cql3ParserRULE_ksName                      = 25
	Cql3ParserRULE_cfName                      = 26
	Cql3ParserRULE_constant                    = 27
	Cql3ParserRULE_fullMapLiteral              = 28
	Cql3ParserRULE_setOrMapLiteral             = 29
	Cql3ParserRULE_setLiteral                  = 30
	Cql3ParserRULE_mapLiteral                  = 31
	Cql3ParserRULE_collectionLiteral           = 32
	Cql3ParserRULE_listLiteral                 = 33
	Cql3ParserRULE_usertypeLiteral             = 34
	Cql3ParserRULE_tupleLiteral                = 35
	Cql3ParserRULE_value                       = 36
	Cql3ParserRULE_functionName                = 37
	Cql3ParserRULE_allowedFunctionName         = 38
	Cql3ParserRULE_function                    = 39
	Cql3ParserRULE_functionArgs                = 40
	Cql3ParserRULE_term                        = 41
	Cql3ParserRULE_termAddition                = 42
	Cql3ParserRULE_termMultiplication          = 43
	Cql3ParserRULE_termGroup                   = 44
	Cql3ParserRULE_simpleTerm                  = 45
	Cql3ParserRULE_properties                  = 46
	Cql3ParserRULE_property                    = 47
	Cql3ParserRULE_propertyValue               = 48
	Cql3ParserRULE_comparatorType              = 49
	Cql3ParserRULE_native_type                 = 50
	Cql3ParserRULE_collection_type             = 51
	Cql3ParserRULE_tuple_type                  = 52
	Cql3ParserRULE_vector_type                 = 53
	Cql3ParserRULE_non_type_ident              = 54
	Cql3ParserRULE_unreserved_keyword          = 55
	Cql3ParserRULE_unreserved_function_keyword = 56
	Cql3ParserRULE_basic_unreserved_keyword    = 57
)

// ICqlStatementsContext is an interface to support dynamic dispatch.
type ICqlStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSt returns the st rule contexts.
	GetSt() ICqlStatementContext

	// SetSt sets the st rule contexts.
	SetSt(ICqlStatementContext)

	// Getter signatures
	EOF() antlr.TerminalNode
	AllCqlStatement() []ICqlStatementContext
	CqlStatement(i int) ICqlStatementContext

	// IsCqlStatementsContext differentiates from other interfaces.
	IsCqlStatementsContext()
}

type CqlStatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	st     ICqlStatementContext
}

func NewEmptyCqlStatementsContext() *CqlStatementsContext {
	var p = new(CqlStatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_cqlStatements
	return p
}

func InitEmptyCqlStatementsContext(p *CqlStatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_cqlStatements
}

func (*CqlStatementsContext) IsCqlStatementsContext() {}

func NewCqlStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlStatementsContext {
	var p = new(CqlStatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_cqlStatements

	return p
}

func (s *CqlStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlStatementsContext) GetSt() ICqlStatementContext { return s.st }

func (s *CqlStatementsContext) SetSt(v ICqlStatementContext) { s.st = v }

func (s *CqlStatementsContext) EOF() antlr.TerminalNode {
	return s.GetToken(Cql3ParserEOF, 0)
}

func (s *CqlStatementsContext) AllCqlStatement() []ICqlStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICqlStatementContext); ok {
			len++
		}
	}

	tst := make([]ICqlStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICqlStatementContext); ok {
			tst[i] = t.(ICqlStatementContext)
			i++
		}
	}

	return tst
}

func (s *CqlStatementsContext) CqlStatement(i int) ICqlStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlStatementContext)
}

func (s *CqlStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCqlStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CqlStatements() (localctx ICqlStatementsContext) {
	localctx = NewCqlStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Cql3ParserRULE_cqlStatements)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)

		var _x = p.CqlStatement()

		localctx.(*CqlStatementsContext).st = _x
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.Match(Cql3ParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(118)

				var _x = p.CqlStatement()

				localctx.(*CqlStatementsContext).st = _x
			}

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__0 {
		{
			p.SetState(124)
			p.Match(Cql3ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
		p.Match(Cql3ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlStatementContext is an interface to support dynamic dispatch.
type ICqlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSt6 returns the st6 rule contexts.
	GetSt6() IUseStatementContext

	// GetSt7 returns the st7 rule contexts.
	GetSt7() ITruncateStatementContext

	// GetSt8 returns the st8 rule contexts.
	GetSt8() ICreateKeyspaceStatementContext

	// GetSt9 returns the st9 rule contexts.
	GetSt9() ICreateTableStatementContext

	// GetSt12 returns the st12 rule contexts.
	GetSt12() IDropTableStatementContext

	// GetSt25 returns the st25 rule contexts.
	GetSt25() ICreateTypeStatementContext

	// SetSt6 sets the st6 rule contexts.
	SetSt6(IUseStatementContext)

	// SetSt7 sets the st7 rule contexts.
	SetSt7(ITruncateStatementContext)

	// SetSt8 sets the st8 rule contexts.
	SetSt8(ICreateKeyspaceStatementContext)

	// SetSt9 sets the st9 rule contexts.
	SetSt9(ICreateTableStatementContext)

	// SetSt12 sets the st12 rule contexts.
	SetSt12(IDropTableStatementContext)

	// SetSt25 sets the st25 rule contexts.
	SetSt25(ICreateTypeStatementContext)

	// Getter signatures
	UseStatement() IUseStatementContext
	TruncateStatement() ITruncateStatementContext
	CreateKeyspaceStatement() ICreateKeyspaceStatementContext
	CreateTableStatement() ICreateTableStatementContext
	DropTableStatement() IDropTableStatementContext
	CreateTypeStatement() ICreateTypeStatementContext

	// IsCqlStatementContext differentiates from other interfaces.
	IsCqlStatementContext()
}

type CqlStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	st6    IUseStatementContext
	st7    ITruncateStatementContext
	st8    ICreateKeyspaceStatementContext
	st9    ICreateTableStatementContext
	st12   IDropTableStatementContext
	st25   ICreateTypeStatementContext
}

func NewEmptyCqlStatementContext() *CqlStatementContext {
	var p = new(CqlStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_cqlStatement
	return p
}

func InitEmptyCqlStatementContext(p *CqlStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_cqlStatement
}

func (*CqlStatementContext) IsCqlStatementContext() {}

func NewCqlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlStatementContext {
	var p = new(CqlStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_cqlStatement

	return p
}

func (s *CqlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlStatementContext) GetSt6() IUseStatementContext { return s.st6 }

func (s *CqlStatementContext) GetSt7() ITruncateStatementContext { return s.st7 }

func (s *CqlStatementContext) GetSt8() ICreateKeyspaceStatementContext { return s.st8 }

func (s *CqlStatementContext) GetSt9() ICreateTableStatementContext { return s.st9 }

func (s *CqlStatementContext) GetSt12() IDropTableStatementContext { return s.st12 }

func (s *CqlStatementContext) GetSt25() ICreateTypeStatementContext { return s.st25 }

func (s *CqlStatementContext) SetSt6(v IUseStatementContext) { s.st6 = v }

func (s *CqlStatementContext) SetSt7(v ITruncateStatementContext) { s.st7 = v }

func (s *CqlStatementContext) SetSt8(v ICreateKeyspaceStatementContext) { s.st8 = v }

func (s *CqlStatementContext) SetSt9(v ICreateTableStatementContext) { s.st9 = v }

func (s *CqlStatementContext) SetSt12(v IDropTableStatementContext) { s.st12 = v }

func (s *CqlStatementContext) SetSt25(v ICreateTypeStatementContext) { s.st25 = v }

func (s *CqlStatementContext) UseStatement() IUseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseStatementContext)
}

func (s *CqlStatementContext) TruncateStatement() ITruncateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITruncateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITruncateStatementContext)
}

func (s *CqlStatementContext) CreateKeyspaceStatement() ICreateKeyspaceStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateKeyspaceStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateKeyspaceStatementContext)
}

func (s *CqlStatementContext) CreateTableStatement() ICreateTableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableStatementContext)
}

func (s *CqlStatementContext) DropTableStatement() IDropTableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropTableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropTableStatementContext)
}

func (s *CqlStatementContext) CreateTypeStatement() ICreateTypeStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTypeStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTypeStatementContext)
}

func (s *CqlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCqlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CqlStatement() (localctx ICqlStatementContext) {
	localctx = NewCqlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Cql3ParserRULE_cqlStatement)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)

			var _x = p.UseStatement()

			localctx.(*CqlStatementContext).st6 = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)

			var _x = p.TruncateStatement()

			localctx.(*CqlStatementContext).st7 = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)

			var _x = p.CreateKeyspaceStatement()

			localctx.(*CqlStatementContext).st8 = _x
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)

			var _x = p.CreateTableStatement()

			localctx.(*CqlStatementContext).st9 = _x
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(136)

			var _x = p.DropTableStatement()

			localctx.(*CqlStatementContext).st12 = _x
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(137)

			var _x = p.CreateTypeStatement()

			localctx.(*CqlStatementContext).st25 = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUseStatementContext is an interface to support dynamic dispatch.
type IUseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKs returns the ks rule contexts.
	GetKs() IKeyspaceNameContext

	// SetKs sets the ks rule contexts.
	SetKs(IKeyspaceNameContext)

	// Getter signatures
	K_USE() antlr.TerminalNode
	KeyspaceName() IKeyspaceNameContext

	// IsUseStatementContext differentiates from other interfaces.
	IsUseStatementContext()
}

type UseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ks     IKeyspaceNameContext
}

func NewEmptyUseStatementContext() *UseStatementContext {
	var p = new(UseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_useStatement
	return p
}

func InitEmptyUseStatementContext(p *UseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_useStatement
}

func (*UseStatementContext) IsUseStatementContext() {}

func NewUseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseStatementContext {
	var p = new(UseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_useStatement

	return p
}

func (s *UseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UseStatementContext) GetKs() IKeyspaceNameContext { return s.ks }

func (s *UseStatementContext) SetKs(v IKeyspaceNameContext) { s.ks = v }

func (s *UseStatementContext) K_USE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_USE, 0)
}

func (s *UseStatementContext) KeyspaceName() IKeyspaceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyspaceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyspaceNameContext)
}

func (s *UseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitUseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) UseStatement() (localctx IUseStatementContext) {
	localctx = NewUseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Cql3ParserRULE_useStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(Cql3ParserK_USE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)

		var _x = p.KeyspaceName()

		localctx.(*UseStatementContext).ks = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMarkerContext is an interface to support dynamic dispatch.
type IMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id rule contexts.
	GetId() INoncol_identContext

	// SetId sets the id rule contexts.
	SetId(INoncol_identContext)

	// Getter signatures
	Noncol_ident() INoncol_identContext
	QMARK() antlr.TerminalNode

	// IsMarkerContext differentiates from other interfaces.
	IsMarkerContext()
}

type MarkerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	id     INoncol_identContext
}

func NewEmptyMarkerContext() *MarkerContext {
	var p = new(MarkerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_marker
	return p
}

func InitEmptyMarkerContext(p *MarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_marker
}

func (*MarkerContext) IsMarkerContext() {}

func NewMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarkerContext {
	var p = new(MarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_marker

	return p
}

func (s *MarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *MarkerContext) GetId() INoncol_identContext { return s.id }

func (s *MarkerContext) SetId(v INoncol_identContext) { s.id = v }

func (s *MarkerContext) Noncol_ident() INoncol_identContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoncol_identContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoncol_identContext)
}

func (s *MarkerContext) QMARK() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQMARK, 0)
}

func (s *MarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitMarker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Marker() (localctx IMarkerContext) {
	localctx = NewMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Cql3ParserRULE_marker)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Match(Cql3ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)

			var _x = p.Noncol_ident()

			localctx.(*MarkerContext).id = _x
		}

	case Cql3ParserQMARK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(Cql3ParserQMARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateKeyspaceStatementContext is an interface to support dynamic dispatch.
type ICreateKeyspaceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKs returns the ks rule contexts.
	GetKs() IKeyspaceNameContext

	// SetKs sets the ks rule contexts.
	SetKs(IKeyspaceNameContext)

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_KEYSPACE() antlr.TerminalNode
	K_WITH() antlr.TerminalNode
	Properties() IPropertiesContext
	KeyspaceName() IKeyspaceNameContext
	IfNotExists() IIfNotExistsContext

	// IsCreateKeyspaceStatementContext differentiates from other interfaces.
	IsCreateKeyspaceStatementContext()
}

type CreateKeyspaceStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ks     IKeyspaceNameContext
}

func NewEmptyCreateKeyspaceStatementContext() *CreateKeyspaceStatementContext {
	var p = new(CreateKeyspaceStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_createKeyspaceStatement
	return p
}

func InitEmptyCreateKeyspaceStatementContext(p *CreateKeyspaceStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_createKeyspaceStatement
}

func (*CreateKeyspaceStatementContext) IsCreateKeyspaceStatementContext() {}

func NewCreateKeyspaceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateKeyspaceStatementContext {
	var p = new(CreateKeyspaceStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_createKeyspaceStatement

	return p
}

func (s *CreateKeyspaceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateKeyspaceStatementContext) GetKs() IKeyspaceNameContext { return s.ks }

func (s *CreateKeyspaceStatementContext) SetKs(v IKeyspaceNameContext) { s.ks = v }

func (s *CreateKeyspaceStatementContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CREATE, 0)
}

func (s *CreateKeyspaceStatementContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEYSPACE, 0)
}

func (s *CreateKeyspaceStatementContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_WITH, 0)
}

func (s *CreateKeyspaceStatementContext) Properties() IPropertiesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertiesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *CreateKeyspaceStatementContext) KeyspaceName() IKeyspaceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyspaceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyspaceNameContext)
}

func (s *CreateKeyspaceStatementContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateKeyspaceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateKeyspaceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateKeyspaceStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCreateKeyspaceStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CreateKeyspaceStatement() (localctx ICreateKeyspaceStatementContext) {
	localctx = NewCreateKeyspaceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Cql3ParserRULE_createKeyspaceStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(Cql3ParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(Cql3ParserK_KEYSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Cql3ParserK_IF {
		{
			p.SetState(150)
			p.IfNotExists()
		}

	}
	{
		p.SetState(153)

		var _x = p.KeyspaceName()

		localctx.(*CreateKeyspaceStatementContext).ks = _x
	}
	{
		p.SetState(154)
		p.Match(Cql3ParserK_WITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Properties()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateTableStatementContext is an interface to support dynamic dispatch.
type ICreateTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCf returns the cf rule contexts.
	GetCf() IColumnFamilyNameContext

	// SetCf sets the cf rule contexts.
	SetCf(IColumnFamilyNameContext)

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_COLUMNFAMILY() antlr.TerminalNode
	TableDefinition() ITableDefinitionContext
	ColumnFamilyName() IColumnFamilyNameContext
	IfNotExists() IIfNotExistsContext

	// IsCreateTableStatementContext differentiates from other interfaces.
	IsCreateTableStatementContext()
}

type CreateTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cf     IColumnFamilyNameContext
}

func NewEmptyCreateTableStatementContext() *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_createTableStatement
	return p
}

func InitEmptyCreateTableStatementContext(p *CreateTableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_createTableStatement
}

func (*CreateTableStatementContext) IsCreateTableStatementContext() {}

func NewCreateTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_createTableStatement

	return p
}

func (s *CreateTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableStatementContext) GetCf() IColumnFamilyNameContext { return s.cf }

func (s *CreateTableStatementContext) SetCf(v IColumnFamilyNameContext) { s.cf = v }

func (s *CreateTableStatementContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CREATE, 0)
}

func (s *CreateTableStatementContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COLUMNFAMILY, 0)
}

func (s *CreateTableStatementContext) TableDefinition() ITableDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableDefinitionContext)
}

func (s *CreateTableStatementContext) ColumnFamilyName() IColumnFamilyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnFamilyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnFamilyNameContext)
}

func (s *CreateTableStatementContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCreateTableStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CreateTableStatement() (localctx ICreateTableStatementContext) {
	localctx = NewCreateTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Cql3ParserRULE_createTableStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(Cql3ParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(Cql3ParserK_COLUMNFAMILY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Cql3ParserK_IF {
		{
			p.SetState(159)
			p.IfNotExists()
		}

	}
	{
		p.SetState(162)

		var _x = p.ColumnFamilyName()

		localctx.(*CreateTableStatementContext).cf = _x
	}
	{
		p.SetState(163)
		p.TableDefinition()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfNotExistsContext is an interface to support dynamic dispatch.
type IIfNotExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_IF() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode

	// IsIfNotExistsContext differentiates from other interfaces.
	IsIfNotExistsContext()
}

type IfNotExistsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistsContext() *IfNotExistsContext {
	var p = new(IfNotExistsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ifNotExists
	return p
}

func InitEmptyIfNotExistsContext(p *IfNotExistsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ifNotExists
}

func (*IfNotExistsContext) IsIfNotExistsContext() {}

func NewIfNotExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistsContext {
	var p = new(IfNotExistsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_ifNotExists

	return p
}

func (s *IfNotExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_IF, 0)
}

func (s *IfNotExistsContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_NOT, 0)
}

func (s *IfNotExistsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_EXISTS, 0)
}

func (s *IfNotExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitIfNotExists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) IfNotExists() (localctx IIfNotExistsContext) {
	localctx = NewIfNotExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Cql3ParserRULE_ifNotExists)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(Cql3ParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(Cql3ParserK_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(Cql3ParserK_EXISTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfExistsContext is an interface to support dynamic dispatch.
type IIfExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_IF() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode

	// IsIfExistsContext differentiates from other interfaces.
	IsIfExistsContext()
}

type IfExistsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExistsContext() *IfExistsContext {
	var p = new(IfExistsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ifExists
	return p
}

func InitEmptyIfExistsContext(p *IfExistsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ifExists
}

func (*IfExistsContext) IsIfExistsContext() {}

func NewIfExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExistsContext {
	var p = new(IfExistsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_ifExists

	return p
}

func (s *IfExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExistsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_IF, 0)
}

func (s *IfExistsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_EXISTS, 0)
}

func (s *IfExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExistsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitIfExists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) IfExists() (localctx IIfExistsContext) {
	localctx = NewIfExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Cql3ParserRULE_ifExists)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(Cql3ParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(Cql3ParserK_EXISTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableDefinitionContext is an interface to support dynamic dispatch.
type ITableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTableColumns() []ITableColumnsContext
	TableColumns(i int) ITableColumnsContext
	K_WITH() antlr.TerminalNode
	AllTableProperty() []ITablePropertyContext
	TableProperty(i int) ITablePropertyContext
	AllK_AND() []antlr.TerminalNode
	K_AND(i int) antlr.TerminalNode

	// IsTableDefinitionContext differentiates from other interfaces.
	IsTableDefinitionContext()
}

type TableDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableDefinitionContext() *TableDefinitionContext {
	var p = new(TableDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableDefinition
	return p
}

func InitEmptyTableDefinitionContext(p *TableDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableDefinition
}

func (*TableDefinitionContext) IsTableDefinitionContext() {}

func NewTableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableDefinitionContext {
	var p = new(TableDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tableDefinition

	return p
}

func (s *TableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TableDefinitionContext) AllTableColumns() []ITableColumnsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableColumnsContext); ok {
			len++
		}
	}

	tst := make([]ITableColumnsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableColumnsContext); ok {
			tst[i] = t.(ITableColumnsContext)
			i++
		}
	}

	return tst
}

func (s *TableDefinitionContext) TableColumns(i int) ITableColumnsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableColumnsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableColumnsContext)
}

func (s *TableDefinitionContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_WITH, 0)
}

func (s *TableDefinitionContext) AllTableProperty() []ITablePropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITablePropertyContext); ok {
			len++
		}
	}

	tst := make([]ITablePropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITablePropertyContext); ok {
			tst[i] = t.(ITablePropertyContext)
			i++
		}
	}

	return tst
}

func (s *TableDefinitionContext) TableProperty(i int) ITablePropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITablePropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITablePropertyContext)
}

func (s *TableDefinitionContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(Cql3ParserK_AND)
}

func (s *TableDefinitionContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_AND, i)
}

func (s *TableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TableDefinition() (localctx ITableDefinitionContext) {
	localctx = NewTableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Cql3ParserRULE_tableDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(Cql3ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.TableColumns()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(174)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-1690538491028439040) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-8527077135) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&144959475565055871) != 0) {
			{
				p.SetState(175)
				p.TableColumns()
			}

		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(Cql3ParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Cql3ParserK_WITH {
		{
			p.SetState(184)
			p.Match(Cql3ParserK_WITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.TableProperty()
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserK_AND {
			{
				p.SetState(186)
				p.Match(Cql3ParserK_AND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(187)
				p.TableProperty()
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableColumnsContext is an interface to support dynamic dispatch.
type ITableColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTableColumnsContext differentiates from other interfaces.
	IsTableColumnsContext()
}

type TableColumnsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableColumnsContext() *TableColumnsContext {
	var p = new(TableColumnsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableColumns
	return p
}

func InitEmptyTableColumnsContext(p *TableColumnsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableColumns
}

func (*TableColumnsContext) IsTableColumnsContext() {}

func NewTableColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableColumnsContext {
	var p = new(TableColumnsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tableColumns

	return p
}

func (s *TableColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *TableColumnsContext) CopyAll(ctx *TableColumnsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TableColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TableColumnDeclarationContext struct {
	TableColumnsContext
	k    IIdentContext
	v    IComparatorTypeContext
	mask IColumnMaskContext
}

func NewTableColumnDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableColumnDeclarationContext {
	var p = new(TableColumnDeclarationContext)

	InitEmptyTableColumnsContext(&p.TableColumnsContext)
	p.parser = parser
	p.CopyAll(ctx.(*TableColumnsContext))

	return p
}

func (s *TableColumnDeclarationContext) GetK() IIdentContext { return s.k }

func (s *TableColumnDeclarationContext) GetV() IComparatorTypeContext { return s.v }

func (s *TableColumnDeclarationContext) GetMask() IColumnMaskContext { return s.mask }

func (s *TableColumnDeclarationContext) SetK(v IIdentContext) { s.k = v }

func (s *TableColumnDeclarationContext) SetV(v IComparatorTypeContext) { s.v = v }

func (s *TableColumnDeclarationContext) SetMask(v IColumnMaskContext) { s.mask = v }

func (s *TableColumnDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableColumnDeclarationContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *TableColumnDeclarationContext) ComparatorType() IComparatorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *TableColumnDeclarationContext) K_STATIC() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_STATIC, 0)
}

func (s *TableColumnDeclarationContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PRIMARY, 0)
}

func (s *TableColumnDeclarationContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEY, 0)
}

func (s *TableColumnDeclarationContext) ColumnMask() IColumnMaskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnMaskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnMaskContext)
}

func (s *TableColumnDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTableColumnDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TableKeyDeclarationContext struct {
	TableColumnsContext
	pk ITablePartitionKeyContext
	c  IIdentContext
}

func NewTableKeyDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableKeyDeclarationContext {
	var p = new(TableKeyDeclarationContext)

	InitEmptyTableColumnsContext(&p.TableColumnsContext)
	p.parser = parser
	p.CopyAll(ctx.(*TableColumnsContext))

	return p
}

func (s *TableKeyDeclarationContext) GetPk() ITablePartitionKeyContext { return s.pk }

func (s *TableKeyDeclarationContext) GetC() IIdentContext { return s.c }

func (s *TableKeyDeclarationContext) SetPk(v ITablePartitionKeyContext) { s.pk = v }

func (s *TableKeyDeclarationContext) SetC(v IIdentContext) { s.c = v }

func (s *TableKeyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableKeyDeclarationContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PRIMARY, 0)
}

func (s *TableKeyDeclarationContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEY, 0)
}

func (s *TableKeyDeclarationContext) TablePartitionKey() ITablePartitionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITablePartitionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITablePartitionKeyContext)
}

func (s *TableKeyDeclarationContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *TableKeyDeclarationContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *TableKeyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTableKeyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TableColumns() (localctx ITableColumnsContext) {
	localctx = NewTableColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Cql3ParserRULE_tableColumns)
	var _la int

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN, Cql3ParserQUOTED_NAME, Cql3ParserIDENT:
		localctx = NewTableColumnDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)

			var _x = p.Ident()

			localctx.(*TableColumnDeclarationContext).k = _x
		}
		{
			p.SetState(196)

			var _x = p.ComparatorType()

			localctx.(*TableColumnDeclarationContext).v = _x
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Cql3ParserK_STATIC {
			{
				p.SetState(197)
				p.Match(Cql3ParserK_STATIC)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Cql3ParserK_MASKED {
			{
				p.SetState(200)

				var _x = p.ColumnMask()

				localctx.(*TableColumnDeclarationContext).mask = _x
			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Cql3ParserK_PRIMARY {
			{
				p.SetState(203)
				p.Match(Cql3ParserK_PRIMARY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(204)
				p.Match(Cql3ParserK_KEY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case Cql3ParserK_PRIMARY:
		localctx = NewTableKeyDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Match(Cql3ParserK_PRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(Cql3ParserK_KEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)

			var _x = p.TablePartitionKey()

			localctx.(*TableKeyDeclarationContext).pk = _x
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserT__3 {
			{
				p.SetState(211)
				p.Match(Cql3ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(212)

				var _x = p.Ident()

				localctx.(*TableKeyDeclarationContext).c = _x
			}

			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(218)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnMaskContext is an interface to support dynamic dispatch.
type IColumnMaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IFunctionNameContext

	// SetName sets the name rule contexts.
	SetName(IFunctionNameContext)

	// Getter signatures
	K_MASKED() antlr.TerminalNode
	K_WITH() antlr.TerminalNode
	ColumnMaskArguments() IColumnMaskArgumentsContext
	FunctionName() IFunctionNameContext
	K_DEFAULT() antlr.TerminalNode

	// IsColumnMaskContext differentiates from other interfaces.
	IsColumnMaskContext()
}

type ColumnMaskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IFunctionNameContext
}

func NewEmptyColumnMaskContext() *ColumnMaskContext {
	var p = new(ColumnMaskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_columnMask
	return p
}

func InitEmptyColumnMaskContext(p *ColumnMaskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_columnMask
}

func (*ColumnMaskContext) IsColumnMaskContext() {}

func NewColumnMaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnMaskContext {
	var p = new(ColumnMaskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_columnMask

	return p
}

func (s *ColumnMaskContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnMaskContext) GetName() IFunctionNameContext { return s.name }

func (s *ColumnMaskContext) SetName(v IFunctionNameContext) { s.name = v }

func (s *ColumnMaskContext) K_MASKED() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MASKED, 0)
}

func (s *ColumnMaskContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_WITH, 0)
}

func (s *ColumnMaskContext) ColumnMaskArguments() IColumnMaskArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnMaskArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnMaskArgumentsContext)
}

func (s *ColumnMaskContext) FunctionName() IFunctionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *ColumnMaskContext) K_DEFAULT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DEFAULT, 0)
}

func (s *ColumnMaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnMaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnMaskContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitColumnMask(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) ColumnMask() (localctx IColumnMaskContext) {
	localctx = NewColumnMaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Cql3ParserRULE_columnMask)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Match(Cql3ParserK_MASKED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(Cql3ParserK_WITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)

			var _x = p.FunctionName()

			localctx.(*ColumnMaskContext).name = _x
		}
		{
			p.SetState(225)
			p.ColumnMaskArguments()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(Cql3ParserK_MASKED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(Cql3ParserK_WITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(Cql3ParserK_DEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnMaskArgumentsContext is an interface to support dynamic dispatch.
type IColumnMaskArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() ITermContext

	// GetCn returns the cn rule contexts.
	GetCn() ITermContext

	// SetC sets the c rule contexts.
	SetC(ITermContext)

	// SetCn sets the cn rule contexts.
	SetCn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsColumnMaskArgumentsContext differentiates from other interfaces.
	IsColumnMaskArgumentsContext()
}

type ColumnMaskArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	c      ITermContext
	cn     ITermContext
}

func NewEmptyColumnMaskArgumentsContext() *ColumnMaskArgumentsContext {
	var p = new(ColumnMaskArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_columnMaskArguments
	return p
}

func InitEmptyColumnMaskArgumentsContext(p *ColumnMaskArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_columnMaskArguments
}

func (*ColumnMaskArgumentsContext) IsColumnMaskArgumentsContext() {}

func NewColumnMaskArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnMaskArgumentsContext {
	var p = new(ColumnMaskArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_columnMaskArguments

	return p
}

func (s *ColumnMaskArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnMaskArgumentsContext) GetC() ITermContext { return s.c }

func (s *ColumnMaskArgumentsContext) GetCn() ITermContext { return s.cn }

func (s *ColumnMaskArgumentsContext) SetC(v ITermContext) { s.c = v }

func (s *ColumnMaskArgumentsContext) SetCn(v ITermContext) { s.cn = v }

func (s *ColumnMaskArgumentsContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ColumnMaskArgumentsContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ColumnMaskArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnMaskArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnMaskArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitColumnMaskArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) ColumnMaskArguments() (localctx IColumnMaskArgumentsContext) {
	localctx = NewColumnMaskArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Cql3ParserRULE_columnMaskArguments)
	var _la int

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _x = p.Term()

			localctx.(*ColumnMaskArgumentsContext).c = _x
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserT__3 {
			{
				p.SetState(236)
				p.Match(Cql3ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(237)

				var _x = p.Term()

				localctx.(*ColumnMaskArgumentsContext).cn = _x
			}

			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(243)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITablePartitionKeyContext is an interface to support dynamic dispatch.
type ITablePartitionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK1 returns the k1 rule contexts.
	GetK1() IIdentContext

	// GetKn returns the kn rule contexts.
	GetKn() IIdentContext

	// SetK1 sets the k1 rule contexts.
	SetK1(IIdentContext)

	// SetKn sets the kn rule contexts.
	SetKn(IIdentContext)

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsTablePartitionKeyContext differentiates from other interfaces.
	IsTablePartitionKeyContext()
}

type TablePartitionKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k1     IIdentContext
	kn     IIdentContext
}

func NewEmptyTablePartitionKeyContext() *TablePartitionKeyContext {
	var p = new(TablePartitionKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tablePartitionKey
	return p
}

func InitEmptyTablePartitionKeyContext(p *TablePartitionKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tablePartitionKey
}

func (*TablePartitionKeyContext) IsTablePartitionKeyContext() {}

func NewTablePartitionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TablePartitionKeyContext {
	var p = new(TablePartitionKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tablePartitionKey

	return p
}

func (s *TablePartitionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *TablePartitionKeyContext) GetK1() IIdentContext { return s.k1 }

func (s *TablePartitionKeyContext) GetKn() IIdentContext { return s.kn }

func (s *TablePartitionKeyContext) SetK1(v IIdentContext) { s.k1 = v }

func (s *TablePartitionKeyContext) SetKn(v IIdentContext) { s.kn = v }

func (s *TablePartitionKeyContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *TablePartitionKeyContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *TablePartitionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TablePartitionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TablePartitionKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTablePartitionKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TablePartitionKey() (localctx ITablePartitionKeyContext) {
	localctx = NewTablePartitionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Cql3ParserRULE_tablePartitionKey)
	var _la int

	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN, Cql3ParserQUOTED_NAME, Cql3ParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)

			var _x = p.Ident()

			localctx.(*TablePartitionKeyContext).k1 = _x
		}

	case Cql3ParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)

			var _x = p.Ident()

			localctx.(*TablePartitionKeyContext).k1 = _x
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserT__3 {
			{
				p.SetState(250)
				p.Match(Cql3ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(251)

				var _x = p.Ident()

				localctx.(*TablePartitionKeyContext).kn = _x
			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(257)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITablePropertyContext is an interface to support dynamic dispatch.
type ITablePropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Property() IPropertyContext
	K_COMPACT() antlr.TerminalNode
	K_STORAGE() antlr.TerminalNode
	K_CLUSTERING() antlr.TerminalNode
	K_ORDER() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	AllTableClusteringOrder() []ITableClusteringOrderContext
	TableClusteringOrder(i int) ITableClusteringOrderContext

	// IsTablePropertyContext differentiates from other interfaces.
	IsTablePropertyContext()
}

type TablePropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTablePropertyContext() *TablePropertyContext {
	var p = new(TablePropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableProperty
	return p
}

func InitEmptyTablePropertyContext(p *TablePropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableProperty
}

func (*TablePropertyContext) IsTablePropertyContext() {}

func NewTablePropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TablePropertyContext {
	var p = new(TablePropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tableProperty

	return p
}

func (s *TablePropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *TablePropertyContext) Property() IPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *TablePropertyContext) K_COMPACT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COMPACT, 0)
}

func (s *TablePropertyContext) K_STORAGE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_STORAGE, 0)
}

func (s *TablePropertyContext) K_CLUSTERING() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CLUSTERING, 0)
}

func (s *TablePropertyContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ORDER, 0)
}

func (s *TablePropertyContext) K_BY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_BY, 0)
}

func (s *TablePropertyContext) AllTableClusteringOrder() []ITableClusteringOrderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableClusteringOrderContext); ok {
			len++
		}
	}

	tst := make([]ITableClusteringOrderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableClusteringOrderContext); ok {
			tst[i] = t.(ITableClusteringOrderContext)
			i++
		}
	}

	return tst
}

func (s *TablePropertyContext) TableClusteringOrder(i int) ITableClusteringOrderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableClusteringOrderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableClusteringOrderContext)
}

func (s *TablePropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TablePropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TablePropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTableProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TableProperty() (localctx ITablePropertyContext) {
	localctx = NewTablePropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Cql3ParserRULE_tableProperty)
	var _la int

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.Property()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.Match(Cql3ParserK_COMPACT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Match(Cql3ParserK_STORAGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(Cql3ParserK_CLUSTERING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(Cql3ParserK_ORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Match(Cql3ParserK_BY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.TableClusteringOrder()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserT__3 {
			{
				p.SetState(269)
				p.Match(Cql3ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(270)
				p.TableClusteringOrder()
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableClusteringOrderContext is an interface to support dynamic dispatch.
type ITableClusteringOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k rule contexts.
	GetK() IIdentContext

	// SetK sets the k rule contexts.
	SetK(IIdentContext)

	// Getter signatures
	Ident() IIdentContext
	K_ASC() antlr.TerminalNode
	K_DESC() antlr.TerminalNode

	// IsTableClusteringOrderContext differentiates from other interfaces.
	IsTableClusteringOrderContext()
}

type TableClusteringOrderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k      IIdentContext
}

func NewEmptyTableClusteringOrderContext() *TableClusteringOrderContext {
	var p = new(TableClusteringOrderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableClusteringOrder
	return p
}

func InitEmptyTableClusteringOrderContext(p *TableClusteringOrderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tableClusteringOrder
}

func (*TableClusteringOrderContext) IsTableClusteringOrderContext() {}

func NewTableClusteringOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableClusteringOrderContext {
	var p = new(TableClusteringOrderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tableClusteringOrder

	return p
}

func (s *TableClusteringOrderContext) GetParser() antlr.Parser { return s.parser }

func (s *TableClusteringOrderContext) GetK() IIdentContext { return s.k }

func (s *TableClusteringOrderContext) SetK(v IIdentContext) { s.k = v }

func (s *TableClusteringOrderContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *TableClusteringOrderContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ASC, 0)
}

func (s *TableClusteringOrderContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DESC, 0)
}

func (s *TableClusteringOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableClusteringOrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableClusteringOrderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTableClusteringOrder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TableClusteringOrder() (localctx ITableClusteringOrderContext) {
	localctx = NewTableClusteringOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Cql3ParserRULE_tableClusteringOrder)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)

		var _x = p.Ident()

		localctx.(*TableClusteringOrderContext).k = _x
	}
	{
		p.SetState(281)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Cql3ParserK_ASC || _la == Cql3ParserK_DESC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateTypeStatementContext is an interface to support dynamic dispatch.
type ICreateTypeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTn returns the tn rule contexts.
	GetTn() IUserTypeNameContext

	// SetTn sets the tn rule contexts.
	SetTn(IUserTypeNameContext)

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_TYPE() antlr.TerminalNode
	AllTypeColumns() []ITypeColumnsContext
	TypeColumns(i int) ITypeColumnsContext
	UserTypeName() IUserTypeNameContext
	IfNotExists() IIfNotExistsContext

	// IsCreateTypeStatementContext differentiates from other interfaces.
	IsCreateTypeStatementContext()
}

type CreateTypeStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tn     IUserTypeNameContext
}

func NewEmptyCreateTypeStatementContext() *CreateTypeStatementContext {
	var p = new(CreateTypeStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_createTypeStatement
	return p
}

func InitEmptyCreateTypeStatementContext(p *CreateTypeStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_createTypeStatement
}

func (*CreateTypeStatementContext) IsCreateTypeStatementContext() {}

func NewCreateTypeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTypeStatementContext {
	var p = new(CreateTypeStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_createTypeStatement

	return p
}

func (s *CreateTypeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTypeStatementContext) GetTn() IUserTypeNameContext { return s.tn }

func (s *CreateTypeStatementContext) SetTn(v IUserTypeNameContext) { s.tn = v }

func (s *CreateTypeStatementContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CREATE, 0)
}

func (s *CreateTypeStatementContext) K_TYPE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TYPE, 0)
}

func (s *CreateTypeStatementContext) AllTypeColumns() []ITypeColumnsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeColumnsContext); ok {
			len++
		}
	}

	tst := make([]ITypeColumnsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeColumnsContext); ok {
			tst[i] = t.(ITypeColumnsContext)
			i++
		}
	}

	return tst
}

func (s *CreateTypeStatementContext) TypeColumns(i int) ITypeColumnsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeColumnsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeColumnsContext)
}

func (s *CreateTypeStatementContext) UserTypeName() IUserTypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserTypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserTypeNameContext)
}

func (s *CreateTypeStatementContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateTypeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTypeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTypeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCreateTypeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CreateTypeStatement() (localctx ICreateTypeStatementContext) {
	localctx = NewCreateTypeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Cql3ParserRULE_createTypeStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(Cql3ParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(Cql3ParserK_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Cql3ParserK_IF {
		{
			p.SetState(285)
			p.IfNotExists()
		}

	}
	{
		p.SetState(288)

		var _x = p.UserTypeName()

		localctx.(*CreateTypeStatementContext).tn = _x
	}
	{
		p.SetState(289)
		p.Match(Cql3ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.TypeColumns()
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(291)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2266999243331862528) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-8527077135) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&144959475565055871) != 0) {
			{
				p.SetState(292)
				p.TypeColumns()
			}

		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(300)
		p.Match(Cql3ParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeColumnsContext is an interface to support dynamic dispatch.
type ITypeColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k rule contexts.
	GetK() IFidentContext

	// GetV returns the v rule contexts.
	GetV() IComparatorTypeContext

	// SetK sets the k rule contexts.
	SetK(IFidentContext)

	// SetV sets the v rule contexts.
	SetV(IComparatorTypeContext)

	// Getter signatures
	Fident() IFidentContext
	ComparatorType() IComparatorTypeContext

	// IsTypeColumnsContext differentiates from other interfaces.
	IsTypeColumnsContext()
}

type TypeColumnsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k      IFidentContext
	v      IComparatorTypeContext
}

func NewEmptyTypeColumnsContext() *TypeColumnsContext {
	var p = new(TypeColumnsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_typeColumns
	return p
}

func InitEmptyTypeColumnsContext(p *TypeColumnsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_typeColumns
}

func (*TypeColumnsContext) IsTypeColumnsContext() {}

func NewTypeColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeColumnsContext {
	var p = new(TypeColumnsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_typeColumns

	return p
}

func (s *TypeColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeColumnsContext) GetK() IFidentContext { return s.k }

func (s *TypeColumnsContext) GetV() IComparatorTypeContext { return s.v }

func (s *TypeColumnsContext) SetK(v IFidentContext) { s.k = v }

func (s *TypeColumnsContext) SetV(v IComparatorTypeContext) { s.v = v }

func (s *TypeColumnsContext) Fident() IFidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFidentContext)
}

func (s *TypeColumnsContext) ComparatorType() IComparatorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *TypeColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeColumnsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTypeColumns(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TypeColumns() (localctx ITypeColumnsContext) {
	localctx = NewTypeColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Cql3ParserRULE_typeColumns)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)

		var _x = p.Fident()

		localctx.(*TypeColumnsContext).k = _x
	}
	{
		p.SetState(303)

		var _x = p.ComparatorType()

		localctx.(*TypeColumnsContext).v = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropTableStatementContext is an interface to support dynamic dispatch.
type IDropTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IColumnFamilyNameContext

	// SetName sets the name rule contexts.
	SetName(IColumnFamilyNameContext)

	// Getter signatures
	K_DROP() antlr.TerminalNode
	K_COLUMNFAMILY() antlr.TerminalNode
	ColumnFamilyName() IColumnFamilyNameContext
	IfExists() IIfExistsContext

	// IsDropTableStatementContext differentiates from other interfaces.
	IsDropTableStatementContext()
}

type DropTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IColumnFamilyNameContext
}

func NewEmptyDropTableStatementContext() *DropTableStatementContext {
	var p = new(DropTableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_dropTableStatement
	return p
}

func InitEmptyDropTableStatementContext(p *DropTableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_dropTableStatement
}

func (*DropTableStatementContext) IsDropTableStatementContext() {}

func NewDropTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTableStatementContext {
	var p = new(DropTableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_dropTableStatement

	return p
}

func (s *DropTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTableStatementContext) GetName() IColumnFamilyNameContext { return s.name }

func (s *DropTableStatementContext) SetName(v IColumnFamilyNameContext) { s.name = v }

func (s *DropTableStatementContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DROP, 0)
}

func (s *DropTableStatementContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COLUMNFAMILY, 0)
}

func (s *DropTableStatementContext) ColumnFamilyName() IColumnFamilyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnFamilyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnFamilyNameContext)
}

func (s *DropTableStatementContext) IfExists() IIfExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropTableStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitDropTableStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) DropTableStatement() (localctx IDropTableStatementContext) {
	localctx = NewDropTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Cql3ParserRULE_dropTableStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(Cql3ParserK_DROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(Cql3ParserK_COLUMNFAMILY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Cql3ParserK_IF {
		{
			p.SetState(307)
			p.IfExists()
		}

	}
	{
		p.SetState(310)

		var _x = p.ColumnFamilyName()

		localctx.(*DropTableStatementContext).name = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITruncateStatementContext is an interface to support dynamic dispatch.
type ITruncateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCf returns the cf rule contexts.
	GetCf() IColumnFamilyNameContext

	// SetCf sets the cf rule contexts.
	SetCf(IColumnFamilyNameContext)

	// Getter signatures
	K_TRUNCATE() antlr.TerminalNode
	ColumnFamilyName() IColumnFamilyNameContext
	K_COLUMNFAMILY() antlr.TerminalNode

	// IsTruncateStatementContext differentiates from other interfaces.
	IsTruncateStatementContext()
}

type TruncateStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cf     IColumnFamilyNameContext
}

func NewEmptyTruncateStatementContext() *TruncateStatementContext {
	var p = new(TruncateStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_truncateStatement
	return p
}

func InitEmptyTruncateStatementContext(p *TruncateStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_truncateStatement
}

func (*TruncateStatementContext) IsTruncateStatementContext() {}

func NewTruncateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TruncateStatementContext {
	var p = new(TruncateStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_truncateStatement

	return p
}

func (s *TruncateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TruncateStatementContext) GetCf() IColumnFamilyNameContext { return s.cf }

func (s *TruncateStatementContext) SetCf(v IColumnFamilyNameContext) { s.cf = v }

func (s *TruncateStatementContext) K_TRUNCATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TRUNCATE, 0)
}

func (s *TruncateStatementContext) ColumnFamilyName() IColumnFamilyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnFamilyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnFamilyNameContext)
}

func (s *TruncateStatementContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COLUMNFAMILY, 0)
}

func (s *TruncateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TruncateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TruncateStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTruncateStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TruncateStatement() (localctx ITruncateStatementContext) {
	localctx = NewTruncateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Cql3ParserRULE_truncateStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(Cql3ParserK_TRUNCATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Cql3ParserK_COLUMNFAMILY {
		{
			p.SetState(313)
			p.Match(Cql3ParserK_COLUMNFAMILY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(316)

		var _x = p.ColumnFamilyName()

		localctx.(*TruncateStatementContext).cf = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) CopyAll(ctx *IdentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentQuotedIdentContext struct {
	IdentContext
	t antlr.Token
}

func NewIdentQuotedIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentQuotedIdentContext {
	var p = new(IdentQuotedIdentContext)

	InitEmptyIdentContext(&p.IdentContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentContext))

	return p
}

func (s *IdentQuotedIdentContext) GetT() antlr.Token { return s.t }

func (s *IdentQuotedIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *IdentQuotedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentQuotedIdentContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *IdentQuotedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitIdentQuotedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentIdentContext struct {
	IdentContext
	t antlr.Token
	k IUnreserved_keywordContext
}

func NewIdentIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentIdentContext {
	var p = new(IdentIdentContext)

	InitEmptyIdentContext(&p.IdentContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentContext))

	return p
}

func (s *IdentIdentContext) GetT() antlr.Token { return s.t }

func (s *IdentIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *IdentIdentContext) GetK() IUnreserved_keywordContext { return s.k }

func (s *IdentIdentContext) SetK(v IUnreserved_keywordContext) { s.k = v }

func (s *IdentIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *IdentIdentContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *IdentIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitIdentIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Cql3ParserRULE_ident)
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserIDENT:
		localctx = NewIdentIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)

			var _m = p.Match(Cql3ParserIDENT)

			localctx.(*IdentIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewIdentQuotedIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*IdentQuotedIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewIdentIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(320)

			var _x = p.Unreserved_keyword()

			localctx.(*IdentIdentContext).k = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFidentContext is an interface to support dynamic dispatch.
type IFidentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFidentContext differentiates from other interfaces.
	IsFidentContext()
}

type FidentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFidentContext() *FidentContext {
	var p = new(FidentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_fident
	return p
}

func InitEmptyFidentContext(p *FidentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_fident
}

func (*FidentContext) IsFidentContext() {}

func NewFidentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FidentContext {
	var p = new(FidentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_fident

	return p
}

func (s *FidentContext) GetParser() antlr.Parser { return s.parser }

func (s *FidentContext) CopyAll(ctx *FidentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FidentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FidentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FIdentQuotedIdentContext struct {
	FidentContext
	t antlr.Token
}

func NewFIdentQuotedIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FIdentQuotedIdentContext {
	var p = new(FIdentQuotedIdentContext)

	InitEmptyFidentContext(&p.FidentContext)
	p.parser = parser
	p.CopyAll(ctx.(*FidentContext))

	return p
}

func (s *FIdentQuotedIdentContext) GetT() antlr.Token { return s.t }

func (s *FIdentQuotedIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *FIdentQuotedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FIdentQuotedIdentContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *FIdentQuotedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitFIdentQuotedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type FIdentIdentContext struct {
	FidentContext
	t antlr.Token
	k IUnreserved_keywordContext
}

func NewFIdentIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FIdentIdentContext {
	var p = new(FIdentIdentContext)

	InitEmptyFidentContext(&p.FidentContext)
	p.parser = parser
	p.CopyAll(ctx.(*FidentContext))

	return p
}

func (s *FIdentIdentContext) GetT() antlr.Token { return s.t }

func (s *FIdentIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *FIdentIdentContext) GetK() IUnreserved_keywordContext { return s.k }

func (s *FIdentIdentContext) SetK(v IUnreserved_keywordContext) { s.k = v }

func (s *FIdentIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FIdentIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *FIdentIdentContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *FIdentIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitFIdentIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Fident() (localctx IFidentContext) {
	localctx = NewFidentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Cql3ParserRULE_fident)
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserIDENT:
		localctx = NewFIdentIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)

			var _m = p.Match(Cql3ParserIDENT)

			localctx.(*FIdentIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewFIdentQuotedIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*FIdentQuotedIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewFIdentIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(325)

			var _x = p.Unreserved_keyword()

			localctx.(*FIdentIdentContext).k = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoncol_identContext is an interface to support dynamic dispatch.
type INoncol_identContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNoncol_identContext differentiates from other interfaces.
	IsNoncol_identContext()
}

type Noncol_identContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoncol_identContext() *Noncol_identContext {
	var p = new(Noncol_identContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_noncol_ident
	return p
}

func InitEmptyNoncol_identContext(p *Noncol_identContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_noncol_ident
}

func (*Noncol_identContext) IsNoncol_identContext() {}

func NewNoncol_identContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Noncol_identContext {
	var p = new(Noncol_identContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_noncol_ident

	return p
}

func (s *Noncol_identContext) GetParser() antlr.Parser { return s.parser }

func (s *Noncol_identContext) CopyAll(ctx *Noncol_identContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Noncol_identContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Noncol_identContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NonColIdentContext struct {
	Noncol_identContext
	i antlr.Token
	k IUnreserved_keywordContext
}

func NewNonColIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NonColIdentContext {
	var p = new(NonColIdentContext)

	InitEmptyNoncol_identContext(&p.Noncol_identContext)
	p.parser = parser
	p.CopyAll(ctx.(*Noncol_identContext))

	return p
}

func (s *NonColIdentContext) GetI() antlr.Token { return s.i }

func (s *NonColIdentContext) SetI(v antlr.Token) { s.i = v }

func (s *NonColIdentContext) GetK() IUnreserved_keywordContext { return s.k }

func (s *NonColIdentContext) SetK(v IUnreserved_keywordContext) { s.k = v }

func (s *NonColIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonColIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *NonColIdentContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *NonColIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitNonColIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type NonColQuotedContext struct {
	Noncol_identContext
	t antlr.Token
}

func NewNonColQuotedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NonColQuotedContext {
	var p = new(NonColQuotedContext)

	InitEmptyNoncol_identContext(&p.Noncol_identContext)
	p.parser = parser
	p.CopyAll(ctx.(*Noncol_identContext))

	return p
}

func (s *NonColQuotedContext) GetT() antlr.Token { return s.t }

func (s *NonColQuotedContext) SetT(v antlr.Token) { s.t = v }

func (s *NonColQuotedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonColQuotedContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *NonColQuotedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitNonColQuoted(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Noncol_ident() (localctx INoncol_identContext) {
	localctx = NewNoncol_identContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Cql3ParserRULE_noncol_ident)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserIDENT:
		localctx = NewNonColIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)

			var _m = p.Match(Cql3ParserIDENT)

			localctx.(*NonColIdentContext).i = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewNonColQuotedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*NonColQuotedContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewNonColIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)

			var _x = p.Unreserved_keyword()

			localctx.(*NonColIdentContext).k = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyspaceNameContext is an interface to support dynamic dispatch.
type IKeyspaceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KsName() IKsNameContext

	// IsKeyspaceNameContext differentiates from other interfaces.
	IsKeyspaceNameContext()
}

type KeyspaceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyspaceNameContext() *KeyspaceNameContext {
	var p = new(KeyspaceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_keyspaceName
	return p
}

func InitEmptyKeyspaceNameContext(p *KeyspaceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_keyspaceName
}

func (*KeyspaceNameContext) IsKeyspaceNameContext() {}

func NewKeyspaceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyspaceNameContext {
	var p = new(KeyspaceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_keyspaceName

	return p
}

func (s *KeyspaceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyspaceNameContext) KsName() IKsNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKsNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKsNameContext)
}

func (s *KeyspaceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyspaceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyspaceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitKeyspaceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) KeyspaceName() (localctx IKeyspaceNameContext) {
	localctx = NewKeyspaceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Cql3ParserRULE_keyspaceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.KsName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnFamilyNameContext is an interface to support dynamic dispatch.
type IColumnFamilyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CfName() ICfNameContext
	KsName() IKsNameContext

	// IsColumnFamilyNameContext differentiates from other interfaces.
	IsColumnFamilyNameContext()
}

type ColumnFamilyNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnFamilyNameContext() *ColumnFamilyNameContext {
	var p = new(ColumnFamilyNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_columnFamilyName
	return p
}

func InitEmptyColumnFamilyNameContext(p *ColumnFamilyNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_columnFamilyName
}

func (*ColumnFamilyNameContext) IsColumnFamilyNameContext() {}

func NewColumnFamilyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnFamilyNameContext {
	var p = new(ColumnFamilyNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_columnFamilyName

	return p
}

func (s *ColumnFamilyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnFamilyNameContext) CfName() ICfNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICfNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICfNameContext)
}

func (s *ColumnFamilyNameContext) KsName() IKsNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKsNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKsNameContext)
}

func (s *ColumnFamilyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnFamilyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnFamilyNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitColumnFamilyName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) ColumnFamilyName() (localctx IColumnFamilyNameContext) {
	localctx = NewColumnFamilyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Cql3ParserRULE_columnFamilyName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(335)
			p.KsName()
		}
		{
			p.SetState(336)
			p.Match(Cql3ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(340)
		p.CfName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUserTypeNameContext is an interface to support dynamic dispatch.
type IUserTypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKs returns the ks rule contexts.
	GetKs() INoncol_identContext

	// GetUt returns the ut rule contexts.
	GetUt() INon_type_identContext

	// SetKs sets the ks rule contexts.
	SetKs(INoncol_identContext)

	// SetUt sets the ut rule contexts.
	SetUt(INon_type_identContext)

	// Getter signatures
	Non_type_ident() INon_type_identContext
	Noncol_ident() INoncol_identContext

	// IsUserTypeNameContext differentiates from other interfaces.
	IsUserTypeNameContext()
}

type UserTypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ks     INoncol_identContext
	ut     INon_type_identContext
}

func NewEmptyUserTypeNameContext() *UserTypeNameContext {
	var p = new(UserTypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_userTypeName
	return p
}

func InitEmptyUserTypeNameContext(p *UserTypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_userTypeName
}

func (*UserTypeNameContext) IsUserTypeNameContext() {}

func NewUserTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserTypeNameContext {
	var p = new(UserTypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_userTypeName

	return p
}

func (s *UserTypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *UserTypeNameContext) GetKs() INoncol_identContext { return s.ks }

func (s *UserTypeNameContext) GetUt() INon_type_identContext { return s.ut }

func (s *UserTypeNameContext) SetKs(v INoncol_identContext) { s.ks = v }

func (s *UserTypeNameContext) SetUt(v INon_type_identContext) { s.ut = v }

func (s *UserTypeNameContext) Non_type_ident() INon_type_identContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INon_type_identContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INon_type_identContext)
}

func (s *UserTypeNameContext) Noncol_ident() INoncol_identContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoncol_identContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoncol_identContext)
}

func (s *UserTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserTypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserTypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitUserTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) UserTypeName() (localctx IUserTypeNameContext) {
	localctx = NewUserTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Cql3ParserRULE_userTypeName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(342)

			var _x = p.Noncol_ident()

			localctx.(*UserTypeNameContext).ks = _x
		}
		{
			p.SetState(343)
			p.Match(Cql3ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(347)

		var _x = p.Non_type_ident()

		localctx.(*UserTypeNameContext).ut = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKsNameContext is an interface to support dynamic dispatch.
type IKsNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKsNameContext differentiates from other interfaces.
	IsKsNameContext()
}

type KsNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKsNameContext() *KsNameContext {
	var p = new(KsNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ksName
	return p
}

func InitEmptyKsNameContext(p *KsNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_ksName
}

func (*KsNameContext) IsKsNameContext() {}

func NewKsNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KsNameContext {
	var p = new(KsNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_ksName

	return p
}

func (s *KsNameContext) GetParser() antlr.Parser { return s.parser }

func (s *KsNameContext) CopyAll(ctx *KsNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KsNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KsNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type KsNameIdentContext struct {
	KsNameContext
	t antlr.Token
	k IUnreserved_keywordContext
}

func NewKsNameIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KsNameIdentContext {
	var p = new(KsNameIdentContext)

	InitEmptyKsNameContext(&p.KsNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*KsNameContext))

	return p
}

func (s *KsNameIdentContext) GetT() antlr.Token { return s.t }

func (s *KsNameIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *KsNameIdentContext) GetK() IUnreserved_keywordContext { return s.k }

func (s *KsNameIdentContext) SetK(v IUnreserved_keywordContext) { s.k = v }

func (s *KsNameIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KsNameIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *KsNameIdentContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *KsNameIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitKsNameIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type KsNameQuotedIdentContext struct {
	KsNameContext
	t antlr.Token
}

func NewKsNameQuotedIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KsNameQuotedIdentContext {
	var p = new(KsNameQuotedIdentContext)

	InitEmptyKsNameContext(&p.KsNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*KsNameContext))

	return p
}

func (s *KsNameQuotedIdentContext) GetT() antlr.Token { return s.t }

func (s *KsNameQuotedIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *KsNameQuotedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KsNameQuotedIdentContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *KsNameQuotedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitKsNameQuotedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type KsNameInvalidBindContext struct {
	KsNameContext
}

func NewKsNameInvalidBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KsNameInvalidBindContext {
	var p = new(KsNameInvalidBindContext)

	InitEmptyKsNameContext(&p.KsNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*KsNameContext))

	return p
}

func (s *KsNameInvalidBindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KsNameInvalidBindContext) QMARK() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQMARK, 0)
}

func (s *KsNameInvalidBindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitKsNameInvalidBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) KsName() (localctx IKsNameContext) {
	localctx = NewKsNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Cql3ParserRULE_ksName)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserIDENT:
		localctx = NewKsNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)

			var _m = p.Match(Cql3ParserIDENT)

			localctx.(*KsNameIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewKsNameQuotedIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*KsNameQuotedIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewKsNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(351)

			var _x = p.Unreserved_keyword()

			localctx.(*KsNameIdentContext).k = _x
		}

	case Cql3ParserQMARK:
		localctx = NewKsNameInvalidBindContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(352)
			p.Match(Cql3ParserQMARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICfNameContext is an interface to support dynamic dispatch.
type ICfNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCfNameContext differentiates from other interfaces.
	IsCfNameContext()
}

type CfNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCfNameContext() *CfNameContext {
	var p = new(CfNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_cfName
	return p
}

func InitEmptyCfNameContext(p *CfNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_cfName
}

func (*CfNameContext) IsCfNameContext() {}

func NewCfNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CfNameContext {
	var p = new(CfNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_cfName

	return p
}

func (s *CfNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CfNameContext) CopyAll(ctx *CfNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CfNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CfNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CfNameIdentContext struct {
	CfNameContext
	t antlr.Token
	k IUnreserved_keywordContext
}

func NewCfNameIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CfNameIdentContext {
	var p = new(CfNameIdentContext)

	InitEmptyCfNameContext(&p.CfNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*CfNameContext))

	return p
}

func (s *CfNameIdentContext) GetT() antlr.Token { return s.t }

func (s *CfNameIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *CfNameIdentContext) GetK() IUnreserved_keywordContext { return s.k }

func (s *CfNameIdentContext) SetK(v IUnreserved_keywordContext) { s.k = v }

func (s *CfNameIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CfNameIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *CfNameIdentContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *CfNameIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCfNameIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type CfNameQuotedIdentContext struct {
	CfNameContext
	t antlr.Token
}

func NewCfNameQuotedIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CfNameQuotedIdentContext {
	var p = new(CfNameQuotedIdentContext)

	InitEmptyCfNameContext(&p.CfNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*CfNameContext))

	return p
}

func (s *CfNameQuotedIdentContext) GetT() antlr.Token { return s.t }

func (s *CfNameQuotedIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *CfNameQuotedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CfNameQuotedIdentContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *CfNameQuotedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCfNameQuotedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type CfNameInvalidBindContext struct {
	CfNameContext
}

func NewCfNameInvalidBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CfNameInvalidBindContext {
	var p = new(CfNameInvalidBindContext)

	InitEmptyCfNameContext(&p.CfNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*CfNameContext))

	return p
}

func (s *CfNameInvalidBindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CfNameInvalidBindContext) QMARK() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQMARK, 0)
}

func (s *CfNameInvalidBindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCfNameInvalidBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CfName() (localctx ICfNameContext) {
	localctx = NewCfNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Cql3ParserRULE_cfName)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserIDENT:
		localctx = NewCfNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)

			var _m = p.Match(Cql3ParserIDENT)

			localctx.(*CfNameIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewCfNameQuotedIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*CfNameQuotedIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewCfNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(357)

			var _x = p.Unreserved_keyword()

			localctx.(*CfNameIdentContext).k = _x
		}

	case Cql3ParserQMARK:
		localctx = NewCfNameInvalidBindContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(358)
			p.Match(Cql3ParserQMARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	DURATION() antlr.TerminalNode
	UUID() antlr.TerminalNode
	HEXNUMBER() antlr.TerminalNode
	K_POSITIVE_INFINITY() antlr.TerminalNode
	K_NEGATIVE_INFINITY() antlr.TerminalNode
	K_POSITIVE_NAN() antlr.TerminalNode
	K_NEGATIVE_NAN() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t      antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetT() antlr.Token { return s.t }

func (s *ConstantContext) SetT(v antlr.Token) { s.t = v }

func (s *ConstantContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(Cql3ParserSTRING_LITERAL, 0)
}

func (s *ConstantContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserINTEGER, 0)
}

func (s *ConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserFLOAT, 0)
}

func (s *ConstantContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserBOOLEAN, 0)
}

func (s *ConstantContext) DURATION() antlr.TerminalNode {
	return s.GetToken(Cql3ParserDURATION, 0)
}

func (s *ConstantContext) UUID() antlr.TerminalNode {
	return s.GetToken(Cql3ParserUUID, 0)
}

func (s *ConstantContext) HEXNUMBER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserHEXNUMBER, 0)
}

func (s *ConstantContext) K_POSITIVE_INFINITY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_POSITIVE_INFINITY, 0)
}

func (s *ConstantContext) K_NEGATIVE_INFINITY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_NEGATIVE_INFINITY, 0)
}

func (s *ConstantContext) K_POSITIVE_NAN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_POSITIVE_NAN, 0)
}

func (s *ConstantContext) K_NEGATIVE_NAN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_NEGATIVE_NAN, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Cql3ParserRULE_constant)
	var _la int

	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)

			var _m = p.Match(Cql3ParserSTRING_LITERAL)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)

			var _m = p.Match(Cql3ParserINTEGER)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(363)

			var _m = p.Match(Cql3ParserFLOAT)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserBOOLEAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)

			var _m = p.Match(Cql3ParserBOOLEAN)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserDURATION:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(365)

			var _m = p.Match(Cql3ParserDURATION)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserUUID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(366)

			var _m = p.Match(Cql3ParserUUID)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserHEXNUMBER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(367)

			var _m = p.Match(Cql3ParserHEXNUMBER)

			localctx.(*ConstantContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_POSITIVE_NAN, Cql3ParserK_NEGATIVE_NAN, Cql3ParserK_POSITIVE_INFINITY, Cql3ParserK_NEGATIVE_INFINITY:
		p.EnterOuterAlt(localctx, 8)
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Cql3ParserK_POSITIVE_NAN, Cql3ParserK_NEGATIVE_NAN:
			{
				p.SetState(368)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Cql3ParserK_POSITIVE_NAN || _la == Cql3ParserK_NEGATIVE_NAN) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		case Cql3ParserK_POSITIVE_INFINITY:
			{
				p.SetState(369)
				p.Match(Cql3ParserK_POSITIVE_INFINITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case Cql3ParserK_NEGATIVE_INFINITY:
			{
				p.SetState(370)
				p.Match(Cql3ParserK_NEGATIVE_INFINITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFullMapLiteralContext is an interface to support dynamic dispatch.
type IFullMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK1 returns the k1 rule contexts.
	GetK1() ITermContext

	// GetV1 returns the v1 rule contexts.
	GetV1() ITermContext

	// GetKn returns the kn rule contexts.
	GetKn() ITermContext

	// GetVn returns the vn rule contexts.
	GetVn() ITermContext

	// SetK1 sets the k1 rule contexts.
	SetK1(ITermContext)

	// SetV1 sets the v1 rule contexts.
	SetV1(ITermContext)

	// SetKn sets the kn rule contexts.
	SetKn(ITermContext)

	// SetVn sets the vn rule contexts.
	SetVn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsFullMapLiteralContext differentiates from other interfaces.
	IsFullMapLiteralContext()
}

type FullMapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k1     ITermContext
	v1     ITermContext
	kn     ITermContext
	vn     ITermContext
}

func NewEmptyFullMapLiteralContext() *FullMapLiteralContext {
	var p = new(FullMapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_fullMapLiteral
	return p
}

func InitEmptyFullMapLiteralContext(p *FullMapLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_fullMapLiteral
}

func (*FullMapLiteralContext) IsFullMapLiteralContext() {}

func NewFullMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullMapLiteralContext {
	var p = new(FullMapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_fullMapLiteral

	return p
}

func (s *FullMapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FullMapLiteralContext) GetK1() ITermContext { return s.k1 }

func (s *FullMapLiteralContext) GetV1() ITermContext { return s.v1 }

func (s *FullMapLiteralContext) GetKn() ITermContext { return s.kn }

func (s *FullMapLiteralContext) GetVn() ITermContext { return s.vn }

func (s *FullMapLiteralContext) SetK1(v ITermContext) { s.k1 = v }

func (s *FullMapLiteralContext) SetV1(v ITermContext) { s.v1 = v }

func (s *FullMapLiteralContext) SetKn(v ITermContext) { s.kn = v }

func (s *FullMapLiteralContext) SetVn(v ITermContext) { s.vn = v }

func (s *FullMapLiteralContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *FullMapLiteralContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FullMapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullMapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullMapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitFullMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) FullMapLiteral() (localctx IFullMapLiteralContext) {
	localctx = NewFullMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Cql3ParserRULE_fullMapLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(Cql3ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2266999243331857780) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-8527077135) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&1142788268006301695) != 0) {
		{
			p.SetState(376)

			var _x = p.Term()

			localctx.(*FullMapLiteralContext).k1 = _x
		}
		{
			p.SetState(377)
			p.Match(Cql3ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)

			var _x = p.Term()

			localctx.(*FullMapLiteralContext).v1 = _x
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserT__3 {
			{
				p.SetState(379)
				p.Match(Cql3ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(380)

				var _x = p.Term()

				localctx.(*FullMapLiteralContext).kn = _x
			}
			{
				p.SetState(381)
				p.Match(Cql3ParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(382)

				var _x = p.Term()

				localctx.(*FullMapLiteralContext).vn = _x
			}

			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(391)
		p.Match(Cql3ParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetOrMapLiteralContext is an interface to support dynamic dispatch.
type ISetOrMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m rule contexts.
	GetM() IMapLiteralContext

	// GetS returns the s rule contexts.
	GetS() ISetLiteralContext

	// SetM sets the m rule contexts.
	SetM(IMapLiteralContext)

	// SetS sets the s rule contexts.
	SetS(ISetLiteralContext)

	// Getter signatures
	MapLiteral() IMapLiteralContext
	SetLiteral() ISetLiteralContext

	// IsSetOrMapLiteralContext differentiates from other interfaces.
	IsSetOrMapLiteralContext()
}

type SetOrMapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	m      IMapLiteralContext
	s      ISetLiteralContext
}

func NewEmptySetOrMapLiteralContext() *SetOrMapLiteralContext {
	var p = new(SetOrMapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_setOrMapLiteral
	return p
}

func InitEmptySetOrMapLiteralContext(p *SetOrMapLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_setOrMapLiteral
}

func (*SetOrMapLiteralContext) IsSetOrMapLiteralContext() {}

func NewSetOrMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetOrMapLiteralContext {
	var p = new(SetOrMapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_setOrMapLiteral

	return p
}

func (s *SetOrMapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *SetOrMapLiteralContext) GetM() IMapLiteralContext { return s.m }

func (s *SetOrMapLiteralContext) GetS() ISetLiteralContext { return s.s }

func (s *SetOrMapLiteralContext) SetM(v IMapLiteralContext) { s.m = v }

func (s *SetOrMapLiteralContext) SetS(v ISetLiteralContext) { s.s = v }

func (s *SetOrMapLiteralContext) MapLiteral() IMapLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapLiteralContext)
}

func (s *SetOrMapLiteralContext) SetLiteral() ISetLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetLiteralContext)
}

func (s *SetOrMapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetOrMapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetOrMapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitSetOrMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) SetOrMapLiteral() (localctx ISetOrMapLiteralContext) {
	localctx = NewSetOrMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Cql3ParserRULE_setOrMapLiteral)
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(393)

			var _x = p.MapLiteral()

			localctx.(*SetOrMapLiteralContext).m = _x
		}

	case Cql3ParserT__3, Cql3ParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)

			var _x = p.SetLiteral()

			localctx.(*SetOrMapLiteralContext).s = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetLiteralContext is an interface to support dynamic dispatch.
type ISetLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTn returns the tn rule contexts.
	GetTn() ITermContext

	// SetTn sets the tn rule contexts.
	SetTn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsSetLiteralContext differentiates from other interfaces.
	IsSetLiteralContext()
}

type SetLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tn     ITermContext
}

func NewEmptySetLiteralContext() *SetLiteralContext {
	var p = new(SetLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_setLiteral
	return p
}

func InitEmptySetLiteralContext(p *SetLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_setLiteral
}

func (*SetLiteralContext) IsSetLiteralContext() {}

func NewSetLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetLiteralContext {
	var p = new(SetLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_setLiteral

	return p
}

func (s *SetLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *SetLiteralContext) GetTn() ITermContext { return s.tn }

func (s *SetLiteralContext) SetTn(v ITermContext) { s.tn = v }

func (s *SetLiteralContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *SetLiteralContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SetLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitSetLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) SetLiteral() (localctx ISetLiteralContext) {
	localctx = NewSetLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Cql3ParserRULE_setLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(397)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)

			var _x = p.Term()

			localctx.(*SetLiteralContext).tn = _x
		}

		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapLiteralContext is an interface to support dynamic dispatch.
type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v rule contexts.
	GetV() ITermContext

	// GetKn returns the kn rule contexts.
	GetKn() ITermContext

	// GetVn returns the vn rule contexts.
	GetVn() ITermContext

	// SetV sets the v rule contexts.
	SetV(ITermContext)

	// SetKn sets the kn rule contexts.
	SetKn(ITermContext)

	// SetVn sets the vn rule contexts.
	SetVn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	v      ITermContext
	kn     ITermContext
	vn     ITermContext
}

func NewEmptyMapLiteralContext() *MapLiteralContext {
	var p = new(MapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_mapLiteral
	return p
}

func InitEmptyMapLiteralContext(p *MapLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_mapLiteral
}

func (*MapLiteralContext) IsMapLiteralContext() {}

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext {
	var p = new(MapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_mapLiteral

	return p
}

func (s *MapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralContext) GetV() ITermContext { return s.v }

func (s *MapLiteralContext) GetKn() ITermContext { return s.kn }

func (s *MapLiteralContext) GetVn() ITermContext { return s.vn }

func (s *MapLiteralContext) SetV(v ITermContext) { s.v = v }

func (s *MapLiteralContext) SetKn(v ITermContext) { s.kn = v }

func (s *MapLiteralContext) SetVn(v ITermContext) { s.vn = v }

func (s *MapLiteralContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *MapLiteralContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) MapLiteral() (localctx IMapLiteralContext) {
	localctx = NewMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Cql3ParserRULE_mapLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(Cql3ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)

		var _x = p.Term()

		localctx.(*MapLiteralContext).v = _x
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(406)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)

			var _x = p.Term()

			localctx.(*MapLiteralContext).kn = _x
		}
		{
			p.SetState(408)
			p.Match(Cql3ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)

			var _x = p.Term()

			localctx.(*MapLiteralContext).vn = _x
		}

		p.SetState(415)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollectionLiteralContext is an interface to support dynamic dispatch.
type ICollectionLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IListLiteralContext

	// GetT returns the t rule contexts.
	GetT() ITermContext

	// GetV returns the v rule contexts.
	GetV() ISetOrMapLiteralContext

	// SetL sets the l rule contexts.
	SetL(IListLiteralContext)

	// SetT sets the t rule contexts.
	SetT(ITermContext)

	// SetV sets the v rule contexts.
	SetV(ISetOrMapLiteralContext)

	// Getter signatures
	ListLiteral() IListLiteralContext
	Term() ITermContext
	SetOrMapLiteral() ISetOrMapLiteralContext

	// IsCollectionLiteralContext differentiates from other interfaces.
	IsCollectionLiteralContext()
}

type CollectionLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      IListLiteralContext
	t      ITermContext
	v      ISetOrMapLiteralContext
}

func NewEmptyCollectionLiteralContext() *CollectionLiteralContext {
	var p = new(CollectionLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_collectionLiteral
	return p
}

func InitEmptyCollectionLiteralContext(p *CollectionLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_collectionLiteral
}

func (*CollectionLiteralContext) IsCollectionLiteralContext() {}

func NewCollectionLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionLiteralContext {
	var p = new(CollectionLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_collectionLiteral

	return p
}

func (s *CollectionLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionLiteralContext) GetL() IListLiteralContext { return s.l }

func (s *CollectionLiteralContext) GetT() ITermContext { return s.t }

func (s *CollectionLiteralContext) GetV() ISetOrMapLiteralContext { return s.v }

func (s *CollectionLiteralContext) SetL(v IListLiteralContext) { s.l = v }

func (s *CollectionLiteralContext) SetT(v ITermContext) { s.t = v }

func (s *CollectionLiteralContext) SetV(v ISetOrMapLiteralContext) { s.v = v }

func (s *CollectionLiteralContext) ListLiteral() IListLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListLiteralContext)
}

func (s *CollectionLiteralContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *CollectionLiteralContext) SetOrMapLiteral() ISetOrMapLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetOrMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetOrMapLiteralContext)
}

func (s *CollectionLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCollectionLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) CollectionLiteral() (localctx ICollectionLiteralContext) {
	localctx = NewCollectionLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Cql3ParserRULE_collectionLiteral)
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)

			var _x = p.ListLiteral()

			localctx.(*CollectionLiteralContext).l = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
			p.Match(Cql3ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)

			var _x = p.Term()

			localctx.(*CollectionLiteralContext).t = _x
		}
		{
			p.SetState(419)

			var _x = p.SetOrMapLiteral()

			localctx.(*CollectionLiteralContext).v = _x
		}
		{
			p.SetState(420)
			p.Match(Cql3ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(422)
			p.Match(Cql3ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(Cql3ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListLiteralContext is an interface to support dynamic dispatch.
type IListLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() ITermContext

	// GetTn returns the tn rule contexts.
	GetTn() ITermContext

	// SetT1 sets the t1 rule contexts.
	SetT1(ITermContext)

	// SetTn sets the tn rule contexts.
	SetTn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsListLiteralContext differentiates from other interfaces.
	IsListLiteralContext()
}

type ListLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t1     ITermContext
	tn     ITermContext
}

func NewEmptyListLiteralContext() *ListLiteralContext {
	var p = new(ListLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_listLiteral
	return p
}

func InitEmptyListLiteralContext(p *ListLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_listLiteral
}

func (*ListLiteralContext) IsListLiteralContext() {}

func NewListLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListLiteralContext {
	var p = new(ListLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_listLiteral

	return p
}

func (s *ListLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ListLiteralContext) GetT1() ITermContext { return s.t1 }

func (s *ListLiteralContext) GetTn() ITermContext { return s.tn }

func (s *ListLiteralContext) SetT1(v ITermContext) { s.t1 = v }

func (s *ListLiteralContext) SetTn(v ITermContext) { s.tn = v }

func (s *ListLiteralContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ListLiteralContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitListLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) ListLiteral() (localctx IListLiteralContext) {
	localctx = NewListLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Cql3ParserRULE_listLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(Cql3ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2266999243331857780) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-8527077135) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&1142788268006301695) != 0) {
		{
			p.SetState(427)

			var _x = p.Term()

			localctx.(*ListLiteralContext).t1 = _x
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Cql3ParserT__3 {
			{
				p.SetState(428)
				p.Match(Cql3ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(429)

				var _x = p.Term()

				localctx.(*ListLiteralContext).tn = _x
			}

			p.SetState(434)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(437)
		p.Match(Cql3ParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUsertypeLiteralContext is an interface to support dynamic dispatch.
type IUsertypeLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK1 returns the k1 rule contexts.
	GetK1() IFidentContext

	// GetV1 returns the v1 rule contexts.
	GetV1() ITermContext

	// GetKn returns the kn rule contexts.
	GetKn() IFidentContext

	// GetVn returns the vn rule contexts.
	GetVn() ITermContext

	// SetK1 sets the k1 rule contexts.
	SetK1(IFidentContext)

	// SetV1 sets the v1 rule contexts.
	SetV1(ITermContext)

	// SetKn sets the kn rule contexts.
	SetKn(IFidentContext)

	// SetVn sets the vn rule contexts.
	SetVn(ITermContext)

	// Getter signatures
	AllFident() []IFidentContext
	Fident(i int) IFidentContext
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsUsertypeLiteralContext differentiates from other interfaces.
	IsUsertypeLiteralContext()
}

type UsertypeLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k1     IFidentContext
	v1     ITermContext
	kn     IFidentContext
	vn     ITermContext
}

func NewEmptyUsertypeLiteralContext() *UsertypeLiteralContext {
	var p = new(UsertypeLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_usertypeLiteral
	return p
}

func InitEmptyUsertypeLiteralContext(p *UsertypeLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_usertypeLiteral
}

func (*UsertypeLiteralContext) IsUsertypeLiteralContext() {}

func NewUsertypeLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UsertypeLiteralContext {
	var p = new(UsertypeLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_usertypeLiteral

	return p
}

func (s *UsertypeLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *UsertypeLiteralContext) GetK1() IFidentContext { return s.k1 }

func (s *UsertypeLiteralContext) GetV1() ITermContext { return s.v1 }

func (s *UsertypeLiteralContext) GetKn() IFidentContext { return s.kn }

func (s *UsertypeLiteralContext) GetVn() ITermContext { return s.vn }

func (s *UsertypeLiteralContext) SetK1(v IFidentContext) { s.k1 = v }

func (s *UsertypeLiteralContext) SetV1(v ITermContext) { s.v1 = v }

func (s *UsertypeLiteralContext) SetKn(v IFidentContext) { s.kn = v }

func (s *UsertypeLiteralContext) SetVn(v ITermContext) { s.vn = v }

func (s *UsertypeLiteralContext) AllFident() []IFidentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFidentContext); ok {
			len++
		}
	}

	tst := make([]IFidentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFidentContext); ok {
			tst[i] = t.(IFidentContext)
			i++
		}
	}

	return tst
}

func (s *UsertypeLiteralContext) Fident(i int) IFidentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFidentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFidentContext)
}

func (s *UsertypeLiteralContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *UsertypeLiteralContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *UsertypeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UsertypeLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UsertypeLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitUsertypeLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) UsertypeLiteral() (localctx IUsertypeLiteralContext) {
	localctx = NewUsertypeLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Cql3ParserRULE_usertypeLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Match(Cql3ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)

		var _x = p.Fident()

		localctx.(*UsertypeLiteralContext).k1 = _x
	}
	{
		p.SetState(441)
		p.Match(Cql3ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)

		var _x = p.Term()

		localctx.(*UsertypeLiteralContext).v1 = _x
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(443)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)

			var _x = p.Fident()

			localctx.(*UsertypeLiteralContext).kn = _x
		}
		{
			p.SetState(445)
			p.Match(Cql3ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)

			var _x = p.Term()

			localctx.(*UsertypeLiteralContext).vn = _x
		}

		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(453)
		p.Match(Cql3ParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITupleLiteralContext is an interface to support dynamic dispatch.
type ITupleLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() ITermContext

	// GetTn returns the tn rule contexts.
	GetTn() ITermContext

	// SetT1 sets the t1 rule contexts.
	SetT1(ITermContext)

	// SetTn sets the tn rule contexts.
	SetTn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsTupleLiteralContext differentiates from other interfaces.
	IsTupleLiteralContext()
}

type TupleLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t1     ITermContext
	tn     ITermContext
}

func NewEmptyTupleLiteralContext() *TupleLiteralContext {
	var p = new(TupleLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tupleLiteral
	return p
}

func InitEmptyTupleLiteralContext(p *TupleLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tupleLiteral
}

func (*TupleLiteralContext) IsTupleLiteralContext() {}

func NewTupleLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleLiteralContext {
	var p = new(TupleLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tupleLiteral

	return p
}

func (s *TupleLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleLiteralContext) GetT1() ITermContext { return s.t1 }

func (s *TupleLiteralContext) GetTn() ITermContext { return s.tn }

func (s *TupleLiteralContext) SetT1(v ITermContext) { s.t1 = v }

func (s *TupleLiteralContext) SetTn(v ITermContext) { s.tn = v }

func (s *TupleLiteralContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *TupleLiteralContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TupleLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTupleLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TupleLiteral() (localctx ITupleLiteralContext) {
	localctx = NewTupleLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Cql3ParserRULE_tupleLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(Cql3ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)

		var _x = p.Term()

		localctx.(*TupleLiteralContext).t1 = _x
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(457)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

			var _x = p.Term()

			localctx.(*TupleLiteralContext).tn = _x
		}

		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(464)
		p.Match(Cql3ParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() IConstantContext

	// GetL returns the l rule contexts.
	GetL() ICollectionLiteralContext

	// GetU returns the u rule contexts.
	GetU() IUsertypeLiteralContext

	// GetT returns the t rule contexts.
	GetT() ITupleLiteralContext

	// GetM returns the m rule contexts.
	GetM() IMarkerContext

	// SetC sets the c rule contexts.
	SetC(IConstantContext)

	// SetL sets the l rule contexts.
	SetL(ICollectionLiteralContext)

	// SetU sets the u rule contexts.
	SetU(IUsertypeLiteralContext)

	// SetT sets the t rule contexts.
	SetT(ITupleLiteralContext)

	// SetM sets the m rule contexts.
	SetM(IMarkerContext)

	// Getter signatures
	Constant() IConstantContext
	CollectionLiteral() ICollectionLiteralContext
	UsertypeLiteral() IUsertypeLiteralContext
	TupleLiteral() ITupleLiteralContext
	K_NULL() antlr.TerminalNode
	Marker() IMarkerContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	c      IConstantContext
	l      ICollectionLiteralContext
	u      IUsertypeLiteralContext
	t      ITupleLiteralContext
	m      IMarkerContext
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) GetC() IConstantContext { return s.c }

func (s *ValueContext) GetL() ICollectionLiteralContext { return s.l }

func (s *ValueContext) GetU() IUsertypeLiteralContext { return s.u }

func (s *ValueContext) GetT() ITupleLiteralContext { return s.t }

func (s *ValueContext) GetM() IMarkerContext { return s.m }

func (s *ValueContext) SetC(v IConstantContext) { s.c = v }

func (s *ValueContext) SetL(v ICollectionLiteralContext) { s.l = v }

func (s *ValueContext) SetU(v IUsertypeLiteralContext) { s.u = v }

func (s *ValueContext) SetT(v ITupleLiteralContext) { s.t = v }

func (s *ValueContext) SetM(v IMarkerContext) { s.m = v }

func (s *ValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) CollectionLiteral() ICollectionLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionLiteralContext)
}

func (s *ValueContext) UsertypeLiteral() IUsertypeLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUsertypeLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUsertypeLiteralContext)
}

func (s *ValueContext) TupleLiteral() ITupleLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleLiteralContext)
}

func (s *ValueContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_NULL, 0)
}

func (s *ValueContext) Marker() IMarkerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarkerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarkerContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Cql3ParserRULE_value)
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)

			var _x = p.Constant()

			localctx.(*ValueContext).c = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(467)

			var _x = p.CollectionLiteral()

			localctx.(*ValueContext).l = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(468)

			var _x = p.UsertypeLiteral()

			localctx.(*ValueContext).u = _x
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(469)

			var _x = p.TupleLiteral()

			localctx.(*ValueContext).t = _x
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(470)
			p.Match(Cql3ParserK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(471)

			var _x = p.Marker()

			localctx.(*ValueContext).m = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKs returns the ks rule contexts.
	GetKs() IKeyspaceNameContext

	// GetF returns the f rule contexts.
	GetF() IAllowedFunctionNameContext

	// SetKs sets the ks rule contexts.
	SetKs(IKeyspaceNameContext)

	// SetF sets the f rule contexts.
	SetF(IAllowedFunctionNameContext)

	// Getter signatures
	AllowedFunctionName() IAllowedFunctionNameContext
	KeyspaceName() IKeyspaceNameContext

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ks     IKeyspaceNameContext
	f      IAllowedFunctionNameContext
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_functionName
	return p
}

func InitEmptyFunctionNameContext(p *FunctionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_functionName
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) GetKs() IKeyspaceNameContext { return s.ks }

func (s *FunctionNameContext) GetF() IAllowedFunctionNameContext { return s.f }

func (s *FunctionNameContext) SetKs(v IKeyspaceNameContext) { s.ks = v }

func (s *FunctionNameContext) SetF(v IAllowedFunctionNameContext) { s.f = v }

func (s *FunctionNameContext) AllowedFunctionName() IAllowedFunctionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllowedFunctionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllowedFunctionNameContext)
}

func (s *FunctionNameContext) KeyspaceName() IKeyspaceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyspaceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyspaceNameContext)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitFunctionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Cql3ParserRULE_functionName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(477)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(474)

			var _x = p.KeyspaceName()

			localctx.(*FunctionNameContext).ks = _x
		}
		{
			p.SetState(475)
			p.Match(Cql3ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(479)

		var _x = p.AllowedFunctionName()

		localctx.(*FunctionNameContext).f = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllowedFunctionNameContext is an interface to support dynamic dispatch.
type IAllowedFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAllowedFunctionNameContext differentiates from other interfaces.
	IsAllowedFunctionNameContext()
}

type AllowedFunctionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowedFunctionNameContext() *AllowedFunctionNameContext {
	var p = new(AllowedFunctionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_allowedFunctionName
	return p
}

func InitEmptyAllowedFunctionNameContext(p *AllowedFunctionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_allowedFunctionName
}

func (*AllowedFunctionNameContext) IsAllowedFunctionNameContext() {}

func NewAllowedFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowedFunctionNameContext {
	var p = new(AllowedFunctionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_allowedFunctionName

	return p
}

func (s *AllowedFunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowedFunctionNameContext) CopyAll(ctx *AllowedFunctionNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AllowedFunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowedFunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllowedFunctionNameIdentContext struct {
	AllowedFunctionNameContext
	f antlr.Token
	u IUnreserved_function_keywordContext
}

func NewAllowedFunctionNameIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllowedFunctionNameIdentContext {
	var p = new(AllowedFunctionNameIdentContext)

	InitEmptyAllowedFunctionNameContext(&p.AllowedFunctionNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*AllowedFunctionNameContext))

	return p
}

func (s *AllowedFunctionNameIdentContext) GetF() antlr.Token { return s.f }

func (s *AllowedFunctionNameIdentContext) SetF(v antlr.Token) { s.f = v }

func (s *AllowedFunctionNameIdentContext) GetU() IUnreserved_function_keywordContext { return s.u }

func (s *AllowedFunctionNameIdentContext) SetU(v IUnreserved_function_keywordContext) { s.u = v }

func (s *AllowedFunctionNameIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowedFunctionNameIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *AllowedFunctionNameIdentContext) Unreserved_function_keyword() IUnreserved_function_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_function_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_function_keywordContext)
}

func (s *AllowedFunctionNameIdentContext) K_TOKEN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TOKEN, 0)
}

func (s *AllowedFunctionNameIdentContext) K_COUNT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COUNT, 0)
}

func (s *AllowedFunctionNameIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitAllowedFunctionNameIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllowedFunctionNameQuotedIdentContext struct {
	AllowedFunctionNameContext
	f antlr.Token
}

func NewAllowedFunctionNameQuotedIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllowedFunctionNameQuotedIdentContext {
	var p = new(AllowedFunctionNameQuotedIdentContext)

	InitEmptyAllowedFunctionNameContext(&p.AllowedFunctionNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*AllowedFunctionNameContext))

	return p
}

func (s *AllowedFunctionNameQuotedIdentContext) GetF() antlr.Token { return s.f }

func (s *AllowedFunctionNameQuotedIdentContext) SetF(v antlr.Token) { s.f = v }

func (s *AllowedFunctionNameQuotedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowedFunctionNameQuotedIdentContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *AllowedFunctionNameQuotedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitAllowedFunctionNameQuotedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) AllowedFunctionName() (localctx IAllowedFunctionNameContext) {
	localctx = NewAllowedFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Cql3ParserRULE_allowedFunctionName)
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserIDENT:
		localctx = NewAllowedFunctionNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)

			var _m = p.Match(Cql3ParserIDENT)

			localctx.(*AllowedFunctionNameIdentContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewAllowedFunctionNameQuotedIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(482)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*AllowedFunctionNameQuotedIdentContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewAllowedFunctionNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(483)

			var _x = p.Unreserved_function_keyword()

			localctx.(*AllowedFunctionNameIdentContext).u = _x
		}

	case Cql3ParserK_TOKEN:
		localctx = NewAllowedFunctionNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(484)
			p.Match(Cql3ParserK_TOKEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_COUNT:
		localctx = NewAllowedFunctionNameIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(485)
			p.Match(Cql3ParserK_COUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f rule contexts.
	GetF() IFunctionNameContext

	// GetArgs returns the args rule contexts.
	GetArgs() IFunctionArgsContext

	// SetF sets the f rule contexts.
	SetF(IFunctionNameContext)

	// SetArgs sets the args rule contexts.
	SetArgs(IFunctionArgsContext)

	// Getter signatures
	FunctionName() IFunctionNameContext
	FunctionArgs() IFunctionArgsContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	f      IFunctionNameContext
	args   IFunctionArgsContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetF() IFunctionNameContext { return s.f }

func (s *FunctionContext) GetArgs() IFunctionArgsContext { return s.args }

func (s *FunctionContext) SetF(v IFunctionNameContext) { s.f = v }

func (s *FunctionContext) SetArgs(v IFunctionArgsContext) { s.args = v }

func (s *FunctionContext) FunctionName() IFunctionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Cql3ParserRULE_function)
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(488)

			var _x = p.FunctionName()

			localctx.(*FunctionContext).f = _x
		}
		{
			p.SetState(489)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)

			var _x = p.FunctionName()

			localctx.(*FunctionContext).f = _x
		}
		{
			p.SetState(493)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)

			var _x = p.FunctionArgs()

			localctx.(*FunctionContext).args = _x
		}
		{
			p.SetState(495)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() ITermContext

	// GetTn returns the tn rule contexts.
	GetTn() ITermContext

	// SetT1 sets the t1 rule contexts.
	SetT1(ITermContext)

	// SetTn sets the tn rule contexts.
	SetTn(ITermContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t1     ITermContext
	tn     ITermContext
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_functionArgs
	return p
}

func InitEmptyFunctionArgsContext(p *FunctionArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_functionArgs
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) GetT1() ITermContext { return s.t1 }

func (s *FunctionArgsContext) GetTn() ITermContext { return s.tn }

func (s *FunctionArgsContext) SetT1(v ITermContext) { s.t1 = v }

func (s *FunctionArgsContext) SetTn(v ITermContext) { s.tn = v }

func (s *FunctionArgsContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Cql3ParserRULE_functionArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)

		var _x = p.Term()

		localctx.(*FunctionArgsContext).t1 = _x
	}
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(500)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)

			var _x = p.Term()

			localctx.(*FunctionArgsContext).tn = _x
		}

		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t rule contexts.
	GetT() ITermAdditionContext

	// SetT sets the t rule contexts.
	SetT(ITermAdditionContext)

	// Getter signatures
	TermAddition() ITermAdditionContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t      ITermAdditionContext
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) GetT() ITermAdditionContext { return s.t }

func (s *TermContext) SetT(v ITermAdditionContext) { s.t = v }

func (s *TermContext) TermAddition() ITermAdditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermAdditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermAdditionContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, Cql3ParserRULE_term)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)

		var _x = p.TermAddition()

		localctx.(*TermContext).t = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermAdditionContext is an interface to support dynamic dispatch.
type ITermAdditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ITermMultiplicationContext

	// GetR returns the r rule contexts.
	GetR() ITermMultiplicationContext

	// SetL sets the l rule contexts.
	SetL(ITermMultiplicationContext)

	// SetR sets the r rule contexts.
	SetR(ITermMultiplicationContext)

	// Getter signatures
	AllTermMultiplication() []ITermMultiplicationContext
	TermMultiplication(i int) ITermMultiplicationContext

	// IsTermAdditionContext differentiates from other interfaces.
	IsTermAdditionContext()
}

type TermAdditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      ITermMultiplicationContext
	r      ITermMultiplicationContext
}

func NewEmptyTermAdditionContext() *TermAdditionContext {
	var p = new(TermAdditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_termAddition
	return p
}

func InitEmptyTermAdditionContext(p *TermAdditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_termAddition
}

func (*TermAdditionContext) IsTermAdditionContext() {}

func NewTermAdditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermAdditionContext {
	var p = new(TermAdditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_termAddition

	return p
}

func (s *TermAdditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TermAdditionContext) GetL() ITermMultiplicationContext { return s.l }

func (s *TermAdditionContext) GetR() ITermMultiplicationContext { return s.r }

func (s *TermAdditionContext) SetL(v ITermMultiplicationContext) { s.l = v }

func (s *TermAdditionContext) SetR(v ITermMultiplicationContext) { s.r = v }

func (s *TermAdditionContext) AllTermMultiplication() []ITermMultiplicationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermMultiplicationContext); ok {
			len++
		}
	}

	tst := make([]ITermMultiplicationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermMultiplicationContext); ok {
			tst[i] = t.(ITermMultiplicationContext)
			i++
		}
	}

	return tst
}

func (s *TermAdditionContext) TermMultiplication(i int) ITermMultiplicationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermMultiplicationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermMultiplicationContext)
}

func (s *TermAdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermAdditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermAdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTermAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TermAddition() (localctx ITermAdditionContext) {
	localctx = NewTermAdditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, Cql3ParserRULE_termAddition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)

		var _x = p.TermMultiplication()

		localctx.(*TermAdditionContext).l = _x
	}
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__10 || _la == Cql3ParserT__11 {
		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Cql3ParserT__10:
			{
				p.SetState(510)
				p.Match(Cql3ParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(511)

				var _x = p.TermMultiplication()

				localctx.(*TermAdditionContext).r = _x
			}

		case Cql3ParserT__11:
			{
				p.SetState(512)
				p.Match(Cql3ParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(513)

				var _x = p.TermMultiplication()

				localctx.(*TermAdditionContext).r = _x
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(518)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermMultiplicationContext is an interface to support dynamic dispatch.
type ITermMultiplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ITermGroupContext

	// GetR returns the r rule contexts.
	GetR() ITermGroupContext

	// SetL sets the l rule contexts.
	SetL(ITermGroupContext)

	// SetR sets the r rule contexts.
	SetR(ITermGroupContext)

	// Getter signatures
	AllTermGroup() []ITermGroupContext
	TermGroup(i int) ITermGroupContext

	// IsTermMultiplicationContext differentiates from other interfaces.
	IsTermMultiplicationContext()
}

type TermMultiplicationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      ITermGroupContext
	r      ITermGroupContext
}

func NewEmptyTermMultiplicationContext() *TermMultiplicationContext {
	var p = new(TermMultiplicationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_termMultiplication
	return p
}

func InitEmptyTermMultiplicationContext(p *TermMultiplicationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_termMultiplication
}

func (*TermMultiplicationContext) IsTermMultiplicationContext() {}

func NewTermMultiplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermMultiplicationContext {
	var p = new(TermMultiplicationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_termMultiplication

	return p
}

func (s *TermMultiplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *TermMultiplicationContext) GetL() ITermGroupContext { return s.l }

func (s *TermMultiplicationContext) GetR() ITermGroupContext { return s.r }

func (s *TermMultiplicationContext) SetL(v ITermGroupContext) { s.l = v }

func (s *TermMultiplicationContext) SetR(v ITermGroupContext) { s.r = v }

func (s *TermMultiplicationContext) AllTermGroup() []ITermGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermGroupContext); ok {
			len++
		}
	}

	tst := make([]ITermGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermGroupContext); ok {
			tst[i] = t.(ITermGroupContext)
			i++
		}
	}

	return tst
}

func (s *TermMultiplicationContext) TermGroup(i int) ITermGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermGroupContext)
}

func (s *TermMultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermMultiplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermMultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTermMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TermMultiplication() (localctx ITermMultiplicationContext) {
	localctx = NewTermMultiplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, Cql3ParserRULE_termMultiplication)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)

		var _x = p.TermGroup()

		localctx.(*TermMultiplicationContext).l = _x
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0 {
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Cql3ParserT__12:
			{
				p.SetState(520)
				p.Match(Cql3ParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(521)

				var _x = p.TermGroup()

				localctx.(*TermMultiplicationContext).r = _x
			}

		case Cql3ParserT__13:
			{
				p.SetState(522)
				p.Match(Cql3ParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(523)

				var _x = p.TermGroup()

				localctx.(*TermMultiplicationContext).r = _x
			}

		case Cql3ParserT__14:
			{
				p.SetState(524)
				p.Match(Cql3ParserT__14)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(525)

				var _x = p.TermGroup()

				localctx.(*TermMultiplicationContext).r = _x
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermGroupContext is an interface to support dynamic dispatch.
type ITermGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t rule contexts.
	GetT() ISimpleTermContext

	// SetT sets the t rule contexts.
	SetT(ISimpleTermContext)

	// Getter signatures
	SimpleTerm() ISimpleTermContext

	// IsTermGroupContext differentiates from other interfaces.
	IsTermGroupContext()
}

type TermGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t      ISimpleTermContext
}

func NewEmptyTermGroupContext() *TermGroupContext {
	var p = new(TermGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_termGroup
	return p
}

func InitEmptyTermGroupContext(p *TermGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_termGroup
}

func (*TermGroupContext) IsTermGroupContext() {}

func NewTermGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermGroupContext {
	var p = new(TermGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_termGroup

	return p
}

func (s *TermGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *TermGroupContext) GetT() ISimpleTermContext { return s.t }

func (s *TermGroupContext) SetT(v ISimpleTermContext) { s.t = v }

func (s *TermGroupContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *TermGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTermGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) TermGroup() (localctx ITermGroupContext) {
	localctx = NewTermGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, Cql3ParserRULE_termGroup)
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserT__1, Cql3ParserT__2, Cql3ParserT__6, Cql3ParserT__8, Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_TOKEN, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_NULL, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_POSITIVE_NAN, Cql3ParserK_NEGATIVE_NAN, Cql3ParserK_POSITIVE_INFINITY, Cql3ParserK_NEGATIVE_INFINITY, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN, Cql3ParserSTRING_LITERAL, Cql3ParserQUOTED_NAME, Cql3ParserINTEGER, Cql3ParserQMARK, Cql3ParserFLOAT, Cql3ParserBOOLEAN, Cql3ParserDURATION, Cql3ParserIDENT, Cql3ParserHEXNUMBER, Cql3ParserUUID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)

			var _x = p.SimpleTerm()

			localctx.(*TermGroupContext).t = _x
		}

	case Cql3ParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.Match(Cql3ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)

			var _x = p.SimpleTerm()

			localctx.(*TermGroupContext).t = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleTermContext is an interface to support dynamic dispatch.
type ISimpleTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v rule contexts.
	GetV() IValueContext

	// GetF returns the f rule contexts.
	GetF() IFunctionContext

	// GetC returns the c rule contexts.
	GetC() IComparatorTypeContext

	// GetT returns the t rule contexts.
	GetT() ISimpleTermContext

	// GetN returns the n rule contexts.
	GetN() INative_typeContext

	// SetV sets the v rule contexts.
	SetV(IValueContext)

	// SetF sets the f rule contexts.
	SetF(IFunctionContext)

	// SetC sets the c rule contexts.
	SetC(IComparatorTypeContext)

	// SetT sets the t rule contexts.
	SetT(ISimpleTermContext)

	// SetN sets the n rule contexts.
	SetN(INative_typeContext)

	// Getter signatures
	Value() IValueContext
	Function() IFunctionContext
	ComparatorType() IComparatorTypeContext
	SimpleTerm() ISimpleTermContext
	K_CAST() antlr.TerminalNode
	K_AS() antlr.TerminalNode
	Native_type() INative_typeContext

	// IsSimpleTermContext differentiates from other interfaces.
	IsSimpleTermContext()
}

type SimpleTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	v      IValueContext
	f      IFunctionContext
	c      IComparatorTypeContext
	t      ISimpleTermContext
	n      INative_typeContext
}

func NewEmptySimpleTermContext() *SimpleTermContext {
	var p = new(SimpleTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_simpleTerm
	return p
}

func InitEmptySimpleTermContext(p *SimpleTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_simpleTerm
}

func (*SimpleTermContext) IsSimpleTermContext() {}

func NewSimpleTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermContext {
	var p = new(SimpleTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_simpleTerm

	return p
}

func (s *SimpleTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermContext) GetV() IValueContext { return s.v }

func (s *SimpleTermContext) GetF() IFunctionContext { return s.f }

func (s *SimpleTermContext) GetC() IComparatorTypeContext { return s.c }

func (s *SimpleTermContext) GetT() ISimpleTermContext { return s.t }

func (s *SimpleTermContext) GetN() INative_typeContext { return s.n }

func (s *SimpleTermContext) SetV(v IValueContext) { s.v = v }

func (s *SimpleTermContext) SetF(v IFunctionContext) { s.f = v }

func (s *SimpleTermContext) SetC(v IComparatorTypeContext) { s.c = v }

func (s *SimpleTermContext) SetT(v ISimpleTermContext) { s.t = v }

func (s *SimpleTermContext) SetN(v INative_typeContext) { s.n = v }

func (s *SimpleTermContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SimpleTermContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *SimpleTermContext) ComparatorType() IComparatorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *SimpleTermContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *SimpleTermContext) K_CAST() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CAST, 0)
}

func (s *SimpleTermContext) K_AS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_AS, 0)
}

func (s *SimpleTermContext) Native_type() INative_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INative_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INative_typeContext)
}

func (s *SimpleTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitSimpleTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) SimpleTerm() (localctx ISimpleTermContext) {
	localctx = NewSimpleTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, Cql3ParserRULE_simpleTerm)
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(536)

			var _x = p.Value()

			localctx.(*SimpleTermContext).v = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)

			var _x = p.Function()

			localctx.(*SimpleTermContext).f = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(538)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)

			var _x = p.ComparatorType()

			localctx.(*SimpleTermContext).c = _x
		}
		{
			p.SetState(540)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(541)

			var _x = p.SimpleTerm()

			localctx.(*SimpleTermContext).t = _x
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(543)
			p.Match(Cql3ParserK_CAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.Match(Cql3ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)

			var _x = p.SimpleTerm()

			localctx.(*SimpleTermContext).t = _x
		}
		{
			p.SetState(546)
			p.Match(Cql3ParserK_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)

			var _x = p.Native_type()

			localctx.(*SimpleTermContext).n = _x
		}
		{
			p.SetState(548)
			p.Match(Cql3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertiesContext is an interface to support dynamic dispatch.
type IPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProperty() []IPropertyContext
	Property(i int) IPropertyContext
	AllK_AND() []antlr.TerminalNode
	K_AND(i int) antlr.TerminalNode

	// IsPropertiesContext differentiates from other interfaces.
	IsPropertiesContext()
}

type PropertiesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesContext() *PropertiesContext {
	var p = new(PropertiesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_properties
	return p
}

func InitEmptyPropertiesContext(p *PropertiesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_properties
}

func (*PropertiesContext) IsPropertiesContext() {}

func NewPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesContext {
	var p = new(PropertiesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_properties

	return p
}

func (s *PropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesContext) AllProperty() []IPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyContext); ok {
			len++
		}
	}

	tst := make([]IPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyContext); ok {
			tst[i] = t.(IPropertyContext)
			i++
		}
	}

	return tst
}

func (s *PropertiesContext) Property(i int) IPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertiesContext) AllK_AND() []antlr.TerminalNode {
	return s.GetTokens(Cql3ParserK_AND)
}

func (s *PropertiesContext) K_AND(i int) antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_AND, i)
}

func (s *PropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitProperties(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Properties() (localctx IPropertiesContext) {
	localctx = NewPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, Cql3ParserRULE_properties)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(552)
		p.Property()
	}
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserK_AND {
		{
			p.SetState(553)
			p.Match(Cql3ParserK_AND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)
			p.Property()
		}

		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k rule contexts.
	GetK() INoncol_identContext

	// GetSimple returns the simple rule contexts.
	GetSimple() IPropertyValueContext

	// GetMap_ returns the map_ rule contexts.
	GetMap_() IFullMapLiteralContext

	// SetK sets the k rule contexts.
	SetK(INoncol_identContext)

	// SetSimple sets the simple rule contexts.
	SetSimple(IPropertyValueContext)

	// SetMap_ sets the map_ rule contexts.
	SetMap_(IFullMapLiteralContext)

	// Getter signatures
	Noncol_ident() INoncol_identContext
	PropertyValue() IPropertyValueContext
	FullMapLiteral() IFullMapLiteralContext

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k      INoncol_identContext
	simple IPropertyValueContext
	map_   IFullMapLiteralContext
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_property
	return p
}

func InitEmptyPropertyContext(p *PropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_property
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) GetK() INoncol_identContext { return s.k }

func (s *PropertyContext) GetSimple() IPropertyValueContext { return s.simple }

func (s *PropertyContext) GetMap_() IFullMapLiteralContext { return s.map_ }

func (s *PropertyContext) SetK(v INoncol_identContext) { s.k = v }

func (s *PropertyContext) SetSimple(v IPropertyValueContext) { s.simple = v }

func (s *PropertyContext) SetMap_(v IFullMapLiteralContext) { s.map_ = v }

func (s *PropertyContext) Noncol_ident() INoncol_identContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoncol_identContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoncol_identContext)
}

func (s *PropertyContext) PropertyValue() IPropertyValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *PropertyContext) FullMapLiteral() IFullMapLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullMapLiteralContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, Cql3ParserRULE_property)
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(560)

			var _x = p.Noncol_ident()

			localctx.(*PropertyContext).k = _x
		}
		{
			p.SetState(561)
			p.Match(Cql3ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)

			var _x = p.PropertyValue()

			localctx.(*PropertyContext).simple = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(564)

			var _x = p.Noncol_ident()

			localctx.(*PropertyContext).k = _x
		}
		{
			p.SetState(565)
			p.Match(Cql3ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)

			var _x = p.FullMapLiteral()

			localctx.(*PropertyContext).map_ = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyValueContext is an interface to support dynamic dispatch.
type IPropertyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() IConstantContext

	// GetU returns the u rule contexts.
	GetU() IUnreserved_keywordContext

	// SetC sets the c rule contexts.
	SetC(IConstantContext)

	// SetU sets the u rule contexts.
	SetU(IUnreserved_keywordContext)

	// Getter signatures
	Constant() IConstantContext
	Unreserved_keyword() IUnreserved_keywordContext

	// IsPropertyValueContext differentiates from other interfaces.
	IsPropertyValueContext()
}

type PropertyValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	c      IConstantContext
	u      IUnreserved_keywordContext
}

func NewEmptyPropertyValueContext() *PropertyValueContext {
	var p = new(PropertyValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_propertyValue
	return p
}

func InitEmptyPropertyValueContext(p *PropertyValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_propertyValue
}

func (*PropertyValueContext) IsPropertyValueContext() {}

func NewPropertyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyValueContext {
	var p = new(PropertyValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_propertyValue

	return p
}

func (s *PropertyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyValueContext) GetC() IConstantContext { return s.c }

func (s *PropertyValueContext) GetU() IUnreserved_keywordContext { return s.u }

func (s *PropertyValueContext) SetC(v IConstantContext) { s.c = v }

func (s *PropertyValueContext) SetU(v IUnreserved_keywordContext) { s.u = v }

func (s *PropertyValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PropertyValueContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *PropertyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitPropertyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) PropertyValue() (localctx IPropertyValueContext) {
	localctx = NewPropertyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, Cql3ParserRULE_propertyValue)
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_POSITIVE_NAN, Cql3ParserK_NEGATIVE_NAN, Cql3ParserK_POSITIVE_INFINITY, Cql3ParserK_NEGATIVE_INFINITY, Cql3ParserSTRING_LITERAL, Cql3ParserINTEGER, Cql3ParserFLOAT, Cql3ParserBOOLEAN, Cql3ParserDURATION, Cql3ParserHEXNUMBER, Cql3ParserUUID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(570)

			var _x = p.Constant()

			localctx.(*PropertyValueContext).c = _x
		}

	case Cql3ParserK_AS, Cql3ParserK_KEY, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_JSON, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(571)

			var _x = p.Unreserved_keyword()

			localctx.(*PropertyValueContext).u = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparatorTypeContext is an interface to support dynamic dispatch.
type IComparatorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS returns the s token.
	GetS() antlr.Token

	// SetS sets the s token.
	SetS(antlr.Token)

	// GetN returns the n rule contexts.
	GetN() INative_typeContext

	// GetC returns the c rule contexts.
	GetC() ICollection_typeContext

	// GetTt returns the tt rule contexts.
	GetTt() ITuple_typeContext

	// GetVc returns the vc rule contexts.
	GetVc() IVector_typeContext

	// GetId returns the id rule contexts.
	GetId() IUserTypeNameContext

	// GetF returns the f rule contexts.
	GetF() IComparatorTypeContext

	// SetN sets the n rule contexts.
	SetN(INative_typeContext)

	// SetC sets the c rule contexts.
	SetC(ICollection_typeContext)

	// SetTt sets the tt rule contexts.
	SetTt(ITuple_typeContext)

	// SetVc sets the vc rule contexts.
	SetVc(IVector_typeContext)

	// SetId sets the id rule contexts.
	SetId(IUserTypeNameContext)

	// SetF sets the f rule contexts.
	SetF(IComparatorTypeContext)

	// Getter signatures
	Native_type() INative_typeContext
	Collection_type() ICollection_typeContext
	Tuple_type() ITuple_typeContext
	Vector_type() IVector_typeContext
	UserTypeName() IUserTypeNameContext
	K_FROZEN() antlr.TerminalNode
	ComparatorType() IComparatorTypeContext
	STRING_LITERAL() antlr.TerminalNode

	// IsComparatorTypeContext differentiates from other interfaces.
	IsComparatorTypeContext()
}

type ComparatorTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	n      INative_typeContext
	c      ICollection_typeContext
	tt     ITuple_typeContext
	vc     IVector_typeContext
	id     IUserTypeNameContext
	f      IComparatorTypeContext
	s      antlr.Token
}

func NewEmptyComparatorTypeContext() *ComparatorTypeContext {
	var p = new(ComparatorTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_comparatorType
	return p
}

func InitEmptyComparatorTypeContext(p *ComparatorTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_comparatorType
}

func (*ComparatorTypeContext) IsComparatorTypeContext() {}

func NewComparatorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorTypeContext {
	var p = new(ComparatorTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_comparatorType

	return p
}

func (s *ComparatorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorTypeContext) GetS() antlr.Token { return s.s }

func (s *ComparatorTypeContext) SetS(v antlr.Token) { s.s = v }

func (s *ComparatorTypeContext) GetN() INative_typeContext { return s.n }

func (s *ComparatorTypeContext) GetC() ICollection_typeContext { return s.c }

func (s *ComparatorTypeContext) GetTt() ITuple_typeContext { return s.tt }

func (s *ComparatorTypeContext) GetVc() IVector_typeContext { return s.vc }

func (s *ComparatorTypeContext) GetId() IUserTypeNameContext { return s.id }

func (s *ComparatorTypeContext) GetF() IComparatorTypeContext { return s.f }

func (s *ComparatorTypeContext) SetN(v INative_typeContext) { s.n = v }

func (s *ComparatorTypeContext) SetC(v ICollection_typeContext) { s.c = v }

func (s *ComparatorTypeContext) SetTt(v ITuple_typeContext) { s.tt = v }

func (s *ComparatorTypeContext) SetVc(v IVector_typeContext) { s.vc = v }

func (s *ComparatorTypeContext) SetId(v IUserTypeNameContext) { s.id = v }

func (s *ComparatorTypeContext) SetF(v IComparatorTypeContext) { s.f = v }

func (s *ComparatorTypeContext) Native_type() INative_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INative_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INative_typeContext)
}

func (s *ComparatorTypeContext) Collection_type() ICollection_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_typeContext)
}

func (s *ComparatorTypeContext) Tuple_type() ITuple_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITuple_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITuple_typeContext)
}

func (s *ComparatorTypeContext) Vector_type() IVector_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_typeContext)
}

func (s *ComparatorTypeContext) UserTypeName() IUserTypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserTypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserTypeNameContext)
}

func (s *ComparatorTypeContext) K_FROZEN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FROZEN, 0)
}

func (s *ComparatorTypeContext) ComparatorType() IComparatorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *ComparatorTypeContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(Cql3ParserSTRING_LITERAL, 0)
}

func (s *ComparatorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitComparatorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) ComparatorType() (localctx IComparatorTypeContext) {
	localctx = NewComparatorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, Cql3ParserRULE_comparatorType)
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(574)

			var _x = p.Native_type()

			localctx.(*ComparatorTypeContext).n = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(575)

			var _x = p.Collection_type()

			localctx.(*ComparatorTypeContext).c = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(576)

			var _x = p.Tuple_type()

			localctx.(*ComparatorTypeContext).tt = _x
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(577)

			var _x = p.Vector_type()

			localctx.(*ComparatorTypeContext).vc = _x
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(578)

			var _x = p.UserTypeName()

			localctx.(*ComparatorTypeContext).id = _x
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(579)
			p.Match(Cql3ParserK_FROZEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(580)
			p.Match(Cql3ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)

			var _x = p.ComparatorType()

			localctx.(*ComparatorTypeContext).f = _x
		}
		{
			p.SetState(582)
			p.Match(Cql3ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(584)

			var _m = p.Match(Cql3ParserSTRING_LITERAL)

			localctx.(*ComparatorTypeContext).s = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INative_typeContext is an interface to support dynamic dispatch.
type INative_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ASCII() antlr.TerminalNode
	K_BIGINT() antlr.TerminalNode
	K_BLOB() antlr.TerminalNode
	K_BOOLEAN() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_DECIMAL() antlr.TerminalNode
	K_DOUBLE() antlr.TerminalNode
	K_DURATION() antlr.TerminalNode
	K_FLOAT() antlr.TerminalNode
	K_INET() antlr.TerminalNode
	K_INT() antlr.TerminalNode
	K_SMALLINT() antlr.TerminalNode
	K_TEXT() antlr.TerminalNode
	K_TIMESTAMP() antlr.TerminalNode
	K_TINYINT() antlr.TerminalNode
	K_UUID() antlr.TerminalNode
	K_VARCHAR() antlr.TerminalNode
	K_VARINT() antlr.TerminalNode
	K_TIMEUUID() antlr.TerminalNode
	K_DATE() antlr.TerminalNode
	K_TIME() antlr.TerminalNode

	// IsNative_typeContext differentiates from other interfaces.
	IsNative_typeContext()
}

type Native_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNative_typeContext() *Native_typeContext {
	var p = new(Native_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_native_type
	return p
}

func InitEmptyNative_typeContext(p *Native_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_native_type
}

func (*Native_typeContext) IsNative_typeContext() {}

func NewNative_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Native_typeContext {
	var p = new(Native_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_native_type

	return p
}

func (s *Native_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Native_typeContext) K_ASCII() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ASCII, 0)
}

func (s *Native_typeContext) K_BIGINT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_BIGINT, 0)
}

func (s *Native_typeContext) K_BLOB() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_BLOB, 0)
}

func (s *Native_typeContext) K_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_BOOLEAN, 0)
}

func (s *Native_typeContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COUNTER, 0)
}

func (s *Native_typeContext) K_DECIMAL() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DECIMAL, 0)
}

func (s *Native_typeContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DOUBLE, 0)
}

func (s *Native_typeContext) K_DURATION() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DURATION, 0)
}

func (s *Native_typeContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FLOAT, 0)
}

func (s *Native_typeContext) K_INET() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_INET, 0)
}

func (s *Native_typeContext) K_INT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_INT, 0)
}

func (s *Native_typeContext) K_SMALLINT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_SMALLINT, 0)
}

func (s *Native_typeContext) K_TEXT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TEXT, 0)
}

func (s *Native_typeContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TIMESTAMP, 0)
}

func (s *Native_typeContext) K_TINYINT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TINYINT, 0)
}

func (s *Native_typeContext) K_UUID() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_UUID, 0)
}

func (s *Native_typeContext) K_VARCHAR() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_VARCHAR, 0)
}

func (s *Native_typeContext) K_VARINT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_VARINT, 0)
}

func (s *Native_typeContext) K_TIMEUUID() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TIMEUUID, 0)
}

func (s *Native_typeContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DATE, 0)
}

func (s *Native_typeContext) K_TIME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TIME, 0)
}

func (s *Native_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Native_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Native_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitNative_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Native_type() (localctx INative_typeContext) {
	localctx = NewNative_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, Cql3ParserRULE_native_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Cql3ParserK_TIMESTAMP || ((int64((_la-117)) & ^0x3f) == 0 && ((int64(1)<<(_la-117))&6553599) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollection_typeContext is an interface to support dynamic dispatch.
type ICollection_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() IComparatorTypeContext

	// GetT2 returns the t2 rule contexts.
	GetT2() IComparatorTypeContext

	// GetT returns the t rule contexts.
	GetT() IComparatorTypeContext

	// SetT1 sets the t1 rule contexts.
	SetT1(IComparatorTypeContext)

	// SetT2 sets the t2 rule contexts.
	SetT2(IComparatorTypeContext)

	// SetT sets the t rule contexts.
	SetT(IComparatorTypeContext)

	// Getter signatures
	K_MAP() antlr.TerminalNode
	AllComparatorType() []IComparatorTypeContext
	ComparatorType(i int) IComparatorTypeContext
	K_LIST() antlr.TerminalNode
	K_SET() antlr.TerminalNode

	// IsCollection_typeContext differentiates from other interfaces.
	IsCollection_typeContext()
}

type Collection_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t1     IComparatorTypeContext
	t2     IComparatorTypeContext
	t      IComparatorTypeContext
}

func NewEmptyCollection_typeContext() *Collection_typeContext {
	var p = new(Collection_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_collection_type
	return p
}

func InitEmptyCollection_typeContext(p *Collection_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_collection_type
}

func (*Collection_typeContext) IsCollection_typeContext() {}

func NewCollection_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_typeContext {
	var p = new(Collection_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_collection_type

	return p
}

func (s *Collection_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_typeContext) GetT1() IComparatorTypeContext { return s.t1 }

func (s *Collection_typeContext) GetT2() IComparatorTypeContext { return s.t2 }

func (s *Collection_typeContext) GetT() IComparatorTypeContext { return s.t }

func (s *Collection_typeContext) SetT1(v IComparatorTypeContext) { s.t1 = v }

func (s *Collection_typeContext) SetT2(v IComparatorTypeContext) { s.t2 = v }

func (s *Collection_typeContext) SetT(v IComparatorTypeContext) { s.t = v }

func (s *Collection_typeContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MAP, 0)
}

func (s *Collection_typeContext) AllComparatorType() []IComparatorTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			len++
		}
	}

	tst := make([]IComparatorTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparatorTypeContext); ok {
			tst[i] = t.(IComparatorTypeContext)
			i++
		}
	}

	return tst
}

func (s *Collection_typeContext) ComparatorType(i int) IComparatorTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *Collection_typeContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_LIST, 0)
}

func (s *Collection_typeContext) K_SET() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_SET, 0)
}

func (s *Collection_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collection_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitCollection_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Collection_type() (localctx ICollection_typeContext) {
	localctx = NewCollection_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, Cql3ParserRULE_collection_type)
	p.SetState(606)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_MAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(589)
			p.Match(Cql3ParserK_MAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.Match(Cql3ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(591)

			var _x = p.ComparatorType()

			localctx.(*Collection_typeContext).t1 = _x
		}
		{
			p.SetState(592)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(593)

			var _x = p.ComparatorType()

			localctx.(*Collection_typeContext).t2 = _x
		}
		{
			p.SetState(594)
			p.Match(Cql3ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(596)
			p.Match(Cql3ParserK_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)
			p.Match(Cql3ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)

			var _x = p.ComparatorType()

			localctx.(*Collection_typeContext).t = _x
		}
		{
			p.SetState(599)
			p.Match(Cql3ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_SET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(601)
			p.Match(Cql3ParserK_SET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.Match(Cql3ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(603)

			var _x = p.ComparatorType()

			localctx.(*Collection_typeContext).t = _x
		}
		{
			p.SetState(604)
			p.Match(Cql3ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITuple_typeContext is an interface to support dynamic dispatch.
type ITuple_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() IComparatorTypeContext

	// GetTn returns the tn rule contexts.
	GetTn() IComparatorTypeContext

	// SetT1 sets the t1 rule contexts.
	SetT1(IComparatorTypeContext)

	// SetTn sets the tn rule contexts.
	SetTn(IComparatorTypeContext)

	// Getter signatures
	K_TUPLE() antlr.TerminalNode
	AllComparatorType() []IComparatorTypeContext
	ComparatorType(i int) IComparatorTypeContext

	// IsTuple_typeContext differentiates from other interfaces.
	IsTuple_typeContext()
}

type Tuple_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t1     IComparatorTypeContext
	tn     IComparatorTypeContext
}

func NewEmptyTuple_typeContext() *Tuple_typeContext {
	var p = new(Tuple_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tuple_type
	return p
}

func InitEmptyTuple_typeContext(p *Tuple_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_tuple_type
}

func (*Tuple_typeContext) IsTuple_typeContext() {}

func NewTuple_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tuple_typeContext {
	var p = new(Tuple_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_tuple_type

	return p
}

func (s *Tuple_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Tuple_typeContext) GetT1() IComparatorTypeContext { return s.t1 }

func (s *Tuple_typeContext) GetTn() IComparatorTypeContext { return s.tn }

func (s *Tuple_typeContext) SetT1(v IComparatorTypeContext) { s.t1 = v }

func (s *Tuple_typeContext) SetTn(v IComparatorTypeContext) { s.tn = v }

func (s *Tuple_typeContext) K_TUPLE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TUPLE, 0)
}

func (s *Tuple_typeContext) AllComparatorType() []IComparatorTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			len++
		}
	}

	tst := make([]IComparatorTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparatorTypeContext); ok {
			tst[i] = t.(IComparatorTypeContext)
			i++
		}
	}

	return tst
}

func (s *Tuple_typeContext) ComparatorType(i int) IComparatorTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *Tuple_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tuple_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tuple_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitTuple_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Tuple_type() (localctx ITuple_typeContext) {
	localctx = NewTuple_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, Cql3ParserRULE_tuple_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		p.Match(Cql3ParserK_TUPLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(609)
		p.Match(Cql3ParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(610)

		var _x = p.ComparatorType()

		localctx.(*Tuple_typeContext).t1 = _x
	}
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Cql3ParserT__3 {
		{
			p.SetState(611)
			p.Match(Cql3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)

			var _x = p.ComparatorType()

			localctx.(*Tuple_typeContext).tn = _x
		}

		p.SetState(617)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(618)
		p.Match(Cql3ParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVector_typeContext is an interface to support dynamic dispatch.
type IVector_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetD returns the d token.
	GetD() antlr.Token

	// SetD sets the d token.
	SetD(antlr.Token)

	// GetT1 returns the t1 rule contexts.
	GetT1() IComparatorTypeContext

	// SetT1 sets the t1 rule contexts.
	SetT1(IComparatorTypeContext)

	// Getter signatures
	K_VECTOR() antlr.TerminalNode
	ComparatorType() IComparatorTypeContext
	INTEGER() antlr.TerminalNode

	// IsVector_typeContext differentiates from other interfaces.
	IsVector_typeContext()
}

type Vector_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t1     IComparatorTypeContext
	d      antlr.Token
}

func NewEmptyVector_typeContext() *Vector_typeContext {
	var p = new(Vector_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_vector_type
	return p
}

func InitEmptyVector_typeContext(p *Vector_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_vector_type
}

func (*Vector_typeContext) IsVector_typeContext() {}

func NewVector_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_typeContext {
	var p = new(Vector_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_vector_type

	return p
}

func (s *Vector_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_typeContext) GetD() antlr.Token { return s.d }

func (s *Vector_typeContext) SetD(v antlr.Token) { s.d = v }

func (s *Vector_typeContext) GetT1() IComparatorTypeContext { return s.t1 }

func (s *Vector_typeContext) SetT1(v IComparatorTypeContext) { s.t1 = v }

func (s *Vector_typeContext) K_VECTOR() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_VECTOR, 0)
}

func (s *Vector_typeContext) ComparatorType() IComparatorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorTypeContext)
}

func (s *Vector_typeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserINTEGER, 0)
}

func (s *Vector_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vector_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitVector_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Vector_type() (localctx IVector_typeContext) {
	localctx = NewVector_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, Cql3ParserRULE_vector_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(620)
		p.Match(Cql3ParserK_VECTOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(621)
		p.Match(Cql3ParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(622)

		var _x = p.ComparatorType()

		localctx.(*Vector_typeContext).t1 = _x
	}
	{
		p.SetState(623)
		p.Match(Cql3ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(624)

		var _m = p.Match(Cql3ParserINTEGER)

		localctx.(*Vector_typeContext).d = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(625)
		p.Match(Cql3ParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INon_type_identContext is an interface to support dynamic dispatch.
type INon_type_identContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNon_type_identContext differentiates from other interfaces.
	IsNon_type_identContext()
}

type Non_type_identContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNon_type_identContext() *Non_type_identContext {
	var p = new(Non_type_identContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_non_type_ident
	return p
}

func InitEmptyNon_type_identContext(p *Non_type_identContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_non_type_ident
}

func (*Non_type_identContext) IsNon_type_identContext() {}

func NewNon_type_identContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Non_type_identContext {
	var p = new(Non_type_identContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_non_type_ident

	return p
}

func (s *Non_type_identContext) GetParser() antlr.Parser { return s.parser }

func (s *Non_type_identContext) CopyAll(ctx *Non_type_identContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Non_type_identContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Non_type_identContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NonTypeQuotedIdentContext struct {
	Non_type_identContext
	t antlr.Token
}

func NewNonTypeQuotedIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NonTypeQuotedIdentContext {
	var p = new(NonTypeQuotedIdentContext)

	InitEmptyNon_type_identContext(&p.Non_type_identContext)
	p.parser = parser
	p.CopyAll(ctx.(*Non_type_identContext))

	return p
}

func (s *NonTypeQuotedIdentContext) GetT() antlr.Token { return s.t }

func (s *NonTypeQuotedIdentContext) SetT(v antlr.Token) { s.t = v }

func (s *NonTypeQuotedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonTypeQuotedIdentContext) QUOTED_NAME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserQUOTED_NAME, 0)
}

func (s *NonTypeQuotedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitNonTypeQuotedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type NonTypeIdentContext struct {
	Non_type_identContext
	i  antlr.Token
	k  IBasic_unreserved_keywordContext
	kk antlr.Token
}

func NewNonTypeIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NonTypeIdentContext {
	var p = new(NonTypeIdentContext)

	InitEmptyNon_type_identContext(&p.Non_type_identContext)
	p.parser = parser
	p.CopyAll(ctx.(*Non_type_identContext))

	return p
}

func (s *NonTypeIdentContext) GetI() antlr.Token { return s.i }

func (s *NonTypeIdentContext) GetKk() antlr.Token { return s.kk }

func (s *NonTypeIdentContext) SetI(v antlr.Token) { s.i = v }

func (s *NonTypeIdentContext) SetKk(v antlr.Token) { s.kk = v }

func (s *NonTypeIdentContext) GetK() IBasic_unreserved_keywordContext { return s.k }

func (s *NonTypeIdentContext) SetK(v IBasic_unreserved_keywordContext) { s.k = v }

func (s *NonTypeIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonTypeIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserIDENT, 0)
}

func (s *NonTypeIdentContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DATE, 0)
}

func (s *NonTypeIdentContext) Basic_unreserved_keyword() IBasic_unreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasic_unreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasic_unreserved_keywordContext)
}

func (s *NonTypeIdentContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEY, 0)
}

func (s *NonTypeIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitNonTypeIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Non_type_ident() (localctx INon_type_identContext) {
	localctx = NewNon_type_identContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, Cql3ParserRULE_non_type_ident)
	var _la int

	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_DATE, Cql3ParserIDENT:
		localctx = NewNonTypeIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(627)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*NonTypeIdentContext).i = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Cql3ParserK_DATE || _la == Cql3ParserIDENT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*NonTypeIdentContext).i = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case Cql3ParserQUOTED_NAME:
		localctx = NewNonTypeQuotedIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(628)

			var _m = p.Match(Cql3ParserQUOTED_NAME)

			localctx.(*NonTypeQuotedIdentContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Cql3ParserK_AS, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		localctx = NewNonTypeIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(629)

			var _x = p.Basic_unreserved_keyword()

			localctx.(*NonTypeIdentContext).k = _x
		}

	case Cql3ParserK_KEY:
		localctx = NewNonTypeIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(630)

			var _m = p.Match(Cql3ParserK_KEY)

			localctx.(*NonTypeIdentContext).kk = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnreserved_keywordContext is an interface to support dynamic dispatch.
type IUnreserved_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k token.
	GetK() antlr.Token

	// SetK sets the k token.
	SetK(antlr.Token)

	// GetU returns the u rule contexts.
	GetU() IUnreserved_function_keywordContext

	// SetU sets the u rule contexts.
	SetU(IUnreserved_function_keywordContext)

	// Getter signatures
	Unreserved_function_keyword() IUnreserved_function_keywordContext
	K_TTL() antlr.TerminalNode
	K_COUNT() antlr.TerminalNode
	K_WRITETIME() antlr.TerminalNode
	K_MAXWRITETIME() antlr.TerminalNode
	K_KEY() antlr.TerminalNode
	K_CAST() antlr.TerminalNode
	K_JSON() antlr.TerminalNode
	K_DISTINCT() antlr.TerminalNode

	// IsUnreserved_keywordContext differentiates from other interfaces.
	IsUnreserved_keywordContext()
}

type Unreserved_keywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	u      IUnreserved_function_keywordContext
	k      antlr.Token
}

func NewEmptyUnreserved_keywordContext() *Unreserved_keywordContext {
	var p = new(Unreserved_keywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_unreserved_keyword
	return p
}

func InitEmptyUnreserved_keywordContext(p *Unreserved_keywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_unreserved_keyword
}

func (*Unreserved_keywordContext) IsUnreserved_keywordContext() {}

func NewUnreserved_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unreserved_keywordContext {
	var p = new(Unreserved_keywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_unreserved_keyword

	return p
}

func (s *Unreserved_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Unreserved_keywordContext) GetK() antlr.Token { return s.k }

func (s *Unreserved_keywordContext) SetK(v antlr.Token) { s.k = v }

func (s *Unreserved_keywordContext) GetU() IUnreserved_function_keywordContext { return s.u }

func (s *Unreserved_keywordContext) SetU(v IUnreserved_function_keywordContext) { s.u = v }

func (s *Unreserved_keywordContext) Unreserved_function_keyword() IUnreserved_function_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_function_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_function_keywordContext)
}

func (s *Unreserved_keywordContext) K_TTL() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TTL, 0)
}

func (s *Unreserved_keywordContext) K_COUNT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COUNT, 0)
}

func (s *Unreserved_keywordContext) K_WRITETIME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_WRITETIME, 0)
}

func (s *Unreserved_keywordContext) K_MAXWRITETIME() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MAXWRITETIME, 0)
}

func (s *Unreserved_keywordContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEY, 0)
}

func (s *Unreserved_keywordContext) K_CAST() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CAST, 0)
}

func (s *Unreserved_keywordContext) K_JSON() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_JSON, 0)
}

func (s *Unreserved_keywordContext) K_DISTINCT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DISTINCT, 0)
}

func (s *Unreserved_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unreserved_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unreserved_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitUnreserved_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Unreserved_keyword() (localctx IUnreserved_keywordContext) {
	localctx = NewUnreserved_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, Cql3ParserRULE_unreserved_keyword)
	var _la int

	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_AS, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TIMESTAMP, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_DATE, Cql3ParserK_TIME, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(633)

			var _x = p.Unreserved_function_keyword()

			localctx.(*Unreserved_keywordContext).u = _x
		}

	case Cql3ParserK_KEY, Cql3ParserK_DISTINCT, Cql3ParserK_COUNT, Cql3ParserK_TTL, Cql3ParserK_CAST, Cql3ParserK_WRITETIME, Cql3ParserK_MAXWRITETIME, Cql3ParserK_JSON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Unreserved_keywordContext).k = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((int64((_la-24)) & ^0x3f) == 0 && ((int64(1)<<(_la-24))&1649267453953) != 0) || ((int64((_la-136)) & ^0x3f) == 0 && ((int64(1)<<(_la-136))&2147483651) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Unreserved_keywordContext).k = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnreserved_function_keywordContext is an interface to support dynamic dispatch.
type IUnreserved_function_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u rule contexts.
	GetU() IBasic_unreserved_keywordContext

	// GetT returns the t rule contexts.
	GetT() INative_typeContext

	// SetU sets the u rule contexts.
	SetU(IBasic_unreserved_keywordContext)

	// SetT sets the t rule contexts.
	SetT(INative_typeContext)

	// Getter signatures
	Basic_unreserved_keyword() IBasic_unreserved_keywordContext
	Native_type() INative_typeContext

	// IsUnreserved_function_keywordContext differentiates from other interfaces.
	IsUnreserved_function_keywordContext()
}

type Unreserved_function_keywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	u      IBasic_unreserved_keywordContext
	t      INative_typeContext
}

func NewEmptyUnreserved_function_keywordContext() *Unreserved_function_keywordContext {
	var p = new(Unreserved_function_keywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_unreserved_function_keyword
	return p
}

func InitEmptyUnreserved_function_keywordContext(p *Unreserved_function_keywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_unreserved_function_keyword
}

func (*Unreserved_function_keywordContext) IsUnreserved_function_keywordContext() {}

func NewUnreserved_function_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unreserved_function_keywordContext {
	var p = new(Unreserved_function_keywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_unreserved_function_keyword

	return p
}

func (s *Unreserved_function_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Unreserved_function_keywordContext) GetU() IBasic_unreserved_keywordContext { return s.u }

func (s *Unreserved_function_keywordContext) GetT() INative_typeContext { return s.t }

func (s *Unreserved_function_keywordContext) SetU(v IBasic_unreserved_keywordContext) { s.u = v }

func (s *Unreserved_function_keywordContext) SetT(v INative_typeContext) { s.t = v }

func (s *Unreserved_function_keywordContext) Basic_unreserved_keyword() IBasic_unreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasic_unreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasic_unreserved_keywordContext)
}

func (s *Unreserved_function_keywordContext) Native_type() INative_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INative_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INative_typeContext)
}

func (s *Unreserved_function_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unreserved_function_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unreserved_function_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitUnreserved_function_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Unreserved_function_keyword() (localctx IUnreserved_function_keywordContext) {
	localctx = NewUnreserved_function_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, Cql3ParserRULE_unreserved_function_keyword)
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Cql3ParserK_AS, Cql3ParserK_KEYS, Cql3ParserK_PER, Cql3ParserK_PARTITION, Cql3ParserK_KEYSPACES, Cql3ParserK_TABLES, Cql3ParserK_CUSTOM, Cql3ParserK_VALUES, Cql3ParserK_TYPE, Cql3ParserK_TYPES, Cql3ParserK_COMPACT, Cql3ParserK_STORAGE, Cql3ParserK_FILTERING, Cql3ParserK_CONTAINS, Cql3ParserK_BETWEEN, Cql3ParserK_GROUP, Cql3ParserK_CLUSTER, Cql3ParserK_INTERNALS, Cql3ParserK_ONLY, Cql3ParserK_ALL, Cql3ParserK_PERMISSION, Cql3ParserK_PERMISSIONS, Cql3ParserK_MBEAN, Cql3ParserK_MBEANS, Cql3ParserK_USER, Cql3ParserK_USERS, Cql3ParserK_ROLE, Cql3ParserK_ROLES, Cql3ParserK_SUPERUSERS, Cql3ParserK_SUPERUSER, Cql3ParserK_NOSUPERUSER, Cql3ParserK_GENERATED, Cql3ParserK_PASSWORD, Cql3ParserK_HASHED, Cql3ParserK_LOGIN, Cql3ParserK_NOLOGIN, Cql3ParserK_OPTIONS, Cql3ParserK_ACCESS, Cql3ParserK_DATACENTERS, Cql3ParserK_CIDRS, Cql3ParserK_IDENTITY, Cql3ParserK_CLUSTERING, Cql3ParserK_EXISTS, Cql3ParserK_MAP, Cql3ParserK_LIST, Cql3ParserK_TUPLE, Cql3ParserK_TRIGGER, Cql3ParserK_STATIC, Cql3ParserK_FROZEN, Cql3ParserK_FUNCTION, Cql3ParserK_FUNCTIONS, Cql3ParserK_AGGREGATE, Cql3ParserK_AGGREGATES, Cql3ParserK_SFUNC, Cql3ParserK_STYPE, Cql3ParserK_FINALFUNC, Cql3ParserK_INITCOND, Cql3ParserK_RETURNS, Cql3ParserK_CALLED, Cql3ParserK_INPUT, Cql3ParserK_LANGUAGE, Cql3ParserK_REPLACE, Cql3ParserK_DEFAULT, Cql3ParserK_UNSET, Cql3ParserK_LIKE, Cql3ParserK_MASKED, Cql3ParserK_UNMASK, Cql3ParserK_SELECT_MASKED, Cql3ParserK_VECTOR, Cql3ParserK_ANN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(637)

			var _x = p.Basic_unreserved_keyword()

			localctx.(*Unreserved_function_keywordContext).u = _x
		}

	case Cql3ParserK_TIMESTAMP, Cql3ParserK_ASCII, Cql3ParserK_BIGINT, Cql3ParserK_BLOB, Cql3ParserK_BOOLEAN, Cql3ParserK_COUNTER, Cql3ParserK_DECIMAL, Cql3ParserK_DOUBLE, Cql3ParserK_DURATION, Cql3ParserK_FLOAT, Cql3ParserK_INET, Cql3ParserK_INT, Cql3ParserK_SMALLINT, Cql3ParserK_TINYINT, Cql3ParserK_TEXT, Cql3ParserK_UUID, Cql3ParserK_VARCHAR, Cql3ParserK_VARINT, Cql3ParserK_TIMEUUID, Cql3ParserK_DATE, Cql3ParserK_TIME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(638)

			var _x = p.Native_type()

			localctx.(*Unreserved_function_keywordContext).t = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBasic_unreserved_keywordContext is an interface to support dynamic dispatch.
type IBasic_unreserved_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k token.
	GetK() antlr.Token

	// SetK sets the k token.
	SetK(antlr.Token)

	// Getter signatures
	K_KEYS() antlr.TerminalNode
	K_AS() antlr.TerminalNode
	K_CLUSTER() antlr.TerminalNode
	K_CLUSTERING() antlr.TerminalNode
	K_COMPACT() antlr.TerminalNode
	K_STORAGE() antlr.TerminalNode
	K_TABLES() antlr.TerminalNode
	K_TYPE() antlr.TerminalNode
	K_TYPES() antlr.TerminalNode
	K_VALUES() antlr.TerminalNode
	K_MAP() antlr.TerminalNode
	K_LIST() antlr.TerminalNode
	K_FILTERING() antlr.TerminalNode
	K_PERMISSION() antlr.TerminalNode
	K_PERMISSIONS() antlr.TerminalNode
	K_KEYSPACES() antlr.TerminalNode
	K_ALL() antlr.TerminalNode
	K_USER() antlr.TerminalNode
	K_USERS() antlr.TerminalNode
	K_ROLE() antlr.TerminalNode
	K_ROLES() antlr.TerminalNode
	K_IDENTITY() antlr.TerminalNode
	K_SUPERUSER() antlr.TerminalNode
	K_SUPERUSERS() antlr.TerminalNode
	K_NOSUPERUSER() antlr.TerminalNode
	K_LOGIN() antlr.TerminalNode
	K_NOLOGIN() antlr.TerminalNode
	K_OPTIONS() antlr.TerminalNode
	K_PASSWORD() antlr.TerminalNode
	K_GENERATED() antlr.TerminalNode
	K_HASHED() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode
	K_CUSTOM() antlr.TerminalNode
	K_TRIGGER() antlr.TerminalNode
	K_CONTAINS() antlr.TerminalNode
	K_INTERNALS() antlr.TerminalNode
	K_ONLY() antlr.TerminalNode
	K_STATIC() antlr.TerminalNode
	K_FROZEN() antlr.TerminalNode
	K_TUPLE() antlr.TerminalNode
	K_FUNCTION() antlr.TerminalNode
	K_FUNCTIONS() antlr.TerminalNode
	K_AGGREGATE() antlr.TerminalNode
	K_AGGREGATES() antlr.TerminalNode
	K_SFUNC() antlr.TerminalNode
	K_STYPE() antlr.TerminalNode
	K_FINALFUNC() antlr.TerminalNode
	K_INITCOND() antlr.TerminalNode
	K_RETURNS() antlr.TerminalNode
	K_LANGUAGE() antlr.TerminalNode
	K_CALLED() antlr.TerminalNode
	K_INPUT() antlr.TerminalNode
	K_LIKE() antlr.TerminalNode
	K_PER() antlr.TerminalNode
	K_PARTITION() antlr.TerminalNode
	K_GROUP() antlr.TerminalNode
	K_DATACENTERS() antlr.TerminalNode
	K_CIDRS() antlr.TerminalNode
	K_ACCESS() antlr.TerminalNode
	K_DEFAULT() antlr.TerminalNode
	K_MBEAN() antlr.TerminalNode
	K_MBEANS() antlr.TerminalNode
	K_REPLACE() antlr.TerminalNode
	K_UNSET() antlr.TerminalNode
	K_MASKED() antlr.TerminalNode
	K_UNMASK() antlr.TerminalNode
	K_SELECT_MASKED() antlr.TerminalNode
	K_VECTOR() antlr.TerminalNode
	K_ANN() antlr.TerminalNode
	K_BETWEEN() antlr.TerminalNode

	// IsBasic_unreserved_keywordContext differentiates from other interfaces.
	IsBasic_unreserved_keywordContext()
}

type Basic_unreserved_keywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	k      antlr.Token
}

func NewEmptyBasic_unreserved_keywordContext() *Basic_unreserved_keywordContext {
	var p = new(Basic_unreserved_keywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_basic_unreserved_keyword
	return p
}

func InitEmptyBasic_unreserved_keywordContext(p *Basic_unreserved_keywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Cql3ParserRULE_basic_unreserved_keyword
}

func (*Basic_unreserved_keywordContext) IsBasic_unreserved_keywordContext() {}

func NewBasic_unreserved_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Basic_unreserved_keywordContext {
	var p = new(Basic_unreserved_keywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cql3ParserRULE_basic_unreserved_keyword

	return p
}

func (s *Basic_unreserved_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Basic_unreserved_keywordContext) GetK() antlr.Token { return s.k }

func (s *Basic_unreserved_keywordContext) SetK(v antlr.Token) { s.k = v }

func (s *Basic_unreserved_keywordContext) K_KEYS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEYS, 0)
}

func (s *Basic_unreserved_keywordContext) K_AS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_AS, 0)
}

func (s *Basic_unreserved_keywordContext) K_CLUSTER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CLUSTER, 0)
}

func (s *Basic_unreserved_keywordContext) K_CLUSTERING() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CLUSTERING, 0)
}

func (s *Basic_unreserved_keywordContext) K_COMPACT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_COMPACT, 0)
}

func (s *Basic_unreserved_keywordContext) K_STORAGE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_STORAGE, 0)
}

func (s *Basic_unreserved_keywordContext) K_TABLES() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TABLES, 0)
}

func (s *Basic_unreserved_keywordContext) K_TYPE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TYPE, 0)
}

func (s *Basic_unreserved_keywordContext) K_TYPES() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TYPES, 0)
}

func (s *Basic_unreserved_keywordContext) K_VALUES() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_VALUES, 0)
}

func (s *Basic_unreserved_keywordContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MAP, 0)
}

func (s *Basic_unreserved_keywordContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_LIST, 0)
}

func (s *Basic_unreserved_keywordContext) K_FILTERING() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FILTERING, 0)
}

func (s *Basic_unreserved_keywordContext) K_PERMISSION() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PERMISSION, 0)
}

func (s *Basic_unreserved_keywordContext) K_PERMISSIONS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PERMISSIONS, 0)
}

func (s *Basic_unreserved_keywordContext) K_KEYSPACES() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_KEYSPACES, 0)
}

func (s *Basic_unreserved_keywordContext) K_ALL() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ALL, 0)
}

func (s *Basic_unreserved_keywordContext) K_USER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_USER, 0)
}

func (s *Basic_unreserved_keywordContext) K_USERS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_USERS, 0)
}

func (s *Basic_unreserved_keywordContext) K_ROLE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ROLE, 0)
}

func (s *Basic_unreserved_keywordContext) K_ROLES() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ROLES, 0)
}

func (s *Basic_unreserved_keywordContext) K_IDENTITY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_IDENTITY, 0)
}

func (s *Basic_unreserved_keywordContext) K_SUPERUSER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_SUPERUSER, 0)
}

func (s *Basic_unreserved_keywordContext) K_SUPERUSERS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_SUPERUSERS, 0)
}

func (s *Basic_unreserved_keywordContext) K_NOSUPERUSER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_NOSUPERUSER, 0)
}

func (s *Basic_unreserved_keywordContext) K_LOGIN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_LOGIN, 0)
}

func (s *Basic_unreserved_keywordContext) K_NOLOGIN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_NOLOGIN, 0)
}

func (s *Basic_unreserved_keywordContext) K_OPTIONS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_OPTIONS, 0)
}

func (s *Basic_unreserved_keywordContext) K_PASSWORD() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PASSWORD, 0)
}

func (s *Basic_unreserved_keywordContext) K_GENERATED() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_GENERATED, 0)
}

func (s *Basic_unreserved_keywordContext) K_HASHED() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_HASHED, 0)
}

func (s *Basic_unreserved_keywordContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_EXISTS, 0)
}

func (s *Basic_unreserved_keywordContext) K_CUSTOM() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CUSTOM, 0)
}

func (s *Basic_unreserved_keywordContext) K_TRIGGER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TRIGGER, 0)
}

func (s *Basic_unreserved_keywordContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CONTAINS, 0)
}

func (s *Basic_unreserved_keywordContext) K_INTERNALS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_INTERNALS, 0)
}

func (s *Basic_unreserved_keywordContext) K_ONLY() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ONLY, 0)
}

func (s *Basic_unreserved_keywordContext) K_STATIC() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_STATIC, 0)
}

func (s *Basic_unreserved_keywordContext) K_FROZEN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FROZEN, 0)
}

func (s *Basic_unreserved_keywordContext) K_TUPLE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_TUPLE, 0)
}

func (s *Basic_unreserved_keywordContext) K_FUNCTION() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FUNCTION, 0)
}

func (s *Basic_unreserved_keywordContext) K_FUNCTIONS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FUNCTIONS, 0)
}

func (s *Basic_unreserved_keywordContext) K_AGGREGATE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_AGGREGATE, 0)
}

func (s *Basic_unreserved_keywordContext) K_AGGREGATES() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_AGGREGATES, 0)
}

func (s *Basic_unreserved_keywordContext) K_SFUNC() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_SFUNC, 0)
}

func (s *Basic_unreserved_keywordContext) K_STYPE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_STYPE, 0)
}

func (s *Basic_unreserved_keywordContext) K_FINALFUNC() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_FINALFUNC, 0)
}

func (s *Basic_unreserved_keywordContext) K_INITCOND() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_INITCOND, 0)
}

func (s *Basic_unreserved_keywordContext) K_RETURNS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_RETURNS, 0)
}

func (s *Basic_unreserved_keywordContext) K_LANGUAGE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_LANGUAGE, 0)
}

func (s *Basic_unreserved_keywordContext) K_CALLED() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CALLED, 0)
}

func (s *Basic_unreserved_keywordContext) K_INPUT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_INPUT, 0)
}

func (s *Basic_unreserved_keywordContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_LIKE, 0)
}

func (s *Basic_unreserved_keywordContext) K_PER() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PER, 0)
}

func (s *Basic_unreserved_keywordContext) K_PARTITION() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_PARTITION, 0)
}

func (s *Basic_unreserved_keywordContext) K_GROUP() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_GROUP, 0)
}

func (s *Basic_unreserved_keywordContext) K_DATACENTERS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DATACENTERS, 0)
}

func (s *Basic_unreserved_keywordContext) K_CIDRS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_CIDRS, 0)
}

func (s *Basic_unreserved_keywordContext) K_ACCESS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ACCESS, 0)
}

func (s *Basic_unreserved_keywordContext) K_DEFAULT() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_DEFAULT, 0)
}

func (s *Basic_unreserved_keywordContext) K_MBEAN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MBEAN, 0)
}

func (s *Basic_unreserved_keywordContext) K_MBEANS() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MBEANS, 0)
}

func (s *Basic_unreserved_keywordContext) K_REPLACE() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_REPLACE, 0)
}

func (s *Basic_unreserved_keywordContext) K_UNSET() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_UNSET, 0)
}

func (s *Basic_unreserved_keywordContext) K_MASKED() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_MASKED, 0)
}

func (s *Basic_unreserved_keywordContext) K_UNMASK() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_UNMASK, 0)
}

func (s *Basic_unreserved_keywordContext) K_SELECT_MASKED() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_SELECT_MASKED, 0)
}

func (s *Basic_unreserved_keywordContext) K_VECTOR() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_VECTOR, 0)
}

func (s *Basic_unreserved_keywordContext) K_ANN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_ANN, 0)
}

func (s *Basic_unreserved_keywordContext) K_BETWEEN() antlr.TerminalNode {
	return s.GetToken(Cql3ParserK_BETWEEN, 0)
}

func (s *Basic_unreserved_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Basic_unreserved_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Basic_unreserved_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cql3Visitor:
		return t.VisitBasic_unreserved_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cql3Parser) Basic_unreserved_keyword() (localctx IBasic_unreserved_keywordContext) {
	localctx = NewBasic_unreserved_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, Cql3ParserRULE_basic_unreserved_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Basic_unreserved_keywordContext).k = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2344686568920317952) != 0) || ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&562949420478991) != 0) || ((int64((_la-142)) & ^0x3f) == 0 && ((int64(1)<<(_la-142))&17137926023) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Basic_unreserved_keywordContext).k = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
