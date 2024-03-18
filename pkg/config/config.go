package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/abhishekghoshhh/gms/pkg/logger"

	"github.com/spf13/viper"
)

const (
	APP           = "app"
	RESOURCES_DIR = "resources"
	ROOT_CONFIG   = "app.yaml"
	APP_PROFILES  = "app.profiles"
)

type Config struct {
	resolvedConfig *viper.Viper
}

func (config *Config) Get(key string) any {
	return config.resolvedConfig.Get(key)
}
func (config *Config) GetString(key string) string {
	return config.resolvedConfig.GetString(key)
}

func New() *Config {
	config := load()
	return &Config{
		resolvedConfig: config,
	}
}
func load() *viper.Viper {
	workingDir, _ := os.Getwd()
	resourcesDir := workingDir + "/" + RESOURCES_DIR
	if _, err := os.Stat(resourcesDir + "/" + ROOT_CONFIG); errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	resolvedConfig := viper.New()
	root := loadConfig(resourcesDir, ROOT_CONFIG)
	addToConfig(resolvedConfig, root)
	if !resolvedConfig.IsSet(APP_PROFILES) {
		return resolvedConfig
	}
	allProfileNames := get(root, APP_PROFILES).(string)
	profiles := strings.Split(allProfileNames, ",")
	if len(profiles) == 0 {
		return resolvedConfig
	}
	for _, profile := range profiles {
		otherConfigFileName := "/app-" + profile + ".yaml"
		if _, err := os.Stat(resourcesDir + otherConfigFileName); !errors.Is(err, os.ErrNotExist) {
			otherConfig := loadConfig(resourcesDir, otherConfigFileName)
			addToConfig(resolvedConfig, otherConfig)
		}
	}
	logger.Info("lodading profiles : " + strings.Join(profiles, ","))
	return resolvedConfig
}
func loadConfig(resourcesDir string, configname string) *viper.Viper {
	c := viper.New()
	c.SetConfigName(configname)
	c.SetConfigType("yaml")
	c.AddConfigPath(resourcesDir)
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return c
}
func addToConfig(resolvedConfig *viper.Viper, config *viper.Viper) {
	keys := config.AllKeys()
	for _, key := range keys {
		val := get(config, key)
		resolvedConfig.Set(key, val)
	}
}

func get(config *viper.Viper, key string) any {
	val := config.Get(key)
	if stringVal, ok := val.(string); ok && strings.HasPrefix(stringVal, "${") && strings.HasSuffix(stringVal, "}") {
		val, _ = strings.CutPrefix(stringVal, "${")
		val, _ = strings.CutSuffix(stringVal, "}")
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
