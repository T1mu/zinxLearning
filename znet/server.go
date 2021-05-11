package znet

import (
	"fmt"
	"net"
	"zinx/zface"
)

// Server IServer的接口实现，定义一个服务器的结构体
type Server struct {
	// 服务器的名称
	Name string
	// 服务器的通讯协议
	IpVersion string
	// 服务器的IP地址
	Ip string
	// 服务器的端口
	Port int
}

// Start 启动服务的实现
func (s *Server) Start() {
	go func() {
		// 通过net.ResolveTCPAddr方法获取TCP地址
		addr, err := net.ResolveTCPAddr(string(s.IpVersion), fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("解析TCPAddr出错:", err)
			return
		}
		// 通过net.listenTCP方法监听是否有连接接入
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("监听TCP地址出错:", err)
			return
		}
		fmt.Printf("[Got Listener...]\n[Wating For Conn...]:\n" +
			"[Server Infomation]Name:%s IpVersion:%s Ip:%s Port:%d\n",
			s.Name, s.IpVersion, s.Ip, s.Port)
		// 已经获取到listener，循环监听数据
		// 注意：读取到一个内容后，会for循环等待下一个conn连接
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("接受信息出错:", err)
				continue
			}
			// 开启go程读取连接信息，并将信息回显给客户端
			go func() {
				for {
					// 读取信息
					buff := make([]byte, 512)
					n, err := conn.Read(buff)
					if err != nil {
						fmt.Println("读数据出错:", err)
						continue
					}
					fmt.Print("[Read Buffers]", string(buff[:n]))
					// 回显功能（写数据）
					if _, err := conn.Write(buff[:n]); err != nil {
						fmt.Println("写数据出错:", err)
						continue
					}
				}
			}()
		}
	}()

}

// Stop 停止服务的实现
func (s *Server) Stop() {
	// TODO 释放资源
}

// Serve 运行服务的实现
func (s *Server) Serve() {
	s.Start()
	// TODO 服务器启动成功后需要做的一些事情
	select {}
}

// NewServer 创建一个Server对象，返回zface.IServer抽象对象
func NewServer(name string) zface.IServer {
	server := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "127.0.0.1",
		Port:      9573,
	}
	return server
}
