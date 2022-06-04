package common_adapters

type OutputPort struct {
	StatusCode int    `json:"-"`
	Data       any    `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
}
