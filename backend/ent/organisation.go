// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/adiwenak/hrapp/ent/organisation"
)

// Organisation is the model entity for the Organisation schema.
type Organisation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganisationQuery when eager-loading is set.
	Edges        OrganisationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganisationEdges holds the relations/edges for other nodes in the graph.
type OrganisationEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e OrganisationEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organisation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organisation.FieldID:
			values[i] = new(sql.NullInt64)
		case organisation.FieldName:
			values[i] = new(sql.NullString)
		case organisation.FieldCreatedAt, organisation.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organisation fields.
func (o *Organisation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organisation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case organisation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case organisation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case organisation.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organisation.
// This includes values selected through modifiers, order, etc.
func (o *Organisation) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Organisation entity.
func (o *Organisation) QueryUsers() *UserQuery {
	return NewOrganisationClient(o.config).QueryUsers(o)
}

// Update returns a builder for updating this Organisation.
// Note that you need to call Organisation.Unwrap() before calling this method if this Organisation
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organisation) Update() *OrganisationUpdateOne {
	return NewOrganisationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organisation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organisation) Unwrap() *Organisation {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organisation is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organisation) String() string {
	var builder strings.Builder
	builder.WriteString("Organisation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Organisations is a parsable slice of Organisation.
type Organisations []*Organisation