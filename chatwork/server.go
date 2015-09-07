package chatwork

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Configuration ChatworkConfig

type ChatworkConfig struct {
	Port  uint16 `toml:"port"`
	Token string `toml:"token"`
}

type ErrorResponse struct {
	Message string
}

type Server struct {
	config ChatworkConfig
}

func NewServer(config ChatworkConfig) *Server {
	return &Server{
		config: config,
	}
}

func (this *Server) Start() error {
	http.HandleFunc("/send", this.Send)

	startMessage := `Chatwork proxy %s Server started at %s
Listening on http://localhost%s
Press Ctrl-C to quit.`

	addr := fmt.Sprintf(":%d", this.config.Port)
	log.Printf(startMessage, VERSION, time.Now().String(), addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}

func (this *Server) Send(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(400)
		response := ErrorResponse{"POST only"}
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
		return
	}
	data := &Message{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	if err != nil {
		writer.WriteHeader(500)
		response := ErrorResponse{err.Error()}
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
		return
	}

	go func(data *Message) {
		request := NewRequest(this.config.Token)
		client := NewClient()
		chatwork := NewChatwork(request, client)
		_, err := chatwork.SendMessage(data)
		if err != nil {
			log.Println(err)
		}
	}(data)

	writer.WriteHeader(202)
}
