package main

import (
	"spl-users/src/config"
	"spl-users/src/controller"
	"spl-users/src/db"
	"spl-users/src/middleware"
	"spl-users/src/queue"
	"spl-users/src/repository"
	"spl-users/src/server"
	"spl-users/src/service"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Setup environment config
		fx.Provide(config.NewEnviromentConfig),
		// Setup background context
		fx.Provide(config.NewContextBackground),
		// Setup auth middleware
		fx.Provide(middleware.NewAuthMiddleware),
		// Setup database
		fx.Provide(db.CreateSqliteConnection),
		// Setup repositories
		fx.Provide(repository.NewUserRepository),
		fx.Provide(repository.NewLocationRepository),
		// Setup services
		fx.Provide(service.NewUserService),
		fx.Provide(service.NewQueueService),
		// Setup controllers
		fx.Provide(controller.NewUserController),
		// Setup Queue
		fx.Provide(queue.NewMapQueue[string]),
		fx.Invoke(func(q *service.QueueService) { go q.Run() }),
		// Setup Fiber server
		fx.Invoke(server.CreateFiberServer),
	).Run()
}
