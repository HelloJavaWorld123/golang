package main

import (
	"fmt"
	"golang/internal/app/service/user"
	"log"
)

func main() {
	vo := user.Vo{Id:1}
	detail, e := vo.Detail()
	if e!=nil{
		log.Printf("查询详情出现异常:%v",e)
	}
	fmt.Println(detail)
}

func init() {
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	//log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
