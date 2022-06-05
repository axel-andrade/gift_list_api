package common_adapters

type OutputPort struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}
