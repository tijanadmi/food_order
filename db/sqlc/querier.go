// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateMeal(ctx context.Context, arg CreateMealParams) (Meal, error)
	DeleteMeal(ctx context.Context, mealid int32) error
	GetMeal(ctx context.Context, mealid int32) (Meal, error)
	GetMealForUpdate(ctx context.Context, mealid int32) (Meal, error)
	GetMealsByCategory(ctx context.Context, category pgtype.Text) ([]Meal, error)
	ListMeals(ctx context.Context) ([]Meal, error)
	UpdateMeal(ctx context.Context, arg UpdateMealParams) (Meal, error)
}

var _ Querier = (*Queries)(nil)