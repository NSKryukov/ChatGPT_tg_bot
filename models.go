package main

type Update struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		Text string `json:"text"`
		Chat struct {
			ChatId int `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type MessageToSend struct {
	Text   string `json:"text"`
	ChatId int    `json:"chat_id"`
}

type ChatRequest struct {
	Model    string `json:"model"`
	Messages [1]struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type СhatResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	}
}
