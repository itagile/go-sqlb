package sqlb

import (
	"strings"
)

// Setter contract to Set value in InsertBuilder and UpdateBuilder
type Setter interface {
	Set(name string, value any)
}

// SQLBuilder contract for building the final SQL
type SQLBuilder interface {
	Build() (query string, args []any)
}

// InsertBuilder generates simple INSERT from values
type InsertBuilder interface {
	Setter
	SQLBuilder
}

type nameValue struct {
	name  string
	value any
}

type sqlData struct {
	table  string
	index  map[string]*nameValue
	values []*nameValue
	engine Engine
}

type insertBuilderData struct {
	*sqlData
}

// NewInsertBuilderWith constructs an InsertBuilder with the provided Engine
func NewInsertBuilder(engine Engine, table string) *insertBuilderData {
	index := map[string]*nameValue{}
	return &insertBuilderData{
		sqlData: &sqlData{
			table:  table,
			index:  index,
			engine: engine,
		},
	}
}

// Set column value in Setter contract
func (s *sqlData) Set(name string, value any) {
	item, exists := s.index[name]
	if exists {
		item.value = value
	} else {
		item = &nameValue{name: name, value: value}
		s.values = append(s.values, item)
		s.index[name] = item
	}
}

// Build INSERT command
func (i *insertBuilderData) Build() (query string, args []any) {
	if i.table == "" || len(i.values) == 0 {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(i.table)
	sb.WriteString(" (")
	last := len(i.values) - 1
	// Appends column names
	for index, item := range i.values {
		sb.WriteString(item.name)
		if index < last {
			sb.WriteString(", ")
		} else {
			sb.WriteRune(')')
		}
	}
	sb.WriteString("\nVALUES (")
	args = make([]any, 0, len(i.values))
	// Appends values
	for index, item := range i.values {
		sb.WriteString(i.engine.Placeholder())
		if index < last {
			sb.WriteString(", ")
		} else {
			sb.WriteRune(')')
		}
		args = append(args, item.value)
	}
	return sb.String(), args
}
