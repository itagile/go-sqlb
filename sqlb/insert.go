package sqlb

import (
	"strings"
)

// Setter contract to Set value in Insert and Update.
type Setter interface {
	Set(name string, value any)
}

// Builder contract for building the final SQL.
type Builder interface {
	Build() (query string, args []any)
}

// Insert generates simple INSERT from values.
type Insert interface {
	Setter
	Builder
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

type insertData struct {
	*sqlData
}

// NewInsertWith constructs an Insert with the provided Engine.
func NewInsert(engine Engine, table string) *insertData {
	index := map[string]*nameValue{}
	return &insertData{
		sqlData: &sqlData{
			table:  table,
			index:  index,
			engine: engine,
		},
	}
}

// Set column value in Setter contract.
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

// Build INSERT command.
func (i *insertData) Build() (query string, args []any) {
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
