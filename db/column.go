package db

import (
	"fmt"
	"strings"
)

type ColumnOption string

const (
	NotNull ColumnOption = "NOT NULL"
)

type Column interface {
	TableDefinition
	Clone(table Table, name string) Column
	ColumnType() string
	Eq(any) Condition
	OptionSQL() string
}

type BasicColumn struct {
	BasicTableDefinition
	options []string
}

func (self *BasicColumn) Init(table Table, name string, options ...ColumnOption) *BasicColumn {
	self.BasicTableDefinition.Init(table, name)
	self.options = make([]string, len(options))

	for i, o := range options {
		self.options[i] = string(o)
	}

	return self
}

func (_ BasicColumn) DefinitionType() string {
	return "COLUMN"
}

func (self BasicColumn) Eq(other any) Condition {
	if c, ok := other.(Column); ok {
		return NewCondition(fmt.Sprintf("%v = %v", self.SQLName(), c.SQLName()))
	}

	return NewCondition(fmt.Sprintf("%v = ?", self.SQLName()), other)
}

func (self BasicColumn) OptionSQL() string {
	if self.options == nil {
		return ""
	}

	return fmt.Sprintf(" %v", strings.Join(self.options, " "))
}

func ColumnCreateSQL(col Column) string {
	return fmt.Sprintf("%v %v%v", TableDefinitionCreateSQL(col), col.ColumnType(), col.OptionSQL())
}

func ColumnDropSQL(col Column) string {
	return TableDefinitionDropSQL(col)
}

func ColumnsSQL(columns ...Column) string {
	sqls := make([]string, len(columns))

	for i, c := range columns {
		sqls[i] = c.SQLName()
	}

	return strings.Join(sqls, "; ")
}
