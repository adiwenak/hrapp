// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldMobileNumber holds the string denoting the mobilenumber field in the database.
	FieldMobileNumber = "mobile_number"
	// FieldFirstName holds the string denoting the firstname field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the lastname field in the database.
	FieldLastName = "last_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNeedPasswordReset holds the string denoting the needpasswordreset field in the database.
	FieldNeedPasswordReset = "need_password_reset"
	// FieldVerificationCode holds the string denoting the verificationcode field in the database.
	FieldVerificationCode = "verification_code"
	// EdgeOrganisation holds the string denoting the organisation edge name in mutations.
	EdgeOrganisation = "organisation"
	// Table holds the table name of the user in the database.
	Table = "users"
	// OrganisationTable is the table that holds the organisation relation/edge.
	OrganisationTable = "users"
	// OrganisationInverseTable is the table name for the Organisation entity.
	// It exists in this package in order to avoid circular dependency with the "organisation" package.
	OrganisationInverseTable = "organisations"
	// OrganisationColumn is the table column denoting the organisation relation/edge.
	OrganisationColumn = "organisation_users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
	FieldUsername,
	FieldMobileNumber,
	FieldFirstName,
	FieldLastName,
	FieldPassword,
	FieldNeedPasswordReset,
	FieldVerificationCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"organisation_users",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByMobileNumber orders the results by the mobileNumber field.
func ByMobileNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobileNumber, opts...).ToFunc()
}

// ByFirstName orders the results by the firstName field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the lastName field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNeedPasswordReset orders the results by the needPasswordReset field.
func ByNeedPasswordReset(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNeedPasswordReset, opts...).ToFunc()
}

// ByVerificationCode orders the results by the verificationCode field.
func ByVerificationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerificationCode, opts...).ToFunc()
}

// ByOrganisationField orders the results by organisation field.
func ByOrganisationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganisationStep(), sql.OrderByField(field, opts...))
	}
}
func newOrganisationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganisationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganisationTable, OrganisationColumn),
	)
}
