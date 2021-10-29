package types

import "net/http"

type APIResponseWriter interface {
	Write(rw http.ResponseWriter, data interface{}, statusCode int)
	Redirect(rw http.ResponseWriter, r *http.Request, uri string)
}

type APIErrorWriter interface {
	Write(rw http.ResponseWriter, err error, statusCode int)
}
