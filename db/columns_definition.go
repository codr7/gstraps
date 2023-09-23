package db

type ColumnsDefinition interface {
	TableDefinition
	Columns() []Column
}

type BasicColumnsDefinition struct {
	BasicTableDefinition
	columns []Column
}

func (self *BasicColumnsDefinition) Init(table Table, name string, columns ...Column) *BasicColumnsDefinition {
	self.BasicTableDefinition.Init(table, name)
	self.columns = columns
	return self
}

func (self *BasicColumnsDefinition) Columns() []Column {
	return self.columns
}
