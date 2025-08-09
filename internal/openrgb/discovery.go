package openrgb

import (
	"fmt"
)

func (c *Client) ListDeviceInfos() ([]DeviceInfo, error) {
	cntResp, err := c.c.RequestControllerCount()
	if err != nil {
		return nil, err
	}
	count := cntResp.Count

	var devices []DeviceInfo

	for i := 0; i < int(count); i++ {
		ctrlRsp, err := c.c.RequestControllerData(uint32(i))
		if err != nil {
			fmt.Printf("Failed to get controller %d: %v", i, err)
			continue
		}
		ctrl := ctrlRsp.Controller

		var deviceInfo DeviceInfo
		deviceInfo.ID = i
		deviceInfo.Name = ctrl.Name
		deviceInfo.Description = ctrl.Description
		deviceInfo.LEDCount = len(ctrl.Leds)
		deviceInfo.Vendor = ctrl.Vendor
		modeNames := make([]string, 0, len(ctrl.Modes))

		for _, mode := range ctrl.Modes {
			modeNames = append(modeNames, mode.ModeName)
		}
		deviceInfo.ModeNames = modeNames
		devices = append(devices, deviceInfo)
	}

	return devices, nil
}

func (c *Client) GetDeviceInfo(deviceId int) (*DeviceInfo, error) {
	ctrlRsp, err := c.c.RequestControllerData(uint32(deviceId))
	if err != nil {
		return nil, fmt.Errorf("failed to get controller %d: %w", deviceId, err)
	}
	ctrl := ctrlRsp.Controller

	deviceInfo := &DeviceInfo{
		ID:          deviceId,
		Name:        ctrl.Name,
		Description: ctrl.Description,
		Vendor:      ctrl.Vendor,
		LEDCount:    len(ctrl.Leds),
	}

	modeNames := make([]string, 0, len(ctrl.Modes))
	for _, mode := range ctrl.Modes {
		modeNames = append(modeNames, mode.ModeName)
	}
	deviceInfo.ModeNames = modeNames

	return deviceInfo, nil
}
