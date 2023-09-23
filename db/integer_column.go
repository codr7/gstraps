package db

type IntegerColumn struct {
	BasicColumn
}

func NewIntegerColumn(table Table, name string) *IntegerColumn {
	return new(IntegerColumn).Init(table, name)
}

func (self *IntegerColumn) Init(table Table, name string) *IntegerColumn {
	self.BasicColumn.Init(table, name)
	return self
}

func (_ IntegerColumn) DataType() string {
	return "INTEGER"
}

func (self *IntegerColumn) Get(record Record) int {
	if v := record.Get(self); v != nil {
		return v.(int)
	}

	return 0
}

func (self *IntegerColumn) Set(record *Record, value int) {
	record.Set(self, value)
}
