package environment_variable

import "time"

type DatabaseConfig struct {
	Driver      string        `yaml:"driver"`
	Master      string        `yaml:"master"`
	MaxIdleTime time.Duration `yaml:"max_idle_time"`
	MaxLifeTime time.Duration `yaml:"max_life_time"`
	MaxIdleConn int           `yaml:"max_idle_conns"`
	MaxOpenConn int           `yaml:"max_open_conns"`
}

type Config struct {
	Env      string                    `yaml:"env"` // env: "local" | "staging" | "production"
	Database map[string]DatabaseConfig `yaml:"database"`
}
