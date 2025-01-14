package entity

import (
	"time"
)

type Transaction struct {
	ID        uint
	AccountID uint
	Amount    int
	Type      Type
	CreatedAt time.Time
}

type Type string

func (t Type) IsValid() string {
	switch t {
	case typeA:
		return "deposit"
	case typeB:
		return "withdraw"

	}
	return ""
}

const (
	typeA Type = "deposit"
	typeB Type = "withdraw"
)
