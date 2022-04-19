package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewDeleteBuilderWithoutWhere(t *testing.T) {
	expected := "DELETE FROM schema.myTable"
	del := sqlb.NewDeleteBuilder(sqlb.DefaultEngine(), "schema.myTable")
	query, args := del.Build()
	require.Equal(t, expected, query)
	require.Empty(t, args)
}

func TestNewDeleteBuilderWithWhere(t *testing.T) {
	expected := `DELETE FROM schema.myTable
WHERE ID = ?`
	del := sqlb.NewDeleteBuilder(sqlb.DefaultEngine(), "schema.myTable")
	del.Where(sqlb.Expr("ID").Eq(1))
	query, args := del.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1}, args)
}

func TestNewDeleteBuilderWithWhereOr(t *testing.T) {
	expected := `DELETE FROM schema.myTable
WHERE Col3 LIKE ? OR Col4 = ?`
	del := sqlb.NewDeleteBuilder(sqlb.DefaultEngine(), "schema.myTable")
	del.WhereOr(sqlb.Expr("Col3").Like("like1"), sqlb.Expr("Col4").Eq(2))
	query, args := del.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{"like1", 2}, args)
}

func TestEmptyDeleteBuilderWhenNoTableName(t *testing.T) {
	expected := ""
	ins := sqlb.NewDeleteBuilder(sqlb.DefaultEngine(), "")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}
