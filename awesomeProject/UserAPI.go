package main

import (
	"github.com/gin-gonic/gin"
	u "ucontroller"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//1.Register
	router.POST("/Register", u.Register)              //帳號註冊
	router.POST("/Login", u.Login)                    //登入驗證
	router.DELETE("/DELETE/:action/:param", u.Delete) //帳號刪除
	router.PUT("/PUT", u.Update)                      //密碼變更
	router.GET("/User/:account", u.Get)               //使用者資訊查詢

	router.Run(":8080")
}
