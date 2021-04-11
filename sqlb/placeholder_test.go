package sqlb_test

import (
	"testing"

	"github.com/itagile/go-sqlb/sqlb"
	"github.com/stretchr/testify/require"
)

func TestDefaultDollarPlaceholder(t *testing.T) {
	param := sqlb.QuestionPlaceholderData
	require.Equal(t, "?", param.Placeholder())
	require.Equal(t, "?", param.Placeholder())
}

func TestNewDollarPlaceholder(t *testing.T) {
	param := sqlb.NewDollarPlaceholder()
	require.Equal(t, "$1", param.Placeholder())
	require.Equal(t, "$2", param.Placeholder())
}

func TestNewColonPlaceholder(t *testing.T) {
	param := sqlb.NewColonPlaceholder()
	require.Equal(t, ":v1", param.Placeholder())
	require.Equal(t, ":v2", param.Placeholder())
}
