package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	tokenOpenAI string
	tokenBot    string
)

func main() {
	workDir, err := os.Executable()
	if err != nil {
		fmt.Println("Can not get current dir to create log files, error: ", err)

		return
	}

	logFile, err := os.Create(filepath.Dir(workDir) + "/bot_error.log")
	if err != nil {
		fmt.Println("Can crete log file, error: ", err)

		return
	}

	logger := log.New(logFile, "INFO:\t", log.Ldate|log.Ltime)

	flag.StringVar(&tokenOpenAI, "key", tokenOpenAI, "OpenAI API Token")
	flag.StringVar(&tokenBot, "token", tokenBot, "Telegram bot token")
	flag.Parse()

	if tokenOpenAI == "" || tokenBot == "" {
		fmt.Println("Please provide openAI and telegram bot token\nBot tg token: ", tokenBot, "\nopenAI token: ", tokenOpenAI)

		return
	}

	fmt.Println("BOOTSTRAP: Starting ChatGPT bot\nBot tg token: ", tokenBot, "\nopenAI token: ", tokenOpenAI)

	botAPI := "https://api.telegram.org/bot"
	botURL := botAPI + tokenBot
	ChatGptApiCompletionsURL := "https://api.openai.com/v1/chat/completions"

	offset := 0
	for true {
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			logger.Println("Can`t get updates from ChatGPTProvider bot, error: ", err)
		}
		for _, update := range updates {
			offset = update.UpdateId + 1
			if update.Message.Text == "/start" {
				update.Message.Text = "Hey, I`m chatGPT bot. Rules of communication with me are simple:\n" +
					"1. When you write a message - I give the answer from chat GPT\n" +
					"2. Remember that it`s advisably to send messages with ~5 seconds period"

				err = sendMessage(update, botURL)
				if err != nil {
					logger.Println("Can`t send message, error: ", err)
				}
			} else {
				gptResp, err := requestToChatGPT(update, tokenOpenAI, ChatGptApiCompletionsURL)
				if err != nil {
					logger.Println("Error in sending request to chatGPT, error: ", err)

					return
				}

				err = sendMessage(gptResp, botURL)
				if err != nil {
					logger.Println("Can`t send message, error:", err)
				}
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}
