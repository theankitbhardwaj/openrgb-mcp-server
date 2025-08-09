package app

import "fmt"

func validateRGB(r, g, b int) error {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return fmt.Errorf("RGB values must be between 0 and 255")
	}
	return nil
}
