CREATE TABLE IF NOT EXISTS jack_black.history
(
    history_id bigserial,
    player_id bigint NOT NULL,
    result_time timestamp with time zone NOT NULL,
    balance bigint NOT NULL,
    result text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT history_pkey PRIMARY KEY (history_id)
);

CREATE INDEX jb_history_player_id ON jack_black.history (player_id);
CREATE INDEX jb_history_result_time ON jack_black.history (result_time);
CREATE INDEX jb_history_result ON jack_black.history (result);
