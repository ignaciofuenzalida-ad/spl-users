package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type FetchStatus string
type Status string

const (
	WAITING   FetchStatus = "WAITING"
	PENDING   FetchStatus = "PENDING"
	ERROR     FetchStatus = "ERROR"
	COMPLETED FetchStatus = "COMPLETED"
)

const (
	EMPTY     Status = "EMPTY"
	NOT_FOUND Status = "NOT_FOUND"
	FOUND     Status = "FOUND"
)

// UserQueue holds the schema definition for the UserQueue entity.
type UserQueue struct {
	ent.Schema
}

// Fields of the UserQueue.
func (UserQueue) Fields() []ent.Field {
	return []ent.Field{
		field.Int("run").Unique().Positive(),
		field.String("verification_digit").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("fetch_status").
			Values(
				string(WAITING),
				string(PENDING),
				string(ERROR),
				string(COMPLETED),
			).Default(string(WAITING)),
		field.Enum("status").
			Values(
				string(EMPTY),
				string(NOT_FOUND),
				string(FOUND),
			).Default(string(EMPTY)),
	}
}

// Edges of the UserQueue.
func (UserQueue) Edges() []ent.Edge {
	return nil
}
