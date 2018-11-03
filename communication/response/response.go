package response

import "encoding/json"

// Response is representing the structure which is send as response
// of the serverless request
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON returns the Response struct as json string
func (response *Response) JSON() (string, error) {
	json, err := json.Marshal(response)
	return string(json), err
}

// NewErrorResponse creates a new Response with status code -1
func NewErrorResponse(err error) string {
	response := Response{
		Status:  -1,
		Message: err.Error(),
	}

	json, _ := response.JSON()
	return json
}

// NewSuccessResponse creates a new Response with status code -0
func NewSuccessResponse(message string, data interface{}) string {
	response := Response{0, message, data}
	json, _ := response.JSON()
	return json
}
