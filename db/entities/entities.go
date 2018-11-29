package entities

// License is the entity which is getting stored
// as license object in the database
type License struct {
	UserID     string `json:"userid,omitempty"`
	Key        string `json:"key,omitempty"`
	ValidUntil int64  `json:"validuntil,omitempty"`
}

// Device is the entity which is getting stored
// as device object in the database
type Device struct {
	UserID     string `json:"userid,omitempty"`
	DeviceID   string `json:"deviceid,omitempty"`
	Devicename string `json:"devicename,omitempty"`
	IAT        int64  `json:"iat,omitempty"`
}

// Value is the entity which is getting stored
// as value object in the database
type Value struct {
	DeviceID string  `json:"deviceid,omitempty"`
	Value    float64 `json:"value,omitempty"`
	IAT      int64   `json:"iat,omitempty"`
}
