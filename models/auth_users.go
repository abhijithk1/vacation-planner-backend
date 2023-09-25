package models

import (
	"errors"
	"html"
	"strings"

	"github.com/abhijithk1/vacation-planner/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error){

	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u,nil
}

func (u *User) BeforeSave(*gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func verifyPassword (hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var tokenString string
	u := User{}

	err = DB.Model(&u).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = verifyPassword(u.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	if (u.Username == "admin") {
		tokenString, err = token.GenerateToken(u.ID,u.Username)
	} else {
		tokenString, err = token.GenerateToken(u.ID)
	}

	if err != nil {
		return "",err
	}

	return tokenString,nil
}

func GetUserID(uid uint) (User, error) {
	
	var u User

	if err = DB.First(&u,uid).Error; err != nil {
		return u,errors.New("user not found")
	}

	return u, nil
}

func GetUsers() ([]User, error) {
	var u []User

	if err = DB.Find(&u).Error; err != nil {
		return u,errors.New("user not found")
	}

	return u, nil
}
