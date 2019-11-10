package user

import (
	"container/list"
	"music/internal/dao/mysql/userdao"
	"music/internal/pkg/modules"
)

/**
用户 先关操作
*/
type Service interface {
	//用户的详情
	Detail() (*modules.User, error)

	//删除指定的用户
	Delete() (bool, error)

	//更新用户信息
	Update() (bool, error)

	//新增
	Insert() (bool, error)

	InsertAndGetId() (int64, error)

	//集合
	List() (list.Element, error)
}

//入参
type UserVo struct {
	Id        int64
	NickName  string
	HeadUrl   string
	Telephone string
	Province  string
	Country   string
	City      string
}

//方法
//接收者 包括值接收者 和 指针接收者
func (userVo UserVo) Delete() (bool, error) {
	panic("implement me")
}

func (userVo UserVo) Update() (bool, error) {
	panic("implement me")
}

func (userVo UserVo) Insert() (bool, error) {
	panic("implement me")
}

func (userVo UserVo) InsertAndGetId() (int64, error) {
	panic("implement me")
}

func (userVo UserVo) List() (list.Element, error) {
	panic("implement me")
}

func (userVo UserVo) Detail() (*modules.User, error) {
	user := userdao.Detail(userVo.Id)
	return user, nil
}
