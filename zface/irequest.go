package zface

/*

 */

type IRequest interface {
//	GetConnection 得到当前连接
	GetConnection() IConnection
//	GetData 得到消息数据
	GetData() []byte
}