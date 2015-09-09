package chatwork

import (
	"fmt"
	"net/url"
	"strconv"
)

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

func (this *Chatwork) SendMessage(message *Message) ([]byte, error) {
	path := fmt.Sprintf("/rooms/%s/messages", strconv.FormatUint(uint64(message.RoomId), 10))

	values := url.Values{}

	body, err := message.Build()
	if err != nil {
		return nil, err
	}
	values.Set("body", body)

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
