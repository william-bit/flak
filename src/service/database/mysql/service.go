package mysql

import (
	"flak/src/service"
	"log"
	"os"
	"os/exec"
)

type MySQLService struct {
	Version string
	Port    int
	Dir     string
	Cmd     *exec.Cmd
}

func New(version, dir string) *MySQLService {
	return &MySQLService{
		Version: version,
		Dir:     dir,
	}
}

// initMySQLDataFolder initializes the MySQL data dir if needed
func (mysql *MySQLService) initMySQLDataFolder(dataDir string) error {
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

func (mysql *MySQLService) Start() {
	// Start MySQL (Update path accordingly)
	if _, err := service.LoadPID("mysql"); err != nil || !service.RunningService(1234) {
		mysql.initMySQLDataFolder(`C:\flak\data\mysql`)
		mysql.Cmd, err = service.StartService("mysql", mysql.Dir, `--console`, `--log_syslog=0`, `--datadir=C:\flak\data\mysql`)
		if err != nil {
			log.Fatalf("Failed to start MySQL: %v", err)
		} else {
			log.Printf("mysqlCmd.Process.Pid: %d", mysql.Cmd.Process.Pid)
		}
	}

}

func (mysql *MySQLService) Stop() {
	service.ShutdownService("MySQL", mysql.Cmd)
}

func (mysql *MySQLService) Status() string {
	if mysql.Cmd != nil && mysql.Cmd.Process != nil {
		return "running"
	}
	return "stopped"
}
