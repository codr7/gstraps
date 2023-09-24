package tests

import (
	"testing"

	"github.com/codr7/gstraps/db"
)

func TestDBRecord(t *testing.T) {
	tbl := db.NewTable("TestRecord")
	col := db.NewIntegerColumn(tbl, "TestRecordColumn")

	c, err := db.DefaultCxOptions().NewCx()

	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.StartTx()

	if err != nil {
		t.Fatal(err)
	}

	rec := db.NewRecord()

	if ok := rec.Null(col); !ok {
		t.Fatalf("Field should be null")
	}

	if ok := rec.Modified(col, tx); ok {
		t.Fatal("Field shouldn't be modified")
	}

	if ok := rec.Stored(col, tx); ok {
		t.Fatal("Field shouldn't be stored")
	}

	col.Set(rec, 42)

	if v := col.Get(*rec); v != 42 {
		t.Fatalf("Wrong value: %v", v)
	}

	if ok := rec.Null(col); ok {
		t.Fatal("Field shouldn't be null")
	}

	if ok := rec.Modified(col, tx); !ok {
		t.Fatal("Field should be modified")
	}

	if ok := rec.Stored(col, tx); ok {
		t.Fatal("Field shouldn't be stored")
	}

	if err = tx.Rollback(); err != nil {
		t.Fatal(err)
	}

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDBTable(t *testing.T) {
	tbl1 := db.NewTable("TestTable1")
	col1 := db.NewIntegerColumn(tbl1, "TestTableColumn1")
	db.NewIndex(tbl1, "TestTableIndex", col1).Unique = true
	key1 := db.NewKey(tbl1, "TestTablePrimaryKey1", col1)
	tbl1.SetPrimaryKey(key1)

	tbl2 := db.NewTable("TestTable2")
	col2 := db.NewIntegerColumn(tbl2, "TestTableColumn2")
	key2 := db.NewKey(tbl2, "TestTablePrimaryKey2", col2)
	tbl2.SetPrimaryKey(key2)
	db.NewForeignKey(tbl2, "TestTableForegnKey", tbl1, col1)

	c, err := db.DefaultCxOptions().NewCx()

	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.StartTx()

	if err != nil {
		t.Fatal(err)
	}

	if err := tbl1.Create(tx); err != nil {
		t.Fatal(err)
	}

	if err := tbl2.Create(tx); err != nil {
		t.Fatal(err)
	}

	if err := tbl2.Drop(tx); err != nil {
		t.Fatal(err)
	}

	if err := tbl1.Drop(tx); err != nil {
		t.Fatal(err)
	}

	if err = tx.Rollback(); err != nil {
		t.Fatal(err)
	}

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDBTxs(t *testing.T) {
	c, err := db.DefaultCxOptions().NewCx()

	if err != nil {
		t.Fatal(err)
	}

	tx1, err := c.StartTx()

	if err != nil {
		t.Fatal(err)
	}

	tx2, err := c.StartTx()

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
		t.Fatal("Commit of rolled back tx")
	}

	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}
