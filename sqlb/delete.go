package sqlb

import (
	"strings"
)

// Delete generates simple UPDATE from values.
type Delete interface {
	WhereBuilder
	Builder
}

type deleteData struct {
	engine Engine
	table  string
	where  *predicateData
}

// NewDelete constructs an Delete with the provided engine and table name.
func NewDelete(engine Engine, table string) *deleteData {
	return &deleteData{
		engine: engine,
		table:  table,
	}
}

// Where for simple where condition initialization with AND operator.
func (d *deleteData) Where(conditions ...Condition) *predicateData {
	d.where = NewAnd(conditions...)
	return d.where
}

// Where for simple where condition initialization with OR operator.
func (d *deleteData) WhereOr(conditions ...Condition) *predicateData {
	d.where = NewOr(conditions...)
	return d.where
}

// Build the UPDATE command.
func (d *deleteData) Build() (query string, args []any) {
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
func (d *deleteData) addWhere(sb *strings.Builder, args []any) []any {
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
