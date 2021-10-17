package seat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpMessage = "/help__cinema__seat — print list of commands\n" +
	"/get__cinema__seat — get a entity\n" +
	"/list__cinema__seat — get a list of your entity\n" +
	"/delete__cinema__seat — delete an existing entity\n" +
	"/new__cinema__seat — create a new entity\n" +
	"/edit__cinema__seat — edit a entity"

func (c *CinemaSeatCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpMessage)
	c.bot.Send(msg)
}
