package entities

// License is the entity which is getting stored
// as license in the database
type License struct {
	UserID     string `json:"userid,omitempty"`
	Key        string `json:"key,omitempty"`
	ValidUntil int64  `json:"validuntil,omitempty"`
}

// Device is the entity which is getting stored
// as device in the database
type Device struct {
	UserID     string `json:"userid,omitempty"`
	DeviceID   string `json:"deviceid,omitempty"`
	Devicename string `json:"devicename,omitempty"`
}
