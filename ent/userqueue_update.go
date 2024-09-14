// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"spl-users/ent/predicate"
	"spl-users/ent/userqueue"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserQueueUpdate is the builder for updating UserQueue entities.
type UserQueueUpdate struct {
	config
	hooks    []Hook
	mutation *UserQueueMutation
}

// Where appends a list predicates to the UserQueueUpdate builder.
func (uqu *UserQueueUpdate) Where(ps ...predicate.UserQueue) *UserQueueUpdate {
	uqu.mutation.Where(ps...)
	return uqu
}

// SetRun sets the "run" field.
func (uqu *UserQueueUpdate) SetRun(i int) *UserQueueUpdate {
	uqu.mutation.ResetRun()
	uqu.mutation.SetRun(i)
	return uqu
}

// SetNillableRun sets the "run" field if the given value is not nil.
func (uqu *UserQueueUpdate) SetNillableRun(i *int) *UserQueueUpdate {
	if i != nil {
		uqu.SetRun(*i)
	}
	return uqu
}

// AddRun adds i to the "run" field.
func (uqu *UserQueueUpdate) AddRun(i int) *UserQueueUpdate {
	uqu.mutation.AddRun(i)
	return uqu
}

// SetVerificationDigit sets the "verification_digit" field.
func (uqu *UserQueueUpdate) SetVerificationDigit(s string) *UserQueueUpdate {
	uqu.mutation.SetVerificationDigit(s)
	return uqu
}

// SetNillableVerificationDigit sets the "verification_digit" field if the given value is not nil.
func (uqu *UserQueueUpdate) SetNillableVerificationDigit(s *string) *UserQueueUpdate {
	if s != nil {
		uqu.SetVerificationDigit(*s)
	}
	return uqu
}

// SetCreatedAt sets the "created_at" field.
func (uqu *UserQueueUpdate) SetCreatedAt(t time.Time) *UserQueueUpdate {
	uqu.mutation.SetCreatedAt(t)
	return uqu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uqu *UserQueueUpdate) SetNillableCreatedAt(t *time.Time) *UserQueueUpdate {
	if t != nil {
		uqu.SetCreatedAt(*t)
	}
	return uqu
}

// SetUpdatedAt sets the "updated_at" field.
func (uqu *UserQueueUpdate) SetUpdatedAt(t time.Time) *UserQueueUpdate {
	uqu.mutation.SetUpdatedAt(t)
	return uqu
}

// SetFetchStatus sets the "fetch_status" field.
func (uqu *UserQueueUpdate) SetFetchStatus(us userqueue.FetchStatus) *UserQueueUpdate {
	uqu.mutation.SetFetchStatus(us)
	return uqu
}

// SetNillableFetchStatus sets the "fetch_status" field if the given value is not nil.
func (uqu *UserQueueUpdate) SetNillableFetchStatus(us *userqueue.FetchStatus) *UserQueueUpdate {
	if us != nil {
		uqu.SetFetchStatus(*us)
	}
	return uqu
}

// SetStatus sets the "status" field.
func (uqu *UserQueueUpdate) SetStatus(u userqueue.Status) *UserQueueUpdate {
	uqu.mutation.SetStatus(u)
	return uqu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uqu *UserQueueUpdate) SetNillableStatus(u *userqueue.Status) *UserQueueUpdate {
	if u != nil {
		uqu.SetStatus(*u)
	}
	return uqu
}

