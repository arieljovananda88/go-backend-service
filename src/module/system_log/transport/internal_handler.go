package transport

import (
	"net/http"

	"go-backend-service/src/module/system_log/config"
	"go-backend-service/src/module/system_log/entity"
	"go-backend-service/src/module/system_log/internal/repository"
	"go-backend-service/src/module/system_log/internal/usecase"
	"go-backend-service/src/utils"

	"github.com/labstack/echo"
)

type InternalSystemLogHandler struct {
	systemLogUsecase usecase.SystemLogUsecase
}

func NewInternalSystemLogHandler(cfg config.SystemLogTransportConfig) *InternalSystemLogHandler {
	systemLogRepository := repository.NewSystemLogRepository(cfg.DBWrite, cfg.DBRead)
	systemLogUsecase := usecase.NewSystemLogUseCase(systemLogRepository)

	return &InternalSystemLogHandler{
		systemLogUsecase: systemLogUsecase,
	}

}

func (t *InternalSystemLogHandler) MountInternal(group *echo.Group) {
	group.POST("/system-log", t.AddSytemLog)
}

func (t *InternalSystemLogHandler) AddSytemLog(c echo.Context) error {
	systemLog := new(entity.SystemLog)

	if err := c.Bind(systemLog); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	newSystemLog, err := t.systemLogUsecase.AddSytemLog(*systemLog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ResponseDetailOutput(false, http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.ResponseDetailOutput(true, http.StatusCreated, "System log added successfully", newSystemLog))
}
