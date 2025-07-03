-- name: GetAssoc :many
SELECT
  *
FROM
  associated;

-- name: CreateAssoc :exec
INSERT INTO
  associated (cpf, name, date_birth, marital_status)
VALUES
  ($1, $2, $3, $4);
