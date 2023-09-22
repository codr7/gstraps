package tests

import (
	"testing"

	"github.com/codr7/gstraps/db"
)

func TestDbRecord(t *testing.T) {
	var tbl db.BasicTable
	tbl.Init("TestDbRecord")

	var col db.IntegerColumn
	col.Init(&tbl, "TestDbRecordColumn")

	var key db.Key
	key.Init(&tbl, "TestDbRecordPrimaryKey", &col)
	tbl.SetPrimaryKey(&key)

	var rec db.Record
	rec.Init()
	col.Set(&rec, 42)

	if v := col.Get(rec); v != 42 {
		t.Fatalf("Wrong value: %v", v)
	}
}
