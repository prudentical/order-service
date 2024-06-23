CREATE TABLE public.positions (
    id bigserial PRIMARY KEY,
    bot_id bigint NOT NULL,
    status character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL
);
CREATE INDEX idx_positions_bot_id ON public.positions USING HASH (bot_id);

CREATE TABLE public.orders (
    id bigserial PRIMARY KEY,
    position_id bigserial NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    internal_id bigint NOT NULL,
    amount character varying(50) NOT NULL,
    filled_amount character varying(50) NOT NULL,
    price character varying(50) NOT NULL,
    type character varying(50) NOT NULL,
    status character varying(50) NOT NULL,
    date_time timestamp with time zone NOT NULL,
    CONSTRAINT fk_orders_positions FOREIGN KEY(position_id) REFERENCES positions(id)
);
CREATE INDEX idx_orders_del ON public.orders USING HASH (deleted_at);