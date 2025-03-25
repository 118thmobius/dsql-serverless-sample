DROP TABLE IF EXISTS app_account;
DROP TABLE IF EXISTS app_tx;

CREATE TABLE app_account(
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    screen_name varchar(255) UNIQUE ,
    balance integer
);

INSERT INTO app_account(user_id,screen_name,balance) VALUES ('11111111-aaaa-1111-aaaa-111111111111','alice',10000);
INSERT INTO app_account(user_id,screen_name,balance) VALUES ('22222222-bbbb-2222-bbbb-222222222222','bob',10000);

CREATE TABLE app_tx(
    tx_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_id UUID,
    to_id UUID,
    amount integer
);

INSERT INTO app_tx(
    from_id,
    to_id,
    amount
) VALUES (
    (SELECT user_id FROM app_account WHERE screen_name='alice'),
    (SELECT user_id FROM app_account WHERE screen_name='bob'),
    500
);

SELECT
    from_user.screen_name AS from_name,
    to_user.screen_name AS to_user,
    tx.amount
FROM
    app_tx tx
JOIN
    app_account from_user ON tx.from_id = from_user.user_id
JOIN
    app_account to_user ON tx.to_id = to_user.user_id;