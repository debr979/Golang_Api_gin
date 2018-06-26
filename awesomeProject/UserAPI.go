package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	u "ucontroller"
)



func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//1.Register
	router.POST("/Register", u.Register)
	router.POST("/Login",u.Login)
	router.DELETE("/DELETE/:action/:param", u.Delete)
    router.PUT("/PUT",u.Update)
	router.GET("/:Account",u.Get)
	router.Run(":8080")
}



