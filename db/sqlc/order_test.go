package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/food_order/util"
)

func createRandomOrder(t *testing.T, customer Customer) Order {

	arg := CreateOrderParams{
		Customerid:  customer.Customerid,
		Orderdate:   time.Now(),
		Totalamount: pgtype.Numeric{Int: util.Float64ToBigInt(util.RandomNumeric(1, 1000)), Exp: -2, Valid: true},
	}

	order, err := testStore.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.Customerid, order.Customerid)
	require.WithinDuration(t, arg.Orderdate, order.Orderdate, time.Second)
	require.Equal(t, arg.Totalamount.Int, order.Totalamount.Int)
	require.Equal(t, arg.Totalamount.Exp, order.Totalamount.Exp)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestCreateOrder(t *testing.T) {
	customer := createRandomCustomer(t)
	createRandomOrder(t, customer)
}

func truncateToMicroseconds(t time.Time) time.Time {
	return t.Truncate(time.Microsecond)
}

func TestGetOrder(t *testing.T) {
	customer := createRandomCustomer(t)
	order1 := createRandomOrder(t, customer)
	order2, err := testStore.GetOrder(context.Background(), order1.Orderid)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.Customerid, order2.Customerid)
	require.WithinDuration(t, order1.Orderdate, order2.Orderdate, time.Second)
	require.Equal(t, order1.Totalamount.Int, order2.Totalamount.Int)
	require.Equal(t, order1.Totalamount.Exp, order2.Totalamount.Exp)
	require.WithinDuration(t, order1.CreatedAt.UTC(), order2.CreatedAt.UTC(), time.Millisecond)
}

func TestUpdateOrder(t *testing.T) {
	customer := createRandomCustomer(t)
	oldOrder := createRandomOrder(t, customer)

	newCustomer := createRandomCustomer(t)
	newOrderdate := time.Now()
	newTotalamount := pgtype.Numeric{Int: util.Float64ToBigInt(util.RandomNumeric(1, 1000)), Exp: -2, Valid: true}

	updatedOrder, err := testStore.UpdateOrder(context.Background(), UpdateOrderParams{
		Orderid:     oldOrder.Orderid,
		Customerid:  newCustomer.Customerid,
		Orderdate:   newOrderdate,
		Totalamount: newTotalamount,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldOrder.Customerid, updatedOrder.Customerid)
	require.Equal(t, newCustomer.Customerid, updatedOrder.Customerid)
	require.WithinDuration(t, oldOrder.Orderdate, updatedOrder.Orderdate, time.Second)
	require.NotEqual(t, oldOrder.Totalamount.Int, updatedOrder.Totalamount.Int)
	require.Equal(t, newTotalamount.Int, updatedOrder.Totalamount.Int)
	require.Equal(t, newTotalamount.Exp, updatedOrder.Totalamount.Exp)
}

func TestDeleteOrder(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	err := testStore.DeleteOrder(context.Background(), order.Orderid)
	require.NoError(t, err)

	deletedOrder, err := testStore.GetOrder(context.Background(), order.Orderid)
	require.Error(t, err)
	require.Empty(t, deletedOrder)
}

func TestListOrders(t *testing.T) {

	orders1, err := testStore.ListOrders(context.Background())
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		customer := createRandomCustomer(t)
		createRandomOrder(t, customer)
	}

	orders2, err := testStore.ListOrders(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, orders2)

	require.Equal(t, len(orders1)+5, len(orders2))
}
