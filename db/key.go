package db

type Key struct {
	BasicConstraint
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
