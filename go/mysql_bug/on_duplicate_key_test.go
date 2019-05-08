package mysql_bug

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
}

func TestOnDuplicateKeyBug(t *testing.T) {
	{
		_, err := DB.Exec("DROP TABLE `test_duplicate_key_bug_table`")
		require.NoError(t, err)
	}
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS `test_duplicate_key_bug_table` (id INTEGER AUTO_INCREMENT,some_data VARCHAR(255) ,created_at TIMESTAMP,PRIMARY KEY(id))")
	require.NoError(t, err)
	createdAt := time.Now().Add(-10 * time.Second)
	// use 'insert ... ON DUPLICATE KEY UPDATE ...' to update data, it also update created_at column
	{
		result, err := DB.Exec("INSERT INTO `test_duplicate_key_bug_table`(some_data,created_at) VALUES(?,?)", "test data", createdAt)
		require.NoError(t, err)
		id, err := result.LastInsertId()
		require.NoError(t, err)
		// mysql's timestamp only storage second's level date, re-find it's real datetime
		row := DB.QueryRow("SELECT created_at FROM `test_duplicate_key_bug_table` WHERE id=? LIMIT 1", id)
		var oldCreatedAt time.Time
		err = row.Scan(&oldCreatedAt)
		require.NoError(t, err)

		result, err = DB.Exec("INSERT INTO `test_duplicate_key_bug_table`(id,some_data,created_at) VALUES(?,?,?) ON DUPLICATE KEY UPDATE some_data = VALUES(some_data)", id, "test data again", time.Now())
		require.NoError(t, err)

		row = DB.QueryRow("SELECT some_data,created_at FROM `test_duplicate_key_bug_table` WHERE id=? LIMIT 1", id)
		var newCreatedAt time.Time
		var newData string
		err = row.Scan(&newData, &newCreatedAt)
		require.NoError(t, err)

		fmt.Printf("old time %s \n", oldCreatedAt)
		fmt.Printf("new time %s \n", newCreatedAt)
		fmt.Printf("new data %s \n", newData)
	}
	// but only update id = VALUES(id) , it work well
	{
		result, err := DB.Exec("INSERT INTO `test_duplicate_key_bug_table`(some_data,created_at) VALUES(?,?)", "test data", createdAt)
		require.NoError(t, err)
		id, err := result.LastInsertId()
		require.NoError(t, err)
		// mysql's timestamp only storage second's level date, re-find it's real datetime
		row := DB.QueryRow("SELECT created_at FROM `test_duplicate_key_bug_table` WHERE id=? LIMIT 1", id)
		var oldCreatedAt time.Time
		err = row.Scan(&oldCreatedAt)
		require.NoError(t, err)

		result, err = DB.Exec("INSERT INTO `test_duplicate_key_bug_table`(id,some_data,created_at) VALUES(?,?,?) ON DUPLICATE KEY UPDATE id = VALUES(id) ", id, "test data again", time.Now())
		require.NoError(t, err)

		row = DB.QueryRow("SELECT some_data,created_at FROM `test_duplicate_key_bug_table` WHERE id=? LIMIT 1", id)
		var newCreatedAt time.Time
		var newData string
		err = row.Scan(&newData, &newCreatedAt)
		require.NoError(t, err)

		fmt.Printf("old time %s \n", oldCreatedAt)
		fmt.Printf("new time %s \n", newCreatedAt)
		fmt.Printf("new data %s \n", newData)
	}
}
