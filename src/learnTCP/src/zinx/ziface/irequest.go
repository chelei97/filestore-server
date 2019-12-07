package ziface

//把客户端请求的连接信息和请求的数据封装到了Request中

type IRequest interface{
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
}
