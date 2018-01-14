package sqldemo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var (
	DB *sql.DB
)

type User struct {
	Name string
	Age  int
}

func TestMain(m *testing.M) {
	for _, d := range []struct {
		Driver      string
		DataSource  string
		TableCreate string
	}{
		{"sqlite3", "test.db", `CREATE TABLE user (id integer PRIMARY KEY AUTOINCREMENT, name varchar(255), age integer, stuff integer)`},
		{"mysql", "root:@tcp(localhost:3306)/toyorm?charset=utf8&parseTime=True&loc=Local", `CREATE TABLE user (id integer AUTO_INCREMENT, name varchar(255), age integer, stuff integer, PRIMARY KEY(id))`},
	} {
		var err error
		DB, err = sql.Open(d.Driver, d.DataSource)
		if err != nil {
			panic(err)
		}
		fmt.Printf("=============== %s ===============\n", d.Driver)
		DB.Exec(`DROP TABLE user`)
		result, err := DB.Exec(d.TableCreate)
		fmt.Printf("%#v", result)
		if err != nil {
			panic(err)
		}
		m.Run()
		if DB != nil {
			DB.Close()
		}
	}

}
