package usecase

import "go-backend-service/src/module/system_log/entity"

type SystemLogRepository interface {
	NewSystemLog(systemLog entity.SystemLog) (output entity.SystemLog, err error)
	GetSystemLog(limit int, offset int) (output []entity.SystemLog, err error)
}
