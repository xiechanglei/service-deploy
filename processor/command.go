package processor

import (
	"fmt"
	"runtime"
	"service-deploy/service"
	shell2 "service-deploy/shell"
)

func ExecCommand(serv service.Service) {
	if runtime.GOOS == "windows" {
		windowsCommand(serv.Dir, serv.Cmd)
	} else {
		linuxCommand(serv.Dir, serv.Cmd)
	}
}

func linuxCommand(dir string, shell string) {
	shell2.ExecuteCommandAndPrintOutPut(fmt.Sprintf(`cd "%s" && %s`, dir, shell))
}

func windowsCommand(dir string, shell string) {
	shell2.ExecuteCommandAndPrintOutPut(fmt.Sprintf(`cd "%s" && %s`, dir, shell))
}
