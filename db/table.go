package db

type Table interface {
	Definition
}

type BasicTable struct {
	BasicDefinition
}
