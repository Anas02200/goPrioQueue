package SDiffApiClient

type TextToImageResponse struct {
	// The generated image in base64 format.
	Images     []string     `json:"images,omitempty"`
	Parameters *interface{} `json:"parameters"`
	Info       string       `json:"info"`
}
