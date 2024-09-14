package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

type Gender string

const (
	MALE    Gender = "MALE"
	FEMALE  Gender = "FEMALE"
	UNKNOWN Gender = "UNKNOWN"
)

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("run").Unique().Positive(),
		field.String("verification_digit").NotEmpty(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("phone_number").Optional(),
		field.Enum("gender").Values(
			string(UNKNOWN),
			string(MALE),
			string(FEMALE),
		).Default(string(UNKNOWN)),
		field.String("marital_status").Optional(),
		field.String("email").Optional(),
		field.String("home_address").Optional(),
		field.String("city").Optional(),
		field.Time("birth_date").Optional(),
		field.Time("expiration_date").Optional(),
		field.String("plant_type").Optional(),
		field.String("emergency_name").Optional(),
		field.String("emergency_number").Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("expose").Default(true),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("locations", Location.Type).Ref("users"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("run").Unique(),
	}
}
