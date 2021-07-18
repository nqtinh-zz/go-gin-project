-- migrate:up

CREATE TABLE public."user"
(
    id         SERIAL PRIMARY KEY,
    username   text UNIQUE NOT NULL,
    password   text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE public.account
(
    id         SERIAL PRIMARY KEY,
    name       Text UNIQUE NOT NULL,
    user_id    int         NOT NULL REFERENCES public.user (id),
    bank       Text        NOT NULL,
    created_at timestamptz DEFAULT NOW()
);

CREATE TABLE public.transaction
(
    id               SERIAL PRIMARY KEY,
    amount           int  NOT NULL,
    account_id       int  NOT NULL REFERENCES public.account (id),
    transaction_type Text NOT NULL,
    created_at       timestamptz DEFAULT NOW(),
    updated_at       timestamptz DEFAULT NOW(),
    deleted_at       timestamptz
);

-- migrate:down

DROP TABLE public.user;
DROP TABLE public.bank;
DROP TABLE public.account;
DROP TABLE public.transaction_type;
DROP TABLE public.transaction;