package pkg

import (
	consts "GrpcMessangerAccServer/pkg/constants"
	"gopkg.in/yaml.v3"
	"os"
	"runtime"
)

type Config struct {
	Port      string `yaml:"Port"`
	IPAddress string `yaml:"IPAddress"`
}

func ReadConfig() *Config {
	var ConfPath string
	if runtime.GOOS == "windows" {
		ConfPath = consts.ConfigPathWindows
	} else if runtime.GOOS == "linux" {
		ConfPath = consts.ConfigPathLinux
	} else {
		ConfPath = consts.ConfigPathLinux // Ну там, дааа, пока что так
	}
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
