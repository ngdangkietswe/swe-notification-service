// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/cdcauthusers"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/emailtemplate"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CdcAuthUsers is the client for interacting with the CdcAuthUsers builders.
	CdcAuthUsers *CdcAuthUsersClient
	// EmailTemplate is the client for interacting with the EmailTemplate builders.
	EmailTemplate *EmailTemplateClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CdcAuthUsers = NewCdcAuthUsersClient(c.config)
	c.EmailTemplate = NewEmailTemplateClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		CdcAuthUsers:  NewCdcAuthUsersClient(cfg),
		EmailTemplate: NewEmailTemplateClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		CdcAuthUsers:  NewCdcAuthUsersClient(cfg),
		EmailTemplate: NewEmailTemplateClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CdcAuthUsers.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CdcAuthUsers.Use(hooks...)
	c.EmailTemplate.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.CdcAuthUsers.Intercept(interceptors...)
	c.EmailTemplate.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CdcAuthUsersMutation:
		return c.CdcAuthUsers.mutate(ctx, m)
	case *EmailTemplateMutation:
		return c.EmailTemplate.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CdcAuthUsersClient is a client for the CdcAuthUsers schema.
type CdcAuthUsersClient struct {
	config
}

// NewCdcAuthUsersClient returns a client for the CdcAuthUsers from the given config.
func NewCdcAuthUsersClient(c config) *CdcAuthUsersClient {
	return &CdcAuthUsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cdcauthusers.Hooks(f(g(h())))`.
func (c *CdcAuthUsersClient) Use(hooks ...Hook) {
	c.hooks.CdcAuthUsers = append(c.hooks.CdcAuthUsers, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cdcauthusers.Intercept(f(g(h())))`.
func (c *CdcAuthUsersClient) Intercept(interceptors ...Interceptor) {
	c.inters.CdcAuthUsers = append(c.inters.CdcAuthUsers, interceptors...)
}

// Create returns a builder for creating a CdcAuthUsers entity.
func (c *CdcAuthUsersClient) Create() *CdcAuthUsersCreate {
	mutation := newCdcAuthUsersMutation(c.config, OpCreate)
	return &CdcAuthUsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CdcAuthUsers entities.
func (c *CdcAuthUsersClient) CreateBulk(builders ...*CdcAuthUsersCreate) *CdcAuthUsersCreateBulk {
	return &CdcAuthUsersCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CdcAuthUsersClient) MapCreateBulk(slice any, setFunc func(*CdcAuthUsersCreate, int)) *CdcAuthUsersCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CdcAuthUsersCreateBulk{err: fmt.Errorf("calling to CdcAuthUsersClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CdcAuthUsersCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CdcAuthUsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CdcAuthUsers.
func (c *CdcAuthUsersClient) Update() *CdcAuthUsersUpdate {
	mutation := newCdcAuthUsersMutation(c.config, OpUpdate)
	return &CdcAuthUsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CdcAuthUsersClient) UpdateOne(cau *CdcAuthUsers) *CdcAuthUsersUpdateOne {
	mutation := newCdcAuthUsersMutation(c.config, OpUpdateOne, withCdcAuthUsers(cau))
	return &CdcAuthUsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CdcAuthUsersClient) UpdateOneID(id uuid.UUID) *CdcAuthUsersUpdateOne {
	mutation := newCdcAuthUsersMutation(c.config, OpUpdateOne, withCdcAuthUsersID(id))
	return &CdcAuthUsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CdcAuthUsers.
func (c *CdcAuthUsersClient) Delete() *CdcAuthUsersDelete {
	mutation := newCdcAuthUsersMutation(c.config, OpDelete)
	return &CdcAuthUsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CdcAuthUsersClient) DeleteOne(cau *CdcAuthUsers) *CdcAuthUsersDeleteOne {
	return c.DeleteOneID(cau.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CdcAuthUsersClient) DeleteOneID(id uuid.UUID) *CdcAuthUsersDeleteOne {
	builder := c.Delete().Where(cdcauthusers.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CdcAuthUsersDeleteOne{builder}
}

// Query returns a query builder for CdcAuthUsers.
func (c *CdcAuthUsersClient) Query() *CdcAuthUsersQuery {
	return &CdcAuthUsersQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCdcAuthUsers},
		inters: c.Interceptors(),
	}
}

// Get returns a CdcAuthUsers entity by its id.
func (c *CdcAuthUsersClient) Get(ctx context.Context, id uuid.UUID) (*CdcAuthUsers, error) {
	return c.Query().Where(cdcauthusers.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CdcAuthUsersClient) GetX(ctx context.Context, id uuid.UUID) *CdcAuthUsers {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CdcAuthUsersClient) Hooks() []Hook {
	return c.hooks.CdcAuthUsers
}

// Interceptors returns the client interceptors.
func (c *CdcAuthUsersClient) Interceptors() []Interceptor {
	return c.inters.CdcAuthUsers
}

func (c *CdcAuthUsersClient) mutate(ctx context.Context, m *CdcAuthUsersMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CdcAuthUsersCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CdcAuthUsersUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CdcAuthUsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CdcAuthUsersDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CdcAuthUsers mutation op: %q", m.Op())
	}
}

// EmailTemplateClient is a client for the EmailTemplate schema.
type EmailTemplateClient struct {
	config
}

// NewEmailTemplateClient returns a client for the EmailTemplate from the given config.
func NewEmailTemplateClient(c config) *EmailTemplateClient {
	return &EmailTemplateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `emailtemplate.Hooks(f(g(h())))`.
func (c *EmailTemplateClient) Use(hooks ...Hook) {
	c.hooks.EmailTemplate = append(c.hooks.EmailTemplate, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `emailtemplate.Intercept(f(g(h())))`.
func (c *EmailTemplateClient) Intercept(interceptors ...Interceptor) {
	c.inters.EmailTemplate = append(c.inters.EmailTemplate, interceptors...)
}

// Create returns a builder for creating a EmailTemplate entity.
func (c *EmailTemplateClient) Create() *EmailTemplateCreate {
	mutation := newEmailTemplateMutation(c.config, OpCreate)
	return &EmailTemplateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EmailTemplate entities.
func (c *EmailTemplateClient) CreateBulk(builders ...*EmailTemplateCreate) *EmailTemplateCreateBulk {
	return &EmailTemplateCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EmailTemplateClient) MapCreateBulk(slice any, setFunc func(*EmailTemplateCreate, int)) *EmailTemplateCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EmailTemplateCreateBulk{err: fmt.Errorf("calling to EmailTemplateClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EmailTemplateCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EmailTemplateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EmailTemplate.
func (c *EmailTemplateClient) Update() *EmailTemplateUpdate {
	mutation := newEmailTemplateMutation(c.config, OpUpdate)
	return &EmailTemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmailTemplateClient) UpdateOne(et *EmailTemplate) *EmailTemplateUpdateOne {
	mutation := newEmailTemplateMutation(c.config, OpUpdateOne, withEmailTemplate(et))
	return &EmailTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmailTemplateClient) UpdateOneID(id uuid.UUID) *EmailTemplateUpdateOne {
	mutation := newEmailTemplateMutation(c.config, OpUpdateOne, withEmailTemplateID(id))
	return &EmailTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EmailTemplate.
func (c *EmailTemplateClient) Delete() *EmailTemplateDelete {
	mutation := newEmailTemplateMutation(c.config, OpDelete)
	return &EmailTemplateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EmailTemplateClient) DeleteOne(et *EmailTemplate) *EmailTemplateDeleteOne {
	return c.DeleteOneID(et.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EmailTemplateClient) DeleteOneID(id uuid.UUID) *EmailTemplateDeleteOne {
	builder := c.Delete().Where(emailtemplate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmailTemplateDeleteOne{builder}
}

// Query returns a query builder for EmailTemplate.
func (c *EmailTemplateClient) Query() *EmailTemplateQuery {
	return &EmailTemplateQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEmailTemplate},
		inters: c.Interceptors(),
	}
}

// Get returns a EmailTemplate entity by its id.
func (c *EmailTemplateClient) Get(ctx context.Context, id uuid.UUID) (*EmailTemplate, error) {
	return c.Query().Where(emailtemplate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmailTemplateClient) GetX(ctx context.Context, id uuid.UUID) *EmailTemplate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmailTemplateClient) Hooks() []Hook {
	return c.hooks.EmailTemplate
}

// Interceptors returns the client interceptors.
func (c *EmailTemplateClient) Interceptors() []Interceptor {
	return c.inters.EmailTemplate
}

func (c *EmailTemplateClient) mutate(ctx context.Context, m *EmailTemplateMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EmailTemplateCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EmailTemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EmailTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EmailTemplateDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EmailTemplate mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		CdcAuthUsers, EmailTemplate []ent.Hook
	}
	inters struct {
		CdcAuthUsers, EmailTemplate []ent.Interceptor
	}
)
