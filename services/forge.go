package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

// url := "http://127.0.0.1:7860/sdapi/v1/txt2img"

// 	payload := ImagePayload{
// 		Prompt:            prompt,
// 		Steps:             steps,
// 		CfgScale:          1,
// 		DistilledCfgScale: 3.5,
// 		Width:             896,
// 		Height:            1152,
// 		SamplerName:       "euler",
// 		// SamplerName: "DPM++ 2M SDE",
// 		Scheduler: "Simple",
// 		// Scheduler:         "Karras",
// 		EnableHr:          false,
// 		Seed:              -1,
// 		DenoisingStrength: 0.5,
// 	}

func (f *Forge) Image2Image(ctx context.Context, request dto.ProcessImageRequest) (string, error) {

	payload := ImagePayload{
		Prompt:            request.Prompt,
		Steps:             30,
		CfgScale:          7,
		Width:             request.Width,
		Height:            request.Height,
		InitImages:        []string{request.Data},
		DenoisingStrength: 0.5,
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

	return result.Images[0], nil
}
