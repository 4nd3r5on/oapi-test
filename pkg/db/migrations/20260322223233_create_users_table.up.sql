CREATE TABLE users (
    id         UUID        PRIMARY KEY DEFAULT uuidv7(),
    username   TEXT,
    locale     TEXT        NOT NULL,
    bio        TEXT        NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE users_identity_oidc (
    user_id UUID   NOT NULL REFERENCES users(id),
    sub     TEXT   UNIQUE NOT NULL 
);

CREATE TABLE users_identity_telegram (
    user_id     UUID   NOT NULL REFERENCES users(id),
    telegram_id BIGINT UNIQUE NOT NULL
);

CREATE UNIQUE INDEX users_username_idx ON users (username);
