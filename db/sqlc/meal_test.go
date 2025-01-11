package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/food_order/util"
)

func createRandomMeal(t *testing.T) Meal {

	arg := CreateMealParams{
		Name:        util.RandomString(20),
		Description: util.RandomString(200),
		Price:       pgtype.Numeric{Int: util.Float64ToBigInt(util.RandomNumeric(1, 100)), Exp: -2, Valid: true},
		Category:    "Dessert",
	}

	meal, err := testStore.CreateMeal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, meal)

	require.Equal(t, arg.Name, meal.Name)
	require.Equal(t, arg.Description, meal.Description)
	require.Equal(t, arg.Price.Int, meal.Price.Int)
	require.Equal(t, arg.Price.Exp, meal.Price.Exp)
	require.Equal(t, arg.Category, meal.Category)
	require.NotZero(t, meal.CreatedAt)

	return meal
}

func TestCreateMeal(t *testing.T) {
	createRandomMeal(t)
}

func TestGetMeal(t *testing.T) {
	meal1 := createRandomMeal(t)
	meal2, err := testStore.GetMeal(context.Background(), meal1.Mealid)
	require.NoError(t, err)
	require.NotEmpty(t, meal2)

	require.Equal(t, meal1.Name, meal2.Name)
	require.Equal(t, meal1.Description, meal2.Description)
	require.Equal(t, meal1.Price.Int, meal2.Price.Int)
	require.Equal(t, meal1.Price.Exp, meal2.Price.Exp)
	require.Equal(t, meal1.Category, meal2.Category)
	require.WithinDuration(t, meal1.CreatedAt, meal2.CreatedAt, time.Second)
}

func TestUpdateMeal(t *testing.T) {
	oldMeal := createRandomMeal(t)

	newName := util.RandomString(20)
	newDescription := util.RandomString(200)
	newPrice := pgtype.Numeric{Int: util.Float64ToBigInt(util.RandomNumeric(1, 100)), Exp: -2, Valid: true}
	newCategory := "Main Course"

	updatedMeal, err := testStore.UpdateMeal(context.Background(), UpdateMealParams{
		Mealid:      oldMeal.Mealid,
		Name:        newName,
		Description: newDescription,
		Price:       newPrice,
		Category:    newCategory,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldMeal.Name, updatedMeal.Name)
	require.Equal(t, newName, updatedMeal.Name)
	require.NotEqual(t, oldMeal.Description, updatedMeal.Description)
	require.Equal(t, newDescription, updatedMeal.Description)
	require.NotEqual(t, oldMeal.Price.Int, updatedMeal.Price.Int)
	require.Equal(t, newPrice.Int, updatedMeal.Price.Int)
	require.Equal(t, newPrice.Exp, updatedMeal.Price.Exp)
	require.NotEqual(t, oldMeal.Category, updatedMeal.Category)
	require.Equal(t, newCategory, updatedMeal.Category)
}

func TestDeleteMeal(t *testing.T) {
	meal := createRandomMeal(t)
	err := testStore.DeleteMeal(context.Background(), meal.Mealid)
	require.NoError(t, err)

	deletedMeal, err := testStore.GetMeal(context.Background(), meal.Mealid)
	require.Error(t, err)
	require.Empty(t, deletedMeal)
}

func TestListMeals(t *testing.T) {

	meals1, err := testStore.ListMeals(context.Background())
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		createRandomMeal(t)
	}

	meals2, err := testStore.ListMeals(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, meals2)

	require.Equal(t, len(meals1)+5, len(meals2))
}
