package sqlb

import "strings"

type ExpressionEq struct {
	Expression
	value any
}

func (e ExpressionEq) Build(engine Engine) (query string, args []any) {
	return BuildExpression(e.Expression, engine, e.buildHandler)
}

func (e ExpressionEq) buildHandler(engine Engine, expression string, sb *strings.Builder) (args []any) {
	sb.WriteString(expression)
	if e.value == nil {
		sb.WriteString(" IS NULL")
	} else {
		sb.WriteString(" = ")
		sb.WriteString(engine.Placeholder())
		args = append(args, e.value)
	}
	return args
}

func (e Expression) Eq(value any) ExpressionBuilder {
	return &ExpressionEq{
		Expression: e,
		value:      value,
	}
}
