package db

type Table interface {
	Definition
	PrimaryKey() *Key
	SetPrimaryKey(*Key)
}

type BasicTable struct {
	BasicDefinition
	definitions []TableDefinition
	primaryKey  *Key
}

func (self *BasicTable) PrimaryKey() *Key {
	return self.primaryKey
}

func (self *BasicTable) SetPrimaryKey(key *Key) {
	self.primaryKey = key
}
