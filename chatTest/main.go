package main

import (
	c1 "chatTest/api/channel/controller"
	c2 "chatTest/api/message/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}



	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", func(c *gin.Context){
		c.HTML(http.StatusOK,"login.html",nil) //로그인 처리
	})
	r.POST("/api/channel", c1.CreateChannel)//채널 생성
	r.GET("/api/channel", c1.GetAllChannelList)//모든 채널 목록 조회 및 채널 생성 폼
	r.GET("/api/channel/:user_id", c1.GetMyChannelList)//내가 참여한 채널 목록 조회
	r.GET("/api/channel/:user_id/:chan_no",c1.GetChannel)//채널 입장
	r.GET("/api/channel/:user_id/:chan_no/msg",c2.GetMessageList)//채널 입장 후 메시지 리스트 조회
	r.POST("/api/channel/:user_id/:chan_no/msg",c2.CreateMessage)//메시지 전송
	r.PUT("/api/channel/:user_id/:chan_no/msg/:msg_no",c2.UpdateMessage)//메시지 수정
	r.DELETE("/api/channel/:user_id/:chan_no/msg/:msg_no",c2.DeleteMessage)//메시지 삭제

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}


