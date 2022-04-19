package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

const likeTest = "like"

func TestSingleLike(t *testing.T) {
	expected := "Col1 LIKE ?"
	expr := sqlb.Expr("Col1").Like(likeTest)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []any{likeTest}, args)
}

func TestMultipleLike(t *testing.T) {
	expected := "(Col1 LIKE ? OR Col2 LIKE ?)"
	expr := sqlb.Expr("Col1", "Col2").Like(likeTest)
	query, args := expr.Build(sqlb.DefaultEngine())
	require.Equal(t, expected, query)
	require.Equal(t, []any{likeTest, likeTest}, args)
}
