package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/JohnnyCPC/reservoir-sampling-go/sks"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(addr, nil)

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				var t, c int
				var mes, mes2 string
				var wf bool
				var sec []int
				switch message.Text {
				case "539", "今彩539":
					t = 39
					c = 5
				case "威力彩":
					t = 38
					c = 6
					sec = []int{1, 2, 3, 4, 5, 6, 7, 8}
					wf = true
				case "大樂透":
					t = 49
					c = 6
				case "雙贏彩":
					t = 24
					c = 12
				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Please Enter Again!")).Do(); err != nil {
						log.Print(err)
					}
					return
				}
				a := make([]int, t)
				for i := range a {
					a[i] = i + 1
				}

				luck := sks.SelectKItems(a, len(a), c)
				sort.Ints(luck)
				mes = fmt.Sprint(luck)

				if wf {
					mes2 = "Second :" + fmt.Sprint(sec[luck[0]%8])
				}

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Lucky Number : "+mes+" "+mes2)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}

}
