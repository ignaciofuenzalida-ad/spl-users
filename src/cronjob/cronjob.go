package cronjob

import (
	"context"
	"fmt"
	"spl-users/src/service"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

func NewUserDelayedCron(lc fx.Lifecycle, userService *service.UserService) *cron.Cron {
	c := cron.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go c.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			_, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()
			c.Stop()
			return nil
		},
	})

	return c
}

func RegisterUserDelayedCronJobs(c *cron.Cron, userService *service.UserService) {
	_, err := c.AddFunc("*/5 * * * *", func() {
		err := userService.CheckDelayedUsers()
		if err != nil {
			fmt.Printf("Error checking delayed users: %v", err)
		}
	})
	if err != nil {
		fmt.Printf("Error adding cron job: %v", err)
	}
}
