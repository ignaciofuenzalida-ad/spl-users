package repository

import (
	"context"
	"fmt"
	"spl-users/ent"
	"spl-users/ent/location"
	"spl-users/ent/predicate"
	"spl-users/ent/schema"
	"spl-users/ent/user"
	"spl-users/ent/userqueue"
	"spl-users/src/dto"
	"spl-users/src/model"
	"time"

	"entgo.io/ent/dialect/sql"
)

type UserRepository struct {
	conn               *ent.Client
	ctx                *context.Context
	locationRepository *LocationRepository
}

func NewUserRepository(
	conn *ent.Client,
	ctx *context.Context,
	locationRepository *LocationRepository,
) *UserRepository {
	return &UserRepository{
		conn:               conn,
		ctx:                ctx,
		locationRepository: locationRepository,
	}
}

func (u *UserRepository) GetUsersBySearch(
	search string,
	limit int,
	locations []string,
) ([]*model.UserSearch, error) {
	predicates := []predicate.User{
		user.Or(
			user.FirstNameContainsFold(search),
			user.LastNameContainsFold(search),
		),
	}

	if len(locations) > 0 {
		predicates = append(predicates, user.HasLocationsWith(location.SlugIn(locations...)))
	}

	users, err := u.conn.User.
		Query().
		Select(
			user.FieldRun,
			user.FieldVerificationDigit,
			user.FieldFirstName,
			user.FieldLastName,
			user.FieldGender,
			user.FieldPlantType).
		Where(predicates...).
		Limit(limit).
		All(*u.ctx)
	if err != nil {
		return nil, err
	}

	return model.EntUsersToUserSearch(users), nil
}

