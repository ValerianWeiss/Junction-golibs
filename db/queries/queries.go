package queries

import "github.com/ValerianWeiss/Junction-golibs/db/entities"

// FindValues request for the db function
type FindValues struct {
	Collection string `json:"collection"`
	Query      struct {
		DeviceID string `json:"deviceid"`
		Time     struct {
			Min int64 `json:"$gt"`
			Max int64 `json:"$lt"`
		} `json:"time"`
	} `json:"query"`
}

// NewFindValuesQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewFindValuesQuery(deviceid string, minTime, maxTime int64) FindValues {
	query := new(FindValues)
	query.Collection = "iotdata"
	query.Query.DeviceID = deviceid
	query.Query.Time.Min = minTime
	query.Query.Time.Max = maxTime
	return *query
}

// FindDevices request for the db function
type FindDevices struct {
	Collection string `json:"collection"`
	Query      struct {
		entities.Device
	} `json:"query"`
}

// NewFindDevicesQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewFindDevicesQuery(device entities.Device) FindDevices {
	query := new(FindDevices)
	query.Collection = "iotdevices"
	query.Query.Device = device
	return *query
}

// FindLicenses request for the db function
type FindLicenses struct {
	Collection string `json:"collection"`
	Query      struct {
		entities.License
	} `json:"query"`
}

// NewFindLicensesQuery creates a new db query struct which can be marshalled to
// JSON and send to the db function
func NewFindLicensesQuery(license entities.License) FindLicenses {
	query := new(FindLicenses)
	query.Collection = "licenses"
	query.Query.License = license
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
