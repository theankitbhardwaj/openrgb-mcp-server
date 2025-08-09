package openrgb

type DeviceInfo struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Vendor      string   `json:"vendor"`
	LEDCount    int      `json:"led_count"`
	ModeNames   []string `json:"supported_modes"`
}
