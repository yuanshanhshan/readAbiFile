package common

type Input struct {
	Indexed      bool   `json:"indexed"`
	InternalType string `json:"internalType"`
	Name         string `json:"owner"`
	Type         string `json:"type"`
}
type JsonBody struct {
	Anonymous bool    `json:"anonymous"`
	Inputs    []Input `json:"inputs"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
}
type JsonBodyArr struct {
	JsonArr []JsonBody `json:"token"`
}
