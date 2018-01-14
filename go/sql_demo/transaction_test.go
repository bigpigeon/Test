package sqldemo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction(t *testing.T) {
	tx, err := DB.Begin()
	if err != nil {
		t.Error(tx)
	}
	defer tx.Rollback()
	//defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO user(name) VALUES(?)")
	if err != nil {
		t.Error(err)
	}
	defer stmt.Close()
	stmt.Exec("pigeon")
	stmt.Exec("bigpigeon")
	var tx_count, db_count int
	err = tx.QueryRow("SELECT COUNT(*) FROM user").Scan(&tx_count)
	assert.Nil(t, err)
	err = DB.QueryRow("SELECT COUNT(*) FROM user").Scan(&db_count)
	t.Logf("transaction count: %d,db count: %d", tx_count, db_count)
	err = tx.Commit()
	assert.Nil(t, err)
	t.Log("transaction commit")
	err = DB.QueryRow("SELECT COUNT(*) FROM user").Scan(&db_count)
	t.Logf("transaction count: %d,db count: %d", tx_count, db_count)
}
