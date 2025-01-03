// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package sql

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createWallet = `-- name: CreateWallet :one
INSERT INTO wallet (id, balance, create_date, update_date)
VALUES (
    $1, $2, $3, $4
)
RETURNING id, balance, create_date, update_date
`

type CreateWalletParams struct {
	ID         uuid.UUID
	Balance    int32
	CreateDate time.Time
	UpdateDate time.Time
}

func (q *Queries) CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, createWallet,
		arg.ID,
		arg.Balance,
		arg.CreateDate,
		arg.UpdateDate,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.CreateDate,
		&i.UpdateDate,
	)
	return i, err
}

const getWallet = `-- name: GetWallet :one
SELECT id, balance, create_date, update_date FROM wallet
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWallet(ctx context.Context, id uuid.UUID) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, getWallet, id)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.CreateDate,
		&i.UpdateDate,
	)
	return i, err
}

const updateBalance = `-- name: UpdateBalance :exec
UPDATE wallet
   SET balance = $2,
       update_date = $3
WHERE id = $1
`

type UpdateBalanceParams struct {
	ID         uuid.UUID
	Balance    int32
	UpdateDate time.Time
}

func (q *Queries) UpdateBalance(ctx context.Context, arg UpdateBalanceParams) error {
	_, err := q.db.ExecContext(ctx, updateBalance, arg.ID, arg.Balance, arg.UpdateDate)
	return err
}
