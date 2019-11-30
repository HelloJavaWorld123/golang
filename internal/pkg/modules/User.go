package modules

import "time"

type User struct {
	Id         int64     `json:"id"`
	NickName   string    `json:"nick_name"`
	Telephone  string    `json:"telephone"`
	HeadUrl    string    `json:"head_url"`
	Province   string    `json:"province"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
