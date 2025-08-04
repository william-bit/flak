package runtime

import (
	"flak/src/service"
	"log"
	"os/exec"
)

var phpCmd *exec.Cmd

func StartPHP() {
	// Start PHP-CGI
	if _, err := service.LoadPID("php"); err != nil || !service.IsProcessRunning(1234) {
		phpCmd, err = service.StartService("php", `C:\flak\bin\php\php-8.4.6-nts-Win32-vs17-x64\php-cgi.exe`, "-b", "127.0.0.1:9003")
		if err != nil {
			log.Fatalf("Failed to start PHP: %v", err)
		} else {
			log.Printf("phpCmd.Process.Pid: %d", phpCmd.Process.Pid)
		}
	}
}

func StopPHP() {
	service.ShutdownService("PHP", phpCmd)
}
