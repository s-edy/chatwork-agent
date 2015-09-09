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
	code    int
	Message string `json:"message"`
}

type Server struct {
	config ChatworkConfig
	chatwork *Chatwork
}

func NewServer(config ChatworkConfig) *Server {
	return &Server{
		config: config,
		chatwork: CreateChatwork(config),
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
	errorResponse := this.assertPOSTMethod(writer, request)
	if errorResponse != nil {
		this.respondError(errorResponse, writer)
		return
	}
	data := &Message{}
	errorResponse = this.parseRequest(data, writer, request)
	if errorResponse != nil {
		this.respondError(errorResponse, writer)
		return
	}

	go func(data *Message) {
		_, err := this.chatwork.SendMessage(data)
		if err != nil {
			log.Println(err)
		}
	}(data)

	writer.WriteHeader(202)
}

func (this *Server) assertPOSTMethod(writer http.ResponseWriter, request *http.Request) *ErrorResponse {
	if request.Method == "POST" {
		return nil
	}

	return &ErrorResponse{
		code: 400,
		Message: "Acceptable HTTP request method is only POST.",
	}
}

func (this *Server) parseRequest(data interface{}, writer http.ResponseWriter, request *http.Request) *ErrorResponse {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	if err == nil {
		return nil
	}

	return &ErrorResponse{
		code: 500,
		Message: err.Error(),
	}
}

func (this *Server) respondError(err *ErrorResponse, writer http.ResponseWriter) {
	writer.WriteHeader(err.code)
	json.NewEncoder(writer).Encode(*err)
}
