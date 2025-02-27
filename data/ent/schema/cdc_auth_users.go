package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CdcAuthUsers holds the schema definition for the CdcAuthUsers entity.
type CdcAuthUsers struct {
	ent.Schema
}

// Fields of the CdcAuthUsers.
func (CdcAuthUsers) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("username").NotEmpty().MaxLen(255),
		field.String("email").NotEmpty().MaxLen(255),
	}
}

// Edges of the CdcAuthUsers.
func (CdcAuthUsers) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the CdcAuthUsers.
func (CdcAuthUsers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "cdc_auth_users",
		},
	}
}
