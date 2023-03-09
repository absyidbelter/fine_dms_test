package model

import "time"

type FileDownload struct {
	Id   int       `json:"id"`
	File File      `json:"file" validate:"required`
	User User      `json:"user" validate:"required`
	Date time.Time `json:"date" validate:"required`
}
