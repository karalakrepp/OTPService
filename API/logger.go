package api

import (
	"time"

	"github.com/sirupsen/logrus"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) *LoggingService {
	return &LoggingService{
		next: next,
	}
}

func (l *LoggingService) twilioSendOTP(phoneNumber string) (code string, err error) {

	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{

			"took":   time.Since(start),
			"err":    err,
			"number": phoneNumber,
			"code":   code,
		}).Info("otpservice")

	}(time.Now())

	return l.next.twilioSendOTP(phoneNumber)
}

func (l *LoggingService) twilioVerifyOTP(phoneNumber string, code string) (err error) {
	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{

			"took":   time.Since(start),
			"err":    err,
			"number": phoneNumber,
			"code":   code,
		}).Info("otpservice")

	}(time.Now())

	return l.next.twilioVerifyOTP(phoneNumber, code)
}
