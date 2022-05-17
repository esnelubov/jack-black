BEGIN;
DELETE FROM jack_black.players
WHERE player_id=1;

DELETE FROM jack_black.balances
WHERE player_id=1;

DELETE FROM jack_black.history
WHERE player_id=1;

DELETE FROM jack_black.sessions
WHERE player_id=1;

COMMIT;