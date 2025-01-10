-- name: CreateMeal :one
INSERT INTO Meals (
    Name,
    Description,
    Price,
    Category,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, DEFAULT, DEFAULT
) RETURNING *;

-- name: GetMeal :one
SELECT * FROM Meals
WHERE MealID = $1
LIMIT 1;

-- name: GetMealForUpdate :one
SELECT * FROM Meals
WHERE MealID = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMeals :many
SELECT * FROM Meals
ORDER BY MealID;

-- name: GetMealsByCategory :many
SELECT * FROM Meals
WHERE Category = $1
ORDER BY MealID;

-- name: UpdateMeal :one
UPDATE Meals
SET
    Name = COALESCE($1, Name),
    Description = COALESCE($2, Description),
    Price = COALESCE($3, Price),
    Category = COALESCE($4, Category),
    updated_at = now()
WHERE
    MealID = $5
RETURNING *;

-- name: DeleteMeal :exec
DELETE FROM Meals
WHERE MealID = $1;