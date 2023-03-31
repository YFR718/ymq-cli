package ymq

type Config struct {
	Addrs string
}

type Client struct {
	Config *Config
}

var myClient *Client

func Init(config *Config) {
	myClient = &Client{Config: config}

}

//func NewProducer(config *Config) (*Client, error) {
//
//	conn, err := net.Dial("tcp", config.Addrs)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Client{Config: config, Conn: conn}, nil
//
//}

//func (c *Client) Close() {
//	c.Conn.Close()
//}
//
//func (c *Client) Publish(topic string, data []byte) error {
//	return nil
//}
//
//func (c *Client) Send(Topic string, Data []byte) error {
//	// 发送数据
//	_, err := c.Conn.Write([]byte(Topic + string(Data))) // 发送数据
//	fmt.Println("发送数据", Topic, string(Data))
//	if err != nil {
//		return err
//	}
//	return nil
//}
