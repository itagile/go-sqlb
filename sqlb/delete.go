package sqlb

import (
	"strings"
)

type Delete struct {
	engine Engine
	table  string
	where  *Condition
}

// NewDelete constructs an Delete with the provided engine and table name.
func NewDelete(engine Engine, table string) *Delete {
	return &Delete{
		engine: engine,
		table:  table,
	}
}

// Where for simple where condition initialization with AND operator.
func (d *Delete) Where(conditions ...ExpressionBuilder) *Condition {
	d.where = NewAnd(conditions...)
	return d.where
}

// Where for simple where condition initialization with OR operator.
func (d *Delete) WhereOr(conditions ...ExpressionBuilder) *Condition {
	d.where = NewOr(conditions...)
	return d.where
}

// Build the UPDATE command.
func (d *Delete) Build() (query string, args []any) {
	if d.table == "" {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("DELETE FROM ")
	sb.WriteString(d.table)
	args = d.addWhere(&sb, args)
	return sb.String(), args
}

// addWhere appends WHERE clause.
func (d *Delete) addWhere(sb *strings.Builder, args []any) []any {
	if d.where == nil {
		return args
	}
	queryWhere, argsWhere := d.where.Build(d.engine)
	if len(queryWhere) == 0 {
		return args
	}
	sb.WriteString("\nWHERE ")
	sb.WriteString(queryWhere)
	return append(args, argsWhere...)
}
