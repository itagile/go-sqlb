package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewUpdateBuilderWithoutWhere(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?,
Col2 = ?`
	upd := sqlb.NewUpdateBuilder(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col2", "2")
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{1, "2"}, args)
}

func TestNewUpdateBuilderWithWhere(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?,
Col2 = ?
WHERE ID = ?`
	upd := sqlb.NewUpdateBuilder(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col2", "2")
	upd.Where(sqlb.Expr("ID").Eq(1))
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{1, "2", 1}, args)
}

func TestNewUpdateBuilderWithWhereOr(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?,
Col2 = ?
WHERE Col3 LIKE ? OR Col4 = ?`
	upd := sqlb.NewUpdateBuilder(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col2", "2")
	upd.WhereOr(sqlb.Expr("Col3").Like("like1"), sqlb.Expr("Col4").Eq(2))
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{1, "2", "like1", 2}, args)
}

func TestEmptyUpdateBuilderWhenNoColumnsSet(t *testing.T) {
	expected := ""
	ins := sqlb.NewUpdateBuilder(sqlb.DefaultEngine(), "schema.myTable")
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestEmptyUpdateBuilderWhenNoTableName(t *testing.T) {
	expected := ""
	ins := sqlb.NewUpdateBuilder(sqlb.DefaultEngine(), "")
	ins.Set("Col1", 1)
	query, args := ins.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestUpdateBuilderWhenValueChanged(t *testing.T) {
	expected := `UPDATE schema.myTable SET
Col1 = ?
WHERE ID = ?`
	upd := sqlb.NewUpdateBuilder(sqlb.DefaultEngine(), "schema.myTable")
	upd.Set("Col1", 1)
	upd.Set("Col1", 2)
	upd.Where(sqlb.Expr("ID").Eq(1))
	query, args := upd.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{2, 1}, args)
}
