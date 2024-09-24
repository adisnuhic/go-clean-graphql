package models

import "time"

// AuthProvider -
type AuthProvider struct {
	Provider  string     `json:"provider"`
	UserID    uint64     `json:"user_id"`
	UID       string     `json:"uid"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Post -
type Post struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserID      uint64     `json:"user_id"`
	User        *User      `json:"user"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// User -
type User struct {
	ID          uint64     `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	IsConfirmed *bool      `json:"is_confirmed,omitempty"`
	AcceptedTos bool       `json:"accepted_tos"`
	Posts       []*Post    `json:"posts"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// Token -
type Token struct {
	ID          uint64     `json:"id"`
	UserID      uint64     `json:"user_id"`
	Token       string     `json:"token"`
	TokenTypeID uint64     `json:"token_type_id"`
	Code        *string    `json:"code,omitempty"`
	ExpiresAt   time.Time  `json:"expires_at"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
