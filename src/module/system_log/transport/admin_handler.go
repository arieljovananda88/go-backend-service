package transport

import (
	"net/http"

	"go-backend-service/src/module/system_log/config"
	"go-backend-service/src/module/system_log/internal/repository"
	"go-backend-service/src/module/system_log/internal/usecase"
	"go-backend-service/src/utils"

	"github.com/labstack/echo"
)

type AdminSystemLogHandler struct {
	systemLogUsecase usecase.SystemLogUsecase
}

func NewAdminSystemLogHandler(cfg config.SystemLogTransportConfig) *AdminSystemLogHandler {
	systemLogRepository := repository.NewSystemLogRepository(cfg.DBWrite, cfg.DBRead)
	systemLogUsecase := usecase.NewSystemLogUseCase(systemLogRepository)

	return &AdminSystemLogHandler{
		systemLogUsecase: systemLogUsecase,
	}

}

func (t *AdminSystemLogHandler) MountAdmin(group *echo.Group) {
	group.GET("/system-logs", t.GetSystemLog)
}

func (t *AdminSystemLogHandler) GetSystemLog(c echo.Context) error {
	param, err := utils.GetLimitOffset(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ResponseDetailOutput(false, http.StatusInternalServerError, err.Error(), nil))
	}

	listGroupLimit, err := t.systemLogUsecase.GetSystemLog(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.ResponseDetailOutput(true, http.StatusOK, "System logs successfully retreived", listGroupLimit))
}
