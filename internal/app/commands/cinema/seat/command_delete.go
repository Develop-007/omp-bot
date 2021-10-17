package seat

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaSeatCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	c.seatService.Remove(idx)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Success deleted",
	)

	c.bot.Send(msg)
}
