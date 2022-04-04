package botsetting

import (
	"go-meme-bot/cache"
	"go-meme-bot/server"

	"github.com/line/line-bot-sdk-go/linebot"
)

func memewListReply(evn *linebot.Event, msg *linebot.TextMessage) {
	str := server.GetKeyWordList(evn.Source.UserID)
	bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage(str)).Do()
}

func getMeMeImage(evn *linebot.Event, msg *linebot.TextMessage) {
	m := server.SearchUrl(evn.Source.UserID)
	url := m[msg.Text[1:]]
	if len(url) > 0 {
		bot.ReplyMessage(evn.ReplyToken, linebot.NewImageMessage(url, url)).Do()
	} else {
		bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage("找不到圖片")).Do()
	}
}

func setMeMeKeyWordTemp(evn *linebot.Event, msg *linebot.TextMessage) {
	go server.BotAddImageKeyWord(evn.Source.UserID, msg.Text[13:])
	bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage("關鍵字 : "+msg.Text[13:]+"; 請在一分鐘內上傳圖片")).Do()
}

func memeImageUpload(evn *linebot.Event, msg *linebot.ImageMessage) {
	res, _ := bot.GetMessageContent(msg.ID).Do()
	data, exist := server.CheckRedisData(evn.Source.UserID)

	if exist {
		go cache.GetRedisKeyWordTemp().Del(data.UserId)
		go cache.GetRedisImageMap().Del(data.UserId)
		imgurRes := server.UploadImgur(res.Content)
		err := server.InsertImageMessage(data, imgurRes)

		if err == nil {
			bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage("上傳成功! 輸入關鍵字#"+data.KeyWord+",即可使用")).Do()
		}
	}
}

func delImageData(evn *linebot.Event, msg *linebot.TextMessage) {
	err := server.DelImageData(evn.Source.UserID, msg.Text[13:])
	if err != nil {
		bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage("刪除失敗")).Do()
		return
	}
	go cache.GetRedisImageMap().Del(evn.Source.UserID)
	bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage("刪除成功")).Do()
}

func memeIllustrate(evn *linebot.Event) {
	reply := "指令說明 \n新增MeMe圖片: MeMe新增+分號(全形)+關鍵字 \n\nex: MeMe新增：Doge\n\n使用MeMe圖片: #+關鍵字\n\nex: #Doge\n\n刪除MeMe圖片: MeMe刪除+分號(全形)+關鍵字\n\nex:MeMe刪除：Doge\n\n查看圖庫關鍵字: MeMe清單"
	bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage(reply)).Do()
}
