package errorhandler

import "github.com/hungrynoodlehead/photos/services/auth-service/utils"

type ErrorHandler struct {
	Config *utils.Config
}

func NewErrorHandler(config *utils.Config) *ErrorHandler {
	return &ErrorHandler{Config: config}
}