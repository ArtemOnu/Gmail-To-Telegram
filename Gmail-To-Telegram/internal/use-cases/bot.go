package bot

import (
	"fmt"
	Logger "gomod/internal/entities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func (b *Bot) GetFile(Html []byte) tgbotapi.FileBytes {
	fileByte := tgbotapi.FileBytes{
		Name:  "mail.html",
		Bytes: Html,
	}
	return fileByte
}

func (b *Bot) Send(Text string,
	ChatID int64,
	Sender string,
	Data string,
	To string,
	Bytes []byte) {

	Logger.Log("send message")
	text := fmt.Sprintf(`📂 <b>Детали сообщения:</b>
• <b>Отправитель:</b> %s
• <b>Получатель:</b> %s
• <b>Дата отправки:</b> %s


<blockquote><b>Отправленный текст:</b>
%s </blockquote>`, Sender, To, Data, Text)
	message := tgbotapi.NewMessage(ChatID, text)
	message.ParseMode = "HTML"

	filemessage := tgbotapi.NewDocument(ChatID, b.GetFile(Bytes))
	b.bot.Send(message)
	b.bot.Send(filemessage)
}

// Иницилизация бота + запуск слушателя событий
// CHATID TEST
func (b *Bot) Initialization(Token string, ChatID int64) error {
	//Иницилизация бота
	var err error
	Logger.Log("Авторизация бота")
	b.bot, err = tgbotapi.NewBotAPI(Token)
	if err != nil {
		Logger.Log("Ошибка создания бота 16 bot.go")
		return err
	}
	//Запуск слушателя в фоне
	//TEST
	go b.Updates(ChatID)
	return nil
}

// TEST
func (b *Bot) Updates(ChatID int64) {
	Logger.Log("запуск апдейтов")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.ChannelPost != nil {
			Logger.Log("сообщение")
			msg := tgbotapi.NewMessage(ChatID, "test")
			b.bot.Send(msg)
		}
	}
}
