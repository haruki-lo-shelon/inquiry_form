package main

import (
    "go-send-contact/handler"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/api/v1/contact", handler.SendMail)
    r.Run()
}
