package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-meme-bot/cache"
	"go-meme-bot/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func BotAddImageKeyWord(userId string, keyWord string) {
	data := &model.UserKeyWord{UserId: userId, KeyWord: keyWord}
	cache.GetRedisKeyWordTemp().Set(data.UserId, *data)
}

func CheckRedisData(key string) (*model.UserKeyWord, bool) {
	var exist bool
	data := cache.GetRedisKeyWordTemp().Get(key)

	if len(data.KeyWord) > 0 {
		exist = true
	}

	return data, exist
}

func UploadImgur(stream io.ReadCloser) *model.ImgurRespones {
	var buf = new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, _ := writer.CreateFormFile("image", "temp")
	io.Copy(part, stream)

	defer writer.Close()
	defer stream.Close()

	req, _ := http.NewRequest("POST", "https://api.imgur.com/3/image", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Client-ID "+os.Getenv("ImgurClientID"))

	cli := &http.Client{}
	res, err := cli.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	respons := new(model.ImgurRespones)

	json.Unmarshal(b, &respons)

	return respons
}

func SearchUrl(userId string) map[string]string {
	m := cache.GetRedisImageMap().Get(userId)
	if len(m) == 0 {
		fmt.Println("Get DB")
		m = GetImageData(userId)
		go cache.GetRedisImageMap().Set(userId, m)
	}

	return m
}

func GetKeyWordList(userId string) string {
	str := "MeMe清單: \n"
	m := SearchUrl(userId)
	for key := range m {
		str += "#" + key + "\n"
	}
	return str
}
