package shell

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func ExecuteCommand(command string) string {
	log.Println("execute command:" + command)
	res := ""
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
	}
	stout, err := cmd.StdoutPipe()
	if err == nil {
		err = cmd.Start()
		if err == nil {
			reader := bufio.NewReader(stout)
			for {
				line, _, err := reader.ReadLine()
				if err != nil || io.EOF == err {
					break
				}
				res += string(line) + "\n"
			}
			cmd.Wait()
		}
	}
	log.Println("get result :" + res)
	return strings.TrimSpace(res)
}

func ExecuteCommandAndPrintOutPut(command string) {
	log.Println("execute command:" + command)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
	}
	stout, err := cmd.StdoutPipe()
	if err == nil {
		err = cmd.Start()
		if err != nil {
			return
		}
		reader := bufio.NewReader(stout)
		for {
			line, _, err := reader.ReadLine()
			if err != nil || io.EOF == err {
				break
			}
			log.Println(string(line))
		}
		err = cmd.Wait()
		if err != nil {
			return
		}
	}

}

func ExecuteShell(command string) {
	log.Println("execute command:" + command)
	command = strings.TrimSpace(command)
	commandArr := strings.Split(command, " ")
	command = commandArr[0]
	args := commandArr[1:]
	cmd := exec.Command(command, args...)
	stout, err := cmd.StdoutPipe()
	if err == nil {
		err = cmd.Start()
		if err != nil {
			return
		}
		reader := bufio.NewReader(stout)
		for {
			line, _, err := reader.ReadLine()
			if err != nil || io.EOF == err {
				break
			}
			log.Println(string(line))
		}
		err = cmd.Wait()
		if err != nil {
			return
		}
	}

}
