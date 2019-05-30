package main

import (
	"ant-chat/cmd/ws"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()

	// index
	app.StaticFile("/chat", "./web/index.html")
	// websocket
	//app.StaticFile("/ws", "./web/ws.html")
	app.GET("/ws/socket", ws.Socket())

	// static files
	app.Static("/static", "./static")

	log.Fatal(app.Run())
}
