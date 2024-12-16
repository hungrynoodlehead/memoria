package handlers

import (
	"errors"
	"net/http"
)

func ErrorHandlerAdapter(next func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)

		var handledError HandlerErr

		if err != nil {
			if errors.As(err, &handledError) {
				http.Error(w, handledError.Error(), handledError.Code())
			} else {
				// TODO: LOGGER
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}
	}
}
