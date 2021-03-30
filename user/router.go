package user

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", FetchAllUsers)
	router.POST("/", SaveUser)
	router.PUT("/:userID", UpdateUser)
	router.GET("/:userID", GetUser)
	router.DELETE("/:userID", DeleteUser)
}

func FetchAllUsers(c *gin.Context) {
	users, err := FetchAll()
	if err != nil {
		log.Fatalln("GetAllUsers error", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
		return
	}
	user, err := Get(uint(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func SaveUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Insert(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := Update(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": result})
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
	}
	err = Delete(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
