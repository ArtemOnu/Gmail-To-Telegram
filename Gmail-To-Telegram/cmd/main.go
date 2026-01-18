package main

import (
	Config "gomod/internal/config"
	Logger "gomod/internal/entities"
	tgbot "gomod/internal/use-cases"
	imap "gomod/internal/use-cases/mailbox"
	"time"
)

func main() {

	Logger.LoggerIni()
	Logger.Log("application start")

	//We get the configuration structure from the file
	var config Config.Config
	err := Config.GetConfig(&config, "../../config.json")
	if err != nil {
		Logger.Log(err.Error())
		panic(err)
	}
	//Create class MailBox
	Logger.Log("Mail create")
	Mailbox := imap.Mail{}
	Mailbox.Connect(config.Mail, config.Password, config.Host)
	Logger.Log("Mail connect")
	defer Mailbox.Disconnect()
	results := Mailbox.MailUpdate()

	//Create class Bot
	bot := tgbot.Bot{}
	//Initialization Bot
	err = bot.Initialization(config.Token, config.ChatID)
	if err != nil {
		panic(err)
	}

	go func() {
		for result := range results {
			time.Sleep(time.Second * 2)
			Logger.Log("new message")
			bot.Send(result, config.ChatID,
				Mailbox.GetAuthor(),
				Mailbox.GetDate(),
				Mailbox.GetBox(),
				Mailbox.GetBody())
		}
	}()

	select {}
}
