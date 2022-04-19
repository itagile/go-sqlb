package sqlb

import "strings"

type Condition interface {
	Build(engine Engine) (query string, args []any)
}

const (
	and = "AND"
	or  = "OR"
)

type predicateData struct {
	operator   string
	conditions []Condition
}

func (p *predicateData) Build(engine Engine) (query string, args []any) {
	if p == nil || len(p.conditions) == 0 {
		return "", nil
	}
	var operatorSearch string
	if p.operator == and {
		operatorSearch = " " + or + " "
	} else {
		operatorSearch = " " + and + " "
	}
	var sb strings.Builder
	args = make([]any, 0, len(p.conditions))
	for _, cond := range p.conditions {
		queryCond, argsCond := cond.Build(engine)
		if queryCond != "" {
			if sb.Len() > 0 {
				sb.WriteRune(' ')
				sb.WriteString(p.operator)
				sb.WriteRune(' ')
			}
			enclose := strings.Contains(queryCond, operatorSearch)
			if enclose {
				sb.WriteRune('(')
			}
			sb.WriteString(queryCond)
			if enclose {
				sb.WriteRune(')')
			}
			args = append(args, argsCond...)
		}
	}
	return sb.String(), args
}

func newPredicateData(operator string, conditions []Condition) *predicateData {
	return &predicateData{
		operator:   operator,
		conditions: conditions,
	}
}

func NewAnd(conditions ...Condition) *predicateData {
	return newPredicateData(and, conditions)
}

func NewOr(conditions ...Condition) *predicateData {
	return newPredicateData(or, conditions)
}

func (p *predicateData) And(conditions ...Condition) *predicateData {
	if p.operator == and {
		p.conditions = append(p.conditions, conditions...)
	} else {
		p.conditions = append(p.conditions, NewAnd(conditions...))
	}
	return p
}

func (p *predicateData) Or(conditions ...Condition) *predicateData {
	if p.operator == or {
		p.conditions = append(p.conditions, conditions...)
	} else {
		p.conditions = append(p.conditions, NewOr(conditions...))
	}
	return p
}
