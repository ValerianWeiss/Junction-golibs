package communication

// PushData represents the API interface of the
// pushdata function
type PushData struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}
