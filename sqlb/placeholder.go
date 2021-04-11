package sqlb

import "fmt"

// ParameterPlaceholder generates de placeholder used in SQL commands
type ParameterPlaceholder interface {
	Placeholder() string
}

// QuestionPlaceholder is a ParameterPlaceholder based on a constant token
type QuestionPlaceholder string

// QuestionPlaceholderData defines the standard placeholder for dbs like mysql
const QuestionPlaceholderData QuestionPlaceholder = "?"

// Get the parameter placeholder
func (t QuestionPlaceholder) Placeholder() string {
	return string(t)
}

// sequencePlaceholderData for placeholders requiring numeric sequence
type sequencePlaceholderData struct {
	prefix  string
	current int
}

// Get the parameter placeholder using this sequence
func (s *sequencePlaceholderData) Placeholder() string {
	s.current++
	return fmt.Sprint(s.prefix, s.current)
}

// NewDollarPlaceholder for PostgreSQL
func NewDollarPlaceholder() ParameterPlaceholder {
	return &sequencePlaceholderData{prefix: "$"}
}

// NewColonPlaceholder for ORACLE
func NewColonPlaceholder() ParameterPlaceholder {
	return &sequencePlaceholderData{prefix: ":v"}
}
