package types

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/errgo.v2/errors"
)

const BadRequestPrefix string = "Bad Request"
const InternalServerErrorPrefix string = "Internal Server Error"
const NotFoundPrefix string = "404 Not Found"
const UnprocessablePrefix string = "Unprocessable Entity"

type APIErrorWriter interface {
	Write(rw http.ResponseWriter, err error, statusCode int)
}

type TerrariumAPIErrorHandler struct{}

func (t *TerrariumAPIErrorHandler) Write(rw http.ResponseWriter, err error, statusCode int) {
	var prefix string = ""
	switch statusCode {
	case http.StatusInternalServerError:
		prefix = InternalServerErrorPrefix
	case http.StatusBadRequest:
		prefix = BadRequestPrefix
	case http.StatusNotFound:
		prefix = NotFoundPrefix
	case http.StatusUnprocessableEntity:
		prefix = UnprocessablePrefix
	default:

	}
	resp := &TerrariumServerResponse{
		Code:    statusCode,
		Message: fmt.Sprintf("%s - %s", prefix, err.Error()),
	}
	jsonData, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("+%v", errors.Wrap(err))
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	rw.Write(jsonData)
}
