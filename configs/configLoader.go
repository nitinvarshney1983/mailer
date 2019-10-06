package configs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ConfigArgs for configurations
type ConfigArgs struct {
	ConfigFilePath []string
	ConfigfileName string
}

const defaultConfigFileName = "appconfigs"

var (
	defaultConfigFilePath = []string{".configs", "$HOME/configs"}
	viperInstance         *viper.Viper
)

// SetUp for configurations
func SetUp(args *ConfigArgs) {
	viperInstance = viper.New()

	filePath := defaultConfigFilePath
	fileName := defaultConfigFileName

	if args != nil {
		if args.ConfigfileName != "" {
			fileName = args.ConfigfileName
		}
		if args.ConfigFilePath != nil && len(args.ConfigFilePath) > 0 {
			filePath = args.ConfigFilePath
		}
	}

	if os.Getenv("CONF_NAME") != "" {
		fileName = os.Getenv("CONF_NAME")
	}

	if os.Getenv("CONF_PATH") != "" {
		filePath = strings.Split(os.Getenv("CONF_PATH"), ",")
	}
	viperInstance.SetConfigName(fileName)

	for _, filePathLocation := range filePath {
		viperInstance.AddConfigPath(filePathLocation)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viperInstance.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

// Get func to the configurations
func Get(key string) interface{} {
	return viperInstance.Get(key)
}
