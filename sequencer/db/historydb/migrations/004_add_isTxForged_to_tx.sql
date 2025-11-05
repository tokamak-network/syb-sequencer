-- +migrate Up
ALTER TABLE tx
    ADD COLUMN IF NOT EXISTS is_tx_forged BOOLEAN NOT NULL DEFAULT FALSE;

-- +migrate Down
ALTER TABLE tx DROP COLUMN IF EXISTS is_tx_forged;