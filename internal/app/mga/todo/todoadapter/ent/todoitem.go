// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/piotr-rusin/playing-with-golang/internal/app/mga/todo/todoadapter/ent/todoitem"
)

// TodoItem is the model entity for the TodoItem schema.
type TodoItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Completed holds the value of the "completed" field.
	Completed bool `json:"completed,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TodoItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case todoitem.FieldCompleted:
			values[i] = new(sql.NullBool)
		case todoitem.FieldID, todoitem.FieldOrder:
			values[i] = new(sql.NullInt64)
		case todoitem.FieldUID, todoitem.FieldTitle:
			values[i] = new(sql.NullString)
		case todoitem.FieldCreatedAt, todoitem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TodoItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TodoItem fields.
func (ti *TodoItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todoitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ti.ID = int(value.Int64)
		case todoitem.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				ti.UID = value.String
			}
		case todoitem.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ti.Title = value.String
			}
		case todoitem.FieldCompleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field completed", values[i])
			} else if value.Valid {
				ti.Completed = value.Bool
			}
		case todoitem.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				ti.Order = int(value.Int64)
			}
		case todoitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ti.CreatedAt = value.Time
			}
		case todoitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ti.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TodoItem.
// Note that you need to call TodoItem.Unwrap() before calling this method if this TodoItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (ti *TodoItem) Update() *TodoItemUpdateOne {
	return (&TodoItemClient{config: ti.config}).UpdateOne(ti)
}

// Unwrap unwraps the TodoItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ti *TodoItem) Unwrap() *TodoItem {
	tx, ok := ti.config.driver.(*txDriver)
	if !ok {
		panic("ent: TodoItem is not a transactional entity")
	}
	ti.config.driver = tx.drv
	return ti
}

// String implements the fmt.Stringer.
func (ti *TodoItem) String() string {
	var builder strings.Builder
	builder.WriteString("TodoItem(")
	builder.WriteString(fmt.Sprintf("id=%v", ti.ID))
	builder.WriteString(", uid=")
	builder.WriteString(ti.UID)
	builder.WriteString(", title=")
	builder.WriteString(ti.Title)
	builder.WriteString(", completed=")
	builder.WriteString(fmt.Sprintf("%v", ti.Completed))
	builder.WriteString(", order=")
	builder.WriteString(fmt.Sprintf("%v", ti.Order))
	builder.WriteString(", created_at=")
	builder.WriteString(ti.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ti.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TodoItems is a parsable slice of TodoItem.
type TodoItems []*TodoItem

func (ti TodoItems) config(cfg config) {
	for _i := range ti {
		ti[_i].config = cfg
	}
}
