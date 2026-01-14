package main

import (
	"fmt"
	Config "gomod/internal/config"
	Logger "gomod/internal/entities"
	tgbot "gomod/internal/use-cases"
	imap "gomod/internal/use-cases/mailbox"
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
	if Mailbox.CheckMessage() {
		Mailbox.Fetch()
	}

	//Create class Bot
	bot := tgbot.Bot{}
	//Initialization Bot
	//CHATID TEST
	err = bot.Initialization(config.Token, config.ChatID)
	if err != nil {
		panic(err)
	}
	fmt.Println(Mailbox.GetFormaBody())
	for true {

	}

}
