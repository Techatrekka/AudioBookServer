package models

import "time"

type Tag struct {
	TagId       int    `json:"tag_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	UserId  int    `json:"user_id"`
	Email   string `json:"email"`
	SsoType string `json:"sso_type"`
}

type Tape struct {
	TapeId      int    `json:"tape_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Synopsis    string `json:"synopsis"`
	IsAudiobook bool   `json:"is_audiobook"`
	Tags        []int  `json:"tags"`
}

type ListeningHistory struct {
	TapeId          int       `json:"tape_id"`
	UserId          int       `json:"user_id"`
	CurrentChapter  int       `json:"current_chapter"`
	ChapterProgress string    `json:"chapter_progress"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Review struct {
	TapeId     int       `json:"tape_id"`
	UserId     int       `json:"user_id"`
	Rating     float64   `json:"rating"`
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"review_date"`
}
