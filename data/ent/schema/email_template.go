package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// EmailTemplate holds the schema definition for the EmailTemplate entity.
type EmailTemplate struct {
	ent.Schema
}

// Fields of the EmailTemplate.
func (EmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("template_key").Unique().NotEmpty().MaxLen(100),
		field.String("name").NotEmpty().MaxLen(255),
		field.String("subject").NotEmpty().MaxLen(255),
		field.Text("body").NotEmpty(),
		field.Bool("is_html").Default(false),
		field.Time("created_at").Immutable().Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the EmailTemplate.
func (EmailTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the EmailTemplate.
func (EmailTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "email_template",
		},
	}
}
