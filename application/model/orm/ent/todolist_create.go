// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mvc-model-sample/application/model/orm/ent/todolist"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodolistCreate is the builder for creating a Todolist entity.
type TodolistCreate struct {
	config
	mutation *TodolistMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (tc *TodolistCreate) SetTitle(s string) *TodolistCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetBody sets the "body" field.
func (tc *TodolistCreate) SetBody(s string) *TodolistCreate {
	tc.mutation.SetBody(s)
	return tc
}

// SetPriority sets the "priority" field.
func (tc *TodolistCreate) SetPriority(i int) *TodolistCreate {
	tc.mutation.SetPriority(i)
	return tc
}

// SetDeleteFlag sets the "delete_flag" field.
func (tc *TodolistCreate) SetDeleteFlag(i int) *TodolistCreate {
	tc.mutation.SetDeleteFlag(i)
	return tc
}

// SetNillableDeleteFlag sets the "delete_flag" field if the given value is not nil.
func (tc *TodolistCreate) SetNillableDeleteFlag(i *int) *TodolistCreate {
	if i != nil {
		tc.SetDeleteFlag(*i)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TodolistCreate) SetCreatedAt(t time.Time) *TodolistCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TodolistCreate) SetNillableCreatedAt(t *time.Time) *TodolistCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// Mutation returns the TodolistMutation object of the builder.
func (tc *TodolistCreate) Mutation() *TodolistMutation {
	return tc.mutation
}

// Save creates the Todolist in the database.
func (tc *TodolistCreate) Save(ctx context.Context) (*Todolist, error) {
	tc.defaults()
	return withHooks[*Todolist, TodolistMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodolistCreate) SaveX(ctx context.Context) *Todolist {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TodolistCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TodolistCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TodolistCreate) defaults() {
	if _, ok := tc.mutation.DeleteFlag(); !ok {
		v := todolist.DefaultDeleteFlag
		tc.mutation.SetDeleteFlag(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := todolist.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TodolistCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Todolist.title"`)}
	}
	if _, ok := tc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Todolist.body"`)}
	}
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Todolist.priority"`)}
	}
	if _, ok := tc.mutation.DeleteFlag(); !ok {
		return &ValidationError{Name: "delete_flag", err: errors.New(`ent: missing required field "Todolist.delete_flag"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Todolist.created_at"`)}
	}
	return nil
}

func (tc *TodolistCreate) sqlSave(ctx context.Context) (*Todolist, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TodolistCreate) createSpec() (*Todolist, *sqlgraph.CreateSpec) {
	var (
		_node = &Todolist{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todolist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todolist.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(todolist.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Body(); ok {
		_spec.SetField(todolist.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := tc.mutation.Priority(); ok {
		_spec.SetField(todolist.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if value, ok := tc.mutation.DeleteFlag(); ok {
		_spec.SetField(todolist.FieldDeleteFlag, field.TypeInt, value)
		_node.DeleteFlag = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(todolist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// TodolistCreateBulk is the builder for creating many Todolist entities in bulk.
type TodolistCreateBulk struct {
	config
	builders []*TodolistCreate
}

// Save creates the Todolist entities in the database.
func (tcb *TodolistCreateBulk) Save(ctx context.Context) ([]*Todolist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Todolist, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TodolistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TodolistCreateBulk) SaveX(ctx context.Context) []*Todolist {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TodolistCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TodolistCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
