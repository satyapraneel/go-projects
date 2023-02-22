package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-asdas")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1asdasd")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvent(bot.CommandEvents())

	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"My yob is 2023"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println(err)
				return
			}
			age := 2023 - yob
			r := fmt.Sprintf("Your age is %d\n", age)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvent(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}
