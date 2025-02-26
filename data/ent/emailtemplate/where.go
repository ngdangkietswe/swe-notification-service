// Code generated by ent, DO NOT EDIT.

package emailtemplate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldID, id))
}

// TemplateKey applies equality check predicate on the "template_key" field. It's identical to TemplateKeyEQ.
func TemplateKey(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldTemplateKey, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldName, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldSubject, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldBody, v))
}

// IsHTML applies equality check predicate on the "is_html" field. It's identical to IsHTMLEQ.
func IsHTML(v bool) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldIsHTML, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldIsDeleted, v))
}

// TemplateKeyEQ applies the EQ predicate on the "template_key" field.
func TemplateKeyEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldTemplateKey, v))
}

// TemplateKeyNEQ applies the NEQ predicate on the "template_key" field.
func TemplateKeyNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldTemplateKey, v))
}

// TemplateKeyIn applies the In predicate on the "template_key" field.
func TemplateKeyIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldTemplateKey, vs...))
}

// TemplateKeyNotIn applies the NotIn predicate on the "template_key" field.
func TemplateKeyNotIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldTemplateKey, vs...))
}

// TemplateKeyGT applies the GT predicate on the "template_key" field.
func TemplateKeyGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldTemplateKey, v))
}

// TemplateKeyGTE applies the GTE predicate on the "template_key" field.
func TemplateKeyGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldTemplateKey, v))
}

// TemplateKeyLT applies the LT predicate on the "template_key" field.
func TemplateKeyLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldTemplateKey, v))
}

// TemplateKeyLTE applies the LTE predicate on the "template_key" field.
func TemplateKeyLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldTemplateKey, v))
}

// TemplateKeyContains applies the Contains predicate on the "template_key" field.
func TemplateKeyContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContains(FieldTemplateKey, v))
}

// TemplateKeyHasPrefix applies the HasPrefix predicate on the "template_key" field.
func TemplateKeyHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasPrefix(FieldTemplateKey, v))
}

// TemplateKeyHasSuffix applies the HasSuffix predicate on the "template_key" field.
func TemplateKeyHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasSuffix(FieldTemplateKey, v))
}

// TemplateKeyEqualFold applies the EqualFold predicate on the "template_key" field.
func TemplateKeyEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEqualFold(FieldTemplateKey, v))
}

// TemplateKeyContainsFold applies the ContainsFold predicate on the "template_key" field.
func TemplateKeyContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContainsFold(FieldTemplateKey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContainsFold(FieldName, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContainsFold(FieldSubject, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldContainsFold(FieldBody, v))
}

// IsHTMLEQ applies the EQ predicate on the "is_html" field.
func IsHTMLEQ(v bool) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldIsHTML, v))
}

// IsHTMLNEQ applies the NEQ predicate on the "is_html" field.
func IsHTMLNEQ(v bool) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldIsHTML, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.FieldNEQ(FieldIsDeleted, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailTemplate) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailTemplate) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EmailTemplate) predicate.EmailTemplate {
	return predicate.EmailTemplate(sql.NotPredicates(p))
}
