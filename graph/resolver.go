package graph

import (
	"github.com/cjpoulsen/wallet/graph/model"
	"github.com/google/uuid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	Wallets map[uuid.UUID]*model.Wallet
}
