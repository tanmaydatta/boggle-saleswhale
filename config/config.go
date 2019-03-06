package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func valid(name string) bool {
	name = strings.ToUpper(name)
	return name == "PROD" || name == "DEV" || name == "LOCAL"
}

func addConfigPaths() {
	configPath := "$GOPATH/src/github.com/tanmaydatta/boggle/config/"
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(configPath + "/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		logrus.WithError(err).Fatal("Fatal error config file")
	}
}

func bindEnv() {
	err := viper.BindEnv("env")

	if err != nil {
		logrus.WithError(err).Fatal("env load failed")
	}
}

func validateEnv(configName string) {
	logrus.WithField("configName", configName).Info()
	if !valid(configName) {
		logrus.Fatal("Please set environment variable ENV as PROD, LOCAL or DEV as applicable")
	}
}

func detectEnvVariablesDynamic() {
	viper.AutomaticEnv()
}

func validateConfig() {
	dictionary := viper.GetString("dictionary_path")
	board := viper.GetString("board_path")
	if dictionary == "" || board == "" {
		logrus.Fatal("dictionary_path and/or board_path is missing from config")
	}
	if _, err := os.Stat(dictionary); err != nil {
		logrus.Fatal("dictionary_path is not valid", err)
	}
	if _, err := os.Stat(board); err != nil {
		logrus.Fatal("board_path is not valid", err)
	}
}

func Load() {
	bindEnv()
	configName := strings.ToLower(viper.GetString("env"))
	if configName == "" {
		configName = "local"
	}
	validateEnv(configName)
	detectEnvVariablesDynamic()
	viper.SetConfigName(configName)
	addConfigPaths()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	viper.SetDefault("SERVER_PORT", port)
	validateConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("Config file changed:", e.Name)
	})
}


