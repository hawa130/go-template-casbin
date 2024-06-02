// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/rs/xid"
)

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Nickname *string
	Username *string
	Email    *string
	Phone    string
	Password string
	ChildIDs []xid.ID
	ParentID *xid.ID
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.Nickname; v != nil {
		m.SetNickname(*v)
	}
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	m.SetPhone(i.Phone)
	m.SetPassword(i.Password)
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	ClearNickname  bool
	Nickname       *string
	ClearUsername  bool
	Username       *string
	ClearEmail     bool
	Email          *string
	Phone          *string
	Password       *string
	ClearChildren  bool
	AddChildIDs    []xid.ID
	RemoveChildIDs []xid.ID
	ClearParent    bool
	ParentID       *xid.ID
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if i.ClearNickname {
		m.ClearNickname()
	}
	if v := i.Nickname; v != nil {
		m.SetNickname(*v)
	}
	if i.ClearUsername {
		m.ClearUsername()
	}
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if i.ClearEmail {
		m.ClearEmail()
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Phone; v != nil {
		m.SetPhone(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if i.ClearChildren {
		m.ClearChildren()
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
