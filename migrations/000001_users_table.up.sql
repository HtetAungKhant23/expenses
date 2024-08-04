CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password bytea NOT NULL,
    activated bool NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version uuid NOT NULL DEFAULT uuid_generate_v4()
)