package user

import (
	"go-gin-mysql-example/database"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Account   string `json:"Account"`
}

func FetchAll() ([]User, error) {
	log.Printf("fetch all users \n")
	users := []User{}
	result := database.DB.Find(&users)
	return users, result.Error
}

func Insert(user *User) error {
	log.Printf("save user: %+v \n", user)
	result := database.DB.Create(&user)
	return result.Error
}

func Get(userId uint) (*User, error) {
	log.Printf("fetch user by userId %d \n", userId)
	var user User
	result := database.DB.First(&user, userId)
	return &user, result.Error
}

func Update(user *User) (*User, error) {
	log.Printf("update user %+v \n", user)
	result := database.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	u, err := Get(user.ID)
	return u, err
}

func Delete(userID uint) error {
	log.Printf("delete user %d \n", userID)
	result := database.DB.Delete(&User{}, userID)
	return result.Error
}
