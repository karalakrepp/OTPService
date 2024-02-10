package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envAccount("ACCOUNT_SID"),
	Password: envAccount("AUTHTOKEN"),
})

type Service interface {
	OTPSender
}
type OTPSender interface {
	twilioSendOTP(string) (string, error)
	twilioVerifyOTP(string, string) error
}

type SMSService struct {
}

func (s *SMSService) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envAccount("SERVICESID"), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
func (app *SMSService) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envAccount("SERVICESID"), params)
	if err != nil {
		return err
	}

	// BREAKING CHANGE IN THE VERIFY API
	// https://www.twilio.com/docs/verify/quickstarts/verify-totp-change-in-api-response-when-authpayload-is-incorrect
	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}
