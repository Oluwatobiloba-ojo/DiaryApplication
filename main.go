package main

import (
	"DiaryApp/src/dto"
	"DiaryApp/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var diaryService service.DiaryService = service.NewDiaryServiceImpl()

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/add", func(context *gin.Context) {
		var object dto.RegisterRequest
		err := context.BindJSON(&object)
		if err != nil {
			return
		}
		diary, err2 := diaryService.CreateDiary(object.Username, object.Password)
		fmt.Println(object)
		if err2 == nil {
			context.JSON(http.StatusOK, gin.H{
				"id":      diary.Id(),
				"message": "Account Created Successfully",
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err2.Error(),
			})
		}
	})

	r.Run()
}
