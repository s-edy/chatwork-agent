package chatwork

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
)

type Request struct {
	token string
}

func NewRequest(token string) *Request {
	return &Request{
		token: token,
	}
}

func (this *Request) Create(method, path string, values url.Values) (*http.Request, error) {
	urlString := ENDPOINT + API_VERSION + path
	request, err := http.NewRequest(method, urlString, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))
	request.Header.Add("X-ChatWorkToken", this.token)

	return request, nil
}
