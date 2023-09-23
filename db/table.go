package db

import (
	"log"
)

type Table interface {
	Definition
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

func (self *BasicTable) PrimaryKey() *Key {
	return self.primaryKey
}

func (self *BasicTable) SetPrimaryKey(key *Key) {
	if key.Table().Name() != self.name {
		log.Fatalf("Key %v does not belong to table %v", key.Name(), self.name)
	}

	self.primaryKey = key
}
