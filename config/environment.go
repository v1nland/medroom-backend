package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const DURATION = time.Second * 5

func LoadSettings(squadName string, appName string) {
	if err := getConfig(); err != nil {
		panic("no config found")
	}

	if len(appName) == 0 {
		panic("no appName to run")
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func getConfig() error {
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key := pair[0]
		value := pair[1]

		// if value is a number
		number, err := strconv.Atoi(value)
		if err == nil {
			viper.Set(key, number)
			continue
		}

		// if value is a boolean
		boolean, err := strconv.ParseBool(value)
		if err == nil {
			viper.Set(key, boolean)
			continue
		}

		// if value is a string
		viper.Set(key, value)
	}

	return nil
}
