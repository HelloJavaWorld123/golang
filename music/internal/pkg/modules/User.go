package modules

import "time"

/**
user 实体
*/
type User struct {
	Id         int64
	NickName   string
	Telephone  string
	HeadUrl    string
	Province   string
	City       string
	Country    string
	CreateDate time.Time
	UpdateDate time.Time
}
