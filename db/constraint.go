package db

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
