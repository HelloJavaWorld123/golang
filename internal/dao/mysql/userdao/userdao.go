package userdao

import (
	"golang/internal/dao/mysql"
	"golang/internal/pkg/modules"
	"log"
)

func Detail(id int64) *modules.User {
	sql := "SELECT * FROM test_user"
	stmt, e := mysql.Template.Prepare(sql)
	checkErr(e)
	rows, e := stmt.Query()
	checkErr(e)
	defer rows.Close()
	user := new(modules.User)
	if err := rows.Scan(&user.Id, &user.NickName, &user.Telephone, &user.HeadUrl, &user.Province, &user.Country, &user.City); err != nill {
		checkErr(e)
	}
	return user
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
