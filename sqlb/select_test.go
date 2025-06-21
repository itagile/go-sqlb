package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewRawSelectWithoutWhere(t *testing.T) {
	expected := `SELECT * FROM TABLE
WHERE Col1 = ? AND Col2 = ?`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), "SELECT * FROM TABLE")
	where := sel.Where()
	where.And(sqlb.Expr("Col1").Eq(1))
	where.And(sqlb.Expr("Col2").Eq("2"))
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestNewRawSelectWithWhere(t *testing.T) {
	expected := `SELECT * FROM TABLE
WHERE Col3 IS NULL
OR Col1 = ? OR Col2 = ?`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT * FROM TABLE
WHERE Col3 IS NULL`)
	where := sel.WhereOr()
	where.Or(sqlb.Expr("Col1").Eq(1))
	where.Or(sqlb.Expr("Col2").Eq("2"))
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestEmptyRawSelect(t *testing.T) {
	expected := ""
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), "")
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestNewRawSelectWithGroupBy(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
GROUP BY Col1, Col2`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.GroupBy("Col1", "Col2")
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestNewRawSelectWithSimpleGroupBy(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
GROUP BY Col1, Col2`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.GroupBy("Col1, Col2")
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestNewRawSelectWithHaving(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
GROUP BY Col1, Col2
HAVING COUNT(*) > 1`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.GroupBy("Col1, Col2")
	sel.Having(sqlb.Expr("COUNT(*) > 1"))
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestNewRawSelectWithHavingOr(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
GROUP BY Col1, Col2
HAVING Col1 = ? OR Col2 LIKE ?`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.GroupBy("Col1, Col2")
	sel.HavingOr(sqlb.Expr("Col1").Eq(1), sqlb.Expr("Col2").Like("2"))
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, "2"}, args)
}

func TestNewRawSelectWithOrderBy(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
ORDER BY Col1, Col2`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.OrderBy("Col1", "Col2")
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestNewRawSelectWithSimpleOrderBy(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
ORDER BY Col1, Col2 DESC`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.OrderBy("Col1, Col2 DESC")
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}

func TestNewRawSelectWithOrderByExplicit(t *testing.T) {
	expected := `SELECT Col1, Col2 FROM TABLE
ORDER BY Col1 DESC, Col2 ASC`
	sel := sqlb.NewRawSelect(sqlb.DefaultEngine(), `SELECT Col1, Col2 FROM TABLE`)
	sel.OrderBy(sqlb.Desc("Col1"), sqlb.Asc("Col2"))
	query, args := sel.Build()
	require.Equal(t, expected, query)
	require.Nil(t, args)
}
