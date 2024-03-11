package repository

import (
	"time"

	"go-backend-service/src/module/system_log/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SystemLogRepo struct {
	DBWrite *gorm.DB
	DBRead  *gorm.DB
}

func NewSystemLogRepository(dbWrite *gorm.DB, dbRead *gorm.DB) *SystemLogRepo {
	return &SystemLogRepo{
		DBWrite: dbWrite,
		DBRead:  dbRead,
	}
}

func (repo *SystemLogRepo) NewSystemLog(systemLog entity.SystemLog) (output entity.SystemLog, err error) {
	systemLog.ID = uuid.New().String()

	// Set the CreatedAt field to the current time
	systemLog.CreatedAt = time.Now()
	result := repo.DBWrite.Create(&systemLog)
	if result.Error != nil {
		return entity.SystemLog{}, result.Error
	}
	return systemLog, nil
}

func (repo *SystemLogRepo) GetSystemLog(limit int, offset int) (output []entity.SystemLog, err error) {
	var systemLogs []entity.SystemLog
	result := repo.DBRead.Limit(limit).Offset(offset).Find(&systemLogs)
	if result.Error != nil {
		return nil, err
	}
	return systemLogs, nil
}
