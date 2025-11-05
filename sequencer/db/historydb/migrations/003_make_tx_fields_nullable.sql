-- +migrate Up
-- Make tx fields nullable to match Go struct with omitempty tags
ALTER TABLE tx ALTER COLUMN position DROP NOT NULL;
ALTER TABLE tx ALTER COLUMN to_idx DROP NOT NULL;
ALTER TABLE tx ALTER COLUMN amount DROP NOT NULL;

-- +migrate Down
-- Revert the changes
ALTER TABLE tx ALTER COLUMN position SET NOT NULL;
ALTER TABLE tx ALTER COLUMN to_idx SET NOT NULL;
ALTER TABLE tx ALTER COLUMN amount SET NOT NULL;
