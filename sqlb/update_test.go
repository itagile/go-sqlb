package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewUpdateWithoutWhere(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?,
Col2 = ?`
	upd := sqlb.NewUpdate(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col2", "2")
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestNewUpdateWithWhere(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?,
Col2 = ?
WHERE ID = ?`
	upd := sqlb.NewUpdate(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col2", "2")
	upd.Where(sqlb.Expr("ID").Eq(1))
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2", 1}, args)
}

func TestNewUpdateWithWhereOr(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?,
Col2 = ?
WHERE Col3 LIKE ? OR Col4 = ?`
	upd := sqlb.NewUpdate(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col2", "2")
	upd.WhereOr(sqlb.Expr("Col3").Like("like1"), sqlb.Expr("Col4").Eq(2))
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2", "like1", 2}, args)
}

func TestEmptyUpdateWhenNoColumnsSet(t *testing.T) {
	expected := ""
	ins := sqlb.NewUpdate(sqlb.DefaultEngine(), "schema.myTable")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestEmptyUpdateWhenNoTableName(t *testing.T) {
	expected := ""
	ins := sqlb.NewUpdate(sqlb.DefaultEngine(), "")
	ins.Set("Col1", 1)
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestUpdateWhenValueChanged(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?
WHERE ID = ?`
	upd := sqlb.NewUpdate(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col1", 2)
	upd.Where(sqlb.Expr("ID").Eq(1))
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{2, 1}, args)
}
