BEGIN;
INSERT INTO jack_black.players(
    login, password_hash)
VALUES ('test', '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08');

INSERT INTO jack_black.balances(
    player_id, balance)
VALUES (1, 1000);

INSERT INTO jack_black.sessions(
    player_id, current_action, last_message_id, player_hand, dealer_hand, deck, bet)
VALUES (1, '', 0, '[]', '[]', '[]', 0);

COMMIT;