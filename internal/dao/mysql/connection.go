package mysql

import (
	"database/sql"
	"fmt"
	"time"

	//数据库驱动 导入
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//全局函数
var Template *sql.DB

type DataBase struct {
	DataSource DataSource `yaml:"database"`
}

type DataSource struct {
	UserName          string `yaml:"username"`
	Password          string `yaml:"password"`
	HostAndIp         string `yaml:"hostAndIp"`
	DataBaseName      string `yaml:"dataBaseName"`
	MaxIdle           int    `yaml:"maxIdle"`
	MaxOpenConnection int    `yaml:"maxOpenConnection"`
	MaxLifeTime       int64  `yaml:"maxLifeTime"`
}

func init() {
	var danseuse DataSource
	yamlFile, e := ioutil.ReadFile("configs/app-test.yml")
	if e != nil {
		log.Println(e)
		log.Printf("读取配置出现异常:{%v}", e)
		log.Fatalf("读取配置出现异常:{%v}", e)
	}
	e = yaml.Unmarshal(yamlFile, danseuse)
	if e != nil {
		fmt.Printf(e.Error())
		log.Fatalf("通过配置文件初始化出现异常:{%v}", e)
	}

	url := danseuse.UserName + ":" + danseuse.Password + "@tcp(" + danseuse.HostAndIp + ")" + "/" + danseuse.DataBaseName
	Template, e := sql.Open("mysql", url)

	if e != nil {
		log.Fatalf("初始化sql链接出现异常:{%+v}", e)
	}
	Template.SetMaxIdleConns(danseuse.MaxIdle)
	Template.SetMaxOpenConns(danseuse.MaxOpenConnection)
	Template.SetConnMaxLifetime(time.Duration(danseuse.MaxLifeTime))
}
