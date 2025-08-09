package openrgb

import "github.com/csutorasa/go-openrgb-sdk"

func (c *Client) ListProfiles() ([]Profile, error) {
	err := c.c.RequestProtocolVersion()
	if err != nil {
		return nil, err
	}
	rplRsp, err := c.c.RequestProfileList()

	if err != nil {
		return nil, err
	}
	profiles := make([]Profile, len(rplRsp.Names))
	for i, name := range rplRsp.Names {
		profiles[i] = Profile{
			Name: name,
		}
	}
	return profiles, nil
}

// TODO: Fix that it doesn't return an error if the profile doesn't exist
func (c *Client) SetProfile(name string) error {
	err := c.c.RequestProtocolVersion()
	if err != nil {
		return err
	}
	rsplReq := &openrgb.RequestLoadProfileRequest{
		ProfileName: name,
	}
	return c.c.RequestLoadProfile(rsplReq)
}
