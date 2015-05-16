package client

import (
	log "github.com/Sirupsen/logrus"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Send(point interface{}) error {
	log.Debugf("sending point: %#v", point)

	return nil
}
