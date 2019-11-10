package main

import (
	"fmt"
	"log"
	"music/internal/dao/mysql"
)

func main() {
	var database mysql.DataBase
	source := database.GetDataSource()
	fmt.Println(source)
}

func init() {
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	//log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
