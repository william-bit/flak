package config

type Component struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Sha256     string   `json:"sha256"`
	Homepage   string   `json:"homepage"`
	URL        string   `json:"url"`
	ExtractDir string   `json:"extractDir"`
	Executable string   `json:"executable"`
	Args       []string `json:"args"`
	License    string   `json:"license"`
	Origin     string   `json:"origin"`
	DataDir    string   `json:"dataDir"`
	Initialize struct {
		InitDataFolder []string `json:"initDataFolder"`
	} `json:"initialize"`
}

type MysqlInitialize struct {
	InitDataFolder []string `json:"initDataFolder"`
}

type Config struct {
	Root       string      `json:"root"`
	AutoStart  bool        `json:"autoStart"`
	AutoUpdate bool        `json:"autoUpdate"`
	ServerMode bool        `json:"serverMode"`
	Service    []Component `json:"service"`
}
