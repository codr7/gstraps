package tests

import (
	"testing"

	"github.com/codr7/gstraps/db"
)

func TestDBRecord(t *testing.T) {
	tbl := db.NewTable("TestRecordTable")
	col := db.NewIntegerColumn(tbl, "TestRecordColumn")
	key := db.NewKey(tbl, "TestDbRecordPrimaryKey", col)
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
