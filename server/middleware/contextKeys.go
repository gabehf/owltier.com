package middleware

type ContextKey string

const (
	ContextKeyUser   = ContextKey("user")
	ContextKeyToken  = ContextKey("token")
	ContextKeyValues = ContextKey("values")
)
