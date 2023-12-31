// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"goBack/ent/tn_bbs"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TN_BBS is the model entity for the TN_BBS schema.
type TN_BBS struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// DelectYn holds the value of the "delect_yn" field.
	DelectYn string `json:"delect_yn,omitempty"`
	// RegDate holds the value of the "reg_date" field.
	RegDate string `json:"reg_date,omitempty"`
	// UdtDate holds the value of the "udt_date" field.
	UdtDate      string `json:"udt_date,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TN_BBS) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tn_bbs.FieldID:
			values[i] = new(sql.NullInt64)
		case tn_bbs.FieldUserID, tn_bbs.FieldUserName, tn_bbs.FieldTitle, tn_bbs.FieldContent, tn_bbs.FieldDelectYn, tn_bbs.FieldRegDate, tn_bbs.FieldUdtDate:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TN_BBS fields.
func (tb *TN_BBS) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tn_bbs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tb.ID = int(value.Int64)
		case tn_bbs.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tb.UserID = value.String
			}
		case tn_bbs.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				tb.UserName = value.String
			}
		case tn_bbs.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				tb.Title = value.String
			}
		case tn_bbs.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				tb.Content = value.String
			}
		case tn_bbs.FieldDelectYn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delect_yn", values[i])
			} else if value.Valid {
				tb.DelectYn = value.String
			}
		case tn_bbs.FieldRegDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reg_date", values[i])
			} else if value.Valid {
				tb.RegDate = value.String
			}
		case tn_bbs.FieldUdtDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field udt_date", values[i])
			} else if value.Valid {
				tb.UdtDate = value.String
			}
		default:
			tb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TN_BBS.
// This includes values selected through modifiers, order, etc.
func (tb *TN_BBS) Value(name string) (ent.Value, error) {
	return tb.selectValues.Get(name)
}

// Update returns a builder for updating this TN_BBS.
// Note that you need to call TN_BBS.Unwrap() before calling this method if this TN_BBS
// was returned from a transaction, and the transaction was committed or rolled back.
func (tb *TN_BBS) Update() *TNBBSUpdateOne {
	return NewTNBBSClient(tb.config).UpdateOne(tb)
}

// Unwrap unwraps the TN_BBS entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tb *TN_BBS) Unwrap() *TN_BBS {
	_tx, ok := tb.config.driver.(*txDriver)
	if !ok {
		panic("ent: TN_BBS is not a transactional entity")
	}
	tb.config.driver = _tx.drv
	return tb
}

// String implements the fmt.Stringer.
func (tb *TN_BBS) String() string {
	var builder strings.Builder
	builder.WriteString("TN_BBS(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tb.ID))
	builder.WriteString("user_id=")
	builder.WriteString(tb.UserID)
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(tb.UserName)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(tb.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(tb.Content)
	builder.WriteString(", ")
	builder.WriteString("delect_yn=")
	builder.WriteString(tb.DelectYn)
	builder.WriteString(", ")
	builder.WriteString("reg_date=")
	builder.WriteString(tb.RegDate)
	builder.WriteString(", ")
	builder.WriteString("udt_date=")
	builder.WriteString(tb.UdtDate)
	builder.WriteByte(')')
	return builder.String()
}

// TN_BBSs is a parsable slice of TN_BBS.
type TN_BBSs []*TN_BBS
