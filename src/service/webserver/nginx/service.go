package nginx

import (
	"flak/src/service"
	"log"
	"os/exec"
)

var nginxCmd *exec.Cmd

func Start() {
	// Start Nginx
	if _, err := service.LoadPID("nginx"); err != nil || !service.RunningProcess(1234) {
		nginxCmd, err = service.StartService("nginx", `C:\flak\bin\nginx\nginx-1.22.0\nginx.exe`, `-c`, `C:\flak\etc\nginx\nginx.conf`)
		if err != nil {
			log.Fatalf("Failed to start Nginx: %v", err)
		} else {
			log.Printf("nginxCmd.Process.Pid: %d", nginxCmd.Process.Pid)
		}
	}
}

func Stop() {
	service.ShutdownService("Nginx", nginxCmd)
}
