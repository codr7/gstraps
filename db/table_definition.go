package db

import (
	"fmt"
)

type TableDefinition interface {
	Definition
	Table() Table
}

type BasicTableDefinition struct {
	BasicDefinition
	table Table
}

func (self *BasicTableDefinition) Init(table Table, name string) *BasicTableDefinition {
	self.table = table
	self.BasicDefinition.Init(name)
	return self
}

func (self *BasicTableDefinition) Table() Table {
	return self.table
}

func TableDefinitionCreateSQL(def TableDefinition) string {
	return fmt.Sprintf("ALTER TABLE %v ADD %v %v",
		def.Table().SQLName(), def.DefinitionType(), def.SQLName())
}

func TableDefinitionDropSQL(def TableDefinition) string {
	return fmt.Sprintf("ALTER TABLE %v DROP %v %v",
		def.Table().SQLName(), def.DefinitionType(), def.SQLName())
}
