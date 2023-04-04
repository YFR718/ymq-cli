package ymq

import (
	"encoding/json"
	"github.com/YFR718/ymq-cli/pkg/common"
)

type Topic struct {
	Name        string
	Partitions  int
	Replication int
	MessageSize int
	Msg         []byte
}

//// 创建Topic类型的结构体，可变参数可以用默认值配置Topic
//func NewTopic(name string, partitions int, replication int, messageSize int) *Topic {
//	return &Topic{
//		Name:        name,
//		Partitions:  partitions,
//		Replication: replication,
//		MessageSize: messageSize,
//	}
//}

// 创建Topic
func (t *Topic) Create() error {
	header := common.Header{Type: common.CREATE_TOPIC}
	body, err := json.Marshal(t)
	println(string(body))
	if err != nil {
		common.PrintError(err)
		return err
	}
	conn, err := NewConnect(myClient.Config)
	if err != nil {
		common.PrintError(err)
		return err
	}
	defer conn.Close()

	msg, err := conn.SendMsg(header, body)
	if err != nil {
		common.PrintError(err)
		return err
	}
	println(string(msg.Body))
	return nil
}

//func (t *Topic) GetTopics() error {
//
//}

// 删除Topic
func (t *Topic) Delete() error {
	header := common.Header{Type: common.DELETE_TOPIC}
	body, err := json.Marshal(t)
	if err != nil {
		common.PrintError(err)
		return err
	}
	conn, err := NewConnect(myClient.Config)
	if err != nil {
		common.PrintError(err)
		return err
	}
	defer conn.Close()

	msg, err := conn.SendMsg(header, body)
	if err != nil {
		common.PrintError(err)
		return err
	}
	println(string(msg.Body))
	return nil
}

// 发送消息
func (t *Topic) Send(msg []byte) error {
	header := common.Header{Type: common.SEND_MESSAGE}
	t.Msg = msg
	body, err := json.Marshal(t)
	if err != nil {
		common.PrintError(err)
		return err
	}

	conn, err := NewConnect(myClient.Config)
	if err != nil {
		common.PrintError(err)
		return err
	}
	defer conn.Close()

	rev, err := conn.SendMsg(header, body)
	if err != nil {
		common.PrintError(err)
		return err
	}
	println(string(rev.Body))
	return nil
}

// 接收消息
func (t *Topic) Receive() error {
	return nil
}
