package ziface

//路由接口， 这里面路由是使用框架者给该连接自定的处理业务的方法

type IRouter interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
