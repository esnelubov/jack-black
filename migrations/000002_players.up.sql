CREATE TABLE IF NOT EXISTS jack_black.players
(
    player_id bigserial,
    login text COLLATE pg_catalog."default" NOT NULL,
    password_hash text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT players_pkey PRIMARY KEY (player_id),
    CONSTRAINT unique_login UNIQUE (login)
);