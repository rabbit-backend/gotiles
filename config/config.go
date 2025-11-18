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

type CacheConfig struct {
	Type       string `json:"type"`
	Connection string `json:"connection"`
}

type GoTilesConfig struct {
	Sources []Source    `json:"sources"`
	Cache   CacheConfig `json:"cache"`
}
