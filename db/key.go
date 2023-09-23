package db

type Key struct {
	BasicConstraint
}

func NewKey(table Table, name string, columns ...Column) *Key {
	return new(Key).Init(table, name, columns...)
}

func (self *Key) Init(table Table, name string, columns ...Column) *Key {
	self.BasicConstraint.Init(table, name, columns...)
	return self
}

func (self *Key) ConstraintType() string {
	if self == self.table.PrimaryKey() {
		return "PRIMARY KEY"
	}

	return "UNIQUE"
}

func (self Key) Create(tx *Transaction) error {
	return tx.ExecSQL(self.CreateSQL())
}

func (self Key) CreateSQL() string {
	return ConstraintCreateSQL(&self)
}

func (self Key) Drop(tx *Transaction) error {
	return tx.ExecSQL(self.DropSQL())
}

func (self Key) DropSQL() string {
	return ConstraintDropSQL(&self)
}
