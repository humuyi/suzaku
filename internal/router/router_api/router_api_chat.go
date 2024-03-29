package router_api

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_chat"
)

func chat(group *gin.RouterGroup) {
	router := group.Group("chat")
	router.POST("send_msg", api_chat.SendMessage)
}
