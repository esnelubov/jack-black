CREATE TABLE IF NOT EXISTS jack_black.sessions
(
    session_id bigserial,
    player_id bigint NOT NULL,
    current_action text COLLATE pg_catalog."default" NOT NULL,
    last_message_id bigint NOT NULL,
    player_hand json NOT NULL,
    dealer_hand json NOT NULL,
    deck json NOT NULL,
    bet bigint NOT NULL,
    CONSTRAINT sessions_pkey PRIMARY KEY (session_id)
);

CREATE INDEX jb_sessions_player_id ON jack_black.sessions (player_id);
