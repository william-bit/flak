package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

const (
	pidDir = "./tmp/pid"
	wwwDir = "./www"
)

func init() {
	_ = os.Mkdir(pidDir, 0755)
}

// Save PID to file
func savePID(name string, pid int) error {
	return os.WriteFile(filepath.Join(pidDir, name+".pid"), []byte(strconv.Itoa(pid)), 0644)
}

// Load PID from file
func loadPID(name string) (int, error) {
	data, err := os.ReadFile(filepath.Join(pidDir, name+".pid"))
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(string(data))
	return pid, err
}

// Check if process is still running
func isProcessRunning(pid int) bool {
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid))
	out, err := cmd.CombinedOutput()
	return err == nil && len(out) > 0
}

// Start service in detached mode
func startService(name, path string, args ...string) (*exec.Cmd, error) {
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

	savePID(name, cmd.Process.Pid)
	log.Printf("%s started with PID: %d", name, cmd.Process.Pid)
	return cmd, nil
}

// Reconnect to existing services
func reconnectToExistingServices() {
	names := []string{"mysql", "php", "nginx"}
	for _, name := range names {
		pid, err := loadPID(name)
		if err != nil {
			continue
		}
		if isProcessRunning(pid) {
			log.Printf("%s is already running with PID %d", name, pid)
		} else {
			log.Printf("%s had PID %d but is no longer running.", name, pid)
		}
	}
}

func main() {
	// Try to reconnect to existing services
	reconnectToExistingServices()

	var mysqlCmd, phpCmd, nginxCmd *exec.Cmd
	var err error

	// Start MySQL (Update path accordingly)
	if _, err := loadPID("mysql"); err != nil || !isProcessRunning(1234) {
		mysqlCmd, err = startService("mysql", `C:\mysql\bin\mysqld.exe`, "--console")
		if err != nil {
			log.Fatalf("Failed to start MySQL: %v", err)
		}
	}

	// Start PHP-CGI
	if _, err := loadPID("php"); err != nil || !isProcessRunning(1234) {
		phpCmd, err = startService("php", `C:\php\php-cgi.exe`, "-b", "127.0.0.1:9000")
		if err != nil {
			log.Fatalf("Failed to start PHP: %v", err)
		}
	}

	// Start Nginx
	if _, err := loadPID("nginx"); err != nil || !isProcessRunning(1234) {
		nginxCmd, err = startService("nginx", `C:\nginx\nginx.exe`)
		if err != nil {
			log.Fatalf("Failed to start Nginx: %v", err)
		}
	}

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down services...")

	// Optional: Add graceful shutdown logic if supported by services
	time.Sleep(3 * time.Second)
}
