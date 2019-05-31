package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Auth struct {
	ID int `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}


func CheckLogin(username, password string) (bool , error) {
	//log.Println(username, password)

	var auth Auth
	err := db.Debug().Where(&Auth{Username:username, Password:password}).First(&auth).Error
	//db.Select("id").Where(Auth{Username:username, Password:password}).First(&auth)
	log.Println(auth)
	if(err !=nil && err != gorm.ErrRecordNotFound) {
		return false, err
	}

	if auth.ID >0 {
		return true, nil
	}

	return false, nil

}
