package db

import (
	"fmt"
)

type ForeignKey struct {
	BasicConstraint
	foreignTable   Table
	foreignColumns []Column
}

func NewForeignKey(table Table, name string, foreignTable Table, columns ...Column) *ForeignKey {
	return new(ForeignKey).Init(table, name, foreignTable, columns...)
}

func (self *ForeignKey) Init(table Table, name string, foreignTable Table, columns ...Column) *ForeignKey {
	self.foreignTable = foreignTable
	self.foreignColumns = columns
	columns = make([]Column, len(columns))

	for i, c := range self.foreignColumns {
		columns[i] = c.Clone(table, fmt.Sprintf("%v%v", name, c.Name()))
	}

	self.BasicConstraint.Init(table, name, columns...)
	table.AddForeignKey(self)
	return self
}

func (_ ForeignKey) ConstraintType() string {
	return "FOREIGN KEY"
}

func (self ForeignKey) Create(tx *Tx) error {
	return tx.ExecSQL(self.CreateSQL())
}

func (self ForeignKey) CreateSQL() string {
	return fmt.Sprintf("%v REFERENCES %v (%v)",
		ConstraintCreateSQL(&self), self.foreignTable.SQLName(), ColumnsSQL(self.foreignColumns...))
}

func (self ForeignKey) Drop(tx *Tx) error {
	return tx.ExecSQL(self.DropSQL())
}

func (self ForeignKey) DropSQL() string {
	return ConstraintDropSQL(&self)
}
