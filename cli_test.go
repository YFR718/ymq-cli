package ymq

import (
	"testing"
)

func Test_Create(t *testing.T) {

	Init(&Config{Addrs: "127.0.0.1:8848"})
	topic := Topic{Name: "topic1"}

	err := topic.Create()
	if err != nil {
		t.Error(err)
	}

	err = topic.Send([]byte("i have a pen."))
	if err != nil {
		t.Error(err)
	}

	err = topic.Send([]byte("i have an apple."))
	if err != nil {
		t.Error(err)
	}

	err = topic.Delete()
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
