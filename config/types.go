package config

type EnvConfig interface {
	GetLogPath() string

	GetListenPort() string

	SetPrev(prev EnvConfig)
}
