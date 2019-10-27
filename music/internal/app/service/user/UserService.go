package user

import (
	"container/list"
	"music/internal/app/dto"
	"music/internal/pkg/modules"
)

/**
用户 先关操作
*/
type Service interface {
	//用户的详情
	Detail(Id int64) (*modules.User, error)

	//删除指定的用户
	Delete(Id int64) (bool, error)

	//更新用户信息
	Update(user modules.User) (bool, error)

	//新增
	Insert(user modules.User) (bool, error)

	InsertAndGetId(user modules.User) (int64, error)

	//集合
	List(dto dto.UserInDTO) (list.Element, error)
}

//接口实现类
type ServiceImpl struct {
}
