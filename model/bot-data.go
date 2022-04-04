package model

import (
	"time"
)

type UserKeyWord struct {
	UserId  string `json:"UserId"`
	KeyWord string `json:"KeyWord"`
}

type ImageData struct {
	UserId     string
	KeyWord    string
	ImageUrl   string
	CreateTime time.Time
}

type ImgurRespones struct {
	Data struct {
		Link string `json:"link"`
		Type string `json:"types"`
	}
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type ImageMap struct {
	KeyWord  string
	ImageUrl string
}
