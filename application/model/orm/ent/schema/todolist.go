package schema

import ("entgo.io/ent"
"entgo.io/ent/schema/field"
"time"
)
// Todolist holds the schema definition for the Todolist entity.
type Todolist struct {
	ent.Schema
}

// Fields of the Todolist.
func (Todolist) Fields() []ent.Field {
	return []ent.Field{
		//ここにcolumnを記載する。
		field.String("title"),
		field.String("body"),
		field.Int("priority"),
		field.Int("delete_flag").Default(0),
		field.Time("created_at").Default(time.Now),
		}
}

// Edges of the Todolist.
func (Todolist) Edges() []ent.Edge {
	return nil
}
