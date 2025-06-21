package sqlb

import (
	"fmt"
	"strings"
)

type Engine interface {
	Placeholderer
	ILike(expression string, like string) (query string, arg string)
}

type defaultEngine struct {
	Placeholderer
}

func DefaultEngine() *defaultEngine {
	return &defaultEngine{
		Placeholderer: QuestionPlaceholder,
	}
}

func (e *defaultEngine) ILike(expression string, like string) (query string, arg string) {
	if expression == "" {
		return "", ""
	}
	return fmt.Sprintf("UPPER(%s) LIKE %s", expression, e.Placeholder()), strings.ToUpper(like)
}

type postgreSQLEngine defaultEngine

func PostgreSQLEngine() *postgreSQLEngine {
	return &postgreSQLEngine{
		Placeholderer: NewDollarPlaceholder(),
	}
}

func (e *postgreSQLEngine) ILike(expression string, like string) (query string, arg string) {
	if expression == "" {
		return "", ""
	}
	return fmt.Sprintf("%s ILIKE %s", expression, e.Placeholder()), like
}

func ORACLEEngine() *defaultEngine {
	return &defaultEngine{
		Placeholderer: NewColonPlaceholder(),
	}
}
