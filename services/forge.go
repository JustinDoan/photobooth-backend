package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"github.com/nfnt/resize"
	_ "image/jpeg"
	_ "image/png"
	"github.com/justindoan/photobooth-backend/dto"
)

var URL = "http://127.0.0.1:7860/sdapi/v1/img2img"

type Forge struct{}

func NewForge() *Forge {
	return &Forge{}
}

type ImagePayload struct {
	Prompt            string   `json:"prompt"`
	Steps             int      `json:"steps"`
	CfgScale          float64  `json:"cfg_scale"`
	DistilledCfgScale float64  `json:"distilled_cfg_scale"`
	Width             int      `json:"width"`
	Height            int      `json:"height"`
	SamplerName       string   `json:"sampler_name"`
	Scheduler         string   `json:"scheduler"`
	EnableHr          bool     `json:"enable_hr"`
	Seed              int      `json:"seed"`
	DenoisingStrength float64  `json:"denoising_strength"`
	InitImages        []string `json:"init_images"`
}

type ImageResponse struct {
	Images []string `json:"images"`
}

func (f *Forge) Image2Image(ctx context.Context, request dto.ProcessImageRequest) (string, error) {
	payload := ImagePayload{
		Prompt:            request.Prompt,
		Steps:             30,
		CfgScale:          request.CfgScale,
		Width:             request.Width,
		Height:            request.Height,
		InitImages:        []string{request.Data},
		DenoisingStrength: request.DenoisingStrength,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling payload: %v", err)
	}

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error! status: %d", resp.StatusCode)
	}

	var result ImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(result.Images) == 0 {
		return "", fmt.Errorf("no image generated")
	}

	// Flip the image upside down
	flippedImage, err := flipImageUpsideDown(result.Images[0])
	if err != nil {
		return "", fmt.Errorf("error flipping image: %v", err)
	}

	return flippedImage, nil
}

func flipImageUpsideDown(encodedImage string) (string, error) {
	// Decode the base64 image
	imageData, err := base64.StdEncoding.DecodeString(encodedImage)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 image: %v", err)
	}

	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return "", fmt.Errorf("error decoding image: %v", err)
	}

	// Flip the image upside down
	flippedImg := imaging.FlipV(img)

	// Encode the flipped image to PNG
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, flippedImg); err != nil {
		return "", fmt.Errorf("error encoding flipped image: %v", err)
	}

	// Return the new base64 encoded image
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
