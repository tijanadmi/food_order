// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO Orders (
    CustomerID,
    OrderDate,
    TotalAmount,
    created_at
) VALUES (
    $1, $2, $3, DEFAULT
) RETURNING orderid, customerid, orderdate, totalamount, created_at
`

type CreateOrderParams struct {
	Customerid  int32          `json:"customerid"`
	Orderdate   time.Time      `json:"orderdate"`
	Totalamount pgtype.Numeric `json:"totalamount"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.Customerid, arg.Orderdate, arg.Totalamount)
	var i Order
	err := row.Scan(
		&i.Orderid,
		&i.Customerid,
		&i.Orderdate,
		&i.Totalamount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM Orders
WHERE OrderID = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderid int32) error {
	_, err := q.db.Exec(ctx, deleteOrder, orderid)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT orderid, customerid, orderdate, totalamount, created_at FROM Orders
WHERE OrderID = $1
LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderid int32) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, orderid)
	var i Order
	err := row.Scan(
		&i.Orderid,
		&i.Customerid,
		&i.Orderdate,
		&i.Totalamount,
		&i.CreatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT orderid, customerid, orderdate, totalamount, created_at FROM Orders
ORDER BY OrderID
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.Orderid,
			&i.Customerid,
			&i.Orderdate,
			&i.Totalamount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrdersByCustomer = `-- name: ListOrdersByCustomer :many
SELECT orderid, customerid, orderdate, totalamount, created_at FROM Orders
WHERE CustomerID = $1
ORDER BY OrderDate
`

func (q *Queries) ListOrdersByCustomer(ctx context.Context, customerid int32) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrdersByCustomer, customerid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.Orderid,
			&i.Customerid,
			&i.Orderdate,
			&i.Totalamount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE Orders
SET
    CustomerID = COALESCE($1, CustomerID),
    OrderDate = COALESCE($2, OrderDate),
    TotalAmount = COALESCE($3, TotalAmount),
    created_at = created_at
WHERE
    OrderID = $4
RETURNING orderid, customerid, orderdate, totalamount, created_at
`

type UpdateOrderParams struct {
	Customerid  int32          `json:"customerid"`
	Orderdate   time.Time      `json:"orderdate"`
	Totalamount pgtype.Numeric `json:"totalamount"`
	Orderid     int32          `json:"orderid"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, updateOrder,
		arg.Customerid,
		arg.Orderdate,
		arg.Totalamount,
		arg.Orderid,
	)
	var i Order
	err := row.Scan(
		&i.Orderid,
		&i.Customerid,
		&i.Orderdate,
		&i.Totalamount,
		&i.CreatedAt,
	)
	return i, err
}
