package znet

import "zinx/zface"

type Request struct {
	conn zface.IConnection
	data []byte
}

func (r *Request) GetConnection() zface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}