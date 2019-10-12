package main

import "fmt"

/*
* var 声明变量
* 变量的位置 包级别 或者 函数级别
* 变量的类型 同函数的参数或者返回值一样 变量的类型 放在后面
 */

/*
*：= 只能使用在函数内
* 在函数内使用时 可用于
 */

var b, c int
var a = 100

func main() {
	//fmt.Print(addVarParam())
	fmt.Print(paramMore())
}

func addVarParam() (paramA, paramB, paramC int) {
	paramA = a
	paramB = b
	paramC = c
	return
}

/*
变量初始化
初始化时 可同时初始化多个变量，而且变量的类型可以由初始化的值推断出
*/
var paramString, paramInt, paramBool = "one", 2, true

func paramMore() (string, int, bool) {
	return paramString, paramInt, paramBool
}
