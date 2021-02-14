package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"

	"github.com/dblogcorp/dbloggy/internal/utils"
	"github.com/dblogcorp/dbloggy/internal/utils/log"
)

func InitConfig(cfgFile string, env, prev EnvConfig) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("./config")
	}

	// 读取配置文件
	if err := readConfigFile(env, prev); err != nil {
		panic(err)
	}

	// 初始化日志
	log.InitLog(env.GetLogPath(), zapcore.InfoLevel)
	log.Infof("Using config file: %s", viper.ConfigFileUsed())
	log.Infof("Config: %+v", env)

	// 监听配置文件变化
}

func readConfigFile(env, prev EnvConfig) error {
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 解析并序列化配置对象
	data, err := yaml.Marshal(viper.AllSettings())
	if err != nil {
		return err
	}

	// 保存 prev 对象
	err = utils.DeepCopy(env, prev)
	if err != nil {
		log.Warnf("Save prev config failed, err: %s", err.Error())
	}
	if err = yaml.Unmarshal(data, env); err != nil {
		return err
	}
	return nil
}