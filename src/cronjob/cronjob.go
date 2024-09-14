package cronjob

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

type CronJob struct {
	CheckDelayedUsers bool
	cron              *cron.Cron
}

func NewUserDelayedCron(lc fx.Lifecycle) *CronJob {
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

	return &CronJob{CheckDelayedUsers: false, cron: c}
}

func RegisterUserDelayedCronJobs(c *CronJob) {
	_, err := c.cron.AddFunc("*/2 * * * *", func() {
		c.CheckDelayedUsers = true
	})
	if err != nil {
		fmt.Printf("Error adding cron job: %v\n", err)
	}
}
