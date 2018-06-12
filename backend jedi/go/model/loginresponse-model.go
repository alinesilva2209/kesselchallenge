package model

import "time"

type LoginResponse struct {
	Result	Result
	Error Error
	Date time.Time
}

type Result struct {
	Status	string
	Token string
}

type Error struct {
	Error string
	Stack string
}
