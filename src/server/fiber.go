package server

import (
	"context"
	"spl-users/src/controller"
	"spl-users/src/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"go.uber.org/fx"
)

func CreateFiberServer(
	lc fx.Lifecycle,
	userController *controller.UserController,
	authMiddleware *middleware.AuthMiddleware,
) {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	// Setup routes
	app.Get("/health", userController.GetUsers)
	app.Get("/api/users", authMiddleware.ValidateAuthHeader, userController.GetUsers)
	app.Get("/api/users/random", authMiddleware.ValidateAuthHeader, userController.GetRandomUsers)
	app.Get("/api/users/statistics", authMiddleware.ValidateAuthHeader, userController.GetQueueUsersStatistics)
	app.Get("/api/users/:run", authMiddleware.ValidateAuthHeader, userController.GetUserByRun)
	app.Post("/api/users/:run", authMiddleware.ValidateAuthHeader, userController.UpdateOrCreateUserByRun)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// TODO: switch the port to an env variable
			go app.Listen(":30001")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
