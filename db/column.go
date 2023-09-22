package db

type Column interface {
	Definition
	Table() Table
}

type BasicColumn struct {
	BasicDefinition
}
