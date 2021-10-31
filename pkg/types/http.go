package types

import "net/http"

// APIResponseWriter is a generic response interface that allows for standardisation of API responses returned by Terrarium
type APIResponseWriter interface {
	Write(rw http.ResponseWriter, data interface{}, statusCode int)
	Redirect(rw http.ResponseWriter, r *http.Request, uri string)
	WriteRaw(rw http.ResponseWriter, data interface{}, statusCode int)
}

// APIErrorWriter is a generic response interface that allows for standardisation of API errors returned by Terrarium
type APIErrorWriter interface {
	Write(rw http.ResponseWriter, err error, statusCode int)
}
