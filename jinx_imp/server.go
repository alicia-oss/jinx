package jinx_imp

import (
	"fmt"
	"github.com/alicia-oss/jinx/jinx_int"
	"github.com/alicia-oss/jinx/pkg/coder"
	"github.com/alicia-oss/jinx/pkg/log"
	"net"
)

func NewServer(name, proto, ip string, port int) jinx_int.IServer {
	return &server{
		name:           name,
		ipVersion:      proto,
		ip:             ip,
		port:           port,
		ipAddr:         fmt.Sprintf("%v:%v", ip, port),
		router:         NewRouter(),
		workerPool:     NewWorkerPool(4, 30),
		connManager:    NewConnManager(),
		coder:          &coder.TlvCoder{MaxPacketSize: 256},
		onCloseHandler: &BaseOnCloseHandler{},
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
	coder          jinx_int.ICoder
	onCloseHandler jinx_int.IOnCloseHandle
}

func (s *server) Start() error {
	cid := -1
	// 开启listen协程
	listen, err := net.Listen(s.ipVersion, s.ipAddr)
	if err != nil {
		log.Error(fmt.Sprintf("服务器 %v 启动失败， Listen err:%v", s.name, err), "Server")
		return err
	}
	log.Info(fmt.Sprintf("server started successfully, name:%v, ip_port:%v", s.name, s.ipAddr), "Server")
	s.workerPool.Start()
	for {
		cid++
		conn, err := listen.Accept()
		if err != nil {
			log.Error(fmt.Sprintf("服务器 %v Accept err:%v", s.name, err), "Server")
			return err
		}
		newConnection := NewConnection(conn, uint32(cid), s.coder, s.router, s.workerPool, s.connManager, s.onCloseHandler)
		// 判断当前是否在服务器关闭中
		if newConnection != nil {
			newConnection.Start()
		}
	}
}

func (s *server) Stop() {
	log.Info(fmt.Sprintf("server:%v stoped......", s.name), ModuleNameServer)
	s.connManager.ClearConn()
}

func (s *server) AddRouter(msgId uint32, router jinx_int.IMsgHandle) error {
	return s.router.AddRouter(msgId, router)
}

func (s *server) SetOnCloseHandler(handle jinx_int.IOnCloseHandle) {
	s.onCloseHandler = handle
}
