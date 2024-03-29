package main

import (
	"log"
	"os"
	"service-deploy/processor"
	"service-deploy/service"
	"sync"
	"time"
)

func main() {
	defer initLog().Close()
	log.Println("start service....")
	services := service.ReadService()
	if services == nil {
		log.Println("no service found")
		return
	}
	asyncCount := 0
	for _, serv := range services {
		if serv.Async {
			asyncCount++
		}
	}

	var wg sync.WaitGroup
	wg.Add(asyncCount)
	for _, serv := range services {
		log.Printf("service name: %s", serv.Name)
		if serv.Async {
			go processService(serv, wg.Done)
		} else {
			processService(serv, nil)
		}
	}
	wg.Wait()
}

func processService(serv service.Service, done func()) {
	if done != nil {
		defer done()
	}
	if serv.Delay > 0 {
		time.Sleep(time.Duration(serv.Delay) * time.Millisecond)
	}
	if serv.Type == "kill" {
		processor.KillProcess(serv)
	} else if serv.Type == "command" {
		processor.ExecCommand(serv)
	} else if serv.Type == "app" {
		processor.RunApp(serv)
	}
}

/**
 * 初始化日志
 */
func initLog() *os.File {
	f, err := os.OpenFile("./service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
	}
	log.SetOutput(f)
	return f
}
