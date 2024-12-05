package main

import (
	"fmt"
	"os"

	tele "gopkg.in/telebot.v4"
)

func main() {
	token := os.Getenv("TOKEN")
	webhookUrl := os.Getenv("WEBHOOK_URL")

	pref := tele.Settings{
		Token: token,
		Poller: &tele.Webhook{
			Listen: ":8080",
			Endpoint: &tele.WebhookEndpoint{
				PublicURL: webhookUrl,
			},
		},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	bot.Handle(tele.OnText, func(c tele.Context) error {
		msg := fmt.Sprintf("echo from: %v", c.Sender())
		return c.Send(msg)
	})

	bot.Start()
}
