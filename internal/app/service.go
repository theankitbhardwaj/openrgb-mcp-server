package app

import (
	"context"
	"fmt"

	"github.com/theankitbhardwaj/openrgb-mcp-server/internal/openrgb"
)

type Service struct {
	rgbClient *openrgb.Client
}

func NewService(rgbClient *openrgb.Client) *Service {
	return &Service{
		rgbClient: rgbClient,
	}
}

func (s *Service) ListDevices(ctx context.Context) ([]openrgb.DeviceInfo, error) {
	return s.rgbClient.ListDeviceInfos()
}

func (s *Service) SetDeviceColor(ctx context.Context, deviceID int, r, g, b int) error {
	if err := validateRGB(r, g, b); err != nil {
		return err
	}

	deviceInfo, err := s.rgbClient.GetDeviceInfo(deviceID)
	if err != nil {
		return err
	}
	return s.rgbClient.SetDeviceColor(*deviceInfo, r, g, b)
}

func (s *Service) SetAllDevicesColor(ctx context.Context, r, g, b int) error {
	if err := validateRGB(r, g, b); err != nil {
		return err
	}
	return s.rgbClient.SetAllDeviceColor(r, g, b)
}

func (s *Service) ListProfiles(ctx context.Context) ([]openrgb.Profile, error) {
	profiles, err := s.rgbClient.ListProfiles()
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *Service) SetProfile(ctx context.Context, profileName string) error {
	if profileName == "" {
		return fmt.Errorf("profile name cannot be empty")
	}
	return s.rgbClient.SetProfile(profileName)
}
