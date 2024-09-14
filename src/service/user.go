package service

import (
	"fmt"
	"spl-users/ent"
	"spl-users/src/config"
	"spl-users/src/dto"
	"spl-users/src/model"
	"spl-users/src/queue"
	"spl-users/src/repository"
)

type UserService struct {
	userRepository  *repository.UserRepository
	config          *config.EnvironmentConfig
	runsQueue       *queue.Queue[string]
	userUpdateQueue *queue.Queue[dto.UpdateUserDto]
}

func NewUserService(
	userRepository *repository.UserRepository,
	config *config.EnvironmentConfig,
	runsQueue *queue.Queue[string],
	userUpdateQueue *queue.Queue[dto.UpdateUserDto],
) *UserService {
	return &UserService{
		userRepository:  userRepository,
		config:          config,
		runsQueue:       runsQueue,
		userUpdateQueue: userUpdateQueue,
	}
}

func (u *UserService) GetAllUsers(search string, limit int, locations []string) ([]*model.UserSearch, error) {
	users, err := u.userRepository.GetUsersBySearch(search, limit, locations)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) GetUserByRun(run int) (*model.User, error) {
	user, err := u.userRepository.GetUserByRun(run)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUserQueueByRun(run int) (*ent.UserQueue, error) {
	user, err := u.userRepository.GetUserQueueByRun(run)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetRandomUsers() *[]string {
	return u.runsQueue.PopMany(u.config.DefaultRandomUsers)
}

func (u *UserService) PushRandomUsers() error {
	for {

		totalElements := len(u.runsQueue.Values)
		if totalElements >= u.config.DefaultRandomUsersQueueSize {
			fmt.Printf("Good total elements: %d\n", totalElements)
			continue
		}
		fmt.Printf("Queue is below limit: %d, fetching more random users...\n", totalElements)
		runs, err := u.userRepository.GetRandomUsers(u.config.DefaultRandomUsersQueueSize)
		if err != nil {
			return err
		}

		u.runsQueue.PushMany(runs)
	}
}

func (u *UserService) UpdateOrCreateUser(run int, data dto.UpdateUserDto) error {
	data.Run = run
	u.userUpdateQueue.Push(data)

	return nil
}

func (u *UserService) GetQueueUsersStatistics() (*model.QueueUsersStatistics, error) {
	data, err := u.userRepository.GetQueueUsersStatistics()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) CheckDelayedUsers() error {
	affected, err := u.userRepository.UpdateDelayedUsers()
	if err != nil {
		return err
	}
	fmt.Printf("Users with delayed PENDING: %d, status updated to WAITING.\n", affected)

	return nil
}
