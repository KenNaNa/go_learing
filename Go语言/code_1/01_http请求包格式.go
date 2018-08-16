package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户的连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	//接收客户端的数据
	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err2)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))

}
