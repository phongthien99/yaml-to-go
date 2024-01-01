package types // LrmApi

type LrmApi struct {
	BasePath string `yaml:"base_path"`
}

// Http
type Http struct {
	BasePath string `yaml:"base_path"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
}

// EnvironmentVariable
type EnvironmentVariable struct {
	Origin Origin `mapstructure:"origin"`
	LrmApi LrmApi `mapstructure:"lrm_api"`
	Http   Http   `mapstructure:"http"`
}

// Origin
type Origin struct {
	BasePath string `yaml:"base_path"`
}
