package main

import (
	"flak/src/service"
	"flak/src/service/database/mysql"
	"flak/src/service/runtime/php"
	"flak/src/service/webserver/nginx"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Try to reconnect to existing services
	service.ResumeServices()

	nginx.Start()
	php.Start()
	mysql.Start()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down services...")

	nginx.Stop()
	php.Stop()
	mysql.Stop()

	// Remove PID files
	os.RemoveAll(service.PidDir)

	// Optional: Add graceful shutdown logic if supported by services
	time.Sleep(3 * time.Second)
}
