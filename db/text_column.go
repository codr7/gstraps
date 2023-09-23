package db

type TextColumn struct {
	BasicColumn
}

func NewTextColumn(table Table, name string) *TextColumn {
	return new(TextColumn).Init(table, name)
}

func (self *TextColumn) Init(table Table, name string) *TextColumn {
	self.BasicColumn.Init(table, name)
	return self
}

func (_ TextColumn) DataType() string {
	return "TEXT"
}

func (self *TextColumn) Get(record Record) string {
	if v := record.Get(self); v != nil {
		return v.(string)
	}

	return ""
}

func (self *TextColumn) Set(record Record, value string) {
	record.Set(self, value)
}
