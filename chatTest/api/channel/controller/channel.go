package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetChannel(c *gin.Context) {
	fmt.Println("채널 입장")
}
func CreateChannel(c *gin.Context){
	fmt.Println("채널 생성")
}
func GetAllChannelList(c *gin.Context) {
	fmt.Println("모든 채널 리스트")
}
func GetMyChannelList(c *gin.Context) {
	fmt.Println("나의 채널 리스트")
}