
SET SYNCHRONOUS_COMMIT = 'off';

CREATE TABLE IF NOT EXISTS users (
    id bigserial not null primary key,
    username varchar,
    created_at timestamptz DEFAULT now()
);


CREATE TABLE IF NOT EXISTS chat (
    id bigserial not null primary key,
    name varchar,
    created_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_users (
    chat_id bigint not null references chat,
    user_id bigint not null references users
);

CREATE TABLE IF NOT EXISTS messages (
    id bigserial not null primary key,
    chat_id bigint not null references chat,
    user_id bigint not null references users,
    text varchar not null,
    created_at timestamptz DEFAULT now()
);

