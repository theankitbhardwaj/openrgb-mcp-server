package openrgb

import (
	"github.com/csutorasa/go-openrgb-sdk"
)

type Client struct {
	c *openrgb.Client
}

func ConnectClient(host string, port int) (*Client, error) {
	cl, err := openrgb.NewClientHostPort(host, port)

	if err != nil {
		return nil, err
	}

	return &Client{cl}, nil
}

func (c *Client) Close() error {
	return c.c.Close()
}
