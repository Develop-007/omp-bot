package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/cinema/subdomain"
)

type CinemaSubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *subdomain.Service
}

func NewCinemaSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *CinemaSubdomainCommander {
	subdomainService := subdomain.NewService()

	return &CinemaSubdomainCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}
}

func (c *CinemaSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CinemaSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CinemaSubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
