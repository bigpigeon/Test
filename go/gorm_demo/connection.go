package gormdemo

import (
	"errors"
	"fmt"
	"io/ioutil"

	"testing"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Config map[string]interface{}
var AvalibeDBs []*gorm.DB

func Connection(driver string) (db *gorm.DB, err error) {
	switch driver {
	case "postgres":
		pconfig := Config["postgres"].(map[string]interface{})
		db, err = gorm.Open("postgres",
			fmt.Sprintf(
				"host=%s user=%s dbname=%s password=%s sslmode=%s",
				pconfig["host"], pconfig["user"], pconfig["dbname"], pconfig["password"], pconfig["sslmode"],
			),
		)
	case "sqlite", "sqlite3":
		sconfig := Config["sqlite"].(map[string]interface{})
		db, err = gorm.Open("sqlite3", sconfig["sqlurl"])
	case "mysql":
		mconfig := Config["mysql"].(map[string]interface{})
		db, err = gorm.Open("mysql",
			fmt.Sprintf(
				"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
				mconfig["user"], mconfig["password"], mconfig["dbname"],
			))
	default:
		err = errors.New("invalid database name")
	}
	return
}

func TraversalTest(t *testing.T, testfunc func(*gorm.DB)) {
	for _, db := range AvalibeDBs {
		t.Logf("start %s test....", db.Dialect().GetName())
		testfunc(db)
	}
}

func init() {
	f, err := ioutil.ReadFile("config.toml")
	IfErrPanic(err)
	err = toml.Unmarshal(f, &Config)
	IfErrPanic(err)
	for _, s := range []string{"sqlite", "mysql", "postgres"} {
		db, err := Connection(s)
		if err == nil {
			AvalibeDBs = append(AvalibeDBs, db)
		} else {
			fmt.Printf("%s db cannot connection: %s", s, err)
		}
	}
}

func IfErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
