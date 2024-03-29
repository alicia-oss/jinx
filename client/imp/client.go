package client_imp

import (
	"fmt"
	client_int "github.com/alicia-oss/jinx/client/int"
	"github.com/alicia-oss/jinx/jinx_int"
	"github.com/alicia-oss/jinx/pkg/coder"
	"github.com/alicia-oss/jinx/pkg/log"
	"net"
	"sync"
	"time"
)

func NewClient(proto, ip string, port int) (client_int.IClient, error) {
	ipAddr := fmt.Sprintf("%s:%v", ip, port)
	conn, err := net.Dial(proto, ipAddr)
	if err != nil {
		log.Error(fmt.Sprintf("client start dial error:%v", err), ModuleNameClient)
		return nil, err
	}
	return &client{
		conn:       conn,
		proto:      proto,
		serverIP:   ip,
		serverPort: port,
		serverAddr: ipAddr,
		closeChan:  make(chan struct{}),
		writeChan:  make(chan []byte, 3),
		coder:      &coder.TlvCoder{MaxPacketSize: 512},
		router:     NewRouter(),
		attr:       &sync.Map{},
		ticker:     time.Tick(3 * time.Second),
	}, nil
}

type client struct {
	conn       net.Conn
	proto      string
	serverIP   string
	serverPort int
	// ip:port
	serverAddr string
	closeChan  chan struct{}
	writeChan  chan []byte
	coder      jinx_int.ICoder
	router     client_int.IRouter
	attr       *sync.Map
	ticker     <-chan time.Time
}

func (c *client) GetTCPConnection() net.Conn {
	return c.conn
}

func (c *client) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *client) Send(data []byte, msgId uint32) error {
	encode, err := c.coder.Encode(data, msgId)
	if err != nil {
		log.Error(fmt.Sprintf("client send encode err, err:%v", err), ModuleNameClient)
		return err
	}
	c.writeChan <- encode
	return nil
}

func (c *client) Start() {

	log.Info(fmt.Sprintf("client started successfully, remote_addr:%v", c.GetRemoteAddr()), ModuleNameClient)
	c.StartReader()
	c.StartWriter()

}

func (c *client) StartReader() {
	go func() {
		for true {
			select {
			case <-c.closeChan:
				log.Info(fmt.Sprintf("client reader closed..., remote_addr:%v", c.GetRemoteAddr()), ModuleNameClient)
				return
			default:
				message, err := c.coder.Decode(c.conn)
				if err != nil {
					log.Error(fmt.Sprintf("client decode error:%v", err), ModuleNameClient)
					c.Close()
					break
				}
				req := NewRequest(message, c)
				handler, ok := c.router.Route(req)
				if !ok {
					log.Error(fmt.Sprintf("conn route error, msg_id:%v, err:%v", req.GetMsgId(), err), ModuleNameClient)
					continue
				}

				handler.PreHandle(req)
				handler.Handle(req)
				handler.PostHandle(req)
			}
		}
	}()
}

func (c *client) StartWriter() {
	go func() {
		for true {
			select {
			case <-c.closeChan:
				log.Info(fmt.Sprintf("client reader closed..., remote_addr:%v", c.GetRemoteAddr()), ModuleNameClient)
				return
			case <-c.ticker:
				log.Info(fmt.Sprintf("ping, remote_addr:%v", c.GetRemoteAddr()), ModuleNameClient)
				err := c.Send(nil, 1)
				if err != nil {
					log.Error(fmt.Sprintf("write err, err:%v", err), ModuleNameClient)
					c.Close()
					return
				}
			case data := <-c.writeChan:
				if _, err := c.conn.Write(data); err != nil {
					log.Error(fmt.Sprintf("write err, err:%v", err), ModuleNameClient)
					c.Close()
					return
				}
			}
		}
	}()

}

func (c *client) Close() {
	log.Info("client closed...", ModuleNameClient)
	close(c.writeChan)
	close(c.closeChan)
}

func (c *client) AddRoute(msg uint32, handle client_int.IMsgHandle) error {
	err := c.router.AddRouter(msg, handle)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) SetAttr(key string, value interface{}) {
	c.attr.Store(key, value)
}

func (c *client) GetAttr(key string) (interface{}, bool) {
	return c.attr.Load(key)
}

func (c *client) DeleteAttr(key string) {
	c.attr.Delete(key)
}
