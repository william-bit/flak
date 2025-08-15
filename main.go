package main

import (
	"flak/src/config"
	"flak/src/service"
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
	data := config.LoadConfig()
	// initTUI(data)
	initServices(data)
}

func initTUI(data config.Config) {
	p := tea.NewProgram(tui.InitScreen(data), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initServices(data config.Config) { // Try to reconnect to existing services

	var services []service.Service = []service.Service{}
	for _, s := range data.Service {
		if (s.Type == "server" || s.Type == "database" || s.Type == "service") && s.AutoStart {
			services = append(services, service.New(data.Root, s))
		}
	}

	for _, s := range services {
		s.Start()
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
