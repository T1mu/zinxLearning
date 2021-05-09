package main

import "code.xf.com/xf/2021-5-9/zinx/znet"

func main() {
	server := znet.NewServer("ServerV0.1")
	server.Serve()
}