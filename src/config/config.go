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
	AuthString              string `env:"AUTH_STRING,required"`
	DefaultRandomUsers      int    `env:"DEFAULT_RANDOM_USERS,required"`
	DebugMode               bool   `env:"DEBUG_MODE"`
	DelayedUsersCronMinutes int    `env:"DELAYED_USERS_CRON_MINUTES,required"`
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
		defaultRandomUsersStr = "50"
	}
	defaultRandomUsers, err := strconv.Atoi(defaultRandomUsersStr)
	if err != nil {
		panic(err)
	}

	envConfig.DefaultRandomUsers = defaultRandomUsers

	// DelayedUsersCronMinutes
	delayedUsersCronMinutesStr := os.Getenv("DELAYED_USERS_CRON_MINUTES")
	if delayedUsersCronMinutesStr == "" {
		delayedUsersCronMinutesStr = "5"
	}
	delayedUsersCronMinutes, err := strconv.Atoi(delayedUsersCronMinutesStr)
	if err != nil {
		panic(err)
	}
	envConfig.DelayedUsersCronMinutes = delayedUsersCronMinutes

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
