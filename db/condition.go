package db

import (
	"fmt"
)

type Condition struct {
	sql    string
	params []any
}

func NewCondition(sql string, params ...any) Condition {
	return Condition{sql, params}
}

func (self Condition) And(other Condition) Condition {
	return NewCondition(fmt.Sprintf("(%v) AND (%v)", self.sql, other.sql),
		append(self.params, other.params...))
}

func (self Condition) Or(other Condition) Condition {
	return NewCondition(fmt.Sprintf("(%v) OR (%v)", self.sql, other.sql),
		append(self.params, other.params...))
}

func And(conds ...Condition) Condition {
	l := conds[0]

	for _, r := range conds[1:] {
		l = l.And(r)
	}

	return l
}

func Or(conds ...Condition) Condition {
	l := conds[0]

	for _, r := range conds[1:] {
		l = l.Or(r)
	}

	return l
}
