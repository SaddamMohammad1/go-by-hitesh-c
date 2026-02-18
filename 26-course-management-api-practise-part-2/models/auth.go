package models

// ContextKey type for auth context keys
type ContextKey string

const (
	ContextKeyAuthorID ContextKey = "author_id"
	ContextKeyRole     ContextKey = "role"
)
