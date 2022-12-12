CREATE TABLE public.orders (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    order_request jsonb NOT NULL DEFAULT '{}'::jsonb,
    rrn text COLLATE pg_catalog."default" NOT NULL,
    order_id text COLLATE pg_catalog."default" NOT NULL,
    seller_id text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT primary_order_seller PRIMARY KEY (order_id, seller_id)
);

COMMENT ON TABLE public.orders IS E'Orders';
COMMENT ON COLUMN public.orders.id IS E'Unique order id';
COMMENT ON COLUMN public.orders.order_request IS E'Order request data';
COMMENT ON COLUMN public.orders.rrn IS E'Reference Retrieval Number';
COMMENT ON COLUMN public.orders.order_id IS E'Order id';
COMMENT ON COLUMN public.orders.seller_id IS E'Seller if (owner credentials)';
ALTER TABLE public.orders OWNER TO orders_user;

