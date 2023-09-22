package db

import (
	"github.com/codr7/gstraps/utils"
)

type Record struct {
	fields utils.Map[Column, any]
}

func (self *Record) Init() *Record {
	self.fields.Init(func(l, r Column) int {
		if c := utils.CompareString(l.Table().Name(), r.Table().Name()); c != 0 {
			return c
		}

		return utils.CompareString(l.Name(), r.Name())
	})

	return self
}

func (self Record) Get(column Column) any {
	if v, ok := self.fields.Find(column); ok {
		return v
	}

	return nil
}

func (self *Record) Set(column Column, value any) {
	self.fields.Upsert(column, value)
}
