package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/karalakrepp/OTPService/API"
)

func main() {
	router := gin.Default()

	sms := api.SMSService{}
	svc := api.NewLoggingService(&sms)

	config := api.NewConfig(router, svc)

	config.Routes()

	router.Run(":8000")

}
