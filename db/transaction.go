package db

import (
	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	connection   *Connection
	imp          pgx.Tx
	storedFields StoredFields
}

func (self *Transaction) Init(connection *Connection, imp pgx.Tx) *Transaction {
	self.connection = connection
	self.imp = imp
	self.storedFields = make(StoredFields)
	return self
}

func (self *Transaction) Commit() error {
	if err := self.imp.Commit(self.connection.cx); err != nil {
		return err
	}

	for f, v := range self.storedFields {
		self.connection.storedFields[f] = v
	}

	return nil
}

func (self *Transaction) ExecSQL(sql string, params ...any) error {
	_, err := self.imp.Exec(self.connection.cx, sql, params...)
	return err
}

func (self *Transaction) Rollback() error {
	return self.imp.Rollback(self.connection.cx)
}
