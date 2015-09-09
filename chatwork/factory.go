package chatwork

func CreateChatwork(config ChatworkConfig) *Chatwork {
	request := NewRequest(config.Token)
	client := NewClient()

	return NewChatwork(request, client)
}