package cinema

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/seat"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CinemaCommander struct {
	bot           *tgbotapi.BotAPI
	seatCommander Commander
}

func NewCinemaCommander(
	bot *tgbotapi.BotAPI,
) *CinemaCommander {
	return &CinemaCommander{
		bot: bot,
		// seatCommander
		seatCommander: seat.NewCinemaSeatCommander(bot),
	}
}

func (c *CinemaCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Seat {
	case "seat":
		c.seatCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CinemaCommander.HandleCallback: unknown seat - %s", callbackPath.Seat)
	}
}

func (c *CinemaCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Seat {
	case "seat":
		c.seatCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CinemaCommander.HandleCommand: unknown seat - %s", commandPath.Seat)
	}
}
