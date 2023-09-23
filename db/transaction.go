package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	cx  context.Context
	imp pgx.Tx
}

func (self *Transaction) Init(cx context.Context, imp pgx.Tx) *Transaction {
	self.cx = cx
	self.imp = imp
	return self
}

func (self *Transaction) Commit() error {
	return self.imp.Commit(self.cx)
}

func (self *Transaction) ExecSQL(sql string, params ...any) error {
	_, err := self.imp.Exec(self.cx, sql, params...)
	return err
}

func (self *Transaction) Rollback() error {
	return self.imp.Rollback(self.cx)
}
