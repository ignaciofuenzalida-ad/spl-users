package service

import (
	"fmt"
	"spl-users/ent"
	"spl-users/ent/schema"
	"spl-users/src/config"
	"spl-users/src/dto"
	"spl-users/src/model"
	"spl-users/src/repository"
	"time"
)

type UserService struct {
	userRepository *repository.UserRepository
	config         *config.EnvironmentConfig
}

func NewUserService(
	userRepository *repository.UserRepository,
	config *config.EnvironmentConfig,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		config:         config,
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
	runs, err := u.userRepository.GetRandomUsers(u.config.DefaultRandomUsers)
	if err != nil {
		return nil
	}
	return runs
}

// DEPRECATED
// func (u *UserService) PushRandomUsers() error {
// 	for {

// 		totalElements := len(u.runsQueue.Values)
// 		if totalElements >= u.config.DefaultRandomUsersQueueSize {
// 			fmt.Printf("Good total elements: %d\n", totalElements)
// 			continue
// 		}
// 		fmt.Printf("Queue is below limit: %d, fetching more random users...\n", totalElements)
// 		runs, err := u.userRepository.GetRandomUsers(u.config.DefaultRandomUsersQueueSize)
// 		if err != nil {
// 			return err
// 		}

// 		u.runsQueue.PushMany(*runs)
// 	}
// }

func (u *UserService) UpdateOrCreateUser(run int, data dto.UpdateUserDto) error {
	if data.FetchStatus == string(schema.ERROR) {
		err := u.userRepository.SetUserQueueError(run)
		if err != nil {
			return err
		}
		return nil
	}

	if data.Status == string(schema.NOT_FOUND) {
		err := u.userRepository.SetUserQueueNotFound(run)
		if err != nil {
			return err
		}
		return nil

	}

	if data.Gender == "" {
		data.Gender = string(schema.UNKNOWN)
	}

	err := u.userRepository.UpdateOrCreateUser(run, data)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetQueueUsersStatistics() (*model.QueueUsersStatistics, error) {
	data, err := u.userRepository.GetQueueUsersStatistics()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) CheckDelayedUsers() {
	for {
		affected, err := u.userRepository.UpdateDelayedUsers(u.config.DelayedUsersCronMinutes)
		if err != nil {
			fmt.Printf("Error during Delayed Users Cron: %s\n", err)
		}

		fmt.Printf("[CRON] Users with delayed PENDING: %d, status updated to WAITING.\n", affected)
		time.Sleep(time.Minute * time.Duration(u.config.DelayedUsersCronMinutes))
	}

}
