
-- +migrate Up
CREATE TABLE wallet (id uuid PRIMARY KEY, balance INT NOT NULL, create_date TIMESTAMP NOT NULL, update_date TIMESTAMP NOT NULL);
-- +migrate Down
DROP TABLE wallet;
