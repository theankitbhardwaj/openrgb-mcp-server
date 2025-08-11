package mcp

type ListDevicesParams struct {
	// No parameters needed for listing devices
}

type SetDeviceColorParams struct {
	DeviceID int `json:"device_id"`
	R        int `json:"r"`
	G        int `json:"g"`
	B        int `json:"b"`
}

type SetAllColorParams struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type ListProfilesParams struct {
	// No parameters needed for listing devices
}

type SetProfileParams struct {
	ProfileName string `json:"profile_name"`
}
