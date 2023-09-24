package db

import (
	"fmt"
)

type Definition interface {
	DefinitionType() string
	Name() string
	//Exists() (bool, error)
	Create(tx *Tx) error
	CreateSQL() string
	Drop(tx *Tx) error
	DropSQL() string
	SQLName() string
}

type BasicDefinition struct {
	name string
}

func (self *BasicDefinition) Init(name string) *BasicDefinition {
	self.name = name
	return self
}

func (self BasicDefinition) Name() string {
	return self.name
}

func (self BasicDefinition) SQLName() string {
	return fmt.Sprintf("\"%v\"", self.name)
}

func DefinitionCreateSQL(def Definition) string {
	return fmt.Sprintf("CREATE %v %v", def.DefinitionType(), def.SQLName())
}

func DefinitionDropSQL(def Definition) string {
	return fmt.Sprintf("DROP %v %v", def.DefinitionType(), def.SQLName())
}
