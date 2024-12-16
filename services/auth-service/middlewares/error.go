package middlewares

import (
	"net/http"

	"github.com/hungrynoodlehead/photos/services/auth-service/utils"
)

type ErrorHandlingMiddleware struct {
	Config *utils.Config
}

func NewErrorHandlingMiddleware(config *utils.Config) *ErrorHandlingMiddleware {
	return &ErrorHandlingMiddleware{Config: config}
}

func (mw *ErrorHandlingMiddleware) ErrorHandler(next func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})
}
