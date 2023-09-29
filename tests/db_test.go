package tests

import (
	"testing"

	"github.com/codr7/gstraps/db"
)

func TestDBRecord(t *testing.T) {
	tbl := db.NewTable("TestRecord")
	col := db.NewIntegerColumn(tbl, "TestRecordColumn")
	pk := db.NewKey(tbl, "TestRecordPrimary", col)
	tbl.SetPrimaryKey(pk)

	c, err := db.DefaultCxOptions().NewCx()

	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.StartTx()

	if err != nil {
		t.Fatal(err)
	}

	if err := tbl.Create(tx); err != nil {
		t.Fatal(err)
	}

	rec := db.NewRecord()

	if ok := rec.Null(col); !ok {
		t.Fatalf("Should be null")
	}

	if ok := rec.Modified(col, tx); ok {
		t.Fatal("Shouldn't be modified")
	}

	if ok := rec.Stored(col, tx); ok {
		t.Fatal("Shouldn't be stored")
	}

	col.Set(rec, 42)

	if v := col.Get(*rec); v != 42 {
		t.Fatalf("Wrong value: %v", v)
	}

	if ok := rec.Null(col); ok {
		t.Fatal("Shouldn't be null")
	}

	if ok := rec.Modified(col, tx); !ok {
		t.Fatal("Should be modified")
	}

	if ok := rec.Stored(col, tx); ok {
		t.Fatal("Shouldn't be stored")
	}

	if err := rec.Store(tbl, tx); err != nil {
		t.Fatal(err)
	}

	if ok := rec.Stored(col, tx); !ok {
		t.Fatal("Should be stored")
	}

	if err := rec.Store(tbl, tx); err != nil {
		t.Fatal(err)
	}

	if ok := rec.Stored(col, tx); !ok {
		t.Fatal("Should be stored")
	}

	if err := rec.Delete(tbl, tx); err != nil {
		t.Fatal(err)
	}

	if ok := rec.Stored(col, tx); ok {
		t.Fatal("Shouldn't not be stored")
	}

	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}

	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDBTable(t *testing.T) {
	tbl1 := db.NewTable("TestTable1")
	tbl1_integer := db.NewIntegerColumn(tbl1, "TestTable1Integer", db.NotNull)
	db.NewTextColumn(tbl1, "TestTable1Text", db.NotNull)
	db.NewIndex(tbl1, "TestTableIndex", true, tbl1_integer)
	tbl1_primary := db.NewKey(tbl1, "TestTable1Primary", tbl1_integer)
	tbl1.SetPrimaryKey(tbl1_primary)

	tbl2 := db.NewTable("TestTable2")
	tbl2_foreign := db.NewForeignKey(tbl2, "TestTable2Foregn", tbl1, tbl1_integer)
	tbl2_primary := db.NewKey(tbl2, "TestTable2Primary", tbl2_foreign.Columns()[0])
	tbl2.SetPrimaryKey(tbl2_primary)

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

func TestDBTx(t *testing.T) {
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
