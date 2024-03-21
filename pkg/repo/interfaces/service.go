package interfaces

import models "service/pkg/models/user"

type UserRepo interface {
	CreateChannel(*models.Channel) error
	CreateOtpKey(string, string) error
	CreateSignUp(*models.Signup) error
	GetByEmail(string) (*models.Channel, error)
	GetByKey(string) (*models.OtpKey, error)
	GetByName(string) (*models.Channel, error)
	GetByPhone(string) (*models.Channel, error)
	GetSignupByPhone(string) (*models.Signup, error)
	Update(*models.Channel) error
	CheckPermission(*models.Channel) (bool,error)
}