// Mutation returns the UserQueueMutation object of the builder.
func (uqu *UserQueueUpdate) Mutation() *UserQueueMutation {
	return uqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uqu *UserQueueUpdate) Save(ctx context.Context) (int, error) {
	uqu.defaults()
	return withHooks(ctx, uqu.sqlSave, uqu.mutation, uqu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uqu *UserQueueUpdate) SaveX(ctx context.Context) int {
	affected, err := uqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uqu *UserQueueUpdate) Exec(ctx context.Context) error {
	_, err := uqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uqu *UserQueueUpdate) ExecX(ctx context.Context) {
	if err := uqu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uqu *UserQueueUpdate) defaults() {
	if _, ok := uqu.mutation.UpdatedAt(); !ok {
		v := userqueue.UpdateDefaultUpdatedAt()
		uqu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uqu *UserQueueUpdate) check() error {
	if v, ok := uqu.mutation.Run(); ok {
		if err := userqueue.RunValidator(v); err != nil {
			return &ValidationError{Name: "run", err: fmt.Errorf(`ent: validator failed for field "UserQueue.run": %w`, err)}
		}
	}
	if v, ok := uqu.mutation.VerificationDigit(); ok {
		if err := userqueue.VerificationDigitValidator(v); err != nil {
			return &ValidationError{Name: "verification_digit", err: fmt.Errorf(`ent: validator failed for field "UserQueue.verification_digit": %w`, err)}
		}
	}
	if v, ok := uqu.mutation.FetchStatus(); ok {
		if err := userqueue.FetchStatusValidator(v); err != nil {
			return &ValidationError{Name: "fetch_status", err: fmt.Errorf(`ent: validator failed for field "UserQueue.fetch_status": %w`, err)}
		}
	}
	if v, ok := uqu.mutation.Status(); ok {
		if err := userqueue.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserQueue.status": %w`, err)}
		}
	}
	return nil
}

func (uqu *UserQueueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uqu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userqueue.Table, userqueue.Columns, sqlgraph.NewFieldSpec(userqueue.FieldID, field.TypeInt))
	if ps := uqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uqu.mutation.Run(); ok {
		_spec.SetField(userqueue.FieldRun, field.TypeInt, value)
	}
	if value, ok := uqu.mutation.AddedRun(); ok {
		_spec.AddField(userqueue.FieldRun, field.TypeInt, value)
	}
	if value, ok := uqu.mutation.VerificationDigit(); ok {
		_spec.SetField(userqueue.FieldVerificationDigit, field.TypeString, value)
	}
	if value, ok := uqu.mutation.CreatedAt(); ok {
		_spec.SetField(userqueue.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uqu.mutation.UpdatedAt(); ok {
		_spec.SetField(userqueue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uqu.mutation.FetchStatus(); ok {
		_spec.SetField(userqueue.FieldFetchStatus, field.TypeEnum, value)
	}
	if value, ok := uqu.mutation.Status(); ok {
		_spec.SetField(userqueue.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userqueue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uqu.mutation.done = true
	return n, nil
}

// UserQueueUpdateOne is the builder for updating a single UserQueue entity.
type UserQueueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserQueueMutation
}

// SetRun sets the "run" field.
func (uquo *UserQueueUpdateOne) SetRun(i int) *UserQueueUpdateOne {
	uquo.mutation.ResetRun()
	uquo.mutation.SetRun(i)
	return uquo
}

// SetNillableRun sets the "run" field if the given value is not nil.
func (uquo *UserQueueUpdateOne) SetNillableRun(i *int) *UserQueueUpdateOne {
	if i != nil {
		uquo.SetRun(*i)
	}
	return uquo
}

// AddRun adds i to the "run" field.
func (uquo *UserQueueUpdateOne) AddRun(i int) *UserQueueUpdateOne {
	uquo.mutation.AddRun(i)
	return uquo
}

// SetVerificationDigit sets the "verification_digit" field.
func (uquo *UserQueueUpdateOne) SetVerificationDigit(s string) *UserQueueUpdateOne {
	uquo.mutation.SetVerificationDigit(s)
	return uquo
}

// SetNillableVerificationDigit sets the "verification_digit" field if the given value is not nil.
func (uquo *UserQueueUpdateOne) SetNillableVerificationDigit(s *string) *UserQueueUpdateOne {
	if s != nil {
		uquo.SetVerificationDigit(*s)
	}
	return uquo
}

// SetCreatedAt sets the "created_at" field.
func (uquo *UserQueueUpdateOne) SetCreatedAt(t time.Time) *UserQueueUpdateOne {
	uquo.mutation.SetCreatedAt(t)
	return uquo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uquo *UserQueueUpdateOne) SetNillableCreatedAt(t *time.Time) *UserQueueUpdateOne {
	if t != nil {
		uquo.SetCreatedAt(*t)
	}
	return uquo
}

// SetUpdatedAt sets the "updated_at" field.
func (uquo *UserQueueUpdateOne) SetUpdatedAt(t time.Time) *UserQueueUpdateOne {
	uquo.mutation.SetUpdatedAt(t)
	return uquo
}

// SetFetchStatus sets the "fetch_status" field.
func (uquo *UserQueueUpdateOne) SetFetchStatus(us userqueue.FetchStatus) *UserQueueUpdateOne {
	uquo.mutation.SetFetchStatus(us)
	return uquo
}

// SetNillableFetchStatus sets the "fetch_status" field if the given value is not nil.
func (uquo *UserQueueUpdateOne) SetNillableFetchStatus(us *userqueue.FetchStatus) *UserQueueUpdateOne {
	if us != nil {
		uquo.SetFetchStatus(*us)
	}
	return uquo
}

// SetStatus sets the "status" field.
func (uquo *UserQueueUpdateOne) SetStatus(u userqueue.Status) *UserQueueUpdateOne {
	uquo.mutation.SetStatus(u)
	return uquo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uquo *UserQueueUpdateOne) SetNillableStatus(u *userqueue.Status) *UserQueueUpdateOne {
	if u != nil {
		uquo.SetStatus(*u)
	}
	return uquo
}

// Mutation returns the UserQueueMutation object of the builder.
func (uquo *UserQueueUpdateOne) Mutation() *UserQueueMutation {
	return uquo.mutation
}

// Where appends a list predicates to the UserQueueUpdate builder.
func (uquo *UserQueueUpdateOne) Where(ps ...predicate.UserQueue) *UserQueueUpdateOne {
	uquo.mutation.Where(ps...)
	return uquo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uquo *UserQueueUpdateOne) Select(field string, fields ...string) *UserQueueUpdateOne {
	uquo.fields = append([]string{field}, fields...)
	return uquo
}

// Save executes the query and returns the updated UserQueue entity.
func (uquo *UserQueueUpdateOne) Save(ctx context.Context) (*UserQueue, error) {
	uquo.defaults()
	return withHooks(ctx, uquo.sqlSave, uquo.mutation, uquo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uquo *UserQueueUpdateOne) SaveX(ctx context.Context) *UserQueue {
	node, err := uquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uquo *UserQueueUpdateOne) Exec(ctx context.Context) error {
	_, err := uquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uquo *UserQueueUpdateOne) ExecX(ctx context.Context) {
	if err := uquo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uquo *UserQueueUpdateOne) defaults() {
	if _, ok := uquo.mutation.UpdatedAt(); !ok {
		v := userqueue.UpdateDefaultUpdatedAt()
		uquo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uquo *UserQueueUpdateOne) check() error {
	if v, ok := uquo.mutation.Run(); ok {
		if err := userqueue.RunValidator(v); err != nil {
			return &ValidationError{Name: "run", err: fmt.Errorf(`ent: validator failed for field "UserQueue.run": %w`, err)}
		}
	}
	if v, ok := uquo.mutation.VerificationDigit(); ok {
		if err := userqueue.VerificationDigitValidator(v); err != nil {
			return &ValidationError{Name: "verification_digit", err: fmt.Errorf(`ent: validator failed for field "UserQueue.verification_digit": %w`, err)}
		}
	}
	if v, ok := uquo.mutation.FetchStatus(); ok {
		if err := userqueue.FetchStatusValidator(v); err != nil {
			return &ValidationError{Name: "fetch_status", err: fmt.Errorf(`ent: validator failed for field "UserQueue.fetch_status": %w`, err)}
		}
	}
	if v, ok := uquo.mutation.Status(); ok {
		if err := userqueue.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserQueue.status": %w`, err)}
		}
	}
	return nil
}

func (uquo *UserQueueUpdateOne) sqlSave(ctx context.Context) (_node *UserQueue, err error) {
	if err := uquo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userqueue.Table, userqueue.Columns, sqlgraph.NewFieldSpec(userqueue.FieldID, field.TypeInt))
	id, ok := uquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserQueue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userqueue.FieldID)
		for _, f := range fields {
			if !userqueue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userqueue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uquo.mutation.Run(); ok {
		_spec.SetField(userqueue.FieldRun, field.TypeInt, value)
	}
	if value, ok := uquo.mutation.AddedRun(); ok {
		_spec.AddField(userqueue.FieldRun, field.TypeInt, value)
	}
	if value, ok := uquo.mutation.VerificationDigit(); ok {
		_spec.SetField(userqueue.FieldVerificationDigit, field.TypeString, value)
	}
	if value, ok := uquo.mutation.CreatedAt(); ok {
		_spec.SetField(userqueue.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uquo.mutation.UpdatedAt(); ok {
		_spec.SetField(userqueue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uquo.mutation.FetchStatus(); ok {
		_spec.SetField(userqueue.FieldFetchStatus, field.TypeEnum, value)
	}
	if value, ok := uquo.mutation.Status(); ok {
		_spec.SetField(userqueue.FieldStatus, field.TypeEnum, value)
	}
	_node = &UserQueue{config: uquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userqueue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uquo.mutation.done = true
	return _node, nil
}
