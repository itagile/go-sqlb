package sqlb

import "strings"

type expresisonEqData struct {
	expressionData
	value interface{}
}

func (e expresisonEqData) Build(engine Engine) (query string, args []interface{}) {
	return BuildExpression(e.expressionData, engine, e.buildHandler)
}

func (e expresisonEqData) buildHandler(engine Engine, expression string, sb *strings.Builder) (args []interface{}) {
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

func (e expressionData) Eq(value interface{}) Condition {
	return &expresisonEqData{
		expressionData: e,
		value:          value,
	}
}
