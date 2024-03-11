package models

import "time"

type SystemLogs struct {
	ID         string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	IDPengguna string    `gorm:"column:id_pengguna" json:"id_pengguna"`
	Action     string    `gorm:"column:action" json:"action"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}
