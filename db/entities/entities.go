package entities

// License is the entity which is getting stored
// as license in the database
type License struct {
	UserID     string `json:"userid"`
	Key        string `json:"key"`
	ValidUntil int64  `json:"validuntil"`
}

// Device is the entity which is getting stored
// as device in the database
type Device struct {
	UserID     string `json:"userid"`
	DeviceID   string `json:"deviceid"`
	Devicename string `json:"devicename"`
}
