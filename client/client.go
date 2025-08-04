package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 连接到 RPC 服务端
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer client.Close()

	// 准备调用 RPC 方法的参数
	args := []int{1, 2, 3}
	var reply int

	// 调用 RPC 方法
	err = client.Call("Arith.Add", &args, &reply)
	if err != nil {
		fmt.Println("调用 RPC 方法失败:", err)
		return
	}

	fmt.Printf("结果: %d\n", reply)
}
