package sqlb

import "strings"

type ExpressionILike struct {
	Expression
	like string
}

func (e ExpressionILike) Build(engine Engine) (query string, args []any) {
	return BuildExpression(e.Expression, engine, e.buildHandler)
}

func (e ExpressionILike) buildHandler(engine Engine, expression string, sb *strings.Builder) (args []any) {
	query, arg := engine.ILike(expression, e.like)
	sb.WriteString(query)
	args = append(args, arg)
	return args
}

func (e Expression) ILike(like string) ExpressionBuilder {
	return &ExpressionILike{
		Expression: e,
		like:       like,
	}
}
