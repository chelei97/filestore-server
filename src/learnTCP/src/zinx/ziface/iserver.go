package ziface

type IServer interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//开启服务器业务的方法
	Serve()
	//路由功能：给当前服务注册一个路由方法
	AddRouter(router IRouter)
}