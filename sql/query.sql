-- name: GetAllFilm :many
SELECT * FROM film;

-- name: GetFilm :one
SELECT * FROM film
WHERE id = $1;

-- name: CreateFilm :one
INSERT INTO film (
  name, title, category_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteFilm :exec
DELETE FROM film
WHERE id = $1;

-- name: UpdateFilm :exec
UPDATE film 
SET name = $2, title = $3, category_id = $4
WHERE id = $1;

-- name: GetAllCategory :many
SELECT * FROM category;

-- name: CreateCategory :one
INSERT INTO category (
  category
) VALUES (
  $1
)
RETURNING *;