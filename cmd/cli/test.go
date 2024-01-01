package types // EnvironmentVariable
type EnvironmentVariable struct {
	Http   Http   `mapstructure:"http"`
	Origin Origin `mapstructure:"origin"`
	LrmApi LrmApi `mapstructure:"lrm_api"`
}

// Http
type Http struct {
	Host     string `yaml:"host"`
	BasePath string `yaml:"base_path"`
	Port     int    `yaml:"port"`
}

// Origin
type Origin struct {
	BasePath string `yaml:"base_path"`
}

// LrmApi
type LrmApi struct {
	BasePath string `yaml:"base_path"`
}

