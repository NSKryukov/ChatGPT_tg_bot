package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	chatGPTAPIToken := os.Getenv("OPENAI_API_KEY")
	botToken := os.Getenv("CHAT_BOT_TOKEN")
	flag.StringVar(&chatGPTAPIToken, "key", chatGPTAPIToken, "OpenAI API Token")
	flag.StringVar(&botToken, "token", botToken, "Telegram bot token")
	flag.Parse()
	botAPI := "https://api.telegram.org/bot"
	botURL := botAPI + botToken
	ChatGptApiCompletionsURL := "https://api.openai.com/v1/chat/completions"

	offset := 0
	for true {
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			fmt.Println("Can`t get updates from ChatGPTProvider bot, error:", err)
		}
		for _, update := range updates {
			offset = update.UpdateId + 1
			if update.Message.Text == "/start" {
				update.Message.Text = "Hey, I`m chatGPT bot. Rules of communication with me are simple:\n" +
					"1. When you write a message - I give the answer from chat GPT\n" +
					"2. Remember that it`s advisably to send messages with ~5 seconds period"
				err = sendMessage(update, botURL)
			} else {
				gptResp, err := requestToChatGPT(update, chatGPTAPIToken, ChatGptApiCompletionsURL)
				if err != nil {
					return
				}
				err = sendMessage(gptResp, botURL)
				if err != nil {
					fmt.Println("Can`t send message, error:", err)
				}
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}
