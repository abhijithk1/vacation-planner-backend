package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Vacation struct {
	gorm.Model
	UserId      uint      `gorm:"not null"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:255;not null"`
	FromDate    time.Time `gorm:"size:255;not null"`
	EndDate     time.Time `gorm:"size:255;not null"`
}

type ListVacation struct {
	Username    string
	Title       string
	Description string
	FromDate    time.Time
	EndDate     time.Time
}

func (v *Vacation) AddVacation() (*Vacation, error) {
	err = DB.Create(&v).Error
	if err != nil {
		return &Vacation{}, err
	}

	return v, nil
}

func ListUpcomingVacationsID(uid uint) ([]Vacation, error) {
	var v []Vacation
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	fmt.Println(now)
	fmt.Println(uid)
	if err = DB.Where("user_id = ? AND from_date > ?", uid, now).Find(&v).Error; err != nil {
		return v, errors.New("user not found")
	}

	return v, nil
}

func ListTeamUpcomingVacations(uid uint) ([]Vacation, error) {
	var v []Vacation
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	fmt.Println(now)
	if err = DB.Where("user_id != ? AND from_date > ?", uid, now).Find(&v).Error; err != nil {
		return v, errors.New("user not found")
	}

	return v, nil
}

func ListTeamMember(v []Vacation) ([]ListVacation,error){
	var data []ListVacation
	for _, vacation := range v {
		var d ListVacation
		vacation_id := vacation.UserId
		d.Title = vacation.Title
		d.Description = vacation.Description
		d.FromDate = vacation.FromDate
		d.EndDate = vacation.EndDate
		u, err := GetUserID(vacation_id)
		if err != nil {
			return nil,err
		}
		d.Username = u.Username
		data = append(data, d)
	}
	return data, nil
}
