package dynamoDB

type Address struct {
	City       string `json:"city"`
	State      string `json:"subdivision"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type DeviceInfo struct {
	DeviceId   string
	DeviceType string
}

// type user models a user's attributes
type User struct {
	UserID    float64 `json:"userId"`
	UserUUID  string  `json:"userUuid"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Gender    string  `json:"gender"`
	Age       int `json:"age"`
	Address   Address
	Devices   []DeviceInfo
}