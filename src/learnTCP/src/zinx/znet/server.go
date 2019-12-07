package znet

import (
	"fmt"
	"learnTCP/src/zinx/utils"
	"learnTCP/src/zinx/ziface"
	"net"
	"time"
)

type Server struct {
	//服务器的名字
	Name string
	//协议类型
	IPVersion string
	//服务器绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
	//由用户绑定的回调router，也就是server注册的连接对应的处理业务
	Router ziface.IRouter
}

//实现接口中的方法

func (s *Server) Start(){
	fmt.Printf("[START] Server listen at IP: %s, Port : %d, is starting\n", s.IP, s.Port)

	//开启一个携程去做服务端的Listen业务
	go func(){
		// 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}

		//已经监听成功
		fmt.Println("start Zinx server ", s.Name, " success, now listening...")

		var cid uint32
		cid = 0
		//启动server网络连接业务
		for {
			//阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept ", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid ++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop(){
	fmt.Println("[STOP] Zinx server, name ", s.Name)
}

func (s *Server) Serve(){
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	s.Start()

	//是否需要做其他事，可以添加阻塞
	for {
		time.Sleep(10 * time.Second)
	}

}

func NewServer(name string) ziface.IServer {
	//初始化全局配置文件
	utils.GlobalObject.Reload()

	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}

	return s
}

//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router

	fmt.Println("Add Router succ! " )
}