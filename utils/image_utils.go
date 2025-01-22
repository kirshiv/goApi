package utils

import (
	"errors"
	"image"
	"net/http"
	_ "image/jpeg"
	_ "image/png"
)

func CalculatePerimeter(url string) (float64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	img, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return 0, errors.New("failed to decode image")
	}

	return 2 * float64(img.Width+img.Height), nil
}
