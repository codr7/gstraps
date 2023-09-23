package db

type Column interface {
	TableDefinition
	DataType() string
}

type BasicColumn struct {
	BasicTableDefinition
}

func (self *BasicColumn) Create(tx *Transaction) error {
	return nil
}

func (self *BasicColumn) Drop(tx *Transaction) error {
	return nil
}
