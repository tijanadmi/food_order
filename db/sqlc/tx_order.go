package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tijanadmi/food_order/util"
)

// OrderTxParams contains the input parameters of the order transaction
type OrderTxParams struct {
	Items []struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	} `json:"items"`
	Customer struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Street     string `json:"street"`
		PostalCode string `json:"postal_code"`
		City       string `json:"city"`
	} `json:"customer"`
}

// TransferTxResult is the result of the transfer transaction

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *SQLStore) OrderTx(ctx context.Context, arg OrderTxParams) (Order, error) {

	var updatedOrder Order

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		customer, err := q.GetCustomerByEmail(ctx, arg.Customer.Email)
		if err != nil {

			if err == pgx.ErrNoRows {
				customer, err = q.CreateCustomer(ctx, CreateCustomerParams{
					Email:       arg.Customer.Email,
					Name:        arg.Customer.Name,
					Street:      arg.Customer.Street,
					Postalcode:  arg.Customer.PostalCode,
					City:        arg.Customer.City,
					Phonenumber: "0648408871",
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		order, err := q.CreateOrder(ctx, CreateOrderParams{
			Customerid:  customer.Customerid,
			Orderdate:   time.Now(),
			Totalamount: pgtype.Numeric{Int: util.Float64ToBigInt(0), Exp: -2, Valid: true},
		})
		if err != nil {
			return err
		}
		totalamount := order.Totalamount

		for _, item := range arg.Items {
			_, err = q.CreateOrderDetail(ctx, CreateOrderDetailParams{
				Orderid:  order.Orderid,
				Mealid:   int32(item.ID),
				Quantity: int32(item.Quantity),
				Price:    pgtype.Numeric{Int: util.Float64ToBigInt(item.Price), Exp: -2, Valid: true},
			})

			totalamount = util.AddNumeric(totalamount, pgtype.Numeric{Int: util.Float64ToBigInt(item.Price), Exp: -2, Valid: true})

			if err != nil {
				return err
			}
		}

		updatedOrder, err = q.UpdateOrder(ctx, UpdateOrderParams{
			Customerid:  order.Customerid,
			Orderdate:   order.Orderdate,
			Totalamount: totalamount,
			Orderid:     order.Orderid,
		})
		return err
	})

	return updatedOrder, err
}
