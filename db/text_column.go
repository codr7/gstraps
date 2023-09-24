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

func (self TextColumn) Clone(table Table, name string) Column {
	return NewTextColumn(table, name)
}

func (_ TextColumn) ColumnType() string {
	return "TEXT"
}

func (self *TextColumn) Create(tx *Tx) error {
	return tx.ExecSQL(self.CreateSQL())
}

func (self *TextColumn) CreateSQL() string {
	return ColumnCreateSQL(self)
}

func (self *TextColumn) Drop(tx *Tx) error {
	return tx.ExecSQL(self.DropSQL())
}

func (self *TextColumn) DropSQL() string {
	return ColumnDropSQL(self)
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
