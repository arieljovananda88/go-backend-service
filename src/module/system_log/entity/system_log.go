package entity

import (
	"time"
)

type SystemLog struct {
	ID         string    `json:"id"`
	IDPengguna string    `json:"id_pengguna"`
	Action     string    `json:"action"`
	CreatedAt  time.Time `json:"created_at"`
}

type ParamValues struct {
	Limit  int
	Offset int
}
