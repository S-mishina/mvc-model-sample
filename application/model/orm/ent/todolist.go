// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mvc-model-sample/application/model/orm/ent/todolist"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Todolist is the model entity for the Todolist schema.
type Todolist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// DeleteFlag holds the value of the "delete_flag" field.
	DeleteFlag int `json:"delete_flag,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todolist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todolist.FieldID, todolist.FieldPriority, todolist.FieldDeleteFlag:
			values[i] = new(sql.NullInt64)
		case todolist.FieldTitle, todolist.FieldBody:
			values[i] = new(sql.NullString)
		case todolist.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Todolist", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todolist fields.
func (t *Todolist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todolist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todolist.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case todolist.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				t.Body = value.String
			}
		case todolist.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = int(value.Int64)
			}
		case todolist.FieldDeleteFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_flag", values[i])
			} else if value.Valid {
				t.DeleteFlag = int(value.Int64)
			}
		case todolist.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Todolist.
// Note that you need to call Todolist.Unwrap() before calling this method if this Todolist
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todolist) Update() *TodolistUpdateOne {
	return (&TodolistClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Todolist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todolist) Unwrap() *Todolist {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todolist is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todolist) String() string {
	var builder strings.Builder
	builder.WriteString("Todolist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(t.Body)
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteString(", ")
	builder.WriteString("delete_flag=")
	builder.WriteString(fmt.Sprintf("%v", t.DeleteFlag))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Todolists is a parsable slice of Todolist.
type Todolists []*Todolist

func (t Todolists) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
