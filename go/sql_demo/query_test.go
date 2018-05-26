package sqldemo

import (
	"testing"
)

func TestQuery(t *testing.T) {
	rows, err := DB.Query("select id,name,stuff from user where id in (?,?,?) AND age > ?", 1, 2, 3, 20)
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		{
			var (
				id    int
				name  string
				stuff *int
			)
			err := rows.Scan(&id, &name, &stuff)
			if err != nil {
				t.Error(err)
			}
			t.Log(id, name, stuff)
		}
		{
			var (
				id    int
				name  *string
				stuff *int
			)
			err := rows.Scan(&id, &name, &stuff)
			if err != nil {
				t.Error(err)
			}
			t.Log(id, *name, stuff)
		}
	}
	err = rows.Err()
	if err != nil {
		t.Error(err)
	}
}

func TestQueryRow(t *testing.T) {
	var (
		id    int
		name  string
		stuff *int
	)
	row := DB.QueryRow("select id,name,stuff from user where age > ?", 100000)
	err := row.Scan(&id, &name, &stuff)
	if err != nil {
		t.Error(err)
	}
}
