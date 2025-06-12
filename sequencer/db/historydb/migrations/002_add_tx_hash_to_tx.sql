-- +migrate Up
ALTER TABLE tx ADD COLUMN tx_hash BYTEA UNIQUE;

-- +migrate Down
ALTER TABLE tx DROP COLUMN tx_hash; 