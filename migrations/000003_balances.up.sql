CREATE TABLE IF NOT EXISTS jack_black.balances
(
    balance_id bigserial,
    player_id bigint NOT NULL,
    balance bigint NOT NULL DEFAULT 0,
    CONSTRAINT balances_pkey PRIMARY KEY (balance_id),
    CONSTRAINT player_id_unique UNIQUE (player_id)
);
