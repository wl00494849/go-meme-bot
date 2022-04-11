package server

import (
	"go-meme-bot/database"
	"go-meme-bot/model"
	"time"
)

func InsertImageMessage(user *model.UserKeyWord, image *model.ImgurRespones) error {
	conn := database.GetDB()
	stmt, _ := conn.Prepare("Insert ImageMessage set userId=?,keyWord=?,imageUrl=?,createTime=?")
	_, err := stmt.Exec(user.UserId, user.KeyWord, image.Data.Link, time.Now())
	return err
}

func GetImageData(userId string) map[string]string {
	m := make(map[string]string)
	conn := database.GetDB()
	rows, _ := conn.Query("Select keyWord ,imageUrl from ImageMessage where userId=?", userId)

	defer rows.Close()

	for rows.Next() {
		var imgMap model.ImageMap
		rows.Scan(&imgMap.KeyWord, &imgMap.ImageUrl)
		m[imgMap.KeyWord] = imgMap.ImageUrl
	}

	return m
}

func DelImageData(userId string, keyWord string) error {
	_, err := database.GetDB().Exec("delete from ImageMessage where userId=? and keyWord=?", userId, keyWord)
	return err
}
