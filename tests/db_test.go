package tests

import (
	"testing"

	"github.com/codr7/gstraps/db"
)

func TestDBRecord(t *testing.T) {
	tbl := db.NewTable("TestRecord")
	col := db.NewIntegerColumn(tbl, "TestRecordColumn")

	c, err := db.DefaultConnectOptions().NewConnection()

	if err != nil {
		t.Fatal(err)
	}

	rec := db.NewRecord(c)

	if ok := rec.Null(col); !ok {
		t.Fatalf("Field should be null")
	}

	if ok := rec.Modified(col); ok {
		t.Fatal("Field shouldn't be modified")
	}

	if ok := rec.Stored(col); ok {
		t.Fatal("Field shouldn't be stored")
	}

	col.Set(rec, 42)

	if v := col.Get(*rec); v != 42 {
		t.Fatalf("Wrong value: %v", v)
	}

	if ok := rec.Null(col); ok {
		t.Fatal("Field shouldn't be null")
	}

	if ok := rec.Modified(col); !ok {
		t.Fatal("Field should be modified")
	}

	if ok := rec.Stored(col); ok {
		t.Fatal("Field shouldn't be stored")
	}

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDBTable(t *testing.T) {
	tbl := db.NewTable("TestTable")
	col := db.NewIntegerColumn(tbl, "TestTableColumn")
	key := db.NewKey(tbl, "TestTablePrimaryKey", col)
	tbl.SetPrimaryKey(key)

	c, err := db.DefaultConnectOptions().NewConnection()

	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.StartTransaction()

	if err != nil {
		t.Fatal(err)
	}

	if err := tbl.Create(tx); err != nil {
		t.Fatal(err)
	}

	if err := tbl.Drop(tx); err != nil {
		t.Fatal(err)
	}

	if err = tx.Rollback(); err != nil {
		t.Fatal(err)
	}

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDBTransactions(t *testing.T) {
	c, err := db.DefaultConnectOptions().NewConnection()

	if err != nil {
		t.Fatal(err)
	}

	tx1, err := c.StartTransaction()

	if err != nil {
		t.Fatal(err)
	}

	tx2, err := c.StartTransaction()

	if err != nil {
		t.Fatal(err)
	}

	if err = tx2.Commit(); err != nil {
		t.Fatal(err)
	}

	if err = tx1.Rollback(); err != nil {
		t.Fatal(err)
	}

	if err = tx1.Commit(); err == nil {
		t.Fatal("Commit of rolled back transaction")
	}

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}
