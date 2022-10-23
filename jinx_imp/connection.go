package jinx_imp

import (
	"fmt"
	"github.com/alicia-oss/jinx/jinx_int"
	"github.com/alicia-oss/jinx/pkg/log"
	"net"
)

func NewConnection(conn net.Conn, connId uint32, coder jinx_int.ICoder, router jinx_int.IRouter, pool jinx_int.IWorkerPool, manager jinx_int.IConnManager) jinx_int.IConnection {
	c := &connection{
		conn:        conn,
		connID:      connId,
		isClosed:    false,
		exitChan:    make(chan struct{}),
		coder:       coder,
		router:      router,
		workerPool:  pool,
		connManager: manager,
		writeChan:   make(chan []byte, 3),
	}
	if err := c.connManager.Add(c); err != nil {
		return nil
	}
	return c
}

type connection struct {
	//当前链接的socket TCP套接字
	conn net.Conn
	//当前链接ID
	connID uint32
	//当前链接是否关闭
	isClosed bool
	//通知当前链接已经退出/停止的channel
	exitChan chan struct{}
	//写数据channel
	writeChan chan []byte
	//编解码器 对应 应用层协议
	coder jinx_int.ICoder
	//路由器
	router jinx_int.IRouter
	//工作线程池
	workerPool jinx_int.IWorkerPool
	//连接管理器
	connManager jinx_int.IConnManager
}

// Start 开启读协程 负责读取数据转换为IRequest
func (c *connection) Start() {
	c.StartWriter()
	c.StartReader()
	log.Info(fmt.Sprintf("connection started successfully, conn_id:%v, remote_addr:%v", c.connID, c.GetRemoteAddr()), ModuleNameConn)

}

func (c *connection) StartReader() {
	go func() {
		for {
			select {
			case <-c.exitChan:
				log.Info(fmt.Sprintf("conn reader closed..., conn_id:%v, remote_addr:%v", c.connID, c.GetRemoteAddr()), ModuleNameConn)
				return
			default:
				message, err := c.coder.Decode(c.conn)
				if err != nil {
					log.Error(fmt.Sprintf("conn decode error, conn_id:%v, err:%v", c.connID, err), ModuleNameConn)
					c.Stop()
					return
				}
				req := NewRequest(message, c)
				handler, ok := c.router.Route(req)
				if !ok {
					log.Error(fmt.Sprintf("conn route error, conn_id:%v, msg_id:%v, err:%v", c.connID, req.GetMsgId(), err), ModuleNameConn)
					continue
				}
				c.workerPool.Submit(handler, req)
			}
		}

	}()
}

// Stop 关闭读协程
func (c *connection) Stop() {
	log.Info(fmt.Sprintf("conn stop, conn_id:%v", c.connID), ModuleNameConn)
	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	close(c.exitChan)
	close(c.writeChan)
	c.isClosed = true
	c.connManager.Remove(c)
}

func (c *connection) GetTCPConnection() net.Conn {
	return c.conn
}

func (c *connection) GetConnID() uint32 {
	return c.connID
}

func (c *connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *connection) Send(data []byte, msgId uint32) error {
	encode, err := c.coder.Encode(data, msgId)
	if err != nil {
		log.Error(fmt.Sprintf("conn send encode err, err:%v", err), ModuleNameConn)
		return err
	}
	c.writeChan <- encode
	return nil
}

func (c *connection) StartWriter() {
	go func() {
		for true {
			select {
			case <-c.exitChan:
				log.Info(fmt.Sprintf("conn writer closed..., conn_id:%v, remote_addr:%v", c.connID, c.GetRemoteAddr()), ModuleNameConn)
				return
			case data := <-c.writeChan:
				if _, err := c.conn.Write(data); err != nil {
					log.Error(fmt.Sprintf("conn:%v write err, err:%v", c.connID, err), ModuleNameConn)
					c.Stop()
					return
				}
			}
		}
	}()
}
