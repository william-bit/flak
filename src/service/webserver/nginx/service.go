package nginx

import (
	"flak/src/service"
	"log"
	"os/exec"
)

type NginxService struct {
	Version string
	Port    int
	Dir     string
	Cmd     *exec.Cmd
}

func New(version, dir string) *NginxService {
	return &NginxService{
		Version: version,
		Dir:     dir,
	}
}

func (nginx *NginxService) Start() {
	// Start Nginx
	if _, err := service.LoadPID("nginx"); err != nil || !service.RunningProcess(1234) {
		nginx.Cmd, err = service.StartService("nginx", nginx.Dir, `-c`, `C:\flak\etc\nginx\nginx.conf`)
		if err != nil {
			log.Fatalf("Failed to start Nginx: %v", err)
		} else {
			log.Printf("nginxCmd.Process.Pid: %d", nginx.Cmd.Process.Pid)
		}
	}
}

func (nginx *NginxService) Stop() {
	service.ShutdownService("Nginx", nginx.Cmd)
}

func (nginx *NginxService) Status() string {
	if nginx.Cmd != nil && nginx.Cmd.Process != nil {
		return "running"
	}
	return "stopped"
}
