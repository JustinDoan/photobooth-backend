// Package dto contains Data Transfer Objects used for API requests and responses
package dto

// ProcessImageRequest represents the request structure for processing an image
type ProcessImageRequest struct {
	Data   string `json:"data"` // Base64 encoded image data
	Name   string `json:"name"` // Name of the image
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Prompt string `json:"prompt"`
}
