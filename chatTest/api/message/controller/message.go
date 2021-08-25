package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context){
	fmt.Println("메시지 입력")
}
func UpdateMessage(c *gin.Context) {
	fmt.Println("메시지 수정")
}
func DeleteMessage(c *gin.Context)  {
	fmt.Println("메시지 삭제")
}
func GetMessageList(c *gin.Context)  {
	fmt.Println("메시지 조회")
}