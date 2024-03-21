package models

import "gorm.io/gorm"

type OtpKey struct {
	Key   string `json:"key"`
	Phone string `json:"phone"`
}

type Channel struct {
	gorm.Model    `json:"-"`
	Id            int    `gorm:"primarykey" json:"id"`
	UserName      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Category      string `json:"Category"`
	DOB           string `json:"dob"`
	ProfilePicURL string `json:"profilepic"`
	Bio           string `json:"bio"`
	Is_Admin      bool   `json:"is_admin"`
	BannerURL     string `json:"banner"`
	Permission    bool   `json:"permission" gorm:"default:true"`
}

type Signup struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Category string `json:"Category"`
	DOB      string `json:"dob"`
}
