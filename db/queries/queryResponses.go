package queries

// FindResponse represents the structure of the response
// body, which is getting send by the db function on /find
type FindResponse struct {
	Status int `json:"status"`
	Data   struct {
		Result []interface{} `json:"result"`
	} `json:"data"`
}
