package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/food_order/util"
)

func createRandomOrderDetail(t *testing.T, order Order, meal Meal) Orderdetail {
	arg := CreateOrderDetailParams{
		Orderid:  order.Orderid,
		Mealid:   meal.Mealid,
		Quantity: int32(util.RandomInt(1, 10)),
		Price:    pgtype.Numeric{Int: util.Float64ToBigInt(util.RandomNumeric(1, 100)), Exp: -2, Valid: true},
	}

	orderDetail, err := testStore.CreateOrderDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail)

	require.Equal(t, arg.Orderid, orderDetail.Orderid)
	require.Equal(t, arg.Mealid, orderDetail.Mealid)
	require.Equal(t, arg.Quantity, orderDetail.Quantity)
	require.Equal(t, arg.Price.Int, orderDetail.Price.Int)
	require.Equal(t, arg.Price.Exp, orderDetail.Price.Exp)
	require.NotZero(t, orderDetail.CreatedAt)

	return orderDetail
}

func TestCreateOrderDetail(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	meal := createRandomMeal(t)
	createRandomOrderDetail(t, order, meal)
}

func TestGetOrderDetail(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	meal := createRandomMeal(t)
	orderDetail1 := createRandomOrderDetail(t, order, meal)
	orderDetail2, err := testStore.GetOrderDetail(context.Background(), orderDetail1.Orderdetailid)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail2)

	require.Equal(t, orderDetail1.Orderid, orderDetail2.Orderid)
	require.Equal(t, orderDetail1.Mealid, orderDetail2.Mealid)
	require.Equal(t, orderDetail1.Quantity, orderDetail2.Quantity)
	require.Equal(t, orderDetail1.Price.Int, orderDetail2.Price.Int)
	require.Equal(t, orderDetail1.Price.Exp, orderDetail2.Price.Exp)
	require.WithinDuration(t, orderDetail1.CreatedAt.UTC(), orderDetail2.CreatedAt.UTC(), time.Millisecond)
}

func TestUpdateOrderDetail(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	meal := createRandomMeal(t)
	oldOrderDetail := createRandomOrderDetail(t, order, meal)

	newQuantity := int32(util.RandomInt(1, 10))
	newPrice := pgtype.Numeric{Int: util.Float64ToBigInt(util.RandomNumeric(1, 100)), Exp: -2, Valid: true}

	updatedOrderDetail, err := testStore.UpdateOrderDetail(context.Background(), UpdateOrderDetailParams{
		Orderdetailid: oldOrderDetail.Orderdetailid,
		Quantity:      int32(newQuantity),
		Price:         newPrice,
	})

	require.NoError(t, err)
	//.NotEqual(t, oldOrderDetail.Mealid, updatedOrderDetail.Mealid)
	require.Equal(t, meal.Mealid, updatedOrderDetail.Mealid)
	require.NotEqual(t, oldOrderDetail.Quantity, updatedOrderDetail.Quantity)
	require.Equal(t, newQuantity, updatedOrderDetail.Quantity)
	require.NotEqual(t, oldOrderDetail.Price.Int, updatedOrderDetail.Price.Int)
	require.Equal(t, newPrice.Int, updatedOrderDetail.Price.Int)
	require.Equal(t, newPrice.Exp, updatedOrderDetail.Price.Exp)
}

func TestDeleteOrderDetail(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	meal := createRandomMeal(t)
	orderDetail := createRandomOrderDetail(t, order, meal)
	err := testStore.DeleteOrderDetail(context.Background(), orderDetail.Orderdetailid)
	require.NoError(t, err)

	deletedOrderDetail, err := testStore.GetOrderDetail(context.Background(), orderDetail.Orderdetailid)
	require.Error(t, err)
	require.Empty(t, deletedOrderDetail)
}

func TestListOrderDetails(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	meal := createRandomMeal(t)

	orderDetails1, err := testStore.ListOrderDetails(context.Background())
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		createRandomOrderDetail(t, order, meal)
	}

	orderDetails2, err := testStore.ListOrderDetails(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, orderDetails2)

	require.Equal(t, len(orderDetails1)+5, len(orderDetails2))
}
