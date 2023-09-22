package db

type Column interface {
	TableDefinition
	DataType() string
}

type BasicColumn struct {
	BasicTableDefinition
}
