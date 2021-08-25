package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetChannel(c *gin.Context) (string, error) {
	fmt.Println("채널 입장")
	return "",nil
}
func CreateChannel(c *gin.Context) (string, error) {
	fmt.Println("채널 생성")
	return "",nil
}
func GetAllChannelList(c *gin.Context) (string, error) {
	fmt.Println("모든 채널 리스트")
	return "",nil
}
func GetMyChannelList(c *gin.Context) (string, error) {
	fmt.Println("나의 채널 리스트")
	return "",nil
}