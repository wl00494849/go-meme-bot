package main

import (
	"flag"
	"go-meme-bot/botsetting"
	"go-meme-bot/cache"
	"go-meme-bot/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	var port string
	flag.StringVar(&port, "p", ":6666", "port")

	if p := os.Getenv("PORT"); len(p) != 0 {
		port = ":" + p
	}

	go botsetting.BotInit()
	go cache.RedisInit()
	go database.DBInit()

	app := gin.Default()
	app.POST("/memebot", botsetting.MemeBotControl)
	app.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"test": "success"})
	})

	err := app.Run()

	if err != nil {
		panic(err)
	}
}
