package db

type ForeignKey struct {
	BasicConstraint
}

func (_ *ForeignKey) ConstraintType() string {
	return "FOREIGN KEY"
}
