package db

import (
	"github.com/codr7/gstraps/utils"
)

type Record struct {
	fields utils.Set[Column, *Field]
}

type Field struct {
	column Column
	value  any
}

func NewRecord() *Record {
	return new(Record).Init()
}

func (self *Record) Init() *Record {
	self.fields.Init(func(l Column, r *Field) int {
		if c := utils.CompareString(l.Table().Name(), r.column.Table().Name()); c != 0 {
			return c
		}

		return utils.CompareString(l.Name(), r.column.Name())
	})

	return self
}

func (self Record) Get(column Column) any {
	if f, ok := self.fields.Find(column); ok {
		return f.value
	}

	return nil
}

func (self Record) Null(column Column) bool {
	return !self.fields.Member(column)
}

func (self *Record) Set(column Column, value any) {
	if i, ok := self.fields.Index(column); ok {
		self.fields.Get(i).value = value
	} else {
		self.fields.Insert(i, &Field{column, value})
	}
}
