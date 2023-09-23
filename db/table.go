package db

import (
	"fmt"
	"log"
)

type Table interface {
	Definition
	AddColumn(Column)
	PrimaryKey() *Key
	SetPrimaryKey(*Key)
}

type BasicTable struct {
	BasicDefinition
	columns    []Column
	primaryKey *Key
}

func NewTable(name string) *BasicTable {
	return new(BasicTable).Init(name)
}

func (self *BasicTable) Init(name string) *BasicTable {
	self.BasicDefinition.Init(name)
	return self
}

func (self *BasicTable) AddColumn(column Column) {
	self.columns = append(self.columns, column)
}

func (self BasicTable) Create(tx *Transaction) error {
	if err := tx.ExecSQL(self.CreateSQL()); err != nil {
		return err
	}

	for _, c := range self.columns {
		if err := c.Create(tx); err != nil {
			return err
		}
	}

	if self.primaryKey != nil {
		if err := self.primaryKey.Create(tx); err != nil {
			return err
		}
	}

	return nil
}

func (self BasicTable) CreateSQL() string {
	return fmt.Sprintf("%v ()", DefinitionCreateSQL(&self))
}

func (self BasicTable) DefinitionType() string {
	return "TABLE"
}

func (self BasicTable) Drop(tx *Transaction) error {
	return tx.ExecSQL(self.DropSQL())
}

func (self BasicTable) DropSQL() string {
	return DefinitionDropSQL(self)
}

func (self BasicTable) PrimaryKey() *Key {
	return self.primaryKey
}

func (self *BasicTable) SetPrimaryKey(key *Key) {
	if key.Table().Name() != self.name {
		log.Fatalf("Key %v does not belong to table %v", key.Name(), self.name)
	}

	self.primaryKey = key
}
