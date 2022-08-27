package usecase

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
