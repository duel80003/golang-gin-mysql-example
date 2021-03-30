package main

import (
	_ "go-gin-mysql-example/database"

	"go-gin-mysql-example/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api")
	user.UserRegister(v1.Group("user"))
	r.Run(":8080")
}
