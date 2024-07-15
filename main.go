package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/JohnnyCPC/reservoir-sampling-go/sks"
	"github.com/JohnnyCPC/tw-lottery-linebot-go/analyze"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type LotteryCombinations struct {
	RepresentHex string `json:"representhex"`
	RepresentBin string `json:"representbin"`
	NGram        int    `json:"ngram"`
	Times        int    `json:"times"`
	Numbers      []int  `json:"numbers"`
}

var bot *linebot.Client
var result map[string]LotteryCombinations

func main() {

	// Open our jsonFile
	jsonFile, err1 := os.Open("./data/539_2007_2024_result.json")
	// if we os.Open returns an error then handle it
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Successfully Opened json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &result)

	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(addr, nil)

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

	cb, err := webhook.ParseRequest(os.Getenv("ChannelSecret"), r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range cb.Events {
		log.Printf("Got event %v", event)
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				var t, c, n int
				var mes, mes2 string
				var wf, is539, anac bool
				var sec, luck []int

				res := strings.Split(message.Text, ",")
				switch res[0] {
				case "分析539":
					anac = true
					is539 = true
				case "539", "今彩539":
					t = 39
					c = 5
					is539 = true
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
					if _, err = bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage("Please Enter Again!")).Do(); err != nil {
						log.Print(err)
					}
					return
				}

				if len(res) < 2 || is539 {
					n = 1
				} else {
					if n, err = strconv.Atoi(res[1]); err != nil {
						n = 1
					} else {
						if n > 5 {
							n = 5
						}
					}
				}

				a := make([]int, t)
				for i := range a {
					a[i] = i + 1
				}

				for j := 0; j < n; j++ {
					luck = sks.SelectKItems(a, len(a), c)
					sort.Ints(luck)
					mes += fmt.Sprint(luck) + ","
				}

				if wf {
					mes2 = "Second Section:" + fmt.Sprint(sks.SelectKItems(sec, 8, n))
				}

				if is539 {
					var target []int

					if anac {
						for h := 1; h < len(res); h++ {
							t, err4 := strconv.Atoi(res[h])
							if err4 != nil || len(res) > 6 {
								if _, err5 := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage("Please Enter Again!")).Do(); err5 != nil {
									log.Print(err5)
								}
								log.Print(err4)
								return
							}
							target = append(target, t)
						}

					} else {
						target = luck
					}

					inputdata := analyze.BuildInputData(target)

					for _, in := range inputdata {
						//fmt.Println(in)
						if val, ok := result[in]; ok {
							if val.NGram >= 3 {
								mes2 += fmt.Sprint(val.NGram) + "-Gram set:" + fmt.Sprint(val.Numbers) + ", Times: " + fmt.Sprint(val.Times) + "\n"
							}
							//fmt.Println(val)
						}
					}

				}

				if _, err = bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage("Lucky Number : "+mes+"\n"+mes2)).Do(); err != nil {
					log.Print(err)
				}

			}
		}
	}

}
