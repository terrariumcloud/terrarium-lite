package types

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/errgo.v2/errors"
)

type TerrariumAPIErrorHandler struct{}

func (t *TerrariumAPIErrorHandler) Write(rw http.ResponseWriter, err error, statusCode int) {
	var prefix string = ""
	switch statusCode {
	case http.StatusInternalServerError:
		prefix = InternalServerErrorPrefix
	case http.StatusBadRequest:
		prefix = BadRequestPrefix
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
