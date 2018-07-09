/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package sql_join

import (
	"fmt"
	"testing"
)

func TestShowTables(t *testing.T) {
	rows, err := TestDB.Query("SHOW TABLES")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			t.Error(err)
		}
		t.Log(name)
		var cname, sqlStr string
		err = TestDB.QueryRow(fmt.Sprintf("SELECT INDEX_SCHEMA,COLUMN_NAME FROM information_schema.statistics ")).Scan(&cname, &sqlStr)
		if err != nil {
			t.Error(err)
		}
		t.Log(cname, sqlStr)
	}
}
