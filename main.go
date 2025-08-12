package main

import (
	"flak/src/service"
	"flak/src/service/database/mysql"
	"flak/src/service/runtime/php"
	"flak/src/service/webserver/nginx"
	"flak/src/tui"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func InitServices() { // Try to reconnect to existing services
	service.ResumeService()

	var services []service.Service = []service.Service{
		nginx.New("1.22.0", `C:\flak\bin\nginx\nginx-1.22.0\nginx.exe`),
		php.New("8.4.6", `C:\flak\bin\php\php-8.4.6-nts-Win32-vs17-x64\php-cgi.exe`),
		mysql.New("5", `C:\flak\bin\mysql\mysql-5.7.43-winx64\bin\mysqld.exe`),
	}

	for _, service := range services {
		service.Start() // ← This is the interface in action!
	}

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down services...")

	for _, service := range services {
		service.Stop() // ← This is the interface in action!
	}

	// Remove PID files
	os.RemoveAll(service.PidDir)

	// Optional: Add graceful shutdown logic if supported by services
	time.Sleep(3 * time.Second)
}
