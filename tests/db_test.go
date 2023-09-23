package tests

import (
	"testing"

	"github.com/codr7/gstraps/db"
)

func TestDBRecord(t *testing.T) {
	tbl := db.NewTable("TestRecord")
	col := db.NewIntegerColumn(tbl, "TestRecordColumn")
	key := db.NewKey(tbl, "TestRecordPrimaryKey", col)
	tbl.SetPrimaryKey(key)
	rec := db.NewRecord()

	if ok := rec.Null(col); !ok {
		t.Fatalf("Should be null")
	}

	col.Set(rec, 42)

	if ok := rec.Null(col); ok {
		t.Fatalf("Shouldn't be null")
	}

	if v := col.Get(*rec); v != 42 {
		t.Fatalf("Wrong value: %v", v)
	}

	c, err := db.DefaultConnectOptions().NewConnection()

	if err != nil {
		t.Fatal(err)
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

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}
