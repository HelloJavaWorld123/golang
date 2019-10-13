package stringUtils

/*
*创建 库文件
* 在可执行的GO 文件的同层目录下创建 包，包名同 Java 包名 小写
* 包里面只是go 文件，文件里面的函数首字母大写 否则不能调用
* 函数库 创建成功后 需要运行 go build 命令，但无任何的输出，除非 运行 go install 命令
* go Install 命令会在GOPath 中 创建pkg,该目录存放编译以后的文件
 */

/*将给定的字符串进行反转*/
func Revers(param string) string {
	stringParamArray := []rune(param)
	for start, stop := 0, len(stringParamArray)-1; start < len(stringParamArray); start, stop = start+1, stop-1 {
		stringParamArray[start], stringParamArray[stop] = stringParamArray[stop], stringParamArray[start]
	}
	return string(stringParamArray)
}
