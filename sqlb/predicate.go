package sqlb

import "strings"

type ExpressionBuilder interface {
	Build(engine Engine) (query string, args []any)
}

type Predicate interface {
	ExpressionBuilder
	// Operator returns the logical operator used in the predicate, e.g., "AND" or "OR".
	Operator() string
	And(conditions ...ExpressionBuilder) *Condition
	Or(conditions ...ExpressionBuilder) *Condition
}

const (
	and = "AND"
	or  = "OR"
)

type Condition struct {
	operator   string
	conditions []ExpressionBuilder
}

func (p *Condition) Operator() string {
	return p.operator
}

func (p *Condition) Build(engine Engine) (query string, args []any) {
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
			p.addCond(&sb, queryCond, operatorSearch)
			args = append(args, argsCond...)
		}
	}
	return sb.String(), args
}

func (p *Condition) addCond(sb *strings.Builder, queryCond string, operatorSearch string) {
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
}

func NewCondition(operator string, conditions []ExpressionBuilder) *Condition {
	return &Condition{
		operator:   operator,
		conditions: conditions,
	}
}

func NewAnd(conditions ...ExpressionBuilder) *Condition {
	return NewCondition(and, conditions)
}

func NewOr(conditions ...ExpressionBuilder) *Condition {
	return NewCondition(or, conditions)
}

func (p *Condition) And(conditions ...ExpressionBuilder) *Condition {
	if p.operator == and {
		p.conditions = append(p.conditions, conditions...)
	} else {
		p.conditions = append(p.conditions, NewAnd(conditions...))
	}
	return p
}

func (p *Condition) Or(conditions ...ExpressionBuilder) *Condition {
	if p.operator == or {
		p.conditions = append(p.conditions, conditions...)
	} else {
		p.conditions = append(p.conditions, NewOr(conditions...))
	}
	return p
}
