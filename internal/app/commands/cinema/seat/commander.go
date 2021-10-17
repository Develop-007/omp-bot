package seat

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/cinema/seat"
)

type CinemaSeatCommander struct {
	bot         *tgbotapi.BotAPI
	seatService seat.SeatService
}

func NewCinemaSeatCommander(
	bot *tgbotapi.BotAPI,
) *CinemaSeatCommander {
	seatService := seat.NewService()

	return &CinemaSeatCommander{
		bot:         bot,
		seatService: seatService,
	}
}

func (c *CinemaSeatCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CinemaSeatCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CinemaSeatCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
