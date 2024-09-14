package service

import (
	"fmt"
	"spl-users/ent/schema"
	"spl-users/src/config"
	"spl-users/src/dto"
	"spl-users/src/queue"
	"spl-users/src/repository"
)

type QueueService struct {
	userRepository  *repository.UserRepository
	config          *config.EnvironmentConfig
	runsQueue       *queue.Queue[string]
	userUpdateQueue *queue.Queue[dto.UpdateUserDto]
}

func NewQueueService(
	userRepository *repository.UserRepository,
	config *config.EnvironmentConfig,
	runsQueue *queue.Queue[string],
	userUpdateQueue *queue.Queue[dto.UpdateUserDto],
) *QueueService {
	return &QueueService{
		userRepository:  userRepository,
		config:          config,
		runsQueue:       runsQueue,
		userUpdateQueue: userUpdateQueue,
	}
}

func (q *QueueService) Run() {
	for {
		// RunsQueue
		totalElements := len(q.runsQueue.Values)
		if totalElements <= q.config.DefaultRandomUsersQueueSize {
			fmt.Printf("Queue is below limit: %d, fetching more random users...\n", totalElements)
			runs, err := q.userRepository.GetRandomUsers(q.config.DefaultRandomUsersQueueSize)
			if err != nil {
				fmt.Println(err)
			} else {
				q.runsQueue.PushMany(runs)
			}
		}

		// UpdateOrCreateUserQueue
		if len(q.userUpdateQueue.Values) > 0 {
			err := q.UpdateOrCreateUser()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (q *QueueService) UpdateOrCreateUser() error {
	data := q.userUpdateQueue.Pop()

	if data.FetchStatus == string(schema.ERROR) {
		err := q.userRepository.SetUserQueueError(data.Run)
		if err != nil {
			return err
		}
		return nil
	}

	if data.Status == string(schema.NOT_FOUND) {
		err := q.userRepository.SetUserQueueNotFound(data.Run)
		if err != nil {
			return err
		}
		return nil

	}

	if data.Gender == "" {
		data.Gender = string(schema.UNKNOWN)
	}

	err := q.userRepository.UpdateOrCreateUser(data.Run, *data)
	if err != nil {
		return err
	}

	return nil

}
