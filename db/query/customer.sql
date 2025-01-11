-- name: CreateCustomer :one
INSERT INTO Customers (
    Email,
    Name,
    Street,
    PostalCode,
    City,
    PhoneNumber,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, DEFAULT, DEFAULT
) RETURNING *;

-- name: GetCustomer :one
SELECT * FROM Customers
WHERE CustomerID = $1
LIMIT 1;

-- name: GetCustomerByEmail :one
SELECT * FROM Customers
WHERE Email = $1
LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM Customers
ORDER BY CustomerID;

-- name: UpdateCustomer :one
UPDATE Customers
SET
    Email = COALESCE($1, Email),
    Name = COALESCE($2, Name),
    Street = COALESCE($3, Street),
    PostalCode = COALESCE($4, PostalCode),
    City = COALESCE($5, City),
    PhoneNumber = COALESCE($6, PhoneNumber),
    updated_at = now()
WHERE
    CustomerID = $7
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM Customers
WHERE CustomerID = $1;
