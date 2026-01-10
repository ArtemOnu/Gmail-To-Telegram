package bot

import (
	Logger "gomod/internal/entities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
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
