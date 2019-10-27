package main

import (
	"GoPipline/pipline"
	"bufio"
	"fmt"
	"os"
)

//文件读写的功能
func main() {
	FileOperator()
}

func FileOperator() {
	const fileName = "small.in"
	const seedNum = 100

	//创建文件
	file, e := os.Create(fileName)
	if e != nil {
		//延期处理该错误
		panic(e)
	}

	//当函数运行完毕后 将文件进行关闭 是个先进后出的栈数据结构
	defer file.Close()

	//创建数据源
	source := pipline.RandomSource(seedNum)
	//将数据写入数据源中
	pipline.WriteSink(bufio.NewWriter(file), source)

	//将上面的文件读出来
	file, e = os.Open(fileName)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	source = pipline.ReaderFile(bufio.NewReader(file))
	for value := range source {
		fmt.Println(value)
	}
}
