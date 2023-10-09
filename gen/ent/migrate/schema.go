// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// JokesColumns holds the columns for the "jokes" table.
	JokesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "explanation", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_jokes", Type: field.TypeString, Nullable: true},
	}
	// JokesTable holds the schema information for the "jokes" table.
	JokesTable = &schema.Table{
		Name:       "jokes",
		Columns:    JokesColumns,
		PrimaryKey: []*schema.Column{JokesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "jokes_users_jokes",
				Columns:    []*schema.Column{JokesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "joke_user_id_title",
				Unique:  true,
				Columns: []*schema.Column{JokesColumns[1], JokesColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "hashed_password", Type: field.TypeString},
		{Name: "fullname", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		JokesTable,
		UsersTable,
	}
)

func init() {
	JokesTable.ForeignKeys[0].RefTable = UsersTable
}