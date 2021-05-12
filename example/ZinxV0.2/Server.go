package main

import "zinx/znet"

func main() {
	server := znet.NewServer("ServerV0.2")
	server.Serve()
}