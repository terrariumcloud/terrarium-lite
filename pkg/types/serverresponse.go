package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TerrariumDataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type TerrariumServerResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ToJSONString(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "   ")
}

func NewTerrariumServerError(message string) ([]byte, error) {
	sr := &TerrariumServerResponse{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("Internal Server Error - %s", message),
	}
	return ToJSONString(sr)
}

func NewTerrariumBadRequest(message string) ([]byte, error) {
	sr := &TerrariumServerResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("Bad Request - %s", message),
	}
	return ToJSONString(sr)
}

func NewTerrariumOK(data interface{}) ([]byte, error) {
	sr := &TerrariumDataResponse{
		Code: http.StatusOK,
		Data: data,
	}
	return ToJSONString(sr)
}
