package model

type RequestBody struct {
	FnamePrefix    string `json:"prefix"`
	Xsplit         int    `json:"xsplit"`
	ImageAsDataURL string `json:"imageAsDataUrl"`
}
