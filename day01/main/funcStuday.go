package main

import (
	"fmt"
)

func main() {
	// 声明函数
	//result := add(1, 10)
	//fmt.Printf(string(result))

	// 包含多个返回值得函数
	//total, param1, param2 := addReturnMoreResults(10, 10)

	//函数包含指定返回值名称的
	//resultOne, resultTwo := resultParamName(11, 11)

	resultOne, resultTwo := singalResultReturn(10, 10)
	fmt.Print(resultOne, resultTwo)
}

/*
函数：
名称
参数（类型相同的参数可以用户逗号隔开声明）
返回值类型（一个函数可以同时返回多个返回值）
*/
func add(x, y int) int {
	return x + y
}

/*
一个函数同时返回多个结果
在使用该函数时 也必须声明相同的内容
声明的返回值 只声明了类型  没有声明返回值名称
*/
func addReturnMoreResults(x, y int) (int, int, int) {
	return x + y, x, y
}

/*
创建函数时 增加返回值得名称
返回值的名称 简单 明了 易懂
直接返回语句 将当前函数所声明的返回值默认全部返回
*/
func resultParamName(x, y int) (resultOne, resultTwo int) {
	resultOne = x + y
	resultTwo = x - y
	return
}

/*函数声明几个返回值 就要返回几个返回值*/
func singalResultReturn(x, y int) (resultOne, resultTwo int) {
	resultTwo = x * y
	return resultTwo, resultOne
}
