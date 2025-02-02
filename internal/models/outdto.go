package models

type Outdto[T any] struct {
	Dat T        `json:"dat"`
	Err []string `json:"err"`
}

type ManGetOutdto struct {
	Output string `json:"output"`
}
