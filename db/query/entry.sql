-- name: CreateEntry :one
INSERT INTO entrys (
    owner, ammount
)VALUES($1,$2)
RETURNING *;