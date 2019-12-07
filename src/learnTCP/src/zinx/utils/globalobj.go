package utils

import (
	"encoding/json"
	"io/ioutil"
	"learnTCP/src/zinx/ziface"
)

type GlobalObj struct {
	TcpServer 	ziface.IServer 	//全局Server对象
	Host 		string			//当前服务器主机IP
	TcpPort 	int				//当前服务器主机监听端口号
	Name 		string			//当前服务器名称
	Version 	string			//当前服务器主机允许的最大链接数

	MaxPacketSize uint32		//都需数据包最大值
	MaxConn 	  int			//当前服务器主机允许最大连接个数
}

//定义一个全局对象
var GlobalObject *GlobalObj

//读取用户配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//提供一个init方法，默认加载
func (g *GlobalObj) init() {
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:    "ZinxServerApp",
		Version: "V0.4",
		TcpPort: 7777,
		Host:    "0.0.0.0",
		MaxConn: 12000,
		MaxPacketSize:4096,
	}

	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}