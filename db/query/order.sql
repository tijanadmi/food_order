-- name: CreateOrder :one
INSERT INTO Orders (
    CustomerID,
    OrderDate,
    TotalAmount,
    created_at
) VALUES (
    $1, $2, $3, DEFAULT
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM Orders
WHERE OrderID = $1
LIMIT 1;

-- name: ListOrders :many
SELECT * FROM Orders
ORDER BY OrderID;

-- name: ListOrdersByCustomer :many
SELECT * FROM Orders
WHERE CustomerID = $1
ORDER BY OrderDate;

-- name: UpdateOrder :one
UPDATE Orders
SET
    CustomerID = COALESCE($1, CustomerID),
    OrderDate = COALESCE($2, OrderDate),
    TotalAmount = COALESCE($3, TotalAmount),
    created_at = created_at
WHERE
    OrderID = $4
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM Orders
WHERE OrderID = $1;
