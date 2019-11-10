package mysql

//mysql 数据库连接的一些配置
import (
	"fmt"
	//数据库驱动 导入
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DataBase struct {
	DataSource DataSource `yaml:"database"`
}

type DataSource struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

func (base *DataBase) GetDataSource() DataSource {
	yamlFile, e := ioutil.ReadFile("configs/app-test.yml")
	if e != nil {
		log.Println(e)
		log.Printf("读取配置出现异常:{%v}", e)
		log.Fatalf("读取配置出现异常:{%v}", e)
	}
	e = yaml.Unmarshal(yamlFile, base)
	if e != nil {
		fmt.Printf(e.Error())
		log.Fatalf("通过配置文件初始化出现异常:{%v}", e)
	}
	return base.DataSource
}
