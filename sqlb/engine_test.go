package sqlb_test

import (
	"strings"
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestDefaultEngineILike(t *testing.T) {
	expected := "UPPER(Col1) LIKE ?"
	engine := sqlb.DefaultEngine()
	query, arg := engine.ILike("Col1", likeTest)
	require.Equal(t, expected, query)
	require.Equal(t, strings.ToUpper(likeTest), arg)
}

func TestDefaultEngineILikeEmptyExpression(t *testing.T) {
	expected := ""
	engine := sqlb.DefaultEngine()
	query, arg := engine.ILike("", likeTest)
	require.Equal(t, expected, query)
	require.Equal(t, expected, arg)
}

func TestPostgreSQLEngineILike(t *testing.T) {
	expected := "Col1 ILIKE $1"
	engine := sqlb.PostgreSQLEngine()
	query, arg := engine.ILike("Col1", likeTest)
	require.Equal(t, expected, query)
	require.Equal(t, likeTest, arg)
}

func TestPostgreSQLEngineILikeEmptyExpression(t *testing.T) {
	expected := ""
	engine := sqlb.PostgreSQLEngine()
	query, arg := engine.ILike("", likeTest)
	require.Equal(t, expected, query)
	require.Equal(t, expected, arg)
}

func TestORACLEEngineILike(t *testing.T) {
	expected := "UPPER(Col1) LIKE :v1"
	engine := sqlb.ORACLEEngine()
	query, arg := engine.ILike("Col1", likeTest)
	require.Equal(t, expected, query)
	require.Equal(t, strings.ToUpper(likeTest), arg)
}
