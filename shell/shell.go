package shell

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func ExecuteCommandAndGetResult(command string) string {
	res := ""
	cmd := buildCommand(command)
	stout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stout)
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		res += string(line) + "\n"
	}
	cmd.Wait()
	return res
}

func ExecuteShellAndGetResult(command string) string {
	res := ""
	cmd := buildShell(command)
	stout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stout)
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		res += string(line) + "\n"
	}
	cmd.Wait()
	return res
}

func ExecuteCommandAndLog(command string) {
	cmd := buildCommand(command)
	stout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stout)
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		log.Println(line)
	}
	cmd.Wait()
}

func ExecuteShellAndLog(command string) {
	cmd := buildShell(command)
	stout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stout)
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		log.Println(line)
	}
	cmd.Wait()
}

func ExecuteCommand(command string) {
	cmd := buildCommand(command)
	cmd.Run()
}

func ExecuteShell(command string) {
	cmd := buildShell(command)
	cmd.Run()
}

func buildCommand(command string) *exec.Cmd {
	command = strings.TrimSpace(command)
	commandArr := strings.Split(command, " ")
	command = commandArr[0]
	args := commandArr[1:]
	return exec.Command(command, args...)
}

func buildShell(command string) *exec.Cmd {
	if runtime.GOOS == "linux" {
		return exec.Command("/bin/sh", "-c", command)
	} else {
		return exec.Command("cmd", "/C", command)
	}
}
