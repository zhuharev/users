package rpc

import (
	"github.com/zhuharev/users"
	"net/rpc"
)

type Client struct {
	Port   string
	client *rpc.Client
}

func NewClient(port string) (*Client, error) {
	c := new(Client)
	c.Port = port

	client, e := rpc.DialHTTP("tcp", ":"+port)
	if e != nil {
		return nil, e
	}
	c.client = client

	return c, nil
}

func (c *Client) Get(id int64) (*users.User, error) {
	req := new(GetRequest)
	resp := GetResponse{}
	e := c.client.Call("Users.Get", req, &resp)
	if e != nil {
		return nil, e
	}

	return resp.User, nil
}
