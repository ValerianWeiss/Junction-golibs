package entities

// License is the entity which is getting stored
// as license in the database
type License struct {
	UserID     string `json:"userid"`
	License    string `json:"license"`
	ValidUntil int64  `json:"validuntil"`
}
