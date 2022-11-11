BEGIN;

CREATE TABLE IF NOT EXISTS user_account(
    id bigserial PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS color(
    id bigserial PRIMARY KEY,
    user_id BIGINT,
    hex VARCHAR(7),
    name VARCHAR(20),

    CONSTRAINT fk_user_account
        FOREIGN KEY(user_id)
            REFERENCES user_account(id)
);

COMMIT;