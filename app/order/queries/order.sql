-- name: AddOrder :exec
INSERT INTO orders (
  order_request,
  rrn,
  order_id,
  seller_id
) VALUES (
  $1,
  $2,
  $3,
  $4
);

-- name: DeleteOrder :one
DELETE FROM orders
WHERE order_id = $1 AND seller_id = $2
RETURNING *;

-- name: GetOrder :one
SELECT
  o.order_request as order_request, o.rrn as rrn, o.order_id as order_id, o.seller_id as seller_id
FROM 
	orders o
WHERE
	o.order_id=$1
AND 
	o.seller_id=$2
LIMIT 1;

-- name: UpdateOrder :one
UPDATE 
  orders
SET 
  rrn = $3
WHERE
  order_id = $1
AND 
  seller_id = $2
RETURNING *;