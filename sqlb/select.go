package sqlb

import (
	"strings"
)

type Select struct {
	rawSelect string
	where     *Condition
	groupBy   []string
	having    *Condition
	orderBy   []string
	args      []any
	engine    Engine
}

// NewSelect constructs an Select with the provided provided engine and select statement.
func NewSelect(engine Engine, rawSelect string, args ...any) *Select {
	return &Select{
		rawSelect: rawSelect,
		args:      args,
		engine:    engine,
	}
}

// Where for simple where condition initialization with AND operator.
func (s *Select) Where(conditions ...ExpressionBuilder) *Condition {
	s.where = NewAnd(conditions...)
	return s.where
}

// WhereOr for simple where condition initialization with OR operator.
func (s *Select) WhereOr(conditions ...ExpressionBuilder) *Condition {
	s.where = NewOr(conditions...)
	return s.where
}

// GroupBy adds columns to GROUP BY statement.
func (s *Select) GroupBy(columns ...string) {
	s.groupBy = append(s.groupBy, columns...)
}

// OrderBy adds columns to ORDER BY statement.
func (s *Select) OrderBy(columns ...string) {
	s.orderBy = append(s.orderBy, columns...)
}

// Adds ORDER keyword to column name for ORDER BY keyword.
func order(column string, order string) string {
	return column + " " + order
}

// Adds ASC keyword to column name for order by.
func Asc(column string) string {
	return order(column, "ASC")
}

// Adds DESC keyword to column name for order by.
func Desc(column string) string {
	return order(column, "DESC")
}

// Having for simple having condition initialization with AND operator.
func (s *Select) Having(conditions ...ExpressionBuilder) *Condition {
	s.having = NewAnd(conditions...)
	return s.having
}

// HavingOr for simple having condition initialization with OR operator.
func (s *Select) HavingOr(conditions ...ExpressionBuilder) *Condition {
	s.having = NewOr(conditions...)
	return s.having
}

// includesClause detects if raw query includes clause. Query parameter comes with ToUpper applied.
func includesClause(rawSelectUpper string, clause string) bool {
	// TODO remove parentheses from rawSelectUpper
	index := strings.LastIndex(rawSelectUpper, clause)
	if index < 0 {
		return false
	}
	whereClause := rawSelectUpper[index+len(clause):]
	// Counts the number of parentheses in WHERE clause
	openCount := strings.Count(whereClause, "(")
	closeCount := strings.Count(whereClause, ")")
	// If parentheses are balanced returns true
	return openCount == closeCount
}

// addClause adds the clause to the sb using the predicate.
func addClause(clause string, p Predicate, engine Engine, sb *strings.Builder,
	args []any, rawSelectUpper string) []any {
	if p == nil {
		return args
	}
	queryClause, argsClause := p.Build(engine)
	if len(queryClause) == 0 {
		return args
	}
	var text string
	if includesClause(rawSelectUpper, clause) {
		text = p.Operator()
	} else {
		text = clause
	}
	sb.WriteRune('\n')
	sb.WriteString(text)
	sb.WriteRune(' ')
	sb.WriteString(queryClause)
	return append(args, argsClause...)
}

func addCommaSeparated(keyword string, slice []string, sb *strings.Builder) {
	if len(slice) > 0 {
		sb.WriteRune('\n')
		sb.WriteString(keyword)
		sb.WriteRune(' ')
		sb.WriteString(strings.Join(slice, ", "))
	}
}

// Build the SELECT command.
func (s *Select) Build() (query string, args []any) {
	if s.rawSelect == "" {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString(strings.TrimSpace(s.rawSelect))
	rawSelectUpper := strings.ToUpper(s.rawSelect)
	args = s.args
	// Appends WHERE clause if not present
	args = addClause("WHERE", s.where, s.engine, &sb, args, rawSelectUpper)
	addCommaSeparated("GROUP BY", s.groupBy, &sb)
	// Appends HAVING clause if not present
	args = addClause("HAVING", s.having, s.engine, &sb, args, rawSelectUpper)
	addCommaSeparated("ORDER BY", s.orderBy, &sb)
	return sb.String(), args
}
