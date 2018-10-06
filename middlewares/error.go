package middlewares

import (
	"errors"
	"fmt"
	"net/http"
)

func ErrorHandler(next httpHandlerFunc) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Something went wrong")
				}

				fmt.Println(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}
