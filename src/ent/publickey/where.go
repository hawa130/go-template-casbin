// Code generated by ent, DO NOT EDIT.

package publickey

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hawa130/computility-cloud/ent/predicate"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldUpdatedAt, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldKey, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldDescription, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldType, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldStatus, v))
}

// ExpiredAt applies equality check predicate on the "expired_at" field. It's identical to ExpiredAtEQ.
func ExpiredAt(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldExpiredAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldUpdatedAt, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContainsFold(FieldKey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContainsFold(FieldType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldContainsFold(FieldStatus, v))
}

// ExpiredAtEQ applies the EQ predicate on the "expired_at" field.
func ExpiredAtEQ(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldEQ(FieldExpiredAt, v))
}

// ExpiredAtNEQ applies the NEQ predicate on the "expired_at" field.
func ExpiredAtNEQ(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNEQ(FieldExpiredAt, v))
}

// ExpiredAtIn applies the In predicate on the "expired_at" field.
func ExpiredAtIn(vs ...time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIn(FieldExpiredAt, vs...))
}

// ExpiredAtNotIn applies the NotIn predicate on the "expired_at" field.
func ExpiredAtNotIn(vs ...time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotIn(FieldExpiredAt, vs...))
}

// ExpiredAtGT applies the GT predicate on the "expired_at" field.
func ExpiredAtGT(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGT(FieldExpiredAt, v))
}

// ExpiredAtGTE applies the GTE predicate on the "expired_at" field.
func ExpiredAtGTE(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldGTE(FieldExpiredAt, v))
}

// ExpiredAtLT applies the LT predicate on the "expired_at" field.
func ExpiredAtLT(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLT(FieldExpiredAt, v))
}

// ExpiredAtLTE applies the LTE predicate on the "expired_at" field.
func ExpiredAtLTE(v time.Time) predicate.PublicKey {
	return predicate.PublicKey(sql.FieldLTE(FieldExpiredAt, v))
}

// ExpiredAtIsNil applies the IsNil predicate on the "expired_at" field.
func ExpiredAtIsNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldIsNull(FieldExpiredAt))
}

// ExpiredAtNotNil applies the NotNil predicate on the "expired_at" field.
func ExpiredAtNotNil() predicate.PublicKey {
	return predicate.PublicKey(sql.FieldNotNull(FieldExpiredAt))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.PublicKey {
	return predicate.PublicKey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.PublicKey {
	return predicate.PublicKey(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PublicKey) predicate.PublicKey {
	return predicate.PublicKey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PublicKey) predicate.PublicKey {
	return predicate.PublicKey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PublicKey) predicate.PublicKey {
	return predicate.PublicKey(sql.NotPredicates(p))
}
