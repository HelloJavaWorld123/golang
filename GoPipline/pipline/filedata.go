package pipline

import (
	"encoding/binary"
	"io"
	"math/rand"
)

//从文件中读取数据并输出

func ReaderFile(reader io.Reader) <-chan int {
	//用户接收 需要输出的数据
	out := make(chan int)

	go func() {
		//初始化一个 buffer用户保存 从文件中读取的数据
		buffer := make([]byte, 8)
		for {
			//第一个返回参数代表 返回的字节数 第二个参数是 返回的错误
			data, err := reader.Read(buffer)
			if data > 0 {
				//将二进制内容 转换为 int类型
				value := binary.BigEndian.Uint64(buffer)
				out <- int(value)
			}

			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out
}

//只进不出的 写文件函数
func WriteSink(writer io.Writer, insource <-chan int) {
	//直接使用for循环读取数据 并写入文件
	for value := range insource {
		//声明缓冲数据集合
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(value))
		_, err := writer.Write(buffer)
		if err != nil {
			break
		}
	}
}

//随机数作为数据源
func RandomSource(seed int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < seed; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
