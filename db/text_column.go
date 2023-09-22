package db

type TextColumn struct {
	BasicColumn
}

func (self *TextColumn) Columns() []Column {
	return []Column{self}
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
