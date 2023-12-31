// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/panforyou/seller-finding/go/ent/migrate"

	"github.com/panforyou/seller-finding/go/ent/seller"
	"github.com/panforyou/seller-finding/go/ent/website"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Seller is the client for interacting with the Seller builders.
	Seller *SellerClient
	// Website is the client for interacting with the Website builders.
	Website *WebsiteClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Seller = NewSellerClient(c.config)
	c.Website = NewWebsiteClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Seller:  NewSellerClient(cfg),
		Website: NewWebsiteClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:  cfg,
		Seller:  NewSellerClient(cfg),
		Website: NewWebsiteClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Seller.
//		Query().
//		Count(ctx)
//
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
	c.Seller.Use(hooks...)
	c.Website.Use(hooks...)
}

// SellerClient is a client for the Seller schema.
type SellerClient struct {
	config
}

// NewSellerClient returns a client for the Seller from the given config.
func NewSellerClient(c config) *SellerClient {
	return &SellerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `seller.Hooks(f(g(h())))`.
func (c *SellerClient) Use(hooks ...Hook) {
	c.hooks.Seller = append(c.hooks.Seller, hooks...)
}

// Create returns a create builder for Seller.
func (c *SellerClient) Create() *SellerCreate {
	mutation := newSellerMutation(c.config, OpCreate)
	return &SellerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Seller entities.
func (c *SellerClient) CreateBulk(builders ...*SellerCreate) *SellerCreateBulk {
	return &SellerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Seller.
func (c *SellerClient) Update() *SellerUpdate {
	mutation := newSellerMutation(c.config, OpUpdate)
	return &SellerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SellerClient) UpdateOne(s *Seller) *SellerUpdateOne {
	mutation := newSellerMutation(c.config, OpUpdateOne, withSeller(s))
	return &SellerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SellerClient) UpdateOneID(id int) *SellerUpdateOne {
	mutation := newSellerMutation(c.config, OpUpdateOne, withSellerID(id))
	return &SellerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Seller.
func (c *SellerClient) Delete() *SellerDelete {
	mutation := newSellerMutation(c.config, OpDelete)
	return &SellerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SellerClient) DeleteOne(s *Seller) *SellerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SellerClient) DeleteOneID(id int) *SellerDeleteOne {
	builder := c.Delete().Where(seller.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SellerDeleteOne{builder}
}

// Query returns a query builder for Seller.
func (c *SellerClient) Query() *SellerQuery {
	return &SellerQuery{
		config: c.config,
	}
}

// Get returns a Seller entity by its id.
func (c *SellerClient) Get(ctx context.Context, id int) (*Seller, error) {
	return c.Query().Where(seller.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SellerClient) GetX(ctx context.Context, id int) *Seller {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SellerClient) Hooks() []Hook {
	return c.hooks.Seller
}

// WebsiteClient is a client for the Website schema.
type WebsiteClient struct {
	config
}

// NewWebsiteClient returns a client for the Website from the given config.
func NewWebsiteClient(c config) *WebsiteClient {
	return &WebsiteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `website.Hooks(f(g(h())))`.
func (c *WebsiteClient) Use(hooks ...Hook) {
	c.hooks.Website = append(c.hooks.Website, hooks...)
}

// Create returns a create builder for Website.
func (c *WebsiteClient) Create() *WebsiteCreate {
	mutation := newWebsiteMutation(c.config, OpCreate)
	return &WebsiteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Website entities.
func (c *WebsiteClient) CreateBulk(builders ...*WebsiteCreate) *WebsiteCreateBulk {
	return &WebsiteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Website.
func (c *WebsiteClient) Update() *WebsiteUpdate {
	mutation := newWebsiteMutation(c.config, OpUpdate)
	return &WebsiteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WebsiteClient) UpdateOne(w *Website) *WebsiteUpdateOne {
	mutation := newWebsiteMutation(c.config, OpUpdateOne, withWebsite(w))
	return &WebsiteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WebsiteClient) UpdateOneID(id int) *WebsiteUpdateOne {
	mutation := newWebsiteMutation(c.config, OpUpdateOne, withWebsiteID(id))
	return &WebsiteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Website.
func (c *WebsiteClient) Delete() *WebsiteDelete {
	mutation := newWebsiteMutation(c.config, OpDelete)
	return &WebsiteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WebsiteClient) DeleteOne(w *Website) *WebsiteDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WebsiteClient) DeleteOneID(id int) *WebsiteDeleteOne {
	builder := c.Delete().Where(website.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WebsiteDeleteOne{builder}
}

// Query returns a query builder for Website.
func (c *WebsiteClient) Query() *WebsiteQuery {
	return &WebsiteQuery{
		config: c.config,
	}
}

// Get returns a Website entity by its id.
func (c *WebsiteClient) Get(ctx context.Context, id int) (*Website, error) {
	return c.Query().Where(website.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WebsiteClient) GetX(ctx context.Context, id int) *Website {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WebsiteClient) Hooks() []Hook {
	return c.hooks.Website
}
