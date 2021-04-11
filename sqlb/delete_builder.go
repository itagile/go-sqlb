package sqlb

import (
	"strings"
)

// DeleteBuilder generates simple UPDATE from values
type DeleteBuilder interface {
	Setter
	SQLBuilder
}

type deleteBuilderData struct {
	engine Engine
	table  string
	where  *predicateData
}

// NewDeleteBuilder constructs an DeleteBuilder with the provided ParameterPlaceholder
func NewDeleteBuilder(engine Engine, table string) *deleteBuilderData {
	return &deleteBuilderData{
		engine: engine,
		table:  table,
	}
}

// Where for simple where condition initialization with AND operator
func (d *deleteBuilderData) Where(conditions ...Condition) *predicateData {
	d.where = NewAnd(conditions...)
	return d.where
}

// Where for simple where condition initialization with OR operator
func (d *deleteBuilderData) WhereOr(conditions ...Condition) *predicateData {
	d.where = NewOr(conditions...)
	return d.where
}

// Build the UPDATE command
func (d *deleteBuilderData) Build() (query string, args []interface{}) {
	if d.table == "" {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("DELETE FROM ")
	sb.WriteString(d.table)
	args = d.addWhere(&sb, args)
	return sb.String(), args
}

// addWhere appends WHERE clause
func (d *deleteBuilderData) addWhere(sb *strings.Builder, args []interface{}) []interface{} {
	queryWhere, argsWhere := d.where.Build(d.engine)
	if len(queryWhere) > 0 {
		sb.WriteString("\nWHERE ")
		sb.WriteString(queryWhere)
		args = append(args, argsWhere...)
	}
	return args
}
