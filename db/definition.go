package db

import (
	"fmt"
)

type Definition interface {
	Name() string
	SQLName() string
	//Exists() (bool, error)
	Create(tx *Transaction) error
	Drop(tx *Transaction) error
}

type BasicDefinition struct {
	name string
}

func (self *BasicDefinition) Init(name string) *BasicDefinition {
	self.name = name
	return self
}

func (self *BasicDefinition) Name() string {
	return self.name
}

func (self *BasicDefinition) SQLName() string {
	return fmt.Sprintf("\"%v\"", self.name)
}
