package chatwork

import (
	"fmt"
	"errors"
)

type Message struct {
	RoomId  uint   `json:"rid"`
	To      []uint `json:"to,omitempty"`
	Message string `json:"message,omitempty"`
	Info    string `json:"info,omitempty"`
	Title   string `json:"title,omitempty"`
}

func (this *Message) Build() (string, error) {
	body, err := this.Body()
	if err != nil {
		return "", err
	}

	envelope := this.Envelope()
	if len(envelope) == 0 {
		return body, nil
	}

	return envelope + "\n" + body, nil
}

func (this *Message) Envelope() string {
	if len(this.To) == 0 {
		return ""
	}
	envelope := ""
	for _, to := range this.To {
		envelope += fmt.Sprintf("[To:%d]", to)
	}

	return envelope
}

func (this *Message) Body() (string, error) {
	if this.Message != "" {
		return this.Message, nil
	}
	if this.Info == "" {
		return "", errors.New("Message is empty")
	}

	title := "";
	if this.Title != "" {
		title = fmt.Sprintf("[title]%s[/title]", this.Title)
	}

	return fmt.Sprintf("[info]%s%s[/info]", title, this.Info), nil
}
