package types

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/errgo.v2/errors"
)

type TerrariumAPIResponseWriter struct{}

func (t *TerrariumAPIResponseWriter) Write(rw http.ResponseWriter, data interface{}, statusCode int) {
	resp := &TerrariumDataResponse{
		Code: statusCode,
		Data: data,
	}
	jsonData, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("+%v", errors.Wrap(err))
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if statusCode != http.StatusNoContent {
		rw.Write(jsonData)
	}
}
