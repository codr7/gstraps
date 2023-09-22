package db

type Index struct {
	BasicColumnsDefinition
}

func (self *Index) Init(table Table, name string, columns ...Column) *Index {
	self.BasicColumnsDefinition.Init(table, name, columns...)
	return self
}
