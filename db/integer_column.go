package db

type IntegerColumn struct {
	BasicColumn
}

func (self *IntegerColumn) Columns() []Column {
	return []Column{self}
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
