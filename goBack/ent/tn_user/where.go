// Code generated by ent, DO NOT EDIT.

package tn_user

import (
	"goBack/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldUserID, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldUserName, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldPassword, v))
}

// DelectYn applies equality check predicate on the "delect_yn" field. It's identical to DelectYnEQ.
func DelectYn(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldDelectYn, v))
}

// RegDate applies equality check predicate on the "reg_date" field. It's identical to RegDateEQ.
func RegDate(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldRegDate, v))
}

// UdtDate applies equality check predicate on the "udt_date" field. It's identical to UdtDateEQ.
func UdtDate(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldUdtDate, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContainsFold(FieldUserID, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContainsFold(FieldUserName, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContainsFold(FieldPassword, v))
}

// DelectYnEQ applies the EQ predicate on the "delect_yn" field.
func DelectYnEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldDelectYn, v))
}

// DelectYnNEQ applies the NEQ predicate on the "delect_yn" field.
func DelectYnNEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldDelectYn, v))
}

// DelectYnIn applies the In predicate on the "delect_yn" field.
func DelectYnIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldDelectYn, vs...))
}

// DelectYnNotIn applies the NotIn predicate on the "delect_yn" field.
func DelectYnNotIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldDelectYn, vs...))
}

// DelectYnGT applies the GT predicate on the "delect_yn" field.
func DelectYnGT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldDelectYn, v))
}

// DelectYnGTE applies the GTE predicate on the "delect_yn" field.
func DelectYnGTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldDelectYn, v))
}

// DelectYnLT applies the LT predicate on the "delect_yn" field.
func DelectYnLT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldDelectYn, v))
}

// DelectYnLTE applies the LTE predicate on the "delect_yn" field.
func DelectYnLTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldDelectYn, v))
}

// DelectYnContains applies the Contains predicate on the "delect_yn" field.
func DelectYnContains(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContains(FieldDelectYn, v))
}

// DelectYnHasPrefix applies the HasPrefix predicate on the "delect_yn" field.
func DelectYnHasPrefix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasPrefix(FieldDelectYn, v))
}

// DelectYnHasSuffix applies the HasSuffix predicate on the "delect_yn" field.
func DelectYnHasSuffix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasSuffix(FieldDelectYn, v))
}

// DelectYnEqualFold applies the EqualFold predicate on the "delect_yn" field.
func DelectYnEqualFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEqualFold(FieldDelectYn, v))
}

// DelectYnContainsFold applies the ContainsFold predicate on the "delect_yn" field.
func DelectYnContainsFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContainsFold(FieldDelectYn, v))
}

// RegDateEQ applies the EQ predicate on the "reg_date" field.
func RegDateEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldRegDate, v))
}

// RegDateNEQ applies the NEQ predicate on the "reg_date" field.
func RegDateNEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldRegDate, v))
}

// RegDateIn applies the In predicate on the "reg_date" field.
func RegDateIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldRegDate, vs...))
}

// RegDateNotIn applies the NotIn predicate on the "reg_date" field.
func RegDateNotIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldRegDate, vs...))
}

// RegDateGT applies the GT predicate on the "reg_date" field.
func RegDateGT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldRegDate, v))
}

// RegDateGTE applies the GTE predicate on the "reg_date" field.
func RegDateGTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldRegDate, v))
}

// RegDateLT applies the LT predicate on the "reg_date" field.
func RegDateLT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldRegDate, v))
}

// RegDateLTE applies the LTE predicate on the "reg_date" field.
func RegDateLTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldRegDate, v))
}

// RegDateContains applies the Contains predicate on the "reg_date" field.
func RegDateContains(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContains(FieldRegDate, v))
}

// RegDateHasPrefix applies the HasPrefix predicate on the "reg_date" field.
func RegDateHasPrefix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasPrefix(FieldRegDate, v))
}

// RegDateHasSuffix applies the HasSuffix predicate on the "reg_date" field.
func RegDateHasSuffix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasSuffix(FieldRegDate, v))
}

// RegDateEqualFold applies the EqualFold predicate on the "reg_date" field.
func RegDateEqualFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEqualFold(FieldRegDate, v))
}

// RegDateContainsFold applies the ContainsFold predicate on the "reg_date" field.
func RegDateContainsFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContainsFold(FieldRegDate, v))
}

// UdtDateEQ applies the EQ predicate on the "udt_date" field.
func UdtDateEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEQ(FieldUdtDate, v))
}

// UdtDateNEQ applies the NEQ predicate on the "udt_date" field.
func UdtDateNEQ(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNEQ(FieldUdtDate, v))
}

// UdtDateIn applies the In predicate on the "udt_date" field.
func UdtDateIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldIn(FieldUdtDate, vs...))
}

// UdtDateNotIn applies the NotIn predicate on the "udt_date" field.
func UdtDateNotIn(vs ...string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldNotIn(FieldUdtDate, vs...))
}

// UdtDateGT applies the GT predicate on the "udt_date" field.
func UdtDateGT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGT(FieldUdtDate, v))
}

// UdtDateGTE applies the GTE predicate on the "udt_date" field.
func UdtDateGTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldGTE(FieldUdtDate, v))
}

// UdtDateLT applies the LT predicate on the "udt_date" field.
func UdtDateLT(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLT(FieldUdtDate, v))
}

// UdtDateLTE applies the LTE predicate on the "udt_date" field.
func UdtDateLTE(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldLTE(FieldUdtDate, v))
}

// UdtDateContains applies the Contains predicate on the "udt_date" field.
func UdtDateContains(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContains(FieldUdtDate, v))
}

// UdtDateHasPrefix applies the HasPrefix predicate on the "udt_date" field.
func UdtDateHasPrefix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasPrefix(FieldUdtDate, v))
}

// UdtDateHasSuffix applies the HasSuffix predicate on the "udt_date" field.
func UdtDateHasSuffix(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldHasSuffix(FieldUdtDate, v))
}

// UdtDateEqualFold applies the EqualFold predicate on the "udt_date" field.
func UdtDateEqualFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldEqualFold(FieldUdtDate, v))
}

// UdtDateContainsFold applies the ContainsFold predicate on the "udt_date" field.
func UdtDateContainsFold(v string) predicate.TN_USER {
	return predicate.TN_USER(sql.FieldContainsFold(FieldUdtDate, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TN_USER) predicate.TN_USER {
	return predicate.TN_USER(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TN_USER) predicate.TN_USER {
	return predicate.TN_USER(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TN_USER) predicate.TN_USER {
	return predicate.TN_USER(sql.NotPredicates(p))
}
