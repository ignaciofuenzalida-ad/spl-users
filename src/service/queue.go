package service

import (
	"fmt"
	"spl-users/src/config"
	"spl-users/src/queue"
	"spl-users/src/repository"
	"time"
)

type QueueService struct {
	userRepository *repository.UserRepository
	config         *config.EnvironmentConfig
	runsQueue      *queue.MapQueue[string]
}

func NewQueueService(
	userRepository *repository.UserRepository,
	config *config.EnvironmentConfig,
	runsQueue *queue.MapQueue[string],
) *QueueService {
	return &QueueService{
		userRepository: userRepository,
		config:         config,
		runsQueue:      runsQueue,
	}
}

func (q *QueueService) Run() {
	for {
		// RunsQueue
		totalElements := len(q.runsQueue.Values)
		if totalElements < q.config.DefaultRandomUsers*30 {
			fmt.Printf("[QUEUE] Below limits: %d, fetching more users...\n", totalElements)
			runs, err := q.userRepository.GetRandomUsers(q.config.DefaultRandomUsers * 30)
			if err != nil {
				fmt.Printf("[QUEUE] Error during users query: %s\n", err)
			}

			q.runsQueue.PushMany(*runs)

		} else {
			fmt.Printf("[QUEUE] Full, waiting...\n")
			time.Sleep(1 * time.Second)
		}

	}
}

// func (q *QueueService) UpdateOrCreateUser() error {
// 	data := q.userUpdateQueue.Pop()

// 	if data.FetchStatus == string(schema.ERROR) {
// 		err := q.userRepository.SetUserQueueError(data.Run)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}

// 	if data.Status == string(schema.NOT_FOUND) {
// 		err := q.userRepository.SetUserQueueNotFound(data.Run)
// 		if err != nil {
// 			return err
// 		}
// 		return nil

// 	}

// 	if data.Gender == "" {
// 		data.Gender = string(schema.UNKNOWN)
// 	}

// 	err := q.userRepository.UpdateOrCreateUser(data.Run, *data)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
