CREATE TABLE public.orders (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    bot_id bigint NOT NULL,
    amount text NOT NULL,
    price text NOT NULL,
    type text NOT NULL,
    date_time timestamp with time zone NOT NULL
);
CREATE INDEX idx_orders_del ON public.orders USING HASH (deleted_at);
CREATE INDEX idx_orders_bot_id ON public.orders USING HASH (bot_id);
