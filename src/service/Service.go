package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

const (
	PidDir = "./tmp/pid"
	wwwDir = "./www"
)

func init() {
	_ = os.Mkdir(PidDir, 0755)
}

// Save PID to file
func SavePID(name string, pid int) error {
	return os.WriteFile(filepath.Join(PidDir, name+".pid"), []byte(strconv.Itoa(pid)), 0644)
}

// Start service in detached mode
func StartService(name, path string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(path, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Starting %s...", name)
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	SavePID(name, cmd.Process.Pid)
	log.Printf("%s started with PID: %d", name, cmd.Process.Pid)
	return cmd, nil
}

// Load PID from file
func LoadPID(name string) (int, error) {
	data, err := os.ReadFile(filepath.Join(PidDir, name+".pid"))
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(string(data))
	return pid, err
}

// Check if process is still running
func IsProcessRunning(pid int) bool {
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid))
	out, err := cmd.CombinedOutput()
	return err == nil && len(out) > 0
}

// Reconnect to existing services
func ReconnectToExistingServices() {
	names := []string{"mysql", "php", "nginx"}
	for _, name := range names {
		pid, err := LoadPID(name)
		if err != nil {
			continue
		}
		if IsProcessRunning(pid) {
			log.Printf("%s is already running with PID %d", name, pid)
		} else {
			log.Printf("%s had PID %d but is no longer running.", name, pid)
		}
	}
}

func ShutdownService(name string, cmd *exec.Cmd) error {
	if cmd != nil && cmd.Process != nil {
		pid := cmd.Process.Pid

		cmd := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error killing process: %s", out)
			return err
		}

		fmt.Printf("Successfully killed process with PID %d\n", pid)
		return nil
	} else {
		return fmt.Errorf("no process found for %s", name)
	}
}
