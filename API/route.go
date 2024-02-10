package api

// modify below
func (app *Config) Routes() {
	app.router.POST("/otp", app.sendSMS())
	app.router.POST("/verifyOTP", app.verifySMS())
}
