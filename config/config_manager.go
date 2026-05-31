package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type ConfigManager struct {
	configMutex sync.RWMutex
	config      *GoUnoConfig
}

func NewConfigManager(
	configPath string,
	env string,
) *ConfigManager {

	configManager := ConfigManager{}

	v := viper.New()
	configManager.setConfigDefaults(v)
	v.AddConfigPath(configPath)
	v.SetConfigName(env)
	v.SetConfigType("yaml")

	v.SetEnvPrefix("GOUNO")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("read config failed, err: %v", err)
	}

	newConfig := GoUnoConfig{}
	if err := v.Unmarshal(&newConfig); err != nil {
		log.Fatalf("unmarshal config failed, err: %v", err)
	}
	configManager.SetConfig(&newConfig)
	return &configManager
}

func (cm *ConfigManager) SetConfig(config *GoUnoConfig) {
	cm.configMutex.Lock()
	defer cm.configMutex.Unlock()
	cm.config = config
}

func (cm *ConfigManager) Config() GoUnoConfig {
	cm.configMutex.RLock()
	defer cm.configMutex.RUnlock()
	return *cm.config
}

func (cm *ConfigManager) setConfigDefaults(v *viper.Viper) {
	// 验证码配置
	v.SetDefault("captcha_type", "math")

	// Web服务器配置
	v.SetDefault("web_server.debug", false)
	v.SetDefault("web_server.address", "0.0.0.0")
	v.SetDefault("web_server.port", "8080")
	v.SetDefault("web_server.idle_timeout", "60s")
	v.SetDefault("web_server.read_timeout", "5s")
	v.SetDefault("web_server.read_header_timeout", "2s")
	v.SetDefault("web_server.write_timeout", "30s")
	v.SetDefault("web_server.request_timeout", "10s")
	v.SetDefault("web_server.rate_limit_per_minute", 100)

	// 数据库配置
	v.SetDefault("database.default", "sqlite")
	v.SetDefault("database.drivers.sqlite.name", "sqlite")
	v.SetDefault("database.drivers.sqlite.driver", "sqlite3")
	v.SetDefault("database.drivers.sqlite.dsn", ":memory:")
	v.SetDefault("database.drivers.sqlite.log_level", 1)
}
