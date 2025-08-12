package php

import (
	"flak/src/service"
	"log"
	"os/exec"
)

type PHPService struct {
	Version string
	Port    int
	Dir     string
	Cmd     *exec.Cmd
}

func New(version, dir string) *PHPService {
	return &PHPService{
		Version: version,
		Dir:     dir,
	}
}

func (php *PHPService) Start() {
	// Start PHP-CGI
	if _, err := service.LoadPID("php"); err != nil || !service.RunningService(1234) {
		php.Cmd, err = service.StartService("php", php.Dir, "-b", "127.0.0.1:9003")
		if err != nil {
			log.Fatalf("Failed to start PHP: %v", err)
		} else {
			log.Printf("phpCmd.Process.Pid: %d", php.Cmd.Process.Pid)
		}
	}
}

func (php *PHPService) Stop() {
	service.ShutdownService("PHP", php.Cmd)
}

func (php *PHPService) Status() string {
	if php.Cmd != nil && php.Cmd.Process != nil {
		return "running"
	}
	return "stopped"
}
