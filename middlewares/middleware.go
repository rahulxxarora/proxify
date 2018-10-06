package middlewares

import "net/http"

type httpHandlerFunc func(http.ResponseWriter, *http.Request)
