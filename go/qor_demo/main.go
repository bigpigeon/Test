package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"net/http"
)

// Create a GORM-backend model
type Role int

const (
	RoleAdmin = Role(iota)
	//社区权限
	RoleCommunity
	//小区权限
	RoleNeighbourhoods
	//楼长权限
	RoleHousing
)

type User struct {
	gorm.Model
	Name string
	Role Role
}

type House struct {
	gorm.Model
	State         string
	City          string
	Community     string
	Neighbourhood string
	Housing       string
	Cell          string
}

type People struct {
	Name    string
	Age     int
	Sex     bool
	Contact string
	Job     string
}

// Create another GORM-backend model
type Resident struct {
	gorm.Model
	Name   string
	Sex    bool
	Age    int
	IdCard string
	//房屋信息
	HouseId int
	House   House
	//民族
	Race string
	//工作单位
	Job string
	//配偶
	Spouse string

	Address     string
	Contact     string
	Description string
}

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.DropTableIfExists(&User{}, &Resident{}, &House{})
	DB.AutoMigrate(&User{}, &Resident{}, &House{})

	// Initalize
	Admin := admin.New(&qor.Config{DB: DB})

	// Allow to use Admin to manage data
	Admin.AddResource(&User{})
	{
		res := Admin.AddResource(&Resident{})
		res.EditAttrs(
			&admin.Section{
				Title: "基本信息",
				Rows: [][]string{
					{"name", "sex", "age"},
					{"id_card", "house", "contact"},
					{"address"},
					{"race", "job"},
					{"spouse"},
				},
			},
			&admin.Section{
				Title: "其他信息",
			},
		)
		res.NewAttrs(res.EditAttrs())
		res.ShowAttrs(res.EditAttrs())
	}
	Admin.AddResource(&House{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
