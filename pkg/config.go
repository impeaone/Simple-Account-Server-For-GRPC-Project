package pkg

import (
	consts "GrpcMessangerAccServer/pkg/constants"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port      string `yaml:"Port"` // Порт: перемножение букв в ascii деленное на количество букв умнож. на 10 (auth -> 13961)
	IPAddress string `yaml:"IPAddress"`
}

func ReadConfig() *Config {
	ConfPath := consts.ConfigPath
	bytes, err := os.ReadFile(ConfPath)
	if err != nil {
		panic(err)
	}
	var config Config
	errUnmarshal := yaml.Unmarshal(bytes, &config)
	if errUnmarshal != nil {
		panic(err)
	}
	return &config
}
