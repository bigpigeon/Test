/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package pd_test

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestPQInsert(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=toyorm sslmode=disable")
	assert.Nil(t, err)

	var id int64
	s := `INSERT INTO "test_foreign_key_table_belong_to"(data) VALUES(?) RETURNING id`
	s = strings.Replace(s, "?", "$1", -1)
	err = db.QueryRow(s, "query row insert").Scan(&id)
	t.Log(id)
	assert.Nil(t, err)
	time.Sleep(1000 * time.Second)
	//_, err = db.Exec(`INSERT INTO "test_foreign_key_table_belong_to"(data) VALUES($1) RETURNING id`, "exec insert")
	//assert.Nil(t, err)

}

func TestPQArray(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres@localhost:5432/data_filter?sslmode=disable")
	assert.NoError(t, err)
	s := `SELECT key FROM audio WHERE key = ANY($1)`
	args := pq.Array([]string{"m_0e313ecc-8bfb-455f-9024-6d6d358d976e", "m_921ec387-3ccd-432c-beb0-b15d68f116e5"})
	id := ""
	rows, err := db.Query(s, args)
	assert.NoError(t, err)
	for rows.Next() {
		err = rows.Scan(&id)
		assert.NoError(t, err)

		t.Log(id)
	}
	rows.Close()
}
