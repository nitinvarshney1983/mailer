package configs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configArgs struct {
	configFilePath []string
	configfileName string
}

const defaultConfigFileName = "appconfigs"

var (
	defaultConfigFilePath = []string{".configs", "$HOME/configs"}
	viperInstance         *viper.Viper
)

func setUp(args *configArgs) {
	viperInstance = viper.New()

	filePath := defaultConfigFilePath
	fileName := defaultConfigFileName

	if args != nil {
		if args.configfileName != "" {
			fileName = args.configfileName
		}
		if args.configFilePath != nil && len(args.configFilePath) > 0 {
			filePath = args.configFilePath
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
