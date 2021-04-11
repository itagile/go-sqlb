package sqlb

import "strings"

type expressionData []string

func Expr(expressions ...string) expressionData {
	return expressions
}

type ExpressionBuilder func(engine Engine, expression string, sb *strings.Builder) []interface{}

func (e expressionData) Build(engine Engine) (query string, args []interface{}) {
	return BuildExpression(e, engine, func(engine Engine, expression string, sb *strings.Builder) []interface{} {
		sb.WriteString(expression)
		return nil
	})
}

func BuildExpression(e expressionData, engine Engine, handler ExpressionBuilder) (query string, args []interface{}) {
	var sb strings.Builder
	args = make([]interface{}, 0, len(e))
	enclose := false
	for _, expression := range e {
		if expression != "" {
			if sb.Len() > 0 {
				sb.WriteString(" OR ")
				enclose = true
			}
			argsHandler := handler(engine, expression, &sb)
			args = append(args, argsHandler...)
		}
	}
	query = sb.String()
	if enclose {
		query = "(" + query + ")"
	}
	return query, args
}
