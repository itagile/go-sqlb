package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewInsert(t *testing.T) {
	expected := `INSERT INTO schema.myTable (Col1, Col2)
VALUES (?, ?)`
	ins := sqlb.NewInsert(sqlb.DefaultEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	ins.Set("Col2", "2")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestNewPostgreSQLInsert(t *testing.T) {
	expected := `INSERT INTO schema.myTable (Col1)
VALUES ($1)`
	ins := sqlb.NewInsert(sqlb.PostgreSQLEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1}, args)
}

func TestNewORAInsert(t *testing.T) {
	expected := `INSERT INTO schema.myTable (Col1, Col2)
VALUES (:v1, :v2)`
	ins := sqlb.NewInsert(sqlb.ORACLEEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	ins.Set("Col2", "2")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestEmptyInsertWhenNoColumnsSet(t *testing.T) {
	expected := ""
	ins := sqlb.NewInsert(sqlb.DefaultEngine(), "schema.myTable")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestEmptyInsertWhenNoTableName(t *testing.T) {
	expected := ""
	ins := sqlb.NewInsert(sqlb.DefaultEngine(), "")
	ins.Set("Col1", 1)
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestInsertWhenValueChanged(t *testing.T) {
	ins := sqlb.NewInsert(sqlb.DefaultEngine(), "schema.myTable")
	ins.Set("Col1", 1)
	ins.Set("Col1", 2)
	_, args := ins.Build()
	require.Equal(t, []any{2}, args)
}
