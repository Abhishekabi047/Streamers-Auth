package service

import (
	"context"
	"errors"
	"fmt"
	models "service/pkg/models/user"
	"service/pkg/pb/auth"
	"service/pkg/repo/interfaces"
	"service/pkg/utils"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	Repo interfaces.UserRepo
	auth.UnimplementedAuthServiceServer
}

func NewUserServer(repo interfaces.UserRepo) auth.AuthServiceServer {
	return &UserServer{
		Repo: repo,
	}
}

func (s *UserServer) Signup(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	var user models.Signup
	copier.Copy(&user, &req)
	if err := utils.Validate(&user); err != nil {
		return nil, err
	}
	var otpKey models.OtpKey
	email, _ := s.Repo.GetByEmail(req.Email)

	if email != nil {
		return nil, errors.New("user with this email already exists")
	}
	phone, err := s.Repo.GetByPhone(req.Phone)
	if err != nil {
		return nil, errors.New("error with server2")
	}
	if phone != nil {
		return nil, errors.New("user with this phone already exists")
	}
	username, err := s.Repo.GetByName(req.Username)
	if err != nil {
		return nil, errors.New("error with server3")
	}
	if username != nil {
		return nil, errors.New("user with this phone already exists")
	}
	age := utils.CalculateAge(req.Dob)
	if age < 18 {
		return nil, errors.New("you must be 18 to SignUp")
	}
	err1 := utils.ValidatePassword(req.Password)
	if err1 != nil {
		return nil, err1
	}
	if req.Password != req.Cpassword {
		return nil, errors.New("Confirm Password error")
	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashpass)
	key, err := utils.SendOtp(req.Phone)
	if err != nil {
		return nil, err
	} else {
		var user models.Signup
		copier.Copy(&user, &req)
		err = s.Repo.CreateSignUp(&user)
		otpKey.Key = key
		otpKey.Phone = req.Phone
		err = s.Repo.CreateOtpKey(key, req.Phone)
		if err != nil {
			return nil, err
		}
		return &auth.SignUpResponse{
			Key: key,
		}, nil
	}

}

func (a *UserServer) Otp(ctx context.Context, req *auth.OtpRequest) (*auth.OtpResponse, error) {
	res, err := a.Repo.GetByKey(req.Key)
	if err != nil {
		return nil, errors.New("error fetching key")
	}
	user, err := a.Repo.GetSignupByPhone(res.Phone)
	if err != nil {
		return nil, errors.New("error fetching phone")
	}
	err = utils.CheckOtp(res.Phone, req.Otp)
	if err != nil {
		return nil, err
	} else {
		newUser := &models.Channel{
			UserName: user.UserName,
			Email:    user.Email,
			Password: user.Password,
			DOB:      user.DOB,
			Category: user.Category,
			Phone:    user.Phone,
		}
		err1 := a.Repo.CreateChannel(newUser)
		if err1 != nil {
			return nil, errors.New("error while creating channel")
		}
	}
	return &auth.OtpResponse{
		Message: "user Sign up succesfull",
	}, nil
}

func (a *UserServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := a.Repo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user with this email not found")
	}
	permission, err := a.Repo.CheckPermission(user)
	if err != nil {
		return nil, err
	}
	if permission == false {
		return nil, errors.New("permission denied")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid password")
	} else {
		return &auth.LoginResponse{
			Id:       int64(user.Id),
			Username: user.UserName,
		}, nil
	}
}

func (a *UserServer) SearchUser(ctx context.Context, req *auth.SearchUserRequest) (*auth.SearchUserResponse, error) {
	list, err := a.Repo.SearchUserName(req.Username, int(req.Limit), int(req.Offset))
	fmt.Println("user",list)
	if err != nil {
		return nil, errors.New("username not found")
	}
	var fulllist []*auth.User

	for _, v := range list {
		fulllist = append(fulllist, &auth.User{
			Username:   v.UserName,
			Profilepic: v.ProfilePicURL,
		})
	}

	return &auth.SearchUserResponse{
		Userdetails: fulllist,
	}, nil
}

func (a *UserServer) UserExists(ctx context.Context, req *auth.UserExistsRequest) (*auth.UserExistsResponse, error) {
	res, err := a.Repo.UserExistsByUsername(req.Username)
	if err != nil {
		return nil, errors.New("error while checking user")
	}
	return &auth.UserExistsResponse{
		Username: res,
	}, nil
}
