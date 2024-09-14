package config

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type EnvironmentConfig struct {
	AuthString                  string `env:"AUTH_STRING,required"`
	DefaultRandomUsers          int    `env:"DEFAULT_RANDOM_USERS,required"`
	DefaultRandomUsersQueueSize int    `env:"DEFAULT_RANDOM_USERS_QUEUE_SIZE"`
	DebugMode                   bool   `env:"DEBUG_MODE"`
}

var envConfig *EnvironmentConfig

func NewEnviromentConfig(lc fx.Lifecycle) *EnvironmentConfig {
	envConfig = &EnvironmentConfig{}

	var err error
	isLocalhost := os.Getenv("IS_LOCALHOST")
	if isLocalhost != "true" {
		err = godotenv.Load()
	}

	if err != nil {
		panic(err)
	}
	// AuthString
	envConfig.AuthString = os.Getenv("AUTH_STRING")

	// DebugMode
	if os.Getenv("DEBUG_MODE") == "true" {
		envConfig.DebugMode = true
	} else {
		envConfig.DebugMode = false
	}

	// DefaultRandomUsers
	defaultRandomUsersStr := os.Getenv("DEFAULT_RANDOM_USERS")
	if defaultRandomUsersStr == "" {
		defaultRandomUsersStr = "10"
	}
	defaultRandomUsers, err := strconv.Atoi(defaultRandomUsersStr)
	if err != nil {
		panic(err)
	}
	envConfig.DefaultRandomUsers = defaultRandomUsers

	// DefaultRandomUsersQueueSize
	defaultRandomUsersQueueSizeStr := os.Getenv("DEFAULT_RANDOM_USERS_QUEUE_SIZE")
	if defaultRandomUsersQueueSizeStr == "" {
		defaultRandomUsersQueueSizeStr = "500"
	}
	defaultRandomUsersQueueSize, err := strconv.Atoi(defaultRandomUsersQueueSizeStr)
	if err != nil {
		panic(err)
	}
	envConfig.DefaultRandomUsersQueueSize = defaultRandomUsersQueueSize

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if err != nil {
				return err
			}
			return nil
		},
	})

	printEnvironmentConfig(*envConfig)
	return envConfig
}

func printEnvironmentConfig(config EnvironmentConfig) {
	v := reflect.ValueOf(config)
	typeOfConfig := v.Type()

	fmt.Println("EnvironmentConfig:")
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("  %s: %v\n", typeOfConfig.Field(i).Name, v.Field(i).Interface())
	}
}

func NewContextBackground() *context.Context {
	ctx := context.Background()
	return &ctx
}
