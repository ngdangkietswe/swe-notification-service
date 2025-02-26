// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/cdcauthusers"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/emailtemplate"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCdcAuthUsers  = "CdcAuthUsers"
	TypeEmailTemplate = "EmailTemplate"
)

// CdcAuthUsersMutation represents an operation that mutates the CdcAuthUsers nodes in the graph.
type CdcAuthUsersMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	username      *string
	email         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*CdcAuthUsers, error)
	predicates    []predicate.CdcAuthUsers
}

var _ ent.Mutation = (*CdcAuthUsersMutation)(nil)

// cdcauthusersOption allows management of the mutation configuration using functional options.
type cdcauthusersOption func(*CdcAuthUsersMutation)

// newCdcAuthUsersMutation creates new mutation for the CdcAuthUsers entity.
func newCdcAuthUsersMutation(c config, op Op, opts ...cdcauthusersOption) *CdcAuthUsersMutation {
	m := &CdcAuthUsersMutation{
		config:        c,
		op:            op,
		typ:           TypeCdcAuthUsers,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCdcAuthUsersID sets the ID field of the mutation.
func withCdcAuthUsersID(id uuid.UUID) cdcauthusersOption {
	return func(m *CdcAuthUsersMutation) {
		var (
			err   error
			once  sync.Once
			value *CdcAuthUsers
		)
		m.oldValue = func(ctx context.Context) (*CdcAuthUsers, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().CdcAuthUsers.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCdcAuthUsers sets the old CdcAuthUsers of the mutation.
func withCdcAuthUsers(node *CdcAuthUsers) cdcauthusersOption {
	return func(m *CdcAuthUsersMutation) {
		m.oldValue = func(context.Context) (*CdcAuthUsers, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CdcAuthUsersMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CdcAuthUsersMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of CdcAuthUsers entities.
func (m *CdcAuthUsersMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CdcAuthUsersMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CdcAuthUsersMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().CdcAuthUsers.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *CdcAuthUsersMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *CdcAuthUsersMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the CdcAuthUsers entity.
// If the CdcAuthUsers object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CdcAuthUsersMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *CdcAuthUsersMutation) ResetUsername() {
	m.username = nil
}

// SetEmail sets the "email" field.
func (m *CdcAuthUsersMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *CdcAuthUsersMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the CdcAuthUsers entity.
// If the CdcAuthUsers object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CdcAuthUsersMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *CdcAuthUsersMutation) ResetEmail() {
	m.email = nil
}

// Where appends a list predicates to the CdcAuthUsersMutation builder.
func (m *CdcAuthUsersMutation) Where(ps ...predicate.CdcAuthUsers) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CdcAuthUsersMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CdcAuthUsersMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.CdcAuthUsers, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CdcAuthUsersMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CdcAuthUsersMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (CdcAuthUsers).
func (m *CdcAuthUsersMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CdcAuthUsersMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.username != nil {
		fields = append(fields, cdcauthusers.FieldUsername)
	}
	if m.email != nil {
		fields = append(fields, cdcauthusers.FieldEmail)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CdcAuthUsersMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case cdcauthusers.FieldUsername:
		return m.Username()
	case cdcauthusers.FieldEmail:
		return m.Email()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CdcAuthUsersMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case cdcauthusers.FieldUsername:
		return m.OldUsername(ctx)
	case cdcauthusers.FieldEmail:
		return m.OldEmail(ctx)
	}
	return nil, fmt.Errorf("unknown CdcAuthUsers field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CdcAuthUsersMutation) SetField(name string, value ent.Value) error {
	switch name {
	case cdcauthusers.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case cdcauthusers.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	}
	return fmt.Errorf("unknown CdcAuthUsers field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CdcAuthUsersMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CdcAuthUsersMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CdcAuthUsersMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown CdcAuthUsers numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CdcAuthUsersMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CdcAuthUsersMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CdcAuthUsersMutation) ClearField(name string) error {
	return fmt.Errorf("unknown CdcAuthUsers nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CdcAuthUsersMutation) ResetField(name string) error {
	switch name {
	case cdcauthusers.FieldUsername:
		m.ResetUsername()
		return nil
	case cdcauthusers.FieldEmail:
		m.ResetEmail()
		return nil
	}
	return fmt.Errorf("unknown CdcAuthUsers field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CdcAuthUsersMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CdcAuthUsersMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CdcAuthUsersMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CdcAuthUsersMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CdcAuthUsersMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CdcAuthUsersMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CdcAuthUsersMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown CdcAuthUsers unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CdcAuthUsersMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown CdcAuthUsers edge %s", name)
}

// EmailTemplateMutation represents an operation that mutates the EmailTemplate nodes in the graph.
type EmailTemplateMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	template_key  *string
	name          *string
	subject       *string
	body          *string
	is_html       *bool
	created_at    *time.Time
	updated_at    *time.Time
	is_deleted    *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*EmailTemplate, error)
	predicates    []predicate.EmailTemplate
}

var _ ent.Mutation = (*EmailTemplateMutation)(nil)

// emailtemplateOption allows management of the mutation configuration using functional options.
type emailtemplateOption func(*EmailTemplateMutation)

// newEmailTemplateMutation creates new mutation for the EmailTemplate entity.
func newEmailTemplateMutation(c config, op Op, opts ...emailtemplateOption) *EmailTemplateMutation {
	m := &EmailTemplateMutation{
		config:        c,
		op:            op,
		typ:           TypeEmailTemplate,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEmailTemplateID sets the ID field of the mutation.
func withEmailTemplateID(id uuid.UUID) emailtemplateOption {
	return func(m *EmailTemplateMutation) {
		var (
			err   error
			once  sync.Once
			value *EmailTemplate
		)
		m.oldValue = func(ctx context.Context) (*EmailTemplate, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().EmailTemplate.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEmailTemplate sets the old EmailTemplate of the mutation.
func withEmailTemplate(node *EmailTemplate) emailtemplateOption {
	return func(m *EmailTemplateMutation) {
		m.oldValue = func(context.Context) (*EmailTemplate, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EmailTemplateMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EmailTemplateMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of EmailTemplate entities.
func (m *EmailTemplateMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EmailTemplateMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EmailTemplateMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().EmailTemplate.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTemplateKey sets the "template_key" field.
func (m *EmailTemplateMutation) SetTemplateKey(s string) {
	m.template_key = &s
}

// TemplateKey returns the value of the "template_key" field in the mutation.
func (m *EmailTemplateMutation) TemplateKey() (r string, exists bool) {
	v := m.template_key
	if v == nil {
		return
	}
	return *v, true
}

// OldTemplateKey returns the old "template_key" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldTemplateKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTemplateKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTemplateKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTemplateKey: %w", err)
	}
	return oldValue.TemplateKey, nil
}

// ResetTemplateKey resets all changes to the "template_key" field.
func (m *EmailTemplateMutation) ResetTemplateKey() {
	m.template_key = nil
}

// SetName sets the "name" field.
func (m *EmailTemplateMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EmailTemplateMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EmailTemplateMutation) ResetName() {
	m.name = nil
}

// SetSubject sets the "subject" field.
func (m *EmailTemplateMutation) SetSubject(s string) {
	m.subject = &s
}

// Subject returns the value of the "subject" field in the mutation.
func (m *EmailTemplateMutation) Subject() (r string, exists bool) {
	v := m.subject
	if v == nil {
		return
	}
	return *v, true
}

// OldSubject returns the old "subject" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldSubject(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSubject is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSubject requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSubject: %w", err)
	}
	return oldValue.Subject, nil
}

// ResetSubject resets all changes to the "subject" field.
func (m *EmailTemplateMutation) ResetSubject() {
	m.subject = nil
}

// SetBody sets the "body" field.
func (m *EmailTemplateMutation) SetBody(s string) {
	m.body = &s
}

// Body returns the value of the "body" field in the mutation.
func (m *EmailTemplateMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "body" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody resets all changes to the "body" field.
func (m *EmailTemplateMutation) ResetBody() {
	m.body = nil
}

// SetIsHTML sets the "is_html" field.
func (m *EmailTemplateMutation) SetIsHTML(b bool) {
	m.is_html = &b
}

// IsHTML returns the value of the "is_html" field in the mutation.
func (m *EmailTemplateMutation) IsHTML() (r bool, exists bool) {
	v := m.is_html
	if v == nil {
		return
	}
	return *v, true
}

// OldIsHTML returns the old "is_html" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldIsHTML(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsHTML is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsHTML requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsHTML: %w", err)
	}
	return oldValue.IsHTML, nil
}

// ResetIsHTML resets all changes to the "is_html" field.
func (m *EmailTemplateMutation) ResetIsHTML() {
	m.is_html = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *EmailTemplateMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EmailTemplateMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EmailTemplateMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *EmailTemplateMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *EmailTemplateMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *EmailTemplateMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetIsDeleted sets the "is_deleted" field.
func (m *EmailTemplateMutation) SetIsDeleted(b bool) {
	m.is_deleted = &b
}

// IsDeleted returns the value of the "is_deleted" field in the mutation.
func (m *EmailTemplateMutation) IsDeleted() (r bool, exists bool) {
	v := m.is_deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDeleted returns the old "is_deleted" field's value of the EmailTemplate entity.
// If the EmailTemplate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailTemplateMutation) OldIsDeleted(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDeleted: %w", err)
	}
	return oldValue.IsDeleted, nil
}

// ResetIsDeleted resets all changes to the "is_deleted" field.
func (m *EmailTemplateMutation) ResetIsDeleted() {
	m.is_deleted = nil
}

// Where appends a list predicates to the EmailTemplateMutation builder.
func (m *EmailTemplateMutation) Where(ps ...predicate.EmailTemplate) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EmailTemplateMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EmailTemplateMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.EmailTemplate, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EmailTemplateMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EmailTemplateMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (EmailTemplate).
func (m *EmailTemplateMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EmailTemplateMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.template_key != nil {
		fields = append(fields, emailtemplate.FieldTemplateKey)
	}
	if m.name != nil {
		fields = append(fields, emailtemplate.FieldName)
	}
	if m.subject != nil {
		fields = append(fields, emailtemplate.FieldSubject)
	}
	if m.body != nil {
		fields = append(fields, emailtemplate.FieldBody)
	}
	if m.is_html != nil {
		fields = append(fields, emailtemplate.FieldIsHTML)
	}
	if m.created_at != nil {
		fields = append(fields, emailtemplate.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, emailtemplate.FieldUpdatedAt)
	}
	if m.is_deleted != nil {
		fields = append(fields, emailtemplate.FieldIsDeleted)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EmailTemplateMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case emailtemplate.FieldTemplateKey:
		return m.TemplateKey()
	case emailtemplate.FieldName:
		return m.Name()
	case emailtemplate.FieldSubject:
		return m.Subject()
	case emailtemplate.FieldBody:
		return m.Body()
	case emailtemplate.FieldIsHTML:
		return m.IsHTML()
	case emailtemplate.FieldCreatedAt:
		return m.CreatedAt()
	case emailtemplate.FieldUpdatedAt:
		return m.UpdatedAt()
	case emailtemplate.FieldIsDeleted:
		return m.IsDeleted()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EmailTemplateMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case emailtemplate.FieldTemplateKey:
		return m.OldTemplateKey(ctx)
	case emailtemplate.FieldName:
		return m.OldName(ctx)
	case emailtemplate.FieldSubject:
		return m.OldSubject(ctx)
	case emailtemplate.FieldBody:
		return m.OldBody(ctx)
	case emailtemplate.FieldIsHTML:
		return m.OldIsHTML(ctx)
	case emailtemplate.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case emailtemplate.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case emailtemplate.FieldIsDeleted:
		return m.OldIsDeleted(ctx)
	}
	return nil, fmt.Errorf("unknown EmailTemplate field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailTemplateMutation) SetField(name string, value ent.Value) error {
	switch name {
	case emailtemplate.FieldTemplateKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTemplateKey(v)
		return nil
	case emailtemplate.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case emailtemplate.FieldSubject:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSubject(v)
		return nil
	case emailtemplate.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case emailtemplate.FieldIsHTML:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsHTML(v)
		return nil
	case emailtemplate.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case emailtemplate.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case emailtemplate.FieldIsDeleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDeleted(v)
		return nil
	}
	return fmt.Errorf("unknown EmailTemplate field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EmailTemplateMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EmailTemplateMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailTemplateMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown EmailTemplate numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EmailTemplateMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EmailTemplateMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EmailTemplateMutation) ClearField(name string) error {
	return fmt.Errorf("unknown EmailTemplate nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EmailTemplateMutation) ResetField(name string) error {
	switch name {
	case emailtemplate.FieldTemplateKey:
		m.ResetTemplateKey()
		return nil
	case emailtemplate.FieldName:
		m.ResetName()
		return nil
	case emailtemplate.FieldSubject:
		m.ResetSubject()
		return nil
	case emailtemplate.FieldBody:
		m.ResetBody()
		return nil
	case emailtemplate.FieldIsHTML:
		m.ResetIsHTML()
		return nil
	case emailtemplate.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case emailtemplate.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case emailtemplate.FieldIsDeleted:
		m.ResetIsDeleted()
		return nil
	}
	return fmt.Errorf("unknown EmailTemplate field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EmailTemplateMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EmailTemplateMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EmailTemplateMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EmailTemplateMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EmailTemplateMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EmailTemplateMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EmailTemplateMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown EmailTemplate unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EmailTemplateMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown EmailTemplate edge %s", name)
}
