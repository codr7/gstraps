package db

import (
	"fmt"
	"strings"
)

type Index struct {
	BasicColumnsDefinition
	Unique bool
}

func NewIndex(table Table, name string, columns ...Column) *Index {
	return new(Index).Init(table, name, columns...)
}

func (self *Index) Init(table Table, name string, columns ...Column) *Index {
	self.BasicColumnsDefinition.Init(table, name, columns...)
	table.AddIndex(self)
	return self
}

func (_ Index) DefinitionType() string {
	return "INDEX"
}

func (self Index) Create(tx *Tx) error {
	return tx.ExecSQL(self.CreateSQL())
}

func (self Index) CreateSQL() string {
	var sql strings.Builder
	sql.WriteString("CREATE")

	if self.Unique {
		sql.WriteString(" UNIQUE")
	}

	fmt.Fprintf(&sql, " INDEX %v ON %v (%v)",
		self.SQLName(), self.table.SQLName(), ColumnsSQL(self.columns...))
	return sql.String()
}

func (self Index) Drop(tx *Tx) error {
	return tx.ExecSQL(self.DropSQL())
}

func (self Index) DropSQL() string {
	return fmt.Sprintf("DROP INDEX %v ON %v", self.SQLName(), self.table.SQLName())
}
