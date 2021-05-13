package zface
type IServer interface {
	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Serve 运行服务器
	Serve()

//	AddRouter 路由功能：给当前的服务增加一个Router方法，供客户端连接使用
	AddRouter(IRouter)
}