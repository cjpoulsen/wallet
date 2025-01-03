package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

import (
	"database/sql"
	sqlc "github.com/cjpoulsen/wallet/sql"
)

type Resolver struct {
	*sqlc.Queries
	DB *sql.DB
}
