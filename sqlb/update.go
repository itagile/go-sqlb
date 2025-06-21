package sqlb

import (
	"strings"
)

// Update generates simple UPDATE from values.
type Update interface {
	Setter
	WhereBuilder
	Builder
}

type updateData struct {
	*sqlData
	where *predicateData
}

// NewUpdate constructs an Update with the provided ParameterPlaceholder.
func NewUpdate(engine Engine, table string) *updateData {
	index := map[string]*nameValue{}
	return &updateData{
		sqlData: &sqlData{
			table:  table,
			index:  index,
			engine: engine,
		},
	}
}

// Where for simple where condition initialization with AND operator.
func (u *updateData) Where(conditions ...Condition) *predicateData {
	u.where = NewAnd(conditions...)
	return u.where
}

// WhereOr for simple where condition initialization with OR operator.
func (u *updateData) WhereOr(conditions ...Condition) *predicateData {
	u.where = NewOr(conditions...)
	return u.where
}

// Build the UPDATE command.
func (u *updateData) Build() (query string, args []any) {
	if u.table == "" || len(u.values) == 0 {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("UPDATE ")
	sb.WriteString(u.table)
	sb.WriteString(" SET\n")
	last := len(u.values) - 1
	// Appends setter for each column.
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

// addWhere appends WHERE clause.
func (u *updateData) addWhere(sb *strings.Builder, args []any) []any {
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
