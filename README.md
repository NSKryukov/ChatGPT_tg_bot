# Chat GPT Telegram bot
[![Golang](https://img.shields.io/github/go-mod/go-version/nskryukov/chatgpt_tg_bot)](https://go.dev/blog/go1.18)
[![OpenAI_API](https://img.shields.io/badge/OpenAI%20API-April%2C%202023-blue)](https://platform.openai.com/docs/guides/chat)
[![Bot_API](https://img.shields.io/badge/Telegram%20Bot%20API-April%2C%202023-blue)](https://core.telegram.org/bots/api)

Telegram bot that responds to user's messages with responses from Chat GPT version **3.5 - turbo**

# Bot functions
1. Command ```/start``` - sends a message to user with a greeting and rules of use
2. Other messages are perceived as requests to ChatGPT API. Bot requests with user`s ```message``` and reponses with ChatGPT answer **without changing it**

# Installation and Deploying

#### Two ways to start bot:

**Important** - You need to have Telegram Bot Token (ask BotFather about it) and Open AI API Key (go to [OpenAI API keys generate link](https://platform.openai.com/account/api-keys)). To specify them add environment variables ```OPENAI_API_KEY``` and ```CHAT_BOT_TOKEN``` to your machine or specify at startup flags ```-key=<OPENAI_API_KEY> -token=<CHAT_BOT_TOKEN>```. Also you can mix it, for example, specify OPENAI_API_KEY as environment variable and CHAT_BOT_TOKEN as flag at startup

* If you have Golang installed on your machine do ```go run .``` in directory with code

* If you have not Golang just download [chatgptprovider.exe](https://github.com/NSKryukov/chatgpt_tg_bot/blob/main/chatgptprovider.exe) file and start it (for Windows OS)

**P.S.** I started it on my VDS server (check it [link to tg bot](https://t.me/chatgptprovider_bot)) on Ubuntu 20.04 LTS that way:

1. ```sudo apt install wine``` install tool "wine" to start .exe files on unix, because my machine do not have Golang compilator
2. ```nohup wine chatgptprovider.exe -token=<my_bot_token> -key=<my_openai_api_key> &``` - it`s run .exe file is in background, will not allow it to finish, create file nohup.out and will write stdout and stderr to it
