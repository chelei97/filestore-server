package main

import "learnZinx/src/znet"

//基于Zinx框架来开发的服务器端应用程序

func main(){
	//创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.1]")
	//启动server
	s.Serve()
}
