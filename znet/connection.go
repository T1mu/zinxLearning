package znet

import (
	"fmt"
	"net"
	"zinx/zface"
)

/*
	连接模块
*/

// 	Connection 连接模块 包含一系列连接对象的属性 并且包含一个HandleAPI回调函数用于处理连接
//	Conn TCPConn（TCP套接字）变量
//	ConnId 连接ID变量
type Connection struct {
	//	当前连接的tcp套接字
	Conn *net.TCPConn
	//	连接的ID
	ConnId int
	//	当前连接的状态
	IsClosed bool
	//	通知退出管道
	ExitChan chan bool
	//	Router
	Router zface.IRouter
}

//	StartReader 读方法 在Connection.Start方法中被调用
//	具体功能为：读取Buffer并将其作为HandleAPI的参数输入
func (c *Connection) StartReader() {
	//	输出当前连接相关信息，连接的ID号和连接客户端信息
	fmt.Printf("\t[Reading]\tAboutConnection:\n\t\t"+
		"Connection ID=%d\n\t\tRemoteAddr=%s\n",
		c.ConnId, c.RemoteAddr().String())
	defer fmt.Printf("\t[Read\tExit]\n\t\tConnection ID=%d\n", c.ConnId)
	defer c.Stop()
	//	读
	for {
		buff := make([]byte, 512)
		cnt, err := c.Conn.Read(buff)
		if err != nil {
			fmt.Println("\t[Reading\tError]", err)
		}
		//	输出读取到的数据流信息，包括数据内容与输出长度
		fmt.Printf("\t[Reading] AboutBuffers:\n"+
			"\t\tBuffers:%s\n\t\tcnt=%d\n",
			buff[:cnt-1], cnt)
		//	集成数据到Request中，作为路由方法的参数
		req := Request{
			conn: c,
			data: buff,
		}
		//	执行路由方法
		go func(request zface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

// 	实现IConnection接口的6个方法

// 	Start 启动连接，让当前的连接准备开始工作
//	包含StartReader读Buffers方法
func (c *Connection) Start() {
	fmt.Printf("[ConnStart]\n\tConnection ID=%d\n", c.ConnId)
	// 读数据
	go c.StartReader()
}

// 	Stop 停止连接，结束当前连接的工作
//	设置连接的IsClosed属性 关闭Conn并退出管道
func (c *Connection) Stop() {
	fmt.Printf("[Stop] Connection ID=%d\n", c.ConnId)
	//	若已经关闭连接
	if c.IsClosed == true {
		return
	}
	// 设置标志位
	c.IsClosed = false
	// 释放资源
	c.Conn.Close()
	close(c.ExitChan)

}

// GetTcpConnection 获取当前连接绑定的socket的connection
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接模块的连接ID
func (c *Connection) GetConnID() int {
	return c.ConnId
}

// RemoteAddr 获取远程客户端的信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 发送数据
func (c *Connection) SendMsg(data []byte) error {
	return nil
}

//	NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, id int,
	router zface.IRouter) zface.IConnection {
	c := &Connection{
		Conn:     conn,
		ConnId:   id,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool),
	}
	return c
}
