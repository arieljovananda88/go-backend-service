package config

import "gorm.io/gorm"

type SystemLogTransportConfig struct {
	DBWrite *gorm.DB
	DBRead  *gorm.DB
}
