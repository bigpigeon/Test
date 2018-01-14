package sqldemo

import (
	"testing"
)

func TestInsert(t *testing.T) {
	stmt, err := DB.Prepare("INSERT INTO user(name, age) VALUES(?, ?)")
	if err != nil {
		t.Error(err)
	}
	for _, d := range []User{{"jia", 20}, {"jiathenine", 21}, {"pigeon", 22}} {
		res, err := stmt.Exec(d.Name, d.Age)
		if err != nil {
			t.Error(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			t.Error(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			t.Error(err)
		}
		t.Logf("ID = %d, affected = %d\n", lastId, rowCnt)
	}

}

//func TestMultipleInsert(t *testing.T) {
//	stmt, err := DB.Prepare("INSERT INTO user(name, age) VALUES((?, ?),(?, ?))")
//	if err != nil {
//		t.Error(err)
//	}
//	stmt.Exec("tom", 30, "jobs", 56)
//}
