package chatwork

import (
	"fmt"
	"net/url"
	"strconv"
)

type Message struct {
	RoomId  uint   `json:"rid"`
	To      []uint `json:"to,omitempty"`
	Re      []uint `json:"re,omitempty"`
	Message string `json:"message,omitempty"`
	Info    string `json:"info,omitempty"`
	Title   string `json:"title,omitempty"`
}

type Chatwork struct {
	request *Request
	client  *Client
}

func NewChatwork(request *Request, client *Client) *Chatwork {
	return &Chatwork{
		request: request,
		client:  client,
	}
}

func (this *Chatwork) SendMessage(data *Message) ([]byte, error) {
	path := fmt.Sprintf("/rooms/%s/messages", strconv.FormatUint(uint64(data.RoomId), 10))

	values := url.Values{}
	values.Set("body", data.Message)

	request, err := this.request.Create("POST", path, values)
	if err != nil {
		return nil, err
	}
	responseBody, err := this.client.Send(request)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
