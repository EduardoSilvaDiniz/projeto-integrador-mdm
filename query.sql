-- name: GetAssociated :many
SELECT
  *
FROM
  associated;

-- name: CreateAssociated :exec
INSERT INTO
  associated (number_card, name)
VALUES
  (?, ?);

-- name: DeleteAssociatedByNumberCard :execresult
DELETE FROM associated
WHERE
  number_card = ?;
