package registry

type Registry struct {
	Schema        string `json:"$schema"`
	SchemaVersion string `json:"$schemaVersion"`
	Php           []struct {
		ID         string      `json:"id"`
		Version    string      `json:"version"`
		Type       string      `json:"type"`
		Arch       string      `json:"arch"`
		Vc         string      `json:"vc"`
		Sha256     interface{} `json:"sha256"`
		Homepage   string      `json:"homepage"`
		URL        string      `json:"url"`
		ExtractDir string      `json:"extractDir"`
		Executable string      `json:"executable"`
		Args       []string    `json:"args"`
		License    string      `json:"license"`
		Origin     string      `json:"origin"`
	} `json:"php"`
	Mysql []struct {
		ID         string      `json:"id"`
		Version    string      `json:"version"`
		Sha256     interface{} `json:"sha256"`
		Homepage   string      `json:"homepage"`
		URL        string      `json:"url"`
		Executable string      `json:"executable"`
		ExtractDir string      `json:"extractDir"`
		DataDir    string      `json:"dataDir"`
		Initialize struct {
			InitDataFolder []string `json:"initDataFolder"`
		} `json:"initialize"`
		Args    []string `json:"args"`
		License string   `json:"license"`
		Origin  string   `json:"origin"`
	} `json:"mysql"`
	Nginx []struct {
		ID         string      `json:"id"`
		Version    string      `json:"version"`
		Sha256     interface{} `json:"sha256"`
		Homepage   string      `json:"homepage"`
		URL        string      `json:"url"`
		Executable string      `json:"executable"`
		ExtractDir string      `json:"extractDir"`
		Args       []string    `json:"args"`
		License    string      `json:"license"`
		Origin     string      `json:"origin"`
	} `json:"nginx"`
}
