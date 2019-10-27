package main

import (
	"GoPipline/pipline"
	"fmt"
)

//cmd目录下按照目录分别创建主函数
//主函数
func main() {
	//array()
	//sortInMemory()
	sortAndMerge()
}

//对多个数据源进行先排序再进行排序
func sortAndMerge() {
	result := pipline.Merge(pipline.InMemorySort(pipline.ArraySource(3, 1, 4, 70, 5, 60, 30, 7, 8)),
		pipline.InMemorySort(pipline.ArraySource(4, 6, 10, 7, 6, 80, 40)))

	for value := range result {
		fmt.Println(value)
	}

}

//对数据源进行内存排序
func sortInMemory() {
	value := pipline.InMemorySort(pipline.ArraySource(1, 30, 100, 39, 40, 2, 5, 4, 9, 8))
	//range 会等待 数据源value的完成 直到有数据可以获取
	for v := range value {
		fmt.Println(v)
	}
}

//简单的以arraysource为数据源 进行读取数据
func array() {
	//使用 pipleline 中的函数 获取打印指定的数据
	//可变长参数 返回的是一个channel
	pip := pipline.ArraySource(2, 1, 4, 5, 6, 7, 9, 10, 100, 80)
	//从 channel 中 源源不断的获取数据 有以下两种方式
	//for{
	//	//第一种方式 如果channel中有数据 则输出num并且ok是true,否则ok为false
	//	if num,ok := <- pip;ok{
	//		fmt.Println(num)
	//	}else{
	//		break
	//	}
	//}
	//第二种方式 使用range 的方式，但是这种方式 被调用放一定要有关闭的动则 否则会一直获取
	for num := range pip {
		fmt.Println(num)
	}
}
