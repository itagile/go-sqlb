package sqlb_test

import (
	"strings"
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestSingleILike(t *testing.T) {
	expected := "UPPER(Col1) LIKE ?"
	expr := sqlb.Expr("Col1").ILike(likeTest)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	likeTestExpected := strings.ToUpper(likeTest)
	require.Equal(t, []interface{}{likeTestExpected}, args)
}

func TestMultipleILike(t *testing.T) {
	expected := "(UPPER(Col1) LIKE ? OR UPPER(Col2) LIKE ?)"
	expr := sqlb.Expr("Col1", "Col2").ILike(likeTest)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	likeTestExpected := strings.ToUpper(likeTest)
	require.Equal(t, []interface{}{likeTestExpected, likeTestExpected}, args)
}
