package seat

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CinemaSeatCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products := c.seatService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "cinema",
		Seat:         "seat",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.bot.Send(msg)
}