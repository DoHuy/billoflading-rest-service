package models

import (
	"time"
)

const (
	IndexChitietdon = "chitiet_don"
	IndexVandonhanhtrinh = "vandon_hanhtrinh"
)

type Chitietdon struct {
	ID          int64 `gorm:"primary_key"`
	LogID       string
	Type        string
	Description string
	ClientID    string
	ClientName  string
	IP          string
	UserAgent   string
	UserID      string
	CreatedAt   *time.Time
}

type Vandonhanhtrinh struct {
	ID          int64 `gorm:"primary_key"`
	LogID       string
	Type        string
	Description string
	ClientID    string
	ClientName  string
	IP          string
	UserAgent   string
	UserID      string
	CreatedAt   *time.Time
}
