package botsetting

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func BotInit() {

	b, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)

	if err != nil {
		panic(err)
	}

	bot = b
}

func MemeBotControl(ctx *gin.Context) {

	events, err := bot.ParseRequest(ctx.Request)

	if err == linebot.ErrInvalidSignature {
		ctx.JSON(400, "")
	} else {
		ctx.JSON(500, "")
	}

	for _, evn := range events {
		if evn.Type == linebot.EventTypeMessage {
			switch msg := evn.Message.(type) {
			case *linebot.TextMessage:
				switch msg.Text {
				case "MeMe清單":
					memewListReply(evn, msg)
				case "MeMe說明":
					memeIllustrate(evn)
				default:
					if msg.Text[:1] == "#" {
						getMeMeImage(evn, msg)
					} else if msg.Text[:13] == "MeMe新增：" {
						setMeMeKeyWordTemp(evn, msg)
					} else if msg.Text[:13] == "MeMe刪除：" {
						delImageData(evn, msg)
					}
				}
			case *linebot.ImageMessage:
				memeImageUpload(evn, msg)
			}
		}
	}
}
