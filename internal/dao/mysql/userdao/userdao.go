package userdao

import (
	"database/sql"
	"music/internal/pkg/modules"
)

func Detail(id int64) *modules.User {
	sql.Open("mysql", "test")
	return &modules.User{Id: 1, NickName: "测试姓名"}
}
