package main

import "learnTCP/src/zinx/znet"

func main(){
	//创建服务器
	s := znet.NewServer("Zinx1")
	//启动服务器
	s.Serve()
}


