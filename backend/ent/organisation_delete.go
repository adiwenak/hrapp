// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adiwenak/hrapp/ent/organisation"
	"github.com/adiwenak/hrapp/ent/predicate"
)

// OrganisationDelete is the builder for deleting a Organisation entity.
type OrganisationDelete struct {
	config
	hooks    []Hook
	mutation *OrganisationMutation
}

// Where appends a list predicates to the OrganisationDelete builder.
func (od *OrganisationDelete) Where(ps ...predicate.Organisation) *OrganisationDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OrganisationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OrganisationDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OrganisationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(organisation.Table, sqlgraph.NewFieldSpec(organisation.FieldID, field.TypeInt))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OrganisationDeleteOne is the builder for deleting a single Organisation entity.
type OrganisationDeleteOne struct {
	od *OrganisationDelete
}

// Where appends a list predicates to the OrganisationDelete builder.
func (odo *OrganisationDeleteOne) Where(ps ...predicate.Organisation) *OrganisationDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OrganisationDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organisation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OrganisationDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}
