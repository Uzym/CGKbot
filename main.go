package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"2bot/app/bot"
	"2bot/app/db"
	"2bot/app/httpI"
	"2bot/app/telegram"
	"2bot/config"
)

func main() {

	fmt.Println("Start prog")

	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	db.DSN = fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", cfg.DB.User, cfg.DB.Name, cfg.DB.Password)
	httpI.Port = ":50051"
	telegram.TelegramAPI = cfg.ApiKeys.Telegram
	bot.QuestionTime = 10 * time.Second

	go telegram.StartTBot()
	go httpI.StartHTTP()
	bot.StartBot()

}
