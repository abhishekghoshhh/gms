package config

import (
	"fmt"
	"os"

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
	config := load()
	return &Config{
		Viper: config,
	}
}
func load() *viper.Viper {
	workingDir, _ := os.Getwd()
	resourcesDir := workingDir + "/" + RESOURCES_DIR
	return loadConfig(resourcesDir, ROOT_CONFIG)

}
func loadConfig(resourcesDir string, configName string) *viper.Viper {
	c := viper.New()
	c.SetConfigName(configName)
	c.SetConfigType("yaml")
	c.AddConfigPath(resourcesDir)
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return c
}
