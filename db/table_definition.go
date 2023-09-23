package db

type TableDefinition interface {
	Definition
	Table() Table
}

type BasicTableDefinition struct {
	BasicDefinition
	table Table
}

func (self *BasicTableDefinition) Init(table Table, name string) *BasicTableDefinition {
	self.table = table
	self.BasicDefinition.Init(name)
	return self
}

func (self *BasicTableDefinition) Table() Table {
	return self.table
}
