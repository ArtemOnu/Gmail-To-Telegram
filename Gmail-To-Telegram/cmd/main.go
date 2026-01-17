package main

import (
	"fmt"
	Config "gomod/internal/config"
	Logger "gomod/internal/entities"
	tgbot "gomod/internal/use-cases"
	imap "gomod/internal/use-cases/mailbox"
	"time"
)

func main() {
	fmt.Println("Application running . . . ")

	Logger.LoggerIni()
	Logger.Log("application start")

	//We get the configuration structure from the file
	var config Config.Config
	err := Config.GetConfig(&config)
	if err != nil {
		panic(err)
	}
	//Create class MailBox
	Logger.Log("Mail create")
	Mailbox := imap.Mail{}
	Mailbox.Connect(config.Mail, config.Password)
	Logger.Log("Mail connect")
	defer Mailbox.Disconnect()
	results := Mailbox.MailUpdate()

	//Create class Bot
	bot := tgbot.Bot{}
	//Initialization Bot
	//CHATID TEST
	err = bot.Initialization(config.Token, config.ChatID)
	if err != nil {
		panic(err)
	}

	for result := range results {
		time.Sleep(time.Second * 2)
		Logger.Log("new message")
		bot.Send(result, config.ChatID, Mailbox.GetAuthor(), Mailbox.GetDate(), Mailbox.GetBox(), Mailbox.GetBody())
	}
}
