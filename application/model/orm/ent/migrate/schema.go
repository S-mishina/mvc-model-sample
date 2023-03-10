// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodolistsColumns holds the columns for the "todolists" table.
	TodolistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString},
		{Name: "priority", Type: field.TypeInt},
		{Name: "delete_flag", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TodolistsTable holds the schema information for the "todolists" table.
	TodolistsTable = &schema.Table{
		Name:       "todolists",
		Columns:    TodolistsColumns,
		PrimaryKey: []*schema.Column{TodolistsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodolistsTable,
	}
)

func init() {
}
