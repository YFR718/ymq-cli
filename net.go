package ymq

import (
	"encoding/binary"
	"github.com/YFR718/ymq-cli/pkg/common"
	"net"
)

type Connect struct {
	Conn net.Conn
}

func NewConnect(config *Config) (*Connect, error) {
	conn, err := net.Dial("tcp", config.Addrs)
	if err != nil {
		return nil, err
	}
	return &Connect{Conn: conn}, nil
}

func (c *Connect) SendMsg(header common.Header, body []byte) (*common.Message, error) {
	// 发送数据
	msg := common.Message{Header: header, Body: body}

	s := msg.Marshal()

	_, err := c.Conn.Write(s)

	// 读取返回
	// 读取4字节长度
	// 读取数据包头，通常包头包含了数据包的长度信息
	length := uint32(0)
	err = binary.Read(c.Conn, binary.BigEndian, &length)
	if err != nil {
		common.PrintError(err)
		return nil, err
	}

	// 读取length长度数据
	// 根据数据包头中的长度信息，从 TCP 连接中读取相应长度的数据
	data := make([]byte, length)
	_, err = c.Conn.Read(data[4:])
	if err != nil {
		common.PrintError(err)
		return nil, err
	}
	binary.BigEndian.PutUint32(data[:4], length)

	message, err := common.Unmarshal(data)
	if err != nil {
		common.PrintError(err)
		return nil, err
	}
	//println("get a message:", string(message.Body))

	return message, nil
}

func (c *Connect) Close() {
	c.Conn.Close()
}
