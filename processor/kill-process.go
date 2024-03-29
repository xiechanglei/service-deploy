package processor

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"service-deploy/service"
	"service-deploy/shell"
	"strconv"
	"strings"
)

// KillProcess 杀进程
func KillProcess(serv service.Service) {
	log.Printf("start to kill process on port %s\n", strconv.Itoa(serv.Port))
	if runtime.GOOS == "windows" {
		windowsKill(serv.Port)
	} else {
		linuxKill(serv.Port)
	}

}

// linux 杀端口对应的进程
func linuxKill(port int) {
	pid := shell.ExecuteShellAndGetResult("lsof -i:" + strconv.Itoa(port) + " | awk '{print $2}' | awk  'NR==2{print}'")
	log.Printf("find process: %s\n", pid)
	shell.ExecuteShell(`kill -9 ` + pid)
}

// windows 杀端口对应的进程
func windowsKill(port int) {
	resStr := shell.ExecuteShellAndGetResult(fmt.Sprintf("netstat -ano -p tcp | findstr %d", port))
	r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr+" ", -1)
	if len(r) > 0 {
		pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err == nil {
			log.Printf("find process: %s\n", strconv.Itoa(pid))
			shell.ExecuteShell(fmt.Sprintf("taskkill -pid %d /F", pid))
		}
	}
}
