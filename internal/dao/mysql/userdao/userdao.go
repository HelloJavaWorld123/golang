package userdao

import (
	"golang/internal/dao/mysql"
	"golang/internal/pkg/modules"
	"log"
)

func Detail(id int64) *modules.User {
	sql := "SELECT * FROM person"
	stmt, e := mysql.Template.Prepare(sql)
	checkErr(e)
	rows, e := stmt.Query()
	checkErr(e)
	defer rows.Close()

	if rows.Next() {
	}
	return &modules.User{Id: 1, NickName: "测试姓名"}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
