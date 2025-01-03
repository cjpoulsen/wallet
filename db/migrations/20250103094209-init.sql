
-- +migrate Up
CREATE TABLE wallet (id uuid, balance int)
-- +migrate Down
DROP TABLE wallet
