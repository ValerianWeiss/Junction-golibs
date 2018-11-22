package queries

// FindResponse represents the structure of the response
// body, which is gettin send by the db function
type FindResponse struct {
	Status int `json:"status`
	Data   struct {
		Result []interface `json:"result"`
	} `json:"data"`
}