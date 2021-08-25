package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) (string, error) {
	fmt.Println("메시지 입력")
	return "",nil
}
func UpdateMessage(c *gin.Context) (string, error) {
	fmt.Println("메시지 수정")
	return "",nil
}
func DeleteMessage(c *gin.Context) (string, error) {
	fmt.Println("메시지 삭제")
	return "",nil
}