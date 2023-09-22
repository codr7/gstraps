package db

type Constraint interface {
	TableDefinition
	ConstraintType() string
}

type BasicConstraint struct {
	BasicTableDefinition
	columns []Column
}

func (self *BasicConstraint) Init(table Table, name string, columns ...Column) *BasicConstraint {
	self.BasicTableDefinition.Init(table, name)
	self.columns = columns
	return self
}
