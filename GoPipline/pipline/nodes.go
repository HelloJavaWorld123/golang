package pipline

import "sort"

//以 数组为数据源,使用channel读取数组中的数据
//参数的形式是边长参数的形式
//func ArraySource(param ...int) chan int{
//简短声明的方式 初始化一个channel
//out := make(chan int)
//使用 goroutine 并发的读取数组中的数据
//go func() {
//使用 range的方式 循环的从 数据源中读取数据
// range循环读取数据时,第一个参数时下标 第二个参数是下标对应的value
// 如果不需要下标时 使用下划线代替
//for _,value := range param{
//	out <- value
//}
//当 读取数据完毕时 将输出流 关闭
//close(out)
//}()
//输出channel中的数据
//return out
//}

//对上面的函数进行优化 通过函数表示出 数据的流向
//加上箭头 表示使用者只能从里面获取数据，是只出不进的函数
func ArraySource(in ...int) <-chan int {
	//通过make初始化 一个 channel
	out := make(chan int)
	go func() {
		//如果只是使用 value一个变量进行接收 则默认是接收的下标
		for _, value := range in {
			out <- value
		}
		//关闭输出channel
		close(out)
	}()
	return out
}

//在增加一个函数
//将channel作为数据源 不断的读取，并以channel的方式输出
func InMemorySort(source <-chan int) <-chan int {
	//初始化一个channel
	out := make(chan int)
	//使用 goroutine 的方式将读取出来的数据 保存在内存中
	go func() {
		//初始化一个分片
		var memory []int
		//简洁的方式声明变量
		for value := range source {
			//对变量进行赋值更新 不能再使用简洁方式接收 否则则是声明新的变量
			memory = append(memory, value)
		}

		//将获取到的数据 进行排序
		sort.Ints(memory)

		//将排序号的内容 输出
		for _, value := range memory {
			out <- value
		}
		//输出完毕后将输出channel关闭
		close(out)
	}()
	return out
}

//对 两个数据输入源 进行合并的函数
//入参是两个channel数据源  输出是一个channel数据源
func Merge(paramSourceOne, paramSourceTwo <-chan int) <-chan int {
	//初始化一个channel管道
	out := make(chan int)
	//针对两个输入源，那个有数据就输出那个 并且有数据的要小于另外一个
	go func() {
		//会等待 两个数据源分别排序好后 再输出
		valueOne, ok1 := <-paramSourceOne
		valueTwo, ok2 := <-paramSourceTwo

		for ok1 || ok2 {
			if !ok2 || (ok1 && valueOne <= valueTwo) {
				out <- valueOne
				//对valueOne 进行重新的赋值
				valueOne, ok1 = <-paramSourceOne
			} else {
				out <- valueTwo
				valueTwo, ok2 = <-paramSourceTwo
			}
			//将 out进行关闭
		}
		close(out)
	}()
	return out
}
