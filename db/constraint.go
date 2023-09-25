package db

import (
	"fmt"
)

type Constraint interface {
	ColumnsDefinition
	ConstraintType() string
}

type BasicConstraint struct {
	BasicColumnsDefinition
}

func (self *BasicConstraint) Init(table Table, name string, columns ...Column) *BasicConstraint {
	self.BasicColumnsDefinition.Init(table, name, columns...)
	return self
}

func (_ BasicConstraint) DefinitionType() string {
	return "CONSTRAINT"
}

func ConstraintCreateSQL(ctr Constraint) string {
	return fmt.Sprintf("%v %v (%v)", TableDefinitionCreateSQL(ctr), ctr.ConstraintType(), ColumnsSQL(ctr.Columns()...))
}

func ConstraintDropSQL(ctr Constraint) string {
	return TableDefinitionDropSQL(ctr)
}
