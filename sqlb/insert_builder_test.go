package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewInsertBuilder(t *testing.T) {
	expected := `INSERT INTO schema.myTable (Col1, Col2)
VALUES (?, ?)`
	ins := sqlb.NewInsertBuilder(sqlb.DefaultEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	ins.Set("Col2", "2")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestNewPostgreSQLInsertBuilder(t *testing.T) {
	expected := `INSERT INTO schema.myTable (Col1)
VALUES ($1)`
	ins := sqlb.NewInsertBuilder(sqlb.PostgreSQLEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1}, args)
}

func TestNewORAInsertBuilder(t *testing.T) {
	expected := `INSERT INTO schema.myTable (Col1, Col2)
VALUES (:v1, :v2)`
	ins := sqlb.NewInsertBuilder(sqlb.ORACLEEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	ins.Set("Col2", "2")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestEmptyInsertBuilderWhenNoColumnsSet(t *testing.T) {
	expected := ""
	ins := sqlb.NewInsertBuilder(sqlb.DefaultEngine(), "schema.myTable")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestEmptyInsertBuilderWhenNoTableName(t *testing.T) {
	expected := ""
	ins := sqlb.NewInsertBuilder(sqlb.DefaultEngine(), "")
	ins.Set("Col1", 1)
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestInsertBuilderWhenValueChanged(t *testing.T) {
	ins := sqlb.NewInsertBuilder(sqlb.DefaultEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	ins.Set("Col1", 2)
	_, args := ins.Build()
	require.Equal(t, []any{2}, args)
}
