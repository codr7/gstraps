package db

import (
	"context"
	"fmt"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	cx  context.Context
	imp *pgx.Conn
}

func Connect(cx context.Context, url string) (*Connection, error) {
	return new(Connection).Init(cx, url)
}

func (self *Connection) Init(cx context.Context, url string) (*Connection, error) {
	var err error

	if self.imp, err = pgx.Connect(cx, url); err != nil {
		return nil, err
	}

	pgxdecimal.Register(self.imp.TypeMap())
	self.cx = cx
	return self, nil
}

func (self *Connection) Close() error {
	return self.imp.Close(self.cx)
}

type ConnectOptions struct {
	Context  context.Context
	Host     string
	Port     int
	Database string
	User     string
	Password string
	SSLMode  bool
}

func (self ConnectOptions) NewConnection() (*Connection, error) {
	return Connect(self.Context,
		fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
			self.Host, self.Port, self.Database, self.User, self.Password))
}

func DefaultConnectOptions() ConnectOptions {
	var o ConnectOptions
	o.Context = context.Background()
	o.Host = "localhost"
	o.Port = 5432
	o.Database = "gstraps"
	o.User = "gstraps"
	o.Password = "gstraps"
	o.SSLMode = false
	return o
}
