package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BotMethodGetUpdates   = "getUpdates"
	BotMethodSendMessage  = "sendMessage"
	BotMethodSendPhoto    = "sendPhoto"
	BotMethodSendDocument = "sendDocument"
)

const (
	BotAPIHost   = "https://api.telegram.org"
	BotAPIMethod = http.MethodPost
)

type botRequest struct {
	ChatID string `json:"chat_id,omitempty"`
}

type Bot struct {
	token string
}

func (b Bot) makeRequest(method string, data interface{}) error {
	var body bytes.Buffer
	if e := json.NewEncoder(&body).Encode(data); e != nil {
		return e
	}
	client := &http.Client{}
	url := fmt.Sprintf("%s/bot%s/%s", BotAPIHost, b.token, method)
	request, e := http.NewRequest(BotAPIMethod, url, &body)
	if e != nil {
		return e
	}

	request.Header.Add("Content-type", "application/json")
	response, e := client.Do(request)
	if e != nil {
		return e
	}
	defer response.Body.Close()
	// read response
	return nil
}

type BotSendMessageRequest struct {
	botRequest
	Text string `json:"text"`
}

func (b Bot) SendMessage(chat, text string) error {
	data := BotSendMessageRequest{
		botRequest: botRequest{chat},
		Text:       text,
	}
	return b.makeRequest(BotMethodSendMessage, data)
}

func NewBot(token string) *Bot {
	return &Bot{
		token: token,
	}
}
