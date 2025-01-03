-- name: GetWallet :one
SELECT * FROM wallet
WHERE id = $1 LIMIT 1;

-- name: CreateWallet :one
INSERT INTO wallet (id, balance, create_date, update_date)
VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateBalance :exec
UPDATE wallet
   SET balance = $2,
       update_date = $3
WHERE id = $1;
