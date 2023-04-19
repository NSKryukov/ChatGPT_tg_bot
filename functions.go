package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func getUpdates(url string, offset int) ([]Update, error) {
	resp, err := http.Get(url + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func requestToChatGPT(Update Update, chatGPTAPIToken string, ChatGptApiCompletionsURL string) (Update, error) {
	var ChatRequest ChatRequest
	ChatRequest.Model = "gpt-3.5-turbo"
	ChatRequest.Messages[0].Role = "user"
	ChatRequest.Messages[0].Content = Update.Message.Text

	buf, err := json.Marshal(ChatRequest)
	if err != nil {
		Update.Message.Text = "Error code, try again later"
		return Update, err
	}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+chatGPTAPIToken)
	req, err := http.NewRequest("POST", ChatGptApiCompletionsURL, bytes.NewBuffer(buf))
	if err != nil {
		Update.Message.Text = "Error while preparing HTTP request, try again later"
		return Update, err
	}
	req.Header = headers
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Update.Message.Text = "Requesting ChatGPT API went wrong, try again later"
		return Update, err
	}

	var СhatResponse СhatResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		Update.Message.Text = "Error serializing ChatGPT answer, try again later"
		return Update, err
	}
	err = json.Unmarshal(bodyBytes, &СhatResponse)
	if err != nil {
		Update.Message.Text = "Error serializing ChatGPT answer, try again later"
		return Update, err
	}

	defer resp.Body.Close()
	if len(СhatResponse.Choices) != 0 {
		Update.Message.Text = СhatResponse.Choices[0].Message.Content
	} else {
		Update.Message.Text = "Error in Open AI response, please wait 5 seconds before your next request"
	}
	return Update, nil
}

func sendMessage(Update Update, botUrl string) error {
	var MessageToSend MessageToSend
	MessageToSend.Text = Update.Message.Text
	MessageToSend.ChatId = Update.Message.Chat.ChatId

	buf, err := json.Marshal(MessageToSend)
	if err != nil {
		return err
	}
	_, err = http.Post(
		botUrl+"/sendMessage",
		"application/json",
		bytes.NewBuffer(buf),
	)
	if err != nil {
		return err
	}
	return nil

}
