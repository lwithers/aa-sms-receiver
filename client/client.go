/*
Package client implements a very simple client for waiting for a message to be
received by the aa-sms-receiver server. It returns that message.
*/
package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is a simple, thread-safe client that allows retrieval of a message.
// It has a basic timeout mechanism.
type Client struct {
	url string
}

// New returns a new client object. Address should be in the form "host:port".
func New(addr string) *Client {
	return &Client{
		url: fmt.Sprintf("http://%s/last-message", addr),
	}
}

// GetMessage will retrieve the last SMS message received by aa-sms-receiver.
// It will wait for a message if none was received; this has a timeout of
// roughly 15 seconds.
func (cl *Client) GetMessage() (string, error) {
	for retry := 0; retry < 15; retry++ {
		resp, err := http.Get(cl.url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			return "", errors.New(resp.Status)
		}
		if len(body) == 0 {
			time.Sleep(time.Second)
			continue
		}

		return string(body), nil
	}

	return "", errors.New("timed out")
}
