package communication

import "encoding/json"

// Response is representing the structure which is send as response
// of the serverless request
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// JSON returns the Response struct as json string
func (response *Response) JSON() string {
	json, err := json.Marshal(response)

	if err != nil {
		panic("Could not marshal response to JSON")
	}

	return string(json)
}

// NewErrorResponse creates a new Response with status code -1
func NewErrorResponse(err error) string {
	response := Response{
		Status:  -1,
		Message: err.Error(),
	}

	return response.JSON()
}

// NewSuccessResponse creates a new Response with status code -0
func NewSuccessResponse(data interface{}) string {
	response := Response{
		Status: 0,
		Data:   data,
	}

	return response.JSON()
}
