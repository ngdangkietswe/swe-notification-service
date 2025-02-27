// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CdcAuthUsersColumns holds the columns for the "cdc_auth_users" table.
	CdcAuthUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Size: 255},
	}
	// CdcAuthUsersTable holds the schema information for the "cdc_auth_users" table.
	CdcAuthUsersTable = &schema.Table{
		Name:       "cdc_auth_users",
		Columns:    CdcAuthUsersColumns,
		PrimaryKey: []*schema.Column{CdcAuthUsersColumns[0]},
	}
	// EmailTemplateColumns holds the columns for the "email_template" table.
	EmailTemplateColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "template_key", Type: field.TypeString, Unique: true, Size: 100},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "subject", Type: field.TypeString, Size: 255},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "is_html", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
	}
	// EmailTemplateTable holds the schema information for the "email_template" table.
	EmailTemplateTable = &schema.Table{
		Name:       "email_template",
		Columns:    EmailTemplateColumns,
		PrimaryKey: []*schema.Column{EmailTemplateColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CdcAuthUsersTable,
		EmailTemplateTable,
	}
)

func init() {
	CdcAuthUsersTable.Annotation = &entsql.Annotation{
		Table: "cdc_auth_users",
	}
	EmailTemplateTable.Annotation = &entsql.Annotation{
		Table: "email_template",
	}
}
