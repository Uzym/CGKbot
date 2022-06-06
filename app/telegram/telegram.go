package telegram

import (
	mybot "2bot/app/bot"
	"2bot/app/db"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var TelegramAPI = ""
var answers map[int64]string

func StartTBot() {
	// Create bot
	bot, err := tgbotapi.NewBotAPI(TelegramAPI)
	if err != nil {
		log.Fatal(err)
	}

	answers = make(map[int64]string)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.Text[0] != '/' {
			if answers[int64(update.Message.Chat.ID)] == "/nil" {
				answers[int64(update.Message.Chat.ID)] = update.Message.Text
			}
		} else {
			if answers[int64(update.Message.Chat.ID)] != "/nil" {
				go checkUpdate(bot, update)
			}
		}
	}

}

func checkUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	reply := ""

	switch update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"/question получить вопрос\n/leaders вывести таблицу лидеров\nВведите имя под которым будете записаны\n")
		bot.Send(msg)

		answers[update.Message.Chat.ID] = "/nil"
		time.Sleep(mybot.QuestionTime)

		userAns := answers[int64(update.Message.Chat.ID)]

		err := db.SetNewUser(update.Message.Chat.ID, userAns)
		if err != nil {
			log.Fatal(err)
		}

		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Готово!")
		bot.Send(msg)

	case "question":
		question, err := mybot.Question(update.Message.Chat.ID)
		if err != nil {
			log.Fatal(err)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, question)
		bot.Send(msg)

		answers[update.Message.Chat.ID] = "/nil"
		time.Sleep(mybot.QuestionTime - time.Second)

		userAns := answers[int64(update.Message.Chat.ID)]

		ok, err := mybot.Answer(update.Message.Chat.ID, userAns, time.Now())
		if err != nil {
			log.Fatal(err)
		}

		if ok {
			reply = "Ответ верный!"
		} else {
			reply = "Ответ неверный!"
		}
		delete(answers, int64(update.Message.Chat.ID))

	case "leaders":
		leaders, err := mybot.Leaderboard()
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(leaders); i++ {
			leader := strings.Split(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintln(leaders[i]), "{", ""), "}", ""), " ")
			reply = reply + leader[0] + ": " + leader[1]
		}
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

	bot.Send(msg)
}
