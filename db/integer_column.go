package db

type IntegerColumn struct {
	BasicColumn
}

func NewIntegerColumn(table Table, name string) *IntegerColumn {
	return new(IntegerColumn).Init(table, name)
}

func (self *IntegerColumn) Init(table Table, name string) *IntegerColumn {
	self.BasicColumn.Init(table, name)
	table.AddColumn(self)
	return self
}

func (_ IntegerColumn) ColumnType() string {
	return "INTEGER"
}

func (self IntegerColumn) Create(tx *Transaction) error {
	return tx.ExecSQL(self.CreateSQL())
}

func (self IntegerColumn) CreateSQL() string {
	return ColumnCreateSQL(&self)
}

func (self IntegerColumn) Drop(tx *Transaction) error {
	return tx.ExecSQL(self.DropSQL())
}

func (self IntegerColumn) DropSQL() string {
	return ColumnDropSQL(&self)
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
