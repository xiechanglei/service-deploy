package processor

import (
	"fmt"
	"service-deploy/service"
	"service-deploy/shell"
)

func ExecCommand(serv service.Service) {
	shell.ExecuteShell(fmt.Sprintf(`cd "%s" && %s`, serv.Dir, serv.Cmd))
}
