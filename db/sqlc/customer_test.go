package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/food_order/util"
)

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		Email:       util.RandomEmail(),
		Name:        util.RandomString(20),
		Street:      util.RandomString(50),
		Postalcode:  util.RandomString(10),
		City:        util.RandomString(20),
		Phonenumber: util.RandomString(15),
	}

	customer, err := testStore.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Street, customer.Street)
	require.Equal(t, arg.Postalcode, customer.Postalcode)
	require.Equal(t, arg.City, customer.City)
	require.Equal(t, arg.Phonenumber, customer.Phonenumber)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testStore.GetCustomer(context.Background(), customer1.Customerid)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Name, customer2.Name)
	require.Equal(t, customer1.Street, customer2.Street)
	require.Equal(t, customer1.Postalcode, customer2.Postalcode)
	require.Equal(t, customer1.City, customer2.City)
	require.Equal(t, customer1.Phonenumber, customer2.Phonenumber)
	require.WithinDuration(t, customer1.CreatedAt, customer2.CreatedAt, time.Second)
}

func TestUpdateCustomer(t *testing.T) {
	oldCustomer := createRandomCustomer(t)

	newEmail := util.RandomEmail()
	newName := util.RandomString(20)
	newStreet := util.RandomString(50)
	newPostalcode := util.RandomString(10)
	newCity := util.RandomString(20)
	newPhonenumber := util.RandomString(15)

	updatedCustomer, err := testStore.UpdateCustomer(context.Background(), UpdateCustomerParams{
		Customerid:  oldCustomer.Customerid,
		Email:       newEmail,
		Name:        newName,
		Street:      newStreet,
		Postalcode:  newPostalcode,
		City:        newCity,
		Phonenumber: newPhonenumber,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldCustomer.Email, updatedCustomer.Email)
	require.Equal(t, newEmail, updatedCustomer.Email)
	require.NotEqual(t, oldCustomer.Name, updatedCustomer.Name)
	require.Equal(t, newName, updatedCustomer.Name)
	require.NotEqual(t, oldCustomer.Street, updatedCustomer.Street)
	require.Equal(t, newStreet, updatedCustomer.Street)
	require.NotEqual(t, oldCustomer.Postalcode, updatedCustomer.Postalcode)
	require.Equal(t, newPostalcode, updatedCustomer.Postalcode)
	require.NotEqual(t, oldCustomer.City, updatedCustomer.City)
	require.Equal(t, newCity, updatedCustomer.City)
	require.NotEqual(t, oldCustomer.Phonenumber, updatedCustomer.Phonenumber)
	require.Equal(t, newPhonenumber, updatedCustomer.Phonenumber)
}

func TestDeleteCustomer(t *testing.T) {
	customer := createRandomCustomer(t)
	err := testStore.DeleteCustomer(context.Background(), customer.Customerid)
	require.NoError(t, err)

	deletedCustomer, err := testStore.GetCustomer(context.Background(), customer.Customerid)
	require.Error(t, err)
	require.Empty(t, deletedCustomer)
}

func TestListCustomers(t *testing.T) {
	customers1, err := testStore.ListCustomers(context.Background())
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		createRandomCustomer(t)
	}

	customers2, err := testStore.ListCustomers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, customers2)

	require.Equal(t, len(customers1)+5, len(customers2))
}
