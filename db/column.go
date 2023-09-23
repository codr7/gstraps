package db

import (
	"fmt"
	"strings"
)

type Column interface {
	TableDefinition
	ColumnType() string
}

type BasicColumn struct {
	BasicTableDefinition
}

func (_ BasicColumn) DefinitionType() string {
	return "COLUMN"
}

func ColumnCreateSQL(col Column) string {
	return fmt.Sprintf("%v %v", TableDefinitionCreateSQL(col), col.ColumnType())
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
