-- name: CreateOrderDetail :one
INSERT INTO OrderDetails (
    OrderID,
    MealID,
    Quantity,
    Price,
    created_at
) VALUES (
    $1, $2, $3, $4, DEFAULT
) RETURNING *;

-- name: GetOrderDetail :one
SELECT * FROM OrderDetails
WHERE OrderDetailID = $1
LIMIT 1;

-- name: ListOrderDetails :many
SELECT * FROM OrderDetails
ORDER BY OrderDetailID;

-- name: ListOrderDetailsByOrder :many
SELECT * FROM OrderDetails
WHERE OrderID = $1
ORDER BY OrderDetailID;

-- name: UpdateOrderDetail :one
UPDATE OrderDetails
SET
    Quantity = COALESCE($1, Quantity),
    Price = COALESCE($2, Price)
WHERE
    OrderDetailID = $3
RETURNING *;

-- name: DeleteOrderDetail :exec
DELETE FROM OrderDetails
WHERE OrderDetailID = $1;
