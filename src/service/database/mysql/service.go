package mysql

import (
	"flak/src/service"
	"log"
	"os"
	"os/exec"
)

var mysqlCmd *exec.Cmd

// initMySQLDataFolder initializes the MySQL data dir if needed
func initMySQLDataFolder(dataDir string) error {
	_, err := os.Stat(dataDir)
	if err != nil {
		cmd := exec.Command(
			`C:\flak\bin\mysql\mysql-5.7.43-winx64\bin\mysqld.exe`,
			`--console`,
			"--initialize",
			"--user=mysql",
			"--datadir="+dataDir,
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		log.Println("Initializing MySQL data directory...")
		return cmd.Run() // This waits until the process finishes
	}
	return nil
}

func Start() {
	// Start MySQL (Update path accordingly)
	if _, err := service.LoadPID("mysql"); err != nil || !service.RunningProcess(1234) {
		initMySQLDataFolder(`C:\flak\data\mysql`)
		mysqlCmd, err = service.StartService("mysql", `C:\flak\bin\mysql\mysql-5.7.43-winx64\bin\mysqld.exe`, `--console`, `--log_syslog=0`, `--datadir=C:\flak\data\mysql`)
		if err != nil {
			log.Fatalf("Failed to start MySQL: %v", err)
		} else {
			log.Printf("mysqlCmd.Process.Pid: %d", mysqlCmd.Process.Pid)
		}
	}

}

func Stop() {
	service.ShutdownService("MySQL", mysqlCmd)
}
