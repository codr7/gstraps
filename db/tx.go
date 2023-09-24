package db

import (
	"github.com/jackc/pgx/v5"
)

type Tx struct {
	StoredValues
	cx  *Cx
	imp pgx.Tx
}

func (self *Tx) Init(cx *Cx, imp pgx.Tx) *Tx {
	self.StoredValues.Init()
	self.cx = cx
	self.imp = imp
	return self
}

func (self *Tx) Commit() error {
	if err := self.imp.Commit(self.cx.cx); err != nil {
		return err
	}

	for f, v := range self.storedValues {
		self.cx.storedValues[f] = v
	}

	return nil
}

func (self *Tx) ExecSQL(sql string, params ...any) error {
	_, err := self.imp.Exec(self.cx.cx, sql, params...)
	return err
}

func (self *Tx) Rollback() error {
	return self.imp.Rollback(self.cx.cx)
}

func (self Tx) StoredValue(field *Field) (any, bool) {
	v, ok := self.StoredValues.StoredValue(field)

	if !ok {
		v, ok = self.cx.StoredValue(field)
	}

	return v, ok
}
