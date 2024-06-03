// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/hawa130/computility-cloud/ent/casbinrule"
	"github.com/hawa130/computility-cloud/ent/publickey"
	"github.com/hawa130/computility-cloud/ent/user"
	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[xid.ID]
	PageInfo       = entgql.PageInfo[xid.ID]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CasbinRuleEdge is the edge representation of CasbinRule.
type CasbinRuleEdge struct {
	Node   *CasbinRule `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// CasbinRuleConnection is the connection containing edges to CasbinRule.
type CasbinRuleConnection struct {
	Edges      []*CasbinRuleEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *CasbinRuleConnection) build(nodes []*CasbinRule, pager *casbinrulePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *CasbinRule
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *CasbinRule {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *CasbinRule {
			return nodes[i]
		}
	}
	c.Edges = make([]*CasbinRuleEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CasbinRuleEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CasbinRulePaginateOption enables pagination customization.
type CasbinRulePaginateOption func(*casbinrulePager) error

// WithCasbinRuleOrder configures pagination ordering.
func WithCasbinRuleOrder(order *CasbinRuleOrder) CasbinRulePaginateOption {
	if order == nil {
		order = DefaultCasbinRuleOrder
	}
	o := *order
	return func(pager *casbinrulePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCasbinRuleOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCasbinRuleFilter configures pagination filter.
func WithCasbinRuleFilter(filter func(*CasbinRuleQuery) (*CasbinRuleQuery, error)) CasbinRulePaginateOption {
	return func(pager *casbinrulePager) error {
		if filter == nil {
			return errors.New("CasbinRuleQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type casbinrulePager struct {
	reverse bool
	order   *CasbinRuleOrder
	filter  func(*CasbinRuleQuery) (*CasbinRuleQuery, error)
}

func newCasbinRulePager(opts []CasbinRulePaginateOption, reverse bool) (*casbinrulePager, error) {
	pager := &casbinrulePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCasbinRuleOrder
	}
	return pager, nil
}

func (p *casbinrulePager) applyFilter(query *CasbinRuleQuery) (*CasbinRuleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *casbinrulePager) toCursor(cr *CasbinRule) Cursor {
	return p.order.Field.toCursor(cr)
}

func (p *casbinrulePager) applyCursors(query *CasbinRuleQuery, after, before *Cursor) (*CasbinRuleQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCasbinRuleOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *casbinrulePager) applyOrder(query *CasbinRuleQuery) *CasbinRuleQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCasbinRuleOrder.Field {
		query = query.Order(DefaultCasbinRuleOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *casbinrulePager) orderExpr(query *CasbinRuleQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCasbinRuleOrder.Field {
			b.Comma().Ident(DefaultCasbinRuleOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to CasbinRule.
func (cr *CasbinRuleQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CasbinRulePaginateOption,
) (*CasbinRuleConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCasbinRulePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if cr, err = pager.applyFilter(cr); err != nil {
		return nil, err
	}
	conn := &CasbinRuleConnection{Edges: []*CasbinRuleEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := cr.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if cr, err = pager.applyCursors(cr, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		cr.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := cr.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	cr = pager.applyOrder(cr)
	nodes, err := cr.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CasbinRuleOrderField defines the ordering field of CasbinRule.
type CasbinRuleOrderField struct {
	// Value extracts the ordering value from the given CasbinRule.
	Value    func(*CasbinRule) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) casbinrule.OrderOption
	toCursor func(*CasbinRule) Cursor
}

// CasbinRuleOrder defines the ordering of CasbinRule.
type CasbinRuleOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *CasbinRuleOrderField `json:"field"`
}

// DefaultCasbinRuleOrder is the default ordering of CasbinRule.
var DefaultCasbinRuleOrder = &CasbinRuleOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CasbinRuleOrderField{
		Value: func(cr *CasbinRule) (ent.Value, error) {
			return cr.ID, nil
		},
		column: casbinrule.FieldID,
		toTerm: casbinrule.ByID,
		toCursor: func(cr *CasbinRule) Cursor {
			return Cursor{ID: cr.ID}
		},
	},
}

// ToEdge converts CasbinRule into CasbinRuleEdge.
func (cr *CasbinRule) ToEdge(order *CasbinRuleOrder) *CasbinRuleEdge {
	if order == nil {
		order = DefaultCasbinRuleOrder
	}
	return &CasbinRuleEdge{
		Node:   cr,
		Cursor: order.Field.toCursor(cr),
	}
}

// PublicKeyEdge is the edge representation of PublicKey.
type PublicKeyEdge struct {
	Node   *PublicKey `json:"node"`
	Cursor Cursor     `json:"cursor"`
}

// PublicKeyConnection is the connection containing edges to PublicKey.
type PublicKeyConnection struct {
	Edges      []*PublicKeyEdge `json:"edges"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

func (c *PublicKeyConnection) build(nodes []*PublicKey, pager *publickeyPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *PublicKey
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *PublicKey {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *PublicKey {
			return nodes[i]
		}
	}
	c.Edges = make([]*PublicKeyEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PublicKeyEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PublicKeyPaginateOption enables pagination customization.
type PublicKeyPaginateOption func(*publickeyPager) error

// WithPublicKeyOrder configures pagination ordering.
func WithPublicKeyOrder(order *PublicKeyOrder) PublicKeyPaginateOption {
	if order == nil {
		order = DefaultPublicKeyOrder
	}
	o := *order
	return func(pager *publickeyPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPublicKeyOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPublicKeyFilter configures pagination filter.
func WithPublicKeyFilter(filter func(*PublicKeyQuery) (*PublicKeyQuery, error)) PublicKeyPaginateOption {
	return func(pager *publickeyPager) error {
		if filter == nil {
			return errors.New("PublicKeyQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type publickeyPager struct {
	reverse bool
	order   *PublicKeyOrder
	filter  func(*PublicKeyQuery) (*PublicKeyQuery, error)
}

func newPublicKeyPager(opts []PublicKeyPaginateOption, reverse bool) (*publickeyPager, error) {
	pager := &publickeyPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPublicKeyOrder
	}
	return pager, nil
}

func (p *publickeyPager) applyFilter(query *PublicKeyQuery) (*PublicKeyQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *publickeyPager) toCursor(pk *PublicKey) Cursor {
	return p.order.Field.toCursor(pk)
}

func (p *publickeyPager) applyCursors(query *PublicKeyQuery, after, before *Cursor) (*PublicKeyQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultPublicKeyOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *publickeyPager) applyOrder(query *PublicKeyQuery) *PublicKeyQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultPublicKeyOrder.Field {
		query = query.Order(DefaultPublicKeyOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *publickeyPager) orderExpr(query *PublicKeyQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPublicKeyOrder.Field {
			b.Comma().Ident(DefaultPublicKeyOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to PublicKey.
func (pk *PublicKeyQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PublicKeyPaginateOption,
) (*PublicKeyConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPublicKeyPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if pk, err = pager.applyFilter(pk); err != nil {
		return nil, err
	}
	conn := &PublicKeyConnection{Edges: []*PublicKeyEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := pk.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if pk, err = pager.applyCursors(pk, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		pk.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := pk.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	pk = pager.applyOrder(pk)
	nodes, err := pk.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// PublicKeyOrderFieldCreatedAt orders PublicKey by created_at.
	PublicKeyOrderFieldCreatedAt = &PublicKeyOrderField{
		Value: func(pk *PublicKey) (ent.Value, error) {
			return pk.CreatedAt, nil
		},
		column: publickey.FieldCreatedAt,
		toTerm: publickey.ByCreatedAt,
		toCursor: func(pk *PublicKey) Cursor {
			return Cursor{
				ID:    pk.ID,
				Value: pk.CreatedAt,
			}
		},
	}
	// PublicKeyOrderFieldUpdatedAt orders PublicKey by updated_at.
	PublicKeyOrderFieldUpdatedAt = &PublicKeyOrderField{
		Value: func(pk *PublicKey) (ent.Value, error) {
			return pk.UpdatedAt, nil
		},
		column: publickey.FieldUpdatedAt,
		toTerm: publickey.ByUpdatedAt,
		toCursor: func(pk *PublicKey) Cursor {
			return Cursor{
				ID:    pk.ID,
				Value: pk.UpdatedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f PublicKeyOrderField) String() string {
	var str string
	switch f.column {
	case PublicKeyOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	case PublicKeyOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f PublicKeyOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *PublicKeyOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("PublicKeyOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *PublicKeyOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *PublicKeyOrderFieldUpdatedAt
	default:
		return fmt.Errorf("%s is not a valid PublicKeyOrderField", str)
	}
	return nil
}

// PublicKeyOrderField defines the ordering field of PublicKey.
type PublicKeyOrderField struct {
	// Value extracts the ordering value from the given PublicKey.
	Value    func(*PublicKey) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) publickey.OrderOption
	toCursor func(*PublicKey) Cursor
}

// PublicKeyOrder defines the ordering of PublicKey.
type PublicKeyOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     *PublicKeyOrderField `json:"field"`
}

// DefaultPublicKeyOrder is the default ordering of PublicKey.
var DefaultPublicKeyOrder = &PublicKeyOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &PublicKeyOrderField{
		Value: func(pk *PublicKey) (ent.Value, error) {
			return pk.ID, nil
		},
		column: publickey.FieldID,
		toTerm: publickey.ByID,
		toCursor: func(pk *PublicKey) Cursor {
			return Cursor{ID: pk.ID}
		},
	},
}

// ToEdge converts PublicKey into PublicKeyEdge.
func (pk *PublicKey) ToEdge(order *PublicKeyOrder) *PublicKeyEdge {
	if order == nil {
		order = DefaultPublicKeyOrder
	}
	return &PublicKeyEdge{
		Node:   pk,
		Cursor: order.Field.toCursor(pk),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := u.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldCreatedAt orders User by created_at.
	UserOrderFieldCreatedAt = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.CreatedAt, nil
		},
		column: user.FieldCreatedAt,
		toTerm: user.ByCreatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.CreatedAt,
			}
		},
	}
	// UserOrderFieldUpdatedAt orders User by updated_at.
	UserOrderFieldUpdatedAt = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.UpdatedAt, nil
		},
		column: user.FieldUpdatedAt,
		toTerm: user.ByUpdatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.UpdatedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.column {
	case UserOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	case UserOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *UserOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *UserOrderFieldUpdatedAt
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}