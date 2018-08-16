package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 4*1024) //4k大小临时缓冲区
	for {
		n, err := sF.Read(buf) //从源文件读取内容
		if err != nil {
			fmt.Println("err = ", err)
			if err == io.EOF { //文件读取完毕
				break
			}
		}
		//往目的文件写，读多少写多少
		dF.Write(buf[:n])
	}

}
