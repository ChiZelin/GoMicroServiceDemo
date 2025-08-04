package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义一个结构体，实现 RPC 方法
type Arith struct{}

// 定义 RPC 方法
func (t *Arith) Add(args *[]int, reply *int) error {
	sum := 0
	for _, v := range *args {
		sum += v
	}
	*reply = sum
	return nil
}

func main() {
	// 注册 RPC 服务
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("RPC 服务正在监听 127.0.0.1:8080...")
	for {
		// 监听客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
