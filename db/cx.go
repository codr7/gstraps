package db

import (
	"context"
	"fmt"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
)

type Cx struct {
	StoredValues
	cx  context.Context
	imp *pgx.Conn
}

func NewCx(cx context.Context, url string) (*Cx, error) {
	return new(Cx).Init(cx, url)
}

func (self *Cx) Init(cx context.Context, url string) (*Cx, error) {
	self.StoredValues.Init()
	var err error

	if self.imp, err = pgx.Connect(cx, url); err != nil {
		return nil, err
	}

	pgxdecimal.Register(self.imp.TypeMap())
	self.cx = cx
	return self, nil
}

func (self *Cx) StartTx() (*Tx, error) {
	imp, err := self.imp.Begin(self.cx)

	if err != nil {
		return nil, err
	}

	return new(Tx).Init(self, imp), nil
}

func (self *Cx) Close() error {
	return self.imp.Close(self.cx)
}

type StoredValues struct {
	storedValues map[*Field]any
}

func (self *StoredValues) Init() *StoredValues {
	self.storedValues = make(map[*Field]any)
	return self
}

func (self StoredValues) StoredValue(field *Field) (any, bool) {
	v, ok := self.storedValues[field]
	return v, ok
}

type CxOptions struct {
	Context  context.Context
	Host     string
	Port     int
	Database string
	User     string
	Password string
	SSLMode  bool
}

func (self CxOptions) NewCx() (*Cx, error) {
	return NewCx(self.Context,
		fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
			self.Host, self.Port, self.Database, self.User, self.Password))
}

func DefaultCxOptions() CxOptions {
	var o CxOptions
	o.Context = context.Background()
	o.Host = "localhost"
	o.Port = 5432
	o.Database = "gstraps"
	o.User = "gstraps"
	o.Password = "gstraps"
	o.SSLMode = false
	return o
}
