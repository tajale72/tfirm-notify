package main

import (
	"github.com/gin-gonic/gin"

	notify "tfirm-notify/internal"
)

func main() {
	r := gin.Default()
	r.POST("/notify", notify.Notify)
	r.Run()
}
