package service

import (
	"flak/src/config"
	"log"
	"os"
	"os/exec"
	"strings"
)

type SetupService struct {
	root    string
	service config.Component
	Cmd     *exec.Cmd
}

func New(root string, service config.Component) *SetupService {
	service.Executable = root + "/" + service.ExtractDir + "/" + service.Executable
	service.Args = replaceInStringArray(service.Args, "${root}", root)
	return &SetupService{
		root:    root,
		service: service,
	}
}

// initMySQLDataFolder initializes the MySQL data dir if needed
func (setup *SetupService) initDataFolder(dataDir string) error {
	_, err := os.Stat(dataDir)
	if err != nil {
		cmd := exec.Command(setup.service.Executable, setup.service.Initialize.InitDataFolder...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		log.Println("Initializing " + setup.service.ID + " data directory...")
		return cmd.Run() // This waits until the process finishes
	}
	return nil
}

func replaceInStringArray(data []string, old string, new string) []string {
	for key, value := range data {
		data[key] = strings.ReplaceAll(value, old, new)
	}
	return data
}

func (setup *SetupService) Start() {
	// Start MySQL (Update path accordingly)
	if pid, err := LoadPID(setup.service.ID); err != nil || !RunningService(pid) {
		if setup.service.DataDir != "" {
			dataDir := setup.root + "/" + setup.service.DataDir
			setup.service.Initialize.InitDataFolder = replaceInStringArray(setup.service.Initialize.InitDataFolder, "${dataDir}", dataDir)
			setup.service.Args = replaceInStringArray(setup.service.Args, "${dataDir}", dataDir)
			setup.initDataFolder(dataDir)
		}
		setup.Cmd, err = StartService(setup.service.ID, setup.service.Executable, setup.service.Args...)
		if err != nil {
			log.Fatalf("Failed to start %s : %v", setup.service.ID, err)
		} else {
			log.Printf(setup.service.ID+".Process.Pid: %d", setup.Cmd.Process.Pid)
		}
	}

}

func (setup *SetupService) Stop() {
	ShutdownService(setup.service.ID, setup.Cmd)
}

func (setup *SetupService) Status() string {
	if setup.Cmd != nil && setup.Cmd.Process != nil {
		return "running"
	}
	return "stopped"
}
