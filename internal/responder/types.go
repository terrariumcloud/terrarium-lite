package responder

type TerrariumDataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type TerrariumServerResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
