package db

type Definition interface {
	Name() string
}

type BasicDefinition struct {
	name string
}
