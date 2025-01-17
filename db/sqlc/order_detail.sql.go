// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order_detail.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrderDetail = `-- name: CreateOrderDetail :one
INSERT INTO OrderDetails (
    OrderID,
    MealID,
    Quantity,
    Price,
    created_at
) VALUES (
    $1, $2, $3, $4, DEFAULT
) RETURNING orderdetailid, orderid, mealid, quantity, price, created_at
`

type CreateOrderDetailParams struct {
	Orderid  int32          `json:"orderid"`
	Mealid   int32          `json:"mealid"`
	Quantity int32          `json:"quantity"`
	Price    pgtype.Numeric `json:"price"`
}

func (q *Queries) CreateOrderDetail(ctx context.Context, arg CreateOrderDetailParams) (Orderdetail, error) {
	row := q.db.QueryRow(ctx, createOrderDetail,
		arg.Orderid,
		arg.Mealid,
		arg.Quantity,
		arg.Price,
	)
	var i Orderdetail
	err := row.Scan(
		&i.Orderdetailid,
		&i.Orderid,
		&i.Mealid,
		&i.Quantity,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOrderDetail = `-- name: DeleteOrderDetail :exec
DELETE FROM OrderDetails
WHERE OrderDetailID = $1
`

func (q *Queries) DeleteOrderDetail(ctx context.Context, orderdetailid int32) error {
	_, err := q.db.Exec(ctx, deleteOrderDetail, orderdetailid)
	return err
}

const getOrderDetail = `-- name: GetOrderDetail :one
SELECT orderdetailid, orderid, mealid, quantity, price, created_at FROM OrderDetails
WHERE OrderDetailID = $1
LIMIT 1
`

func (q *Queries) GetOrderDetail(ctx context.Context, orderdetailid int32) (Orderdetail, error) {
	row := q.db.QueryRow(ctx, getOrderDetail, orderdetailid)
	var i Orderdetail
	err := row.Scan(
		&i.Orderdetailid,
		&i.Orderid,
		&i.Mealid,
		&i.Quantity,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const listOrderDetails = `-- name: ListOrderDetails :many
SELECT orderdetailid, orderid, mealid, quantity, price, created_at FROM OrderDetails
ORDER BY OrderDetailID
`

func (q *Queries) ListOrderDetails(ctx context.Context) ([]Orderdetail, error) {
	rows, err := q.db.Query(ctx, listOrderDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Orderdetail{}
	for rows.Next() {
		var i Orderdetail
		if err := rows.Scan(
			&i.Orderdetailid,
			&i.Orderid,
			&i.Mealid,
			&i.Quantity,
			&i.Price,
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

const listOrderDetailsByOrder = `-- name: ListOrderDetailsByOrder :many
SELECT orderdetailid, orderid, mealid, quantity, price, created_at FROM OrderDetails
WHERE OrderID = $1
ORDER BY OrderDetailID
`

func (q *Queries) ListOrderDetailsByOrder(ctx context.Context, orderid int32) ([]Orderdetail, error) {
	rows, err := q.db.Query(ctx, listOrderDetailsByOrder, orderid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Orderdetail{}
	for rows.Next() {
		var i Orderdetail
		if err := rows.Scan(
			&i.Orderdetailid,
			&i.Orderid,
			&i.Mealid,
			&i.Quantity,
			&i.Price,
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

const updateOrderDetail = `-- name: UpdateOrderDetail :one
UPDATE OrderDetails
SET
    Quantity = COALESCE($1, Quantity),
    Price = COALESCE($2, Price)
WHERE
    OrderDetailID = $3
RETURNING orderdetailid, orderid, mealid, quantity, price, created_at
`

type UpdateOrderDetailParams struct {
	Quantity      int32          `json:"quantity"`
	Price         pgtype.Numeric `json:"price"`
	Orderdetailid int32          `json:"orderdetailid"`
}

func (q *Queries) UpdateOrderDetail(ctx context.Context, arg UpdateOrderDetailParams) (Orderdetail, error) {
	row := q.db.QueryRow(ctx, updateOrderDetail, arg.Quantity, arg.Price, arg.Orderdetailid)
	var i Orderdetail
	err := row.Scan(
		&i.Orderdetailid,
		&i.Orderid,
		&i.Mealid,
		&i.Quantity,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}
