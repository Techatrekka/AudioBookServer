package model

type Audiobook struct {
	AudioFile    string `json:"audioFile"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Synopsis     string `json:"synopsis"`
	Id           string `json:"id"`
	IconLocation string `json:"iconLocation"`
}
