package sqlb

import (
	"strings"
)

// UpdateBuilder generates simple UPDATE from values
type UpdateBuilder interface {
	Setter
	WhereBuilder
	SQLBuilder
}

type updateBuilderData struct {
	*sqlData
	where *predicateData
}

// NewUpdateBuilder constructs an UpdateBuilder with the provided ParameterPlaceholder
func NewUpdateBuilder(engine Engine, table string) *updateBuilderData {
	index := map[string]*nameValue{}
	return &updateBuilderData{
		sqlData: &sqlData{
			table:  table,
			index:  index,
			engine: engine,
		},
	}
}

// Where for simple where condition initialization with AND operator
func (u *updateBuilderData) Where(conditions ...Condition) *predicateData {
	u.where = NewAnd(conditions...)
	return u.where
}

// WhereOr for simple where condition initialization with OR operator
func (u *updateBuilderData) WhereOr(conditions ...Condition) *predicateData {
	u.where = NewOr(conditions...)
	return u.where
}

// Build the UPDATE command
func (u *updateBuilderData) Build() (query string, args []any) {
	if u.table == "" || len(u.values) == 0 {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("UPDATE ")
	sb.WriteString(u.table)
	sb.WriteString(" SET\n")
	last := len(u.values) - 1
	// Appends setter for each column
	args = make([]any, 0, len(u.values))
	for index, item := range u.values {
		sb.WriteString(item.name)
		sb.WriteString(" = ?")
		if index < last {
			sb.WriteRune(',')
			sb.WriteRune('\n')
		}
		args = append(args, item.value)
	}
	args = u.addWhere(&sb, args)
	return sb.String(), args
}

// addWhere appends WHERE clause
func (u *updateBuilderData) addWhere(sb *strings.Builder, args []any) []any {
	if u.where == nil {
		return args
	}
	queryWhere, argsWhere := u.where.Build(u.engine)
	if len(queryWhere) == 0 {
		return args
	}
	sb.WriteString("\nWHERE ")
	sb.WriteString(queryWhere)
	return append(args, argsWhere...)
}
