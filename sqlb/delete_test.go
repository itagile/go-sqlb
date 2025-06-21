package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewDeleteWithoutWhere(t *testing.T) {
	expected := "DELETE FROM schema.myTable"
	del := sqlb.NewDelete(sqlb.DefaultEngine(), "schema.myTable")
	query, args := del.Build()
	require.Equal(t, expected, query)
	require.Empty(t, args)
}

func TestNewDeleteWithWhere(t *testing.T) {
	expected := `DELETE FROM schema.myTable
WHERE ID = ?`
	del := sqlb.NewDelete(sqlb.DefaultEngine(), "schema.myTable")
	del.Where(sqlb.Expr("ID").Eq(1))
	query, args := del.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1}, args)
}

func TestNewDeleteWithWhereOr(t *testing.T) {
	expected := `DELETE FROM schema.myTable
WHERE Col3 LIKE ? OR Col4 = ?`
	del := sqlb.NewDelete(sqlb.DefaultEngine(), "schema.myTable")
	del.WhereOr(sqlb.Expr("Col3").Like("like1"), sqlb.Expr("Col4").Eq(2))
	query, args := del.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{"like1", 2}, args)
}

func TestEmptyDeleteWhenNoTableName(t *testing.T) {
	expected := ""
	ins := sqlb.NewDelete(sqlb.DefaultEngine(), "")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}
