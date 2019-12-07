package ziface

import "net"

type IConnection interface {
	//启动连接，使当前连接开始工作
	Start()
	//停止连接，结束当前连接状态
	Stop()
	//从当前连接获取原始socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获得当前连接ID
	GetConnID() uint32
	//获取远程客户端地址信息
	RemoteAddr() net.Addr
}

//定义一个统一处理链接业务的接口，第一个参数是socket原生连接，后两个是请求的数据
type HandFunc func(*net.TCPConn, []byte, int) error
