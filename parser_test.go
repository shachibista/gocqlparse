package cql

import (
	"math"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/uuid"
	"github.com/shachibista/gocqlparse/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	input    string
	expected any
	hasError bool
	skip     string
	test     func(*testing.T, testCase, any)
}

func TestParseNativeType(t *testing.T) {
	types := []string{
		"ascii",
		"bigint",
		"blob",
		"boolean",
		"counter",
		"decimal",
		"double",
		"duration",
		"float",
		"inet",
		"int",
		"smallint",
		"text",
		"timestamp",
		"tinyint",
		"uuid",
		"varchar",
		"varint",
		"timeuuid",
		"date",
		"time",
	}

	for _, nt := range types {
		p := NewParser(nt)
		st, err := parseRule[*ast.NativeType](func() antlr.ParseTree { return p.Native_type() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, nt, st.Name)

		p = NewParser(strings.ToUpper(nt))
		st, err = parseRule[*ast.NativeType](func() antlr.ParseTree { return p.Native_type() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, nt, st.Name)
	}
}

func TestParseCollectionType(t *testing.T) {
	types := []testCase{
		{
			input:    "list<int>",
			expected: &ast.ListType{Elem: &ast.NativeType{Name: "int"}},
		},
		{
			input:    "set<int>",
			expected: &ast.SetType{Elem: &ast.NativeType{Name: "int"}},
		},
		{
			input: "map<int,ascii>",
			expected: &ast.MapType{
				Key:   &ast.NativeType{Name: "int"},
				Value: &ast.NativeType{Name: "ascii"},
			},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.Collection_type() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseTupleType(t *testing.T) {
	types := []testCase{
		{
			input: "tuple<int>",
			expected: &ast.TupleType{
				Elements: []ast.GenericType{
					&ast.NativeType{Name: "int"},
				},
			},
		},
		{
			input: "tuple<int,ascii>",
			expected: &ast.TupleType{
				Elements: []ast.GenericType{
					&ast.NativeType{Name: "int"},
					&ast.NativeType{Name: "ascii"},
				},
			},
		},
		{
			input: "tuple<int,ascii,int>",
			expected: &ast.TupleType{
				Elements: []ast.GenericType{
					&ast.NativeType{Name: "int"},
					&ast.NativeType{Name: "ascii"},
					&ast.NativeType{Name: "int"},
				},
			},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.Tuple_type() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseVectorType(t *testing.T) {
	types := []testCase{
		{
			input: "vector<int,5>",
			expected: &ast.VectorType{
				Elem:       &ast.NativeType{Name: "int"},
				Dimensions: 5,
			},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.Vector_type() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseUserTypeName(t *testing.T) {
	types := []testCase{
		{
			input: "a.b",
			expected: &ast.ObjectRef{
				Keyspace: ast.UnquotedIdentifier("a"),
				Name:     ast.UnquotedIdentifier("b"),
			},
		},
		{
			input: "b",
			expected: &ast.ObjectRef{
				Keyspace: nil,
				Name:     ast.UnquotedIdentifier("b"),
			},
		},
		{
			input: "A.B",
			expected: &ast.ObjectRef{
				Keyspace: ast.UnquotedIdentifier("A"),
				Name:     ast.UnquotedIdentifier("B"),
			},
		},
		{
			input: `"A"."B"`,
			expected: &ast.ObjectRef{
				Keyspace: ast.QuotedIdentifier("A"),
				Name:     ast.QuotedIdentifier("B"),
			},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.UserTypeName() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseNonTypeIdent(t *testing.T) {
	types := []testCase{
		{
			input:    "hello",
			expected: ast.UnquotedIdentifier("hello"),
			hasError: false,
		},
		{
			input:    "KEY",
			expected: ast.UnquotedIdentifier("KEY"),
			hasError: false,
		},
		{
			input:    "MAP",
			expected: ast.UnquotedIdentifier("MAP"),
			hasError: false,
		},
		{
			input:    `"Date"`,
			expected: ast.QuotedIdentifier("Date"),
		},
	}

	for t := range ast.ReservedTypeNames {
		types = append(types, testCase{
			input:    t,
			expected: ast.UnquotedIdentifier(t),
			hasError: true,
		})
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[ast.Identifier](func() antlr.ParseTree { return p.Non_type_ident() })

		if tc.hasError {
			require.NotNil(t, err, "expected error for non-type ident %s", tc.input)
		} else {
			require.Nil(t, err)
			assert.Equal(t, tc.expected, st)
		}
	}
}

func TestParseFrozenType(t *testing.T) {
	types := []testCase{
		{
			input: "frozen<int>",
			expected: &ast.FrozenType{
				Elem: &ast.NativeType{
					Name: "int",
				},
			},
		},
		{
			input: "frozen<list<int>>",
			expected: &ast.FrozenType{
				Elem: &ast.ListType{
					Elem: &ast.NativeType{
						Name: "int",
					},
				},
			},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.ComparatorType() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseCustomType(t *testing.T) {
	t.Skip("custom types are not supported yet, don't know how to test them")
}

func TestParseComparatorType(t *testing.T) {
	types := []testCase{
		{
			input: "list<map<int,ascii>>",
			expected: &ast.ListType{
				Elem: &ast.MapType{
					Key:   &ast.NativeType{Name: "int"},
					Value: &ast.NativeType{Name: "ascii"},
				},
			},
		},
		{
			input: "map<int,list<ascii>>",
			expected: &ast.MapType{
				Key: &ast.NativeType{Name: "int"},
				Value: &ast.ListType{
					Elem: &ast.NativeType{Name: "ascii"},
				},
			},
		},
		{
			input: "list<map<int,list<ascii>>>",
			expected: &ast.ListType{
				Elem: &ast.MapType{
					Key: &ast.NativeType{Name: "int"},
					Value: &ast.ListType{
						Elem: &ast.NativeType{Name: "ascii"},
					},
				},
			},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.ComparatorType() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseTableColumnDeclaration(t *testing.T) {
	decls := []testCase{
		{
			input: "col1 int",
			expected: &ast.TableColumn{
				Name: ast.UnquotedIdentifier("col1"),
				Type: &ast.NativeType{Name: "int"},
			},
		},
		{
			input: "col1 int static",
			expected: &ast.TableColumn{
				Name:   ast.UnquotedIdentifier("col1"),
				Type:   &ast.NativeType{Name: "int"},
				Static: true,
			},
		},
		{
			input: "col1 int primary key",
			expected: &ast.TableColumn{
				Name:               ast.UnquotedIdentifier("col1"),
				Type:               &ast.NativeType{Name: "int"},
				ImplicitPrimaryKey: true,
			},
		},
		{
			input: "col1 int static primary key",
			expected: &ast.TableColumn{
				Name:               ast.UnquotedIdentifier("col1"),
				Type:               &ast.NativeType{Name: "int"},
				Static:             true,
				ImplicitPrimaryKey: true,
			},
		},
	}

	for _, tc := range decls {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.TableColumns() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseTableKeyDeclaration(t *testing.T) {
	decls := []testCase{
		{
			input: "primary key (col1)",
			expected: &ast.TableKeys{
				Partition: []ast.Identifier{
					ast.UnquotedIdentifier("col1"),
				},
			},
		},
		{
			input: "primary key (col1, col2)",
			expected: &ast.TableKeys{
				Partition: []ast.Identifier{
					ast.UnquotedIdentifier("col1"),
				},
				Clustering: []ast.Identifier{
					ast.UnquotedIdentifier("col2"),
				},
			},
		},
		{
			input: "primary key ((col1, col2))",
			expected: &ast.TableKeys{
				Partition: []ast.Identifier{
					ast.UnquotedIdentifier("col1"),
					ast.UnquotedIdentifier("col2"),
				},
			},
		},
		{
			input: "primary key ((col1), col2)",
			expected: &ast.TableKeys{
				Partition: []ast.Identifier{
					ast.UnquotedIdentifier("col1"),
				},
				Clustering: []ast.Identifier{
					ast.UnquotedIdentifier("col2"),
				},
			},
		},
		{
			input: "primary key ((col1, col2), col3, col4)",
			expected: &ast.TableKeys{
				Partition: []ast.Identifier{
					ast.UnquotedIdentifier("col1"),
					ast.UnquotedIdentifier("col2"),
				},
				Clustering: []ast.Identifier{
					ast.UnquotedIdentifier("col3"),
					ast.UnquotedIdentifier("col4"),
				},
			},
		},
	}

	for _, tc := range decls {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.TableColumns() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseTablePartitionKeys(t *testing.T) {
	decls := []testCase{
		{
			input: "col1",
			expected: []ast.Identifier{
				ast.UnquotedIdentifier("col1"),
			},
		},
		{
			input: "(col1)",
			expected: []ast.Identifier{
				ast.UnquotedIdentifier("col1"),
			},
		},
		{
			input: "(col1, col2)",
			expected: []ast.Identifier{
				ast.UnquotedIdentifier("col1"),
				ast.UnquotedIdentifier("col2"),
			},
		},
	}

	for _, tc := range decls {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.TablePartitionKey() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseColumnFamilyName(t *testing.T) {
	types := []testCase{
		{
			input:    "a.b",
			expected: &ast.ObjectRef{ast.UnquotedIdentifier("a"), ast.UnquotedIdentifier("b")},
		},
		{
			input:    "b",
			expected: &ast.ObjectRef{nil, ast.UnquotedIdentifier("b")},
		},
		{
			input:    "A.B",
			expected: &ast.ObjectRef{ast.UnquotedIdentifier("A"), ast.UnquotedIdentifier("B")},
		},
		{
			input:    `"A"."B"`,
			expected: &ast.ObjectRef{ast.QuotedIdentifier("A"), ast.QuotedIdentifier("B")},
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.ColumnFamilyName() })
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expected, st)
	}
}

func TestParseConstant(t *testing.T) {
	consts := []testCase{
		{
			input:    "$$hello$$",
			expected: "hello",
		},
		{
			input:    "$$with$signs$$",
			expected: "with$signs",
		},
		{
			input: `$$
			BEGIN
    RETURN ($1 ~ $q$[\t\r\n\v\\]$q$);
END;
$$`,
			expected: `
			BEGIN
    RETURN ($1 ~ $q$[\t\r\n\v\\]$q$);
END;
`,
		},
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "12432",
			expected: 12432,
		},
		{
			input:    "-323",
			expected: -323,
		},
		{
			input:    "3.14",
			expected: 3.14,
		},
		{
			input:    "3e10",
			expected: 3e10,
		},
		{
			input:    "true",
			expected: true,
		},
		{
			input:    "false",
			expected: false,
		},
		{
			input:    "TRUE",
			expected: true,
		},
		{
			input:    "FALSE",
			expected: false,
		},
		{
			input: "P9348842955W",
			skip:  "TODO: duration not implemented",
		},
		{
			input:    "00000000-0000-0000-0000-000000000001",
			expected: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		},
		{
			input:    "0x30",
			expected: []byte{0x30},
		},
		{
			input:    "0x30efd3",
			expected: []byte{0x30, 0xef, 0xd3},
		},
		{
			input:    "nan",
			expected: math.NaN(),
		},
		{
			input:    "-nan",
			expected: -math.NaN(),
		},
		{
			input:    "infinity",
			expected: math.Inf(1),
		},
		{
			input:    "-infinity",
			expected: math.Inf(-1),
		},
	}

	for _, tc := range consts {
		if tc.skip != "" {
			t.Log(tc.skip)
			continue
		}

		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.Constant() })
		if err != nil {
			t.Fatal(err)
		}

		switch tc.expected.(type) {
		case float64, float32:
			assert.InDelta(t, tc.expected, st, 0.0001)
		default:
			assert.Equal(t, tc.expected, st)
		}
	}
}

func TestParseFunctionName(t *testing.T) {
	cases := []testCase{
		{
			input: "foo",
			expected: &ast.ObjectRef{
				Name: ast.UnquotedIdentifier("foo"),
			},
		},
		{
			input: "foo.bar",
			expected: &ast.ObjectRef{
				Keyspace: ast.UnquotedIdentifier("foo"),
				Name:     ast.UnquotedIdentifier("bar"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.FunctionName() })
}

func TestParseStatementCreateTable(t *testing.T) {
	cases := []testCase{
		{
			input: "create table test (id int)",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
			},
		},
		{
			input: "create table ks.test (id int)",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks"),
					Name:     ast.UnquotedIdentifier("test"),
				},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
			},
		},
		{
			input: "create table if not exists test (id int)",
			expected: &ast.CreateTableStatement{
				IfNotExists: true,
				ColumnFamily: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
			},
		},
		{
			input: "create table test (id int primary key)",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
				Columns: []*ast.TableColumn{
					{
						Name:               ast.UnquotedIdentifier("id"),
						Type:               &ast.NativeType{"int"},
						ImplicitPrimaryKey: true,
					},
				},
				PartitionKeys: []ast.Identifier{
					ast.UnquotedIdentifier("id"),
				},
			},
		},
		{
			input:    "create table test (id int primary key, primary key (id))",
			hasError: true,
		},
		{
			input:    "create table test (id int primary key, value int primary key)",
			hasError: true,
		},
		{
			input: "create table test (id int, primary key(id))",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
				PartitionKeys: []ast.Identifier{ast.UnquotedIdentifier("id")},
			},
		},
		{
			input: "create table test (id int, col1 ascii, col2 ascii, primary key (id, col1))",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
					{
						Name: ast.UnquotedIdentifier("col1"),
						Type: &ast.NativeType{"ascii"},
					},
					{
						Name: ast.UnquotedIdentifier("col2"),
						Type: &ast.NativeType{"ascii"},
					},
				},
				PartitionKeys:  []ast.Identifier{ast.UnquotedIdentifier("id")},
				ClusteringKeys: []ast.Identifier{ast.UnquotedIdentifier("col1")},
			},
		},
		{
			input: "create table test (id int, col1 ascii, col2 ascii, primary key (id, col1, col2))",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
					{
						Name: ast.UnquotedIdentifier("col1"),
						Type: &ast.NativeType{"ascii"},
					},
					{
						Name: ast.UnquotedIdentifier("col2"),
						Type: &ast.NativeType{"ascii"},
					},
				},
				PartitionKeys:  []ast.Identifier{ast.UnquotedIdentifier("id")},
				ClusteringKeys: []ast.Identifier{ast.UnquotedIdentifier("col1"), ast.UnquotedIdentifier("col2")},
			},
		},
		{
			input: "create table test (id int, col1 ascii, col2 ascii, primary key ((id, col1), col2))",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
					{
						Name: ast.UnquotedIdentifier("col1"),
						Type: &ast.NativeType{"ascii"},
					},
					{
						Name: ast.UnquotedIdentifier("col2"),
						Type: &ast.NativeType{"ascii"},
					},
				},
				PartitionKeys:  []ast.Identifier{ast.UnquotedIdentifier("id"), ast.UnquotedIdentifier("col1")},
				ClusteringKeys: []ast.Identifier{ast.UnquotedIdentifier("col2")},
			},
		},
		{
			input: `CREATE TABLE patients (
				id timeuuid PRIMARY KEY,
				name text MASKED WITH mask_inner(1, null),
				birth date MASKED WITH mask_default()
			);`,
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("patients")},
				Columns: []*ast.TableColumn{
					{
						Name:               ast.UnquotedIdentifier("id"),
						Type:               &ast.NativeType{"timeuuid"},
						ImplicitPrimaryKey: true,
					},
					{
						Name: ast.UnquotedIdentifier("name"),
						Type: &ast.NativeType{"text"},
						Mask: &ast.Function{
							Name: &ast.ObjectRef{
								Name: ast.UnquotedIdentifier("mask_inner"),
							},
							Arguments: []ast.Term{
								1, ast.NullLiteral,
							},
						},
					},
					{
						Name: ast.UnquotedIdentifier("birth"),
						Type: &ast.NativeType{"date"},
						Mask: &ast.Function{
							Name: &ast.ObjectRef{
								Name: ast.UnquotedIdentifier("mask_default"),
							},
						},
					},
				},
				PartitionKeys: []ast.Identifier{
					ast.UnquotedIdentifier("id"),
				},
			},
		},
		{
			input: "create table test (id int) with compact storage",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
				CompactStorage: true,
			},
		},
		{
			input: "create table test (id int) with compact storage and clustering order by (id desc)",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
				CompactStorage: true,
				ClusteringOrder: []*ast.ColumnOrder{
					{
						Name:      ast.UnquotedIdentifier("id"),
						Ascending: false,
					},
				},
			},
		},
		{
			input: "create table test (id int) with replication = {'class': 'NetworkTopologyStrategy', 'DC1': '3', 'DC2': '3'} AND durable_writes = true",
			expected: &ast.CreateTableStatement{
				ColumnFamily: &ast.ObjectRef{Name: ast.UnquotedIdentifier("test")},
				Columns: []*ast.TableColumn{
					{
						Name: ast.UnquotedIdentifier("id"),
						Type: &ast.NativeType{"int"},
					},
				},
				Properties: []*ast.Property{
					{
						Key: ast.UnquotedIdentifier("replication"),
						Value: &ast.MapLiteral{
							"class": "NetworkTopologyStrategy",
							"DC1":   "3",
							"DC2":   "3",
						},
					},
					{
						Key:   ast.UnquotedIdentifier("durable_writes"),
						Value: true,
					},
				},
			},
		},
	}

	for _, tc := range cases {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.CreateTableStatement() })

		if tc.hasError {
			require.NotNil(t, err, "expected error for non-type ident %s", tc.input)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tc.expected, st)
		}
	}
}

func TestParseStatementCreateType(t *testing.T) {
	cases := []testCase{
		{
			input: "create type test (f1 int, f2 ascii)",
			expected: &ast.CreateTypeStatement{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
				Fields: []*ast.TypeField{
					{
						Name: ast.UnquotedIdentifier("f1"),
						Type: &ast.NativeType{"int"},
					},
					{
						Name: ast.UnquotedIdentifier("f2"),
						Type: &ast.NativeType{"ascii"},
					},
				},
			},
		},
		{
			input: "create type if not exists test (f1 int, f2 ascii)",
			expected: &ast.CreateTypeStatement{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
				IfNotExists: true,
				Fields: []*ast.TypeField{
					{
						Name: ast.UnquotedIdentifier("f1"),
						Type: &ast.NativeType{"int"},
					},
					{
						Name: ast.UnquotedIdentifier("f2"),
						Type: &ast.NativeType{"ascii"},
					},
				},
			},
		},
		{
			input:    "create type test (f1 test)",
			expected: nil,
			hasError: true,
		},
		{
			input:    "create type ks.test (f1 ks.test)",
			expected: nil,
			hasError: true,
		},
		{
			input: "create type ks1.test (f1 ks2.test)",
			expected: &ast.CreateTypeStatement{
				Name: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks1"),
					Name:     ast.UnquotedIdentifier("test"),
				},
				Fields: []*ast.TypeField{
					{
						Name: ast.UnquotedIdentifier("f1"),
						Type: &ast.ObjectRef{
							Keyspace: ast.UnquotedIdentifier("ks2"),
							Name:     ast.UnquotedIdentifier("test"),
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		p := NewParser(tc.input)
		st, err := parseRule[any](func() antlr.ParseTree { return p.CreateTypeStatement() })

		if tc.hasError {
			require.NotNil(t, err, "expected error for non-type ident %s", tc.input)
		} else {
			require.Nil(t, err)
			assert.Equal(t, tc.expected, st)
		}
	}
}

func TestParseStatementUse(t *testing.T) {
	cases := []testCase{
		{
			input:    "use test",
			expected: &ast.UseStatement{Keyspace: ast.UnquotedIdentifier("test")},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.UseStatement() })
}

func TestParseColumnMask(t *testing.T) {
	cases := []testCase{
		{
			input: "masked with default",
			expected: &ast.Function{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("mask_default"),
				},
			},
		},
		{
			input: "masked with mask_null('Alice')",
			expected: &ast.Function{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("mask_null"),
				},
				Arguments: []ast.Term{
					"Alice",
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.ColumnMask() })
}

func TestParseTerm(t *testing.T) {
	cases := []testCase{
		{
			input: "1 + 2",
			expected: &ast.TermOperation{
				Operator: ast.AdditionTermOperator,
				Left:     1,
				Right:    2,
			},
		},
		{
			input: "1 - 2",
			expected: &ast.TermOperation{
				Operator: ast.SubtractionTermOperator,
				Left:     1,
				Right:    2,
			},
		},
		{
			input: "1 / 2",
			expected: &ast.TermOperation{
				Operator: ast.DivisionTermOperator,
				Left:     1,
				Right:    2,
			},
		},
		{
			input: "1 % 2",
			expected: &ast.TermOperation{
				Operator: ast.RemainderTermOperator,
				Left:     1,
				Right:    2,
			},
		},
		{
			input: "1 * 2",
			expected: &ast.TermOperation{
				Operator: ast.MultiplicationTermOperator,
				Left:     1,
				Right:    2,
			},
		},
		{
			input:    "-1",
			expected: -1,
		},
		{
			input: "f()",
			expected: &ast.Function{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("f"),
				},
				Arguments: nil,
			},
		},
		{
			input: "mask_null(123)",
			expected: &ast.Function{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("mask_null"),
				},
				Arguments: []ast.Term{123},
			},
		},
		{
			input: "mask_null('Alice')",
			expected: &ast.Function{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("mask_null"),
				},
				Arguments: []ast.Term{"Alice"},
			},
		},
		{
			input: "(vector<int, 3>) [1, 2, 3]",
			expected: &ast.TermOperation{
				Operator: ast.LiteralCastTermOperator,
				Left: &ast.VectorType{
					Elem:       &ast.NativeType{"int"},
					Dimensions: 3,
				},
				Right: []any{1, 2, 3},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.Term() })
}

func TestParseValue(t *testing.T) {
	cases := []testCase{
		{
			input:    "{}",
			expected: &ast.SetLiteral{},
		},
		{
			input:    "{ 1, 2, 3 }",
			expected: &ast.SetLiteral{1, 2, 3},
		},
		{
			input: "{'age': 30, 'city': 'New York', 'status': 'active'}",
			expected: &ast.MapLiteral{
				"age":    30,
				"city":   "New York",
				"status": "active",
			},
		},
		{
			input:    "NULL",
			expected: ast.NullLiteral,
		},
		{
			input: "{ birthday : '1993-06-18', nationality : 'New Zealand', weight : null, height : null }",
			expected: &ast.UserTypeLiteral{
				ast.UnquotedIdentifier("birthday"):    "1993-06-18",
				ast.UnquotedIdentifier("nationality"): "New Zealand",
				ast.UnquotedIdentifier("weight"):      nil,
				ast.UnquotedIdentifier("height"):      nil,
			},
			test: func(t *testing.T, tc testCase, actual any) {
				assert.ObjectsAreEqual(tc.expected, actual)
			},
		},
		{
			input: ":hello",
			expected: &ast.Marker{
				Name: ast.UnquotedIdentifier("hello"),
			},
		},
		{
			input:    "?",
			expected: &ast.Marker{},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.Value() })
}

func TestParseTableClusteringOrder(t *testing.T) {
	cases := []testCase{
		{
			input: "a asc",
			expected: &ast.ColumnOrder{
				Name:      ast.UnquotedIdentifier("a"),
				Ascending: true,
			},
		},
		{
			input: "a desc",
			expected: &ast.ColumnOrder{
				Name:      ast.UnquotedIdentifier("a"),
				Ascending: false,
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.TableClusteringOrder() })
}

func TestParseStatementDropTable(t *testing.T) {
	cases := []testCase{
		{
			input: "drop table ks.test",
			expected: &ast.DropTableStatement{
				ColumnFamily: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks"),
					Name:     ast.UnquotedIdentifier("test"),
				},
			},
		},
		{
			input: "drop table if exists ks.test",
			expected: &ast.DropTableStatement{
				IfExists: true,
				ColumnFamily: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks"),
					Name:     ast.UnquotedIdentifier("test"),
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropTableStatement() })
}

func TestParseStatementTruncate(t *testing.T) {
	cases := []testCase{
		{
			input: "truncate ks.test",
			expected: &ast.TruncateStatement{
				ColumnFamily: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks"),
					Name:     ast.UnquotedIdentifier("test"),
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.TruncateStatement() })
}

func TestParseStatementCreateKeyspace(t *testing.T) {
	cases := []testCase{
		{
			input: `CREATE KEYSPACE cycling
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  }`,
			expected: &ast.CreateKeyspaceStatement{
				Name: ast.UnquotedIdentifier("cycling"),
				Properties: []*ast.Property{
					{
						Key: ast.UnquotedIdentifier("REPLICATION"),
						Value: &ast.MapLiteral{
							"class":              "SimpleStrategy",
							"replication_factor": 1,
						},
					},
				},
			},
		},
		{
			input: `CREATE KEYSPACE if not exists cycling
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  }`,
			expected: &ast.CreateKeyspaceStatement{
				IfNotExists: true,
				Name:        ast.UnquotedIdentifier("cycling"),
				Properties: []*ast.Property{
					{
						Key: ast.UnquotedIdentifier("REPLICATION"),
						Value: &ast.MapLiteral{
							"class":              "SimpleStrategy",
							"replication_factor": 1,
						},
					},
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.CreateKeyspaceStatement() })
}

func TestParseStatementDropType(t *testing.T) {
	cases := []testCase{
		{
			input: "drop type ks.test",
			expected: &ast.DropTypeStatement{
				Name: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks"),
					Name:     ast.UnquotedIdentifier("test"),
				},
			},
		},
		{
			input: "drop type test",
			expected: &ast.DropTypeStatement{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
			},
		},
		{
			input: "drop type if exists test",
			expected: &ast.DropTypeStatement{
				IfExists: true,
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropTypeStatement() })
}

func TestParseStatementListUsers(t *testing.T) {
	cases := []testCase{
		{
			input:    "list users",
			expected: &ast.ListUsersStatement{},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.ListUsersStatement() })
}

func TestParseStatementListSuperUsers(t *testing.T) {
	cases := []testCase{
		{
			input:    "list superusers",
			expected: &ast.ListSuperUsersStatement{},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.ListSuperUsersStatement() })
}

func TestParseStatementDropKeyspace(t *testing.T) {
	cases := []testCase{
		{
			input: "drop keyspace test",
			expected: &ast.DropKeyspaceStatement{
				Name: ast.UnquotedIdentifier("test"),
			},
		},
		{
			input: "drop keyspace if exists test",
			expected: &ast.DropKeyspaceStatement{
				IfExists: true,
				Name:     ast.UnquotedIdentifier("test"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropKeyspaceStatement() })
}

func TestParseIdentity(t *testing.T) {
	types := []testCase{
		{
			input:    "hello",
			expected: ast.UnquotedIdentifier("hello"),
			hasError: false,
		},
		{
			input:    "'hello'",
			expected: ast.QuotedIdentifier("hello"),
			hasError: false,
		},
		{
			input:    `"hello"`,
			hasError: true,
		},
	}

	for _, tc := range types {
		p := NewParser(tc.input)
		st, err := parseRule[ast.Identifier](func() antlr.ParseTree { return p.Identity() })

		if tc.hasError {
			require.NotNil(t, err, "expected error for non-type ident %s", tc.input)
		} else {
			require.Nil(t, err)
			assert.Equal(t, tc.expected, st)
		}
	}
}

func TestParseStatementDropIdentity(t *testing.T) {
	cases := []testCase{
		{
			input: "drop identity test",
			expected: &ast.DropIdentityStatement{
				Identity: ast.UnquotedIdentifier("test"),
			},
		},
		{
			input: "drop identity if exists test",
			expected: &ast.DropIdentityStatement{
				IfExists: true,
				Identity: ast.UnquotedIdentifier("test"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropIdentityStatement() })
}

func TestParseStatementAddIdentity(t *testing.T) {
	cases := []testCase{
		{
			input: "add identity 'user' to role admin",
			expected: &ast.AddIdentityStatement{
				Identity: ast.QuotedIdentifier("user"),
				Role:     ast.UnquotedIdentifier("admin"),
			},
		},
		{
			input: "add identity if not exists 'user' to role admin",
			expected: &ast.AddIdentityStatement{
				IfNotExists: true,
				Identity:    ast.QuotedIdentifier("user"),
				Role:        ast.UnquotedIdentifier("admin"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.AddIdentityStatement() })
}

func TestParseStatementDropUser(t *testing.T) {
	cases := []testCase{
		{
			input: "drop user test",
			expected: &ast.DropUserStatement{
				Username: ast.UnquotedIdentifier("test"),
			},
		},
		{
			input: "drop user if exists test",
			expected: &ast.DropUserStatement{
				IfExists: true,
				Username: ast.UnquotedIdentifier("test"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropUserStatement() })
}

func TestParseStatementDropIndex(t *testing.T) {
	cases := []testCase{
		{
			input: "drop index test",
			expected: &ast.DropIndexStatement{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
			},
		},
		{
			input: "drop index ks.test",
			expected: &ast.DropIndexStatement{
				Name: &ast.ObjectRef{
					Keyspace: ast.UnquotedIdentifier("ks"),
					Name:     ast.UnquotedIdentifier("test"),
				},
			},
		},
		{
			input: "drop index if exists test",
			expected: &ast.DropIndexStatement{
				IfExists: true,
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropIndexStatement() })
}

func TestParseStatementDropMaterializedView(t *testing.T) {
	cases := []testCase{
		{
			input: "drop materialized view test",
			expected: &ast.DropMaterializedViewStatement{
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
			},
		},
		{
			input: "drop materialized view if exists test",
			expected: &ast.DropMaterializedViewStatement{
				IfExists: true,
				Name: &ast.ObjectRef{
					Name: ast.UnquotedIdentifier("test"),
				},
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.DropMaterializedViewStatement() })
}

func TestParseStatementGrantRole(t *testing.T) {
	cases := []testCase{
		{
			input: "grant admin to test",
			expected: &ast.GrantRoleStatement{
				Role:    ast.UnquotedIdentifier("admin"),
				Grantee: ast.UnquotedIdentifier("test"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.GrantRoleStatement() })
}

func TestParseStatementRevokeRole(t *testing.T) {
	cases := []testCase{
		{
			input: "revoke admin from test",
			expected: &ast.RevokeRoleStatement{
				Role:    ast.UnquotedIdentifier("admin"),
				Revokee: ast.UnquotedIdentifier("test"),
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.RevokeRoleStatement() })
}

func TestParseStatementListRoles(t *testing.T) {
	cases := []testCase{
		{
			input: "list roles",
			expected: &ast.ListRolesStatement{
				Recursive: true,
			},
		},
		{
			input: "list roles of test",
			expected: &ast.ListRolesStatement{
				Role:      ast.UnquotedIdentifier("test"),
				Recursive: true,
			},
		},
		{
			input: "list roles of test norecursive",
			expected: &ast.ListRolesStatement{
				Role:      ast.UnquotedIdentifier("test"),
				Recursive: false,
			},
		},
	}

	testEqual[any](t, cases, func(p *Parser) antlr.ParseTree { return p.ListRolesStatement() })
}

func testEqual[T any](t *testing.T, cases []testCase, parser func(*Parser) antlr.ParseTree) {
	t.Helper()

	for _, tc := range cases {
		if tc.skip != "" {
			t.Log(tc.skip)
			continue
		}
		p := NewParser(tc.input)
		st, err := parseRule[T](func() antlr.ParseTree { return parser(p) })
		if err != nil {
			t.Fatal(err)
		}

		if tc.test == nil {
			assert.Equal(t, tc.expected, st)
		} else {
			tc.test(t, tc, st)
		}
	}
}
