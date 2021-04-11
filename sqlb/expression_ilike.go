package sqlb

import "strings"

type expresisonILikeData struct {
	expressionData
	like string
}

func (e expresisonILikeData) Build(engine Engine) (query string, args []interface{}) {
	return BuildExpression(e.expressionData, engine, e.buildHandler)
}

func (e expresisonILikeData) buildHandler(engine Engine, expression string, sb *strings.Builder) (args []interface{}) {
	query, arg := engine.ILike(expression, e.like)
	sb.WriteString(query)
	args = append(args, arg)
	return args
}

func (e expressionData) ILike(like string) Condition {
	return &expresisonILikeData{
		expressionData: e,
		like:           like,
	}
}
