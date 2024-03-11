package handler

import (
	"go-backend-service/src/database"
	systemLogConfig "go-backend-service/src/module/system_log/config"
	systemLogTransport "go-backend-service/src/module/system_log/transport"
)

type Service struct {
	InternalSystemLogHandler *systemLogTransport.InternalSystemLogHandler
	AdminSytemLogHandler     *systemLogTransport.AdminSystemLogHandler
}

func MakeHandler() *Service {

	dbWrite := database.GetInstance()
	dbRead := database.GetInstance()

	InternalSystemLogHandler := systemLogTransport.NewInternalSystemLogHandler(systemLogConfig.SystemLogTransportConfig{
		DBWrite: dbWrite,
		DBRead:  dbRead,
	})
	AdminSystemLogHandler := systemLogTransport.NewAdminSystemLogHandler(systemLogConfig.SystemLogTransportConfig{
		DBWrite: dbWrite,
		DBRead:  dbRead,
	})

	return &Service{
		InternalSystemLogHandler: InternalSystemLogHandler,
		AdminSytemLogHandler:     AdminSystemLogHandler,
	}
}
