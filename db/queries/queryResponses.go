package queries

// FindResponse represents the structure of the response
// body, which is getting send by the db function on /find
type FindResponse struct {
	Status int `json:"status"`
	Data   struct {
		Result []interface{} `json:"result"`
	} `json:"data"`
}

// InsertResponse represents the structure of the response
// body, which is getting send by the db function on /insert
type InsertResponse struct {
	Result struct {
		Ok          int `json:"ok"`
		RecordCount int `json:"n"`
	} `json:"result"`
	InsertedRecords []interface{} `json:"ops"`
}
