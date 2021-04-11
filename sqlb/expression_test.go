package sqlb_test

import (
	"fmt"
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestSingleExpr(t *testing.T) {
	expected := "Col1 = Col2"
	expr := sqlb.Expr(expected)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Empty(t, args)
}

func TestMultipleExpr(t *testing.T) {
	expr1 := "Col1 = Col2"
	expr2 := "Col3 = Col4"
	expected := fmt.Sprintf("(%s OR %s)", expr1, expr2)
	expr := sqlb.Expr(expected)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Empty(t, args)
}
