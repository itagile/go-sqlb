package sqlb

import "strings"

type Condition interface {
	Build(engine Engine) (query string, args []any)
}

type Predicate interface {
	Condition
	// Operator returns the logical operator used in the predicate, e.g., "AND" or "OR".
	Operator() string
	And(conditions ...Condition) *predicateData
	Or(conditions ...Condition) *predicateData
}

const (
	and = "AND"
	or  = "OR"
)

type WhereBuilder interface {
	// Where for simple where condition initialization with AND operator.
	Where(conditions ...Condition) *predicateData
	// WhereOr for simple where condition initialization with OR operator.
	WhereOr(conditions ...Condition) *predicateData
}

type predicateData struct {
	operator   string
	conditions []Condition
}

func (p *predicateData) Operator() string {
	return p.operator
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
			p.addCond(&sb, queryCond, operatorSearch)
			args = append(args, argsCond...)
		}
	}
	return sb.String(), args
}

func (p *predicateData) addCond(sb *strings.Builder, queryCond string, operatorSearch string) {
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
