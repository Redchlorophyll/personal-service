package response

type GeneralResponse struct {
	ErrorCode  *int   `json:"error_code"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
