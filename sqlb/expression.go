package sqlb

import "strings"

type Expression []string

func Expr(expressions ...string) Expression {
	return expressions
}

type ExpressionHandler func(engine Engine, expression string, sb *strings.Builder) []any

func (e Expression) Build(engine Engine) (query string, args []any) {
	return BuildExpression(e, engine, func(engine Engine, expression string, sb *strings.Builder) []any {
		sb.WriteString(expression)
		return nil
	})
}

func BuildExpression(e Expression, engine Engine, handler ExpressionHandler) (query string, args []any) {
	var sb strings.Builder
	args = make([]any, 0, len(e))
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
