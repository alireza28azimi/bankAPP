package entity

import "time"

type Card struct {
	UserID         uint
	AccountID      uint
	CardNumber     string
	ExpirationDate time.Time
	Status         Status
}

type Status string

func (p Status) IsValid() string {
	switch p {
	case StatusA:
		return "Available"
	case StatusB:
		return "Blocked"
	case StatusC:
		return "Expired"
	}
	return ""
}

const (
	StatusA Status = "Available"
	StatusB Status = "Blocked"
	StatusC Status = "Expired"
)
