package db

import (
	"fmt"
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

func (self Record) Condition(table Table, tx *Tx) (Condition, error) {
	var conds []Condition
	var err error

	for _, c := range table.PrimaryKey().Columns() {
		v, ok := self.StoredValue(c, tx)

		if !ok {
			err = fmt.Errorf("Not stored: %v", c.QualifiedName())
		}

		conds = append(conds, c.Eq(v))
	}

	return And(conds...), err
}

func (self Record) Delete(table Table, tx *Tx) error {
	cond, err := self.Condition(table, tx)

	if err != nil {
		return err
	}

	if err := table.Delete(cond, tx); err != nil {
		return err
	}

	for _, c := range table.Columns() {
		if f := self.getField(c); f != nil {
			tx.DeleteStoredValue(f)
		}
	}

	return nil
}

func (self Record) Fields() []*Field {
	return self.fields.Items()
}

func (self Record) getField(column Column) *Field {
	if f, ok := self.fields.Find(column); ok {
		return f
	}

	return nil
}

func (self Record) Get(column Column) any {
	if f := self.getField(column); f != nil {
		return f.value
	}

	return nil
}

func (self Record) Modified(column Column, tx *Tx) bool {
	if f := self.getField(column); f != nil {
		if v, ok := tx.StoredValue(f); !ok || v != f.value {
			return true
		}
	}

	return false
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

func (self Record) Store(table Table, tx *Tx) error {
	stored := false

	for _, c := range table.Columns() {
		if self.Stored(c, tx) {
			stored = true
			break
		}
	}

	if stored {
		cond, err := self.Condition(table, tx)

		if err != nil {
			return err
		}

		if err := table.Update(self, cond, tx); err != nil {
			return err
		}
	} else {
		if err := table.Insert(self, tx); err != nil {
			return err
		}
	}

	for _, c := range table.Columns() {
		tx.StoreValue(self.getField(c))
	}

	return nil
}

func (self Record) Stored(column Column, tx *Tx) bool {
	if f := self.getField(column); f != nil {
		if _, ok := tx.StoredValue(f); ok {
			return true
		}
	}

	return false
}

func (self Record) StoredValue(column Column, tx *Tx) (any, bool) {
	if f := self.getField(column); f != nil {
		return tx.StoredValue(f)
	}

	return nil, false
}
