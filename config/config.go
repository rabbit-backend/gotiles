package config

type DBConnection struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Source struct {
	Name       string       `json:"name"`
	Connection DBConnection `json:"connection"`
	Type       string       `json:"type"`
}

type GoTilesConfig struct {
	Sources []Source `json:"sources"`
}