func (u *UserRepository) GetUserByRun(run int) (*model.User, error) {
	user, err := u.conn.User.
		Query().
		Where(user.RunEQ(run)).
		WithLocations().
		First(*u.ctx)

	if ent.IsNotFound(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return model.EntUserToUser(user), nil
}

func (u *UserRepository) GetUserQueueByRun(run int) (*ent.UserQueue, error) {
	user, err := u.conn.UserQueue.
		Query().
		Where(userqueue.RunEQ(run)).
		First(*u.ctx)

	if ent.IsNotFound(err) {
		return user, nil
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) GetRandomUsers(limit int) ([]string, error) {
	tx, err := u.conn.Tx(*u.ctx)
	if err != nil {
		return nil, err
	}

	users, err := tx.UserQueue.Query().
		Where(userqueue.StatusEQ(userqueue.Status(schema.EMPTY))).
		Where(userqueue.FetchStatusEQ(userqueue.FetchStatus(schema.WAITING))).
		Order(sql.OrderByRand()).
		Limit(limit).
		All(*u.ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var userIdentifiers []string
	var usersRun []int
	for _, user := range users {
		userIdentifiers = append(
			userIdentifiers,
			fmt.Sprintf("%d-%s", user.Run, user.VerificationDigit),
		)
		usersRun = append(usersRun, user.Run)
	}

	_, err = tx.UserQueue.Update().
		Where(userqueue.RunIn(usersRun...)).
		SetFetchStatus(userqueue.FetchStatus(schema.PENDING)).
		Save(*u.ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return userIdentifiers, nil
}

func (u *UserRepository) SetUserQueueError(run int) error {
	_, err := u.conn.UserQueue.Update().
		Where(userqueue.RunEQ(run)).
		SetFetchStatus(userqueue.FetchStatus(schema.ERROR)).
		Save(*u.ctx)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) SetUserQueueNotFound(run int) error {
	_, err := u.conn.UserQueue.Update().
		Where(userqueue.RunEQ(run)).
		SetStatus(userqueue.Status(schema.NOT_FOUND)).
		SetFetchStatus(userqueue.FetchStatus(schema.COMPLETED)).
		Save(*u.ctx)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateOrCreateUser(run int, data dto.UpdateUserDto) error {
	tx, err := u.conn.Tx(*u.ctx)
	if err != nil {
		return err
	}

	queueUser, err := tx.UserQueue.
		Query().
		Where(userqueue.RunEQ(run)).
		Only(*u.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create or Update Locations, and remove it from the user if exist
	var locations []*ent.Location
	if len(data.Locations) >= 1 {
		result, err := u.locationRepository.CreateLocations(data.Locations, tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		locations = result

	}

	// Remove locations if user exists
	foundUser, err := tx.User.
		Query().
		Where(user.RunEQ(run)).
		First(*u.ctx)

	if foundUser != nil {
		_, err := tx.User.
			UpdateOne(foundUser).
			ClearLocations().
			Save(*u.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if err != nil && !ent.IsNotFound(err) {
		tx.Rollback()
		return err
	}

	err = tx.User.
		Create().
		SetRun(queueUser.Run).
		SetVerificationDigit(queueUser.VerificationDigit).
		SetFirstName(data.FirstName).
		SetLastName(data.LastName).
		SetPhoneNumber(data.PhoneNumber).
		SetEmail(data.Email).
		SetGender(user.Gender(data.Gender)).
		SetHomeAddress(data.HomeAddress).
		SetCity(data.City).
		SetBirthDate(data.BirthDate).
		SetExpirationDate(data.ExpirationDate).
		SetPlantType(data.PlantType).
		SetEmergencyName(data.EmergencyName).
		SetEmergencyNumber(data.EmergencyNumber).
		SetMaritalStatus(data.MaritalStatus).
		AddLocations(locations...).
		OnConflict().
		UpdateNewValues().
		Exec(*u.ctx)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.UserQueue.UpdateOne(queueUser).
		Where(userqueue.RunEQ(run)).
		SetStatus(userqueue.Status(schema.FOUND)).
		SetFetchStatus(userqueue.FetchStatus(schema.COMPLETED)).
		Save(*u.ctx)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (u *UserRepository) GetQueueUsersStatistics() (*model.QueueUsersStatistics, error) {
	tx, err := u.conn.Tx(*u.ctx)
	if err != nil {
		return nil, err
	}

	// General Statistics
	totalFound, err := tx.User.
		Query().
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalNotFound, err := tx.UserQueue.
		Query().
		Where(userqueue.StatusEQ(userqueue.StatusNOT_FOUND)).
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalWithError, err := tx.UserQueue.
		Query().
		Where(userqueue.FetchStatusEQ(userqueue.FetchStatusERROR)).
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Last Hour Statistics
	oneHourAgo := time.Now().Add(-1 * time.Hour)

	totalFound1Hr, err := tx.User.
		Query().
		Where(
			user.CreatedAtGTE(oneHourAgo),
		).
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalNotFound1Hr, err := tx.UserQueue.
		Query().
		Where(
			userqueue.StatusEQ(userqueue.StatusNOT_FOUND),
			userqueue.UpdatedAtGTE(oneHourAgo),
		).
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalWithError1Hr, err := tx.UserQueue.
		Query().
		Where(
			userqueue.FetchStatusEQ(userqueue.FetchStatusERROR),
			userqueue.UpdatedAtGTE(oneHourAgo),
		).
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Total Waiting
	totalWaiting, err := tx.UserQueue.
		Query().
		Where(userqueue.FetchStatusEQ(userqueue.FetchStatusWAITING)).
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	total, err := tx.UserQueue.
		Query().
		Count(*u.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	statistics := &model.QueueUsersStatistics{
		General: model.Details{
			Found:     totalFound,
			NotFound:  totalNotFound,
			WithError: totalWithError,
		},
		LastHour: model.Details{
			Found:     totalFound1Hr,
			NotFound:  totalNotFound1Hr,
			WithError: totalWithError1Hr,
		},
		Waiting: totalWaiting,
		Total:   total,
	}

	return statistics, nil

}

func (u *UserRepository) UpdateDelayedUsers() (int, error) {
	fiveMinutesAgo := time.Now().Add(-1 * time.Minute * 5)

	result, err := u.conn.UserQueue.
		Update().
		Where(
			userqueue.FetchStatusEQ(userqueue.FetchStatusPENDING),
			userqueue.UpdatedAtLTE(fiveMinutesAgo),
		).
		SetFetchStatus(userqueue.FetchStatusWAITING).
		Save(*u.ctx)
	if err != nil {
		return 0, err
	}

	return result, nil
}
