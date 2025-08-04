package main

import (
	"flak/src/service"
	"flak/src/service/database"
	"flak/src/service/runtime"
	"flak/src/service/webserver"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Try to reconnect to existing services
	service.ReconnectToExistingServices()

	webserver.StartNginx()
	runtime.StartPHP()
	database.StartMySql()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down services...")

	webserver.StopNginx()
	runtime.StopPHP()
	database.StopMySql()

	// Remove PID files
	os.RemoveAll(service.PidDir)

	// Optional: Add graceful shutdown logic if supported by services
	time.Sleep(3 * time.Second)
}
