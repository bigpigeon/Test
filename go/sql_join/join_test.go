/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package sql_join

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TestDB     *sql.DB
	TestDriver string
)

func TestJoin(t *testing.T) {
	//query := "SELECT t1.id,t1.created_at,t1.updated_at,t1.deleted_at,t1.name,t1.belong_to_id,t2.id,t2.name,t2.test_preload_table_id  " +
	//	"FROM test_preload_table as t1 " +
	//	"JOIN (test_preload_table_one_to_many as t2,test_preload_table_one_to_one as t3)" +
	//	" ON (t1.id = t2.test_preload_table_id AND t1.id = t3.test_preload_table_id)"
	query := "SELECT t1.id,t1.created_at,t1.updated_at,t1.deleted_at,t1.name,t1.belong_to_id,t2.id,t2.name,t2.test_preload_table_id,t3.id  " +
		"FROM test_preload_table as `t1` " +
		"JOIN test_preload_table_one_to_many as t2 ON t1.id = t2.test_preload_table_id " +
		"JOIN test_preload_table_one_to_one as t3 ON t1.id = t3.test_preload_table_id " +
		"WHERE t3.id > 2"

		//" ON (t1.id = t3.test_preload_table_id)"
	t.Log("query", query)
	rows, err := TestDB.Query(query)
	if err != nil {
		panic(err)
	}
	//assert.Nil(t, err)
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			t.Error(err)
		}
		t.Log(columns)
		args := make([]interface{}, len(columns))
		for i := range columns {
			var data interface{}
			args[i] = &data
		}
		err = rows.Scan(args...)
		if err != nil {
			t.Error(err)
		}
		buff, err := json.Marshal(args)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(buff))
	}

	rows.Close()
}

func TestMain(m *testing.M) {
	var err error
	TestDriver = "mysql"
	TestDB, err = sql.Open(TestDriver, "root:@tcp(localhost:3306)/toyorm?charset=utf8&parseTime=True")
	//TestDriver = "postgres"
	//TestDB, err = sql.Open(TestDriver, "user=postgres dbname=toyorm sslmode=disable")
	if err != nil {
		panic(err)
	}
	m.Run()
}
