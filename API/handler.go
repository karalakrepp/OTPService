package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karalakrepp/OTPService/data"
)

// Service is the Interface that contains any service function
// router is a gin router
type Config struct {
	router  *gin.Engine
	service Service
}

var (
	appTimeout = 7 * time.Second
)

func NewConfig(router *gin.Engine, service Service) *Config {

	return &Config{
		router:  router,
		service: service,
	}
}

// POST handler for SMS Sending. this func include service func logging or service
func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		//validate json data is correct
		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		//this is the service func
		_, err := app.service.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
	}
}

// POST handler for SMS Sending. this func include service func logging or service

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		//validate json data is correct
		app.validateBody(c, &payload)

		//define new data
		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}
		//use service func
		err := app.service.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ", err)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		//encoding
		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}
