package db

type ForeignKey struct {
	BasicConstraint
}

func (self *ForeignKey) DataType() string {
	return "FOREIGN KEY"
}
