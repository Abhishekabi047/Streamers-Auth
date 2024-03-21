package utils

import (
	"errors"
	"fmt"
	"service/pkg/config"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTH_TOKEN  string
	VERIFY_SERVICE_SID string
	client             *twilio.RestClient
)


func Init() {
	config, _ := config.LoadConfig()
	TWILIO_ACCOUNT_SID = config.Key1
	TWILIO_AUTH_TOKEN = config.Key2
	VERIFY_SERVICE_SID = config.Key3
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
}

func SendOtp(phone string) (string, error) {
	cfg, _ := config.LoadConfig()
	TWILIO_ACCOUNT_SID = cfg.Key1
	TWILIO_AUTH_TOKEN = cfg.Key2
	VERIFY_SERVICE_SID = cfg.Key3
	

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

	to := "+91" + phone
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	fmt.Println("service",params)
	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("Otp failed to generate")
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		return *resp.Sid, nil
	}
}

func CheckOtp(phone, code string) error {
	cfg, _ := config.LoadConfig()
	TWILIO_ACCOUNT_SID = cfg.Key1
	TWILIO_AUTH_TOKEN = cfg.Key2
	VERIFY_SERVICE_SID = cfg.Key3
	

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
	to := "+91" + phone
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Invalid otp")
	} else if *resp.Status == "approved" {
		return nil
	} else {
		return errors.New("Invalid otp")
	}
}
