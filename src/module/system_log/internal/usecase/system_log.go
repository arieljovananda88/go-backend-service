package usecase

import (
	"go-backend-service/src/module/system_log/entity"
	"go-backend-service/src/utils"
)

type SystemLogUsecase interface {
	AddSytemLog(systemLog entity.SystemLog) (entity.SystemLog, error)
	GetSystemLog(param utils.LimitOffset) ([]entity.SystemLog, error)
}

type SystemLogUc struct {
	systemLogRepo SystemLogRepository
}

func NewSystemLogUseCase(systemLogRepo SystemLogRepository) *SystemLogUc {
	return &SystemLogUc{
		systemLogRepo: systemLogRepo,
	}
}

func (uc *SystemLogUc) AddSytemLog(systemLog entity.SystemLog) (entity.SystemLog, error) {
	newSystemLog, err := uc.systemLogRepo.NewSystemLog(systemLog)

	if err != nil {
		return newSystemLog, err
	}

	return newSystemLog, nil

}

func (uc *SystemLogUc) GetSystemLog(param utils.LimitOffset) ([]entity.SystemLog, error) {
	return uc.systemLogRepo.GetSystemLog(param.Limit, param.Offset)
}
