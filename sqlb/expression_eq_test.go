package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestSingleEq(t *testing.T) {
	expected := "Col1 = ?"
	expr := sqlb.Expr("Col1").Eq(1)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []any{1}, args)
}

func TestMultipleEq(t *testing.T) {
	expected := "(Col1 = ? OR Col2 = ?)"
	expr := sqlb.Expr("Col1", "Col2").Eq(1)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []any{1, 1}, args)
}

func TestSingleEqIsNull(t *testing.T) {
	expected := "Col1 IS NULL"
	expr := sqlb.Expr("Col1").Eq(nil)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Empty(t, args)
}

func TestMultipleEqIsNull(t *testing.T) {
	expected := "(Col1 IS NULL OR Col2 IS NULL)"
	expr := sqlb.Expr("Col1", "Col2").Eq(nil)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Empty(t, args)
}
