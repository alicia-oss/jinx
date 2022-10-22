package jinx_int

type IConnManager interface {
	Add(conn IConnection) error            //添加链接
	Remove(conn IConnection)               //删除连接
	Get(connID uint32) (IConnection, bool) //利用ConnID获取链接
	Len() int                              //获取当前连接数
	ClearConn()                            //删除并停止所有链接
}
