package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfig() {
	if err := viper.UnmarshalKey("cache", &cacheConfig); err!= nil {
		log.Fatal(err)
	}

	if err := viper.UnmarshalKey("redis", &redisConfig); err!= nil {
		log.Fatal(err)
	}

	if err := viper.UnmarshalKey("general", &generalConfig); err!= nil {
		log.Fatal(err)
	}
}

func initCmd(cmd *cobra.Command) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err!= nil {
		if _, ok := err.(viper.ConfigFileNotFoundError);!ok {
			log.Fatal(err)
		}
		viper.SetConfigFile("config.yml")
		if err := viper.ReadInConfig(); err!= nil {
			log.Fatal(err)
		}
	}

	initConfig()

	cmd.Flags().StringP("config", "c", "config.yml", "config file")
	cmd.Flags().BoolP("help", "h", false, "show help")
}

func getCacheConfig() *CacheConfig {
	return &cacheConfig
}

func getRedisConfig() *RedisConfig {
	return &redisConfig
}

func getGeneralConfig() *GeneralConfig {
	return &generalConfig
}

func getEnvVar(key string, def string) string {
	if value := os.Getenv(key); value!= "" {
		return value
	}
	return def
}

func getBoolEnvVar(key string, def bool) bool {
	if value := getEnvVar(key, "false"); value!= "" {
		if value == "true" {
			return true
		}
	}
	return def
}

func getStrSliceEnvVar(key string, def []string) []string {
	if value := getEnvVar(key, ""); value!= "" {
		return strings.Split(value, ",")
	}
	return def
}

func getStrMapEnvVar(key string, def map[string]string) map[string]string {
	if value := getEnvVar(key, ""); value!= "" {
		m := make(map[string]string)
		for _, pair := range strings.Split(value, ",") {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
		return m
	}
	return def
}

func getStrSliceFlag(cmd *cobra.Command, key string, def []string) []string {
	if value, err := cmd.Flags().GetStringSlice(key); err == nil {
		return value
	}
	return def
}

func getStrMapFlag(cmd *cobra.Command, key string, def map[string]string) map[string]string {
	if value, err := cmd.Flags().GetStringMapStringSlice(key); err == nil {
		m := make(map[string]string)
		for k, v := range value {
			m[k] = v[0]
		}
		return m
	}
	return def
}

func getBoolFlag(cmd *cobra.Command, key string, def bool) bool {
	if value, err := cmd.Flags().GetBool(key); err == nil {
		return value
	}
	return def
}

type CacheConfig struct {
	TTL int `mapstructure:"ttl"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type GeneralConfig struct {
	LogLevel string `mapstructure:"log_level"`
}

var (
	cacheConfig CacheConfig
	redisConfig  RedisConfig
	generalConfig GeneralConfig
)