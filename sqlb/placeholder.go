package sqlb

import "fmt"

// Placeholderer generates de placeholder used in SQL commands
type Placeholderer interface {
	Placeholder() string
}

// StringPlaceholder is a Placeholderer based on a constant token
type StringPlaceholder string

// QuestionPlaceholder defines the standard placeholder for dbs like mysql
const QuestionPlaceholder StringPlaceholder = "?"

// Get the parameter placeholder
func (t StringPlaceholder) Placeholder() string {
	return string(t)
}

// SequencePlaceholder for placeholders requiring numeric sequence
type SequencePlaceholder struct {
	prefix  string
	current int
}

// Get the parameter placeholder using this sequence
func (s *SequencePlaceholder) Placeholder() string {
	s.current++
	return fmt.Sprint(s.prefix, s.current)
}

// NewDollarPlaceholder for PostgreSQL
func NewDollarPlaceholder() Placeholderer {
	return &SequencePlaceholder{prefix: "$"}
}

// NewColonPlaceholder for ORACLE
func NewColonPlaceholder() Placeholderer {
	return &SequencePlaceholder{prefix: ":v"}
}
