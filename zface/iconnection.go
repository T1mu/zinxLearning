package zface

import "net"

// 定义连接模块的抽象层

type IConnection interface {
	// Start 启动连接，让当前的连接准备开始工作
	Start()
	// Stop 停止连接，结束当前连接的工作
	Stop()
	// GetTcpConnection 获取当前连接绑定的socket的connection
	GetTcpConnection() *net.TCPConn
	// GetConnID 获取当前连接模块的连接ID
	GetConnID() int
	// RemoteAddr 获取远程客户端的信息
	RemoteAddr() net.Addr
	// SendMsg 发送数据
	SendMsg(data []byte) error
}

// HandleFunc 定义一个处理业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error