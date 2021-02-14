package sso

import "github.com/dblogcorp/dbloggy/config"

var (
	envConfig   *EnvConfig
)

func GetConfig() *EnvConfig {
	return envConfig
}

type EnvConfig struct {
	LogPath   string `yaml:"log_path"`
	ListenPort string `yaml:"listen_port"`

	prevConfig *EnvConfig
}

func (env *EnvConfig) GetLogPath() string {
	return env.LogPath
}

func (env *EnvConfig) GetListenPort() string {
	return env.ListenPort
}

func (env *EnvConfig) Prev() *EnvConfig {
	return env.prevConfig
}

func (env *EnvConfig) SetPrev(prev config.EnvConfig) {
	env.prevConfig = prev.(*EnvConfig)
}
