package db

import (
	"context"
	"fmt"
	"spl-users/ent"
	"spl-users/src/config"

	"go.uber.org/fx"
)

func CreateSqliteConnection(lc fx.Lifecycle, envConfig *config.EnvironmentConfig, ctx *context.Context) *ent.Client {
	options := []ent.Option{}
	if envConfig.DebugMode {
		options = append(options, ent.Debug())
	}
	conn, connError := ent.Open("sqlite3", "file:./database/sportlife.db?_fk=1", options...)
	// Run the auto migration tool.
	migrationError := conn.Schema.Create(*ctx)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if connError != nil {
				fmt.Printf("Failed opening connection to sqlite: %v", connError)
				return connError
			}
			if migrationError != nil {
				fmt.Printf("Failed to execute the migration tool: %v", migrationError)
				return migrationError
			}
			return nil
		},
		OnStop: func(context.Context) error {
			conn.Close()
			return nil
		},
	})

	return conn
}
