package processor

import (
	"service-deploy/service"
	"service-deploy/shell"
)

func RunApp(serv service.Service) {
	shell.ExecuteShell(serv.Cmd)
}
