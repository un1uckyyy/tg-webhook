package main

import (
	tele "gopkg.in/telebot.v4"
	"os"
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
		return c.Send(c.Sender())
	})

	bot.Start()
}
