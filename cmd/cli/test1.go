package types // EnvironmentVariable
type EnvironmentVariable struct {
	Origin Origin `mapstructure:"origin"`
	LrmApi LrmApi `mapstructure:"lrm_api"`
	Http   Http   `mapstructure:"http"`
}

// Origin
type Origin struct {
	BasePath string `mapstructure:"base_path"`
}

// LrmApi
type LrmApi struct {
	BasePath string `mapstructure:"base_path"`
}

// Http
type Http struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
}

