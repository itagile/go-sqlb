package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestNewAndEmpty(t *testing.T) {
	expected := ""
	pre := sqlb.NewAnd()
	query, args := pre.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Empty(t, args)
}

func TestNewAndSingle(t *testing.T) {
	expected := "Col1 = ?"
	pre := sqlb.NewAnd(sqlb.Expr("Col1").Eq(1))
	query, args := pre.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{1}, args)
}

func TestNewAndMixedOr(t *testing.T) {
	expected := "Col1 = ? AND (Col2 = ? OR Col3 = ?)"
	pre := sqlb.NewAnd()
	pre.And(sqlb.Expr("Col1").Eq(1))
	pre.Or(sqlb.Expr("Col2").Eq(2), sqlb.Expr("Col3").Eq(3))
	query, args := pre.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{1, 2, 3}, args)
}

func TestNewOrMixedAnd(t *testing.T) {
	expected := "Col1 = ? OR (Col2 = ? AND Col3 = ?)"
	pre := sqlb.NewOr()
	pre.Or(sqlb.Expr("Col1").Eq(1))
	pre.And(sqlb.Expr("Col2").Eq(2), sqlb.Expr("Col3").Eq(3))
	query, args := pre.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []interface{}{1, 2, 3}, args)
}
