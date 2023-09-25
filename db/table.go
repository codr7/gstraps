package db

import (
	"fmt"
	"log"
)

type Table interface {
	Definition
	AddColumn(Column)
	AddDefinition(Definition)
	AddForeignKey(*ForeignKey)
	AddIndex(*Index)
	PrimaryKey() *Key
	SetPrimaryKey(*Key)
}

type BasicTable struct {
	BasicDefinition
	columns     []Column
	definitions []Definition
	foreignKeys []*ForeignKey
	indexes     []*Index
	primaryKey  *Key
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
	self.AddDefinition(column)
}

func (self *BasicTable) AddDefinition(definition Definition) {
	self.definitions = append(self.definitions, definition)
}

func (self *BasicTable) AddForeignKey(key *ForeignKey) {
	self.foreignKeys = append(self.foreignKeys, key)
	self.AddDefinition(key)
}

func (self *BasicTable) AddIndex(index *Index) {
	self.indexes = append(self.indexes, index)
	self.AddDefinition(index)
}

func (self BasicTable) Create(tx *Tx) error {
	if err := tx.ExecSQL(self.CreateSQL()); err != nil {
		return err
	}

	for _, d := range self.definitions {
		if err := d.Create(tx); err != nil {
			return err
		}
	}

	return nil
}

func (self BasicTable) CreateSQL() string {
	return fmt.Sprintf("%v ()", DefinitionCreateSQL(&self))
}

func (_ BasicTable) DefinitionType() string {
	return "TABLE"
}

func (self BasicTable) Drop(tx *Tx) error {
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
