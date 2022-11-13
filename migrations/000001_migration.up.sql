BEGIN;

CREATE TABLE IF NOT EXISTS user_account(
    id bigserial PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS color(
    id bigserial PRIMARY KEY,
    user_id BIGINT,
    hex VARCHAR(7),
    name VARCHAR(20),
    CONSTRAINT unique_color UNIQUE (user_id, hex),
    CONSTRAINT unique_name UNIQUE (user_id, name),

    CONSTRAINT fk_user_account
        FOREIGN KEY(user_id)
            REFERENCES user_account(id)
);

COMMIT;