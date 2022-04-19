package sqlb

import "strings"

type expresisonLikeData struct {
	expressionData
	like string
}

func (e expresisonLikeData) Build(engine Engine) (query string, args []any) {
	return BuildExpression(e.expressionData, engine, e.buildHandler)
}

func (e expresisonLikeData) buildHandler(engine Engine, expression string, sb *strings.Builder) (args []any) {
	sb.WriteString(expression)
	sb.WriteString(" LIKE ")
	sb.WriteString(engine.Placeholder())
	args = append(args, e.like)
	return args
}

func (e expressionData) Like(like string) Condition {
	return &expresisonLikeData{
		expressionData: e,
		like:           like,
	}
}
