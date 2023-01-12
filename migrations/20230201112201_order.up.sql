CREATE TABLE IF NOT EXISTS public."order" (
	id bigserial NOT NULL,
	create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp,
	lifetime interval,
	partner_id bigint NOT NULL,
	merchant_id bigint NOT NULL,
	shop_id bigint NOT NULL,
	merchant_order_id text NOT NULL,
	payment_type integer NOT NULL DEFAULT 0,
	amount int8 NOT NULL,
	currency varchar(3) NOT NULL,
	items jsonb NOT NULL DEFAULT '{}'::jsonb,
	description text,
	user_data jsonb NOT NULL DEFAULT '{}'::jsonb,
	buyer jsonb NOT NULL DEFAULT '{}'::jsonb
);
ALTER TABLE public."order" OWNER TO orders_user;