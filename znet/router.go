package znet

import "zinx/zface"

// 	BaseRouter 实现router时，可先嵌入这个BaseRouter
//	应当继承这个BaseRouter
type BaseRouter struct {

}

func (br *BaseRouter) PreHandle(request zface.IRequest)  {
	
}
func (br *BaseRouter) Handle(request zface.IRequest)  {
	
}
func (br *BaseRouter) PostHandle(request zface.IRequest)  {
	
}