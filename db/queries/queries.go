package queries

import "github.com/ValerianWeiss/Junction-golibs/db/entities"

// FindValues request for the db function
type FindValues struct {
	Collection string `json:"collection"`
	Query      struct {
		DeviceID  string `json:"deviceid"`
		Timestamp struct {
			Min int64 `json:"$gt"`
			Max int64 `json:"$lt"`
		} `json:"timestamp"`
	} `json:"query"`
}

// NewFindValuesQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewFindValuesQuery(deviceid string, min, max int64) FindValues {
	query := new(FindValues)
	query.Collection = "iotdevices"
	query.Query.DeviceID = deviceid
	query.Query.Timestamp.Min = min
	query.Query.Timestamp.Max = max
	return *query
}

// FindDeviceID request for the db function
type FindDeviceID struct {
	Collection string `json:"collection"`
	Query      struct {
		UserID     string `json:"userid"`
		Devicename string `json:"devicename"`
	} `json:"query"`
}

// NewFindDeviceIDQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewFindDeviceIDQuery(userid, devicename string) FindDeviceID {
	query := new(FindDeviceID)
	query.Collection = "iotdevices"
	query.Query.UserID = userid
	query.Query.Devicename = devicename
	return *query
}

// FindDevices request for the db function
type FindDevices struct {
	Collection string `json:"collection"`
	Query      struct {
		UserID string `json:"userid"`
	} `json:"query"`
}

// NewFindDevicesQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewFindDevicesQuery(userid string) FindDevices {
	query := new(FindDevices)
	query.Collection = "iotdevices"
	query.Query.UserID = userid
	return *query
}

// StoreLicenses request for the db function
type StoreLicenses struct {
	Collection string             `json:"collection"`
	Records    []entities.License `json:"records"` // db function is expecting the value as array
}

// NewStoreLicenseQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewStoreLicenseQuery(license entities.License) StoreLicenses {
	return StoreLicenses{"licenses", []entities.License{license}}
}

// StoreDevices request for the db function
type StoreDevices struct {
	Collection string            `json:"collection"`
	Records    []entities.Device `json:"records"` // db function is expecting the value as array
}

// NewStoreDeviceQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewStoreDeviceQuery(device entities.Device) StoreDevices {
	return StoreDevices{"iotdevices", []entities.Device{device}}
}
