package sqlb

import "strings"

type ExpressionLike struct {
	Expression
	like string
}

func (e ExpressionLike) Build(engine Engine) (query string, args []any) {
	return BuildExpression(e.Expression, engine, e.buildHandler)
}

func (e ExpressionLike) buildHandler(engine Engine, expression string, sb *strings.Builder) (args []any) {
	sb.WriteString(expression)
	sb.WriteString(" LIKE ")
	sb.WriteString(engine.Placeholder())
	args = append(args, e.like)
	return args
}

func (e Expression) Like(like string) ExpressionBuilder {
	return &ExpressionLike{
		Expression: e,
		like:       like,
	}
}
