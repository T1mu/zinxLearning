package zface

/*
	路由模块，用于处理不同消息
	路由里的消息均是IRequest
 */

type IRouter interface {
//	在处理业务之前的钩子方法 Hook
	PreHandle(request IRequest)
//	在处理业务的主钩子方法 Hook
	Handle(request IRequest)
//	在处理业务之后的钩子方法 Hook
	PostHandle(request IRequest)
}