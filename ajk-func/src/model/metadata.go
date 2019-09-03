package model

type Metadata struct {
	OriginURL string     `json:"originalImageUrl"`
	Emojis    [][]*Emoji `json:"emojis"`
}
