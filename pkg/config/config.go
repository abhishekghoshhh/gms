package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	RESOURCES_DIR = "resources"
	ROOT_CONFIG   = "app.yaml"
)

type Config struct {
	*viper.Viper
}

func (*Config) FromEnv(key string) string {
	return os.Getenv(key)
}

func (*Config) FromEnvOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultVal
}
func (config *Config) FromEnvOrConfig(envKey, configKey string) string {
	val := os.Getenv(envKey)
	if val != "" {
		return val
	}
	return config.GetString(configKey)
}

func New() *Config {
	workingDir, _ := os.Getwd()
	config := loadConfig(workingDir+"/"+RESOURCES_DIR, ROOT_CONFIG)
	return &Config{
		Viper: config,
	}
}
func loadConfig(resourcesDir string, configName string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(configName)
	config.SetConfigType("yaml")
	config.AddConfigPath(resourcesDir)
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	keys := config.AllKeys()
	for _, key := range keys {
		val := get(config, key)
		config.Set(key, val)
	}
	return config
}
func get(config *viper.Viper, key string) any {
	val := config.Get(key)
	if stringVal, ok := val.(string); ok && strings.HasPrefix(stringVal, "${") && strings.HasSuffix(stringVal, "}") {
		stringVal, _ := strings.CutPrefix(stringVal, "${")
		stringVal, _ = strings.CutSuffix(stringVal, "}")
		parts := strings.Split(stringVal, ":")
		valFromEnv := os.Getenv(parts[0])
		if valFromEnv == "" && len(parts) == 1 {
			panic("no default value for key " + key + ", no environment variable set for " + parts[0])
		}
		if valFromEnv != "" {
			return valFromEnv
		}
		return parts[1]
	}
	return val
}
