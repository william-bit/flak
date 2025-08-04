package service

import (
	"os"
	"path/filepath"
	"strconv"
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

// Load PID from file
func LoadPID(name string) (int, error) {
	data, err := os.ReadFile(filepath.Join(PidDir, name+".pid"))
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(string(data))
	return pid, err
}
