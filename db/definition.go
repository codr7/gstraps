package db

type Definition interface {
	Name() string
	//Exists() (bool, error)
	//Create() error
	//Drop() error
}

type BasicDefinition struct {
	name string
}

func (self *BasicDefinition) Init(name string) *BasicDefinition {
	self.name = name
	return self
}

func (self *BasicDefinition) Name() string {
	return self.name
}
