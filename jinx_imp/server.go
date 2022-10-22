package jinx_imp

import (
	"fmt"
	"jinx/jinx_int"
	"jinx/pkg/coder"
	"jinx/pkg/log"
	"net"
)

func NewServer(name, proto, ip string, port int) jinx_int.IServer {
	return &server{
		name:        name,
		ipVersion:   proto,
		ip:          ip,
		port:        port,
		ipAddr:      fmt.Sprintf("%v:%v", ip, port),
		router:      NewRouter(),
		workerPool:  NewWorkerPool(4, 30),
		connManager: NewConnManager(),
		coder:       &coder.TlvCoder{MaxPacketSize: 256},
	}
}

type server struct {
	//服务器的名称
	name string
	//tcp4 or other
	ipVersion string
	//服务绑定的IP地址
	ip string
	//服务绑定的端口
	port int
	//ip:port
	ipAddr string
	//路由器 传给connction
	router jinx_int.IRouter
	//工作线程池 传给connction
	workerPool jinx_int.IWorkerPool
	//连接管理器
	connManager jinx_int.IConnManager
	//coder 传给connction
	coder jinx_int.ICoder
}

func (s *server) Start() {
	go func() {
		s.workerPool.Start()
		cid := -1
		// 开启listen协程
		listen, err := net.Listen(s.ipVersion, s.ipAddr)
		if err != nil {
			log.Error(fmt.Sprintf("服务器 %v 启动失败， Listen err:%v", s.name, err), "Server")
			return
		}
		log.Info(fmt.Sprintf("server started successfully, name:%v, ip_port:%v", s.name, s.ipAddr), "Server")
		for {
			cid++
			conn, err := listen.Accept()
			if err != nil {
				log.Error(fmt.Sprintf("服务器 %v Accept err:%v", s.name, err), "Server")
				return
			}
			newConnection := NewConnection(conn, uint32(cid), s.coder, s.router, s.workerPool, s.connManager)
			// 判断当前是否在服务器关闭中
			if newConnection != nil {
				newConnection.Start()
			}
		}
	}()

}

func (s *server) Stop() {
	log.Info(fmt.Sprintf("server:%v stoped......", s.name), ModuleNameServer)
	s.connManager.ClearConn()
}

func (s *server) AddRouter(msgId uint32, router jinx_int.IMsgHandle) error {
	return s.router.AddRouter(msgId, router)
}
