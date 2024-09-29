package main

import (
  "io"
  "os"
  "github.com/line/line-bot-sdk-go/v7/linebot"
  "net/http"
  "log"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)


func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // create new linebot client
  bot, err := linebot.New(
	  os.Getenv("LINE_BOT_CHANNEL_SECRET"),
	  os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
  )
  if err != nil {
    log.Fatal(err)
  }
  // get events from request
  events, err := bot.ParseRequest(r)
  if err != nil {
    log.Fatal(err)
  }
  // handle each event
  for _, event := range events {
    switch event.Type {
    case linebot.EventTypeMessage:
      // handle message event
      switch message := event.Message.(type) {
      case *linebot.TextMessage:
        // reply text message
        if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
          log.Fatal(err)
        }
      }
    }
  }
  io.WriteString(w, "completed")
	})

  lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
