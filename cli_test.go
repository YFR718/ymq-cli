package ymq

import (
	"testing"
	"time"
)

func Test_Create(t *testing.T) {

	cli, err := NewProducer(&Config{Addrs: "127.0.0.1:8848"})
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()

	err = cli.Send("test", []byte("hello"))
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * 1)
	err = cli.Send("Close", []byte{})
	if err != nil {
		t.Error(err)
	}

}

func Test_Create_Topic(t *testing.T) {
	topic := Topic{Name: "topic1"}
	err := topic.Create()
	if err != nil {
		t.Error(err)
	}
	err = topic.Delete()
	if err != nil {
		t.Error(err)
	}
}
