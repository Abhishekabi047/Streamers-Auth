package repo

import (
	"errors"
	"log"
	models "service/pkg/models/user"
	"service/pkg/repo/interfaces"

	"gorm.io/gorm"
)

type UserRepo struct{
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo{
	return &UserRepo{
		DB: db,
	}
}

func (a *UserRepo)GetByEmail(email string) (*models.Channel,error) {
	var channel models.Channel
	res:=a.DB.Where(&models.Channel{Email: email}).First(&channel)
	if res.Error != nil{
		if errors.Is(res.Error,gorm.ErrRecordNotFound) {
			return nil,nil
		}
		return nil,res.Error
	}
	return &channel,nil

}

func (a *UserRepo) GetByPhone(phone string) (*models.Channel,error) {
	var channel models.Channel
	res:=a.DB.Where(&models.Channel{Phone: phone}).First(&channel)
	if res.Error != nil{
		if errors.Is(res.Error,gorm.ErrRecordNotFound) {
			return nil,nil
		}
		return nil,res.Error
	}
	return &channel,nil
}

func (a *UserRepo) GetByName(username string) (*models.Channel,error) {
	var channel models.Channel
	res:=a.DB.Where(&models.Channel{UserName: username}).First(&channel)
	if res.Error != nil{
		if errors.Is(res.Error,gorm.ErrRecordNotFound) {
			return nil,nil
		}
		return nil,res.Error
	}
	return &channel,nil
}

func (a *UserRepo)CreateChannel(channel *models.Channel) error{
	return a.DB.Create(channel).Error
}

func (a *UserRepo) CreateOtpKey(key,phone string) error{
	var otpKey models.OtpKey
	otpKey.Key=key
	otpKey.Phone=phone
	if err :=a.DB.Create(&otpKey).Error;err != nil{
		log.Print("error creating otp key",err)
		return err
	}
	return nil
}

func (a *UserRepo) CreateSignUp(user *models.Signup) error{
	return a.DB.Create(&user).Error 
}

func (a *UserRepo) GetByKey(key string) (*models.OtpKey,error) {
	var otpKey models.OtpKey
	result := a.DB.Where(&models.OtpKey{Key: key}).First(&otpKey)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &otpKey, nil
}

func (a *UserRepo) GetSignupByPhone(phone string) (*models.Signup,error){
	var user models.Signup
	result :=a.DB.Where(&models.Signup{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("error fetching data")
	}
	return &user, nil

}

func (a *UserRepo) Update(user *models.Channel) error{
	return a.DB.Updates(user).Error
}

func (a *UserRepo) CheckPermission(user *models.Channel) (bool,error) {
	result := a.DB.Where(&models.Channel{Phone: user.Phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.New("error in fetching block detail")
	}
	permission := user.Permission
	return permission, nil
}