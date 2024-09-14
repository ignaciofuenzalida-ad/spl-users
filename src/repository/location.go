package repository

import (
	"context"
	"spl-users/ent"
	"spl-users/ent/location"
	"spl-users/src/utils"
)

type LocationRepository struct {
	conn *ent.Client
	ctx  *context.Context
}

func NewLocationRepository(conn *ent.Client, ctx *context.Context) *LocationRepository {
	return &LocationRepository{conn: conn, ctx: ctx}
}

func (l *LocationRepository) GetAllLocations() ([]*ent.Location, error) {
	locations, err := l.conn.Location.Query().All(*l.ctx)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (l *LocationRepository) CreateLocations(locations []string, tx *ent.Tx) ([]*ent.Location, error) {
	var slugedLocations = make([]string, 0)
	for _, location := range locations {
		slugedLocations = append(slugedLocations, utils.Slugify(location))
	}

	var toCreateLocations = make([]*ent.LocationCreate, len(locations))
	for i, location := range locations {
		toCreateLocations[i] = l.conn.Location.
			Create().
			SetValue(location).
			SetSlug(slugedLocations[i])
	}

	if tx != nil {
		err := tx.Location.
			CreateBulk(toCreateLocations...).
			OnConflict().
			UpdateNewValues().
			Exec(*l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

	} else {
		err := l.conn.Location.
			CreateBulk(toCreateLocations...).
			OnConflict().
			UpdateNewValues().
			Exec(*l.ctx)
		if err != nil {
			return nil, err
		}
	}

	if tx != nil {
		result, err := tx.Location.
			Query().
			Where(location.SlugIn(slugedLocations...)).
			All(*l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		return result, nil
	}

	result, err := l.conn.Location.
		Query().
		Where(location.SlugIn(slugedLocations...)).
		All(*l.ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

}
