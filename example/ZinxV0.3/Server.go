package main

import (
	"fmt"
	"zinx/zface"
	"zinx/znet"
)

type pingRouter struct {
	znet.BaseRouter
}

func (pr *pingRouter) PreHandle(request zface.IRequest)  {
	conn := request.GetConnection().GetTcpConnection()
	_, err := conn.Write([]byte("[pingRouter] Pre\n"))
	if err != nil {
		fmt.Println("[pingRouter]Error", err)
	}
}
func (pr *pingRouter) Handle(request zface.IRequest)  {
	conn := request.GetConnection().GetTcpConnection()
	_, err := conn.Write([]byte("[pingRouter] Main\n"))
	if err != nil {
		fmt.Println("[pingRouter]Error", err)
	}
}
func (pr *pingRouter) PostHandle(request zface.IRequest)  {
	conn := request.GetConnection().GetTcpConnection()
	_, err := conn.Write([]byte("[pingRouter] Post\n"))
	if err != nil {
		fmt.Println("[pingRouter]Error", err)
	}
}
func main() {
	server := znet.NewServer("ServerV0.3")
	server.AddRouter(&pingRouter{})
	server.Serve()
}