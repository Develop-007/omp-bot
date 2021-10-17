package seat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/cinema/seat"
)

func (c *CinemaSeatCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	c.seatService.Add(seat.Seat{Title: args})
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Success add seat",
	)

	c.bot.Send(msg)
}
