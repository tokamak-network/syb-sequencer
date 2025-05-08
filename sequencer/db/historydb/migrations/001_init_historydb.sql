-- +migrate Up

-- NOTE: We use "DECIMAL(78,0)" to encode go *big.Int types.  All the *big.Int
-- that we deal with represent a value in the SNARK field, which is an integer
-- of 256 bits.  `log(2**256, 10) = 77.06`: that is, a 256 bit number can have
-- at most 78 digits, so we use this value to specify the precision in the
-- PostgreSQL DECIMAL guaranteeing that we will never lose precision.

-- Create sequence for tx_item_id
CREATE SEQUENCE IF NOT EXISTS tx_item_id;

-- Create tx scheme
CREATE TABLE IF NOT EXISTS tx (
    item_id INTEGER PRIMARY KEY DEFAULT nextval('tx_item_id'),
    batch_num BIGINT NOT NULL,
    position INT NOT NULL,
    type VARCHAR(40) NOT NULL,
    from_idx BIGINT,
    from_eth_addr STRING,
    to_idx BIGINT NOT NULL,
    to_eth_addr STRING,
    amount DECIMAL(78,0) NOT NULL
);

-- Create batch scheme
CREATE TABLE IF NOT EXISTS batch (
    item_id SERIAL PRIMARY KEY,
    account_root DECIMAL(78,0) NOT NULL,
    vouch_root DECIMAL(78,0) NOT NULL,
    score_root DECIMAL(78,0) NOT NULL
);

-- Create account scheme
CREATE TABLE IF NOT EXISTS account (
    item_id SERIAL,
    idx BIGINT PRIMARY KEY,
    eth_addr STRING NOT NULL,
    balance DECIMAL(78,0) NOT NULL,
    score DECIMAL(78,0) NOT NULL,
    score_siblings DECIMAL(78,0)[NLEVELS] NOT NULL
);

-- Create vouch scheme
CREATE TABLE IF NOT EXISTS vouch (
    idx BIGINT PRIMARY KEY,
    from_idx BIGINT,
    from_eth_addr STRING,
    to_idx BIGINT,
    to_eth_addr STRING
);

-- +migrate Down
DROP TABLE IF EXISTS tx;
DROP TABLE IF EXISTS batch;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS vouch;