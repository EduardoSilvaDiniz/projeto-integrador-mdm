-- ASSOCIATED
-- name: GetAssociated :many
SELECT
  *
FROM
  associated;

-- name: GetAssociatedByNumberCard :one
SELECT
  *
FROM
  associated
WHERE
  number_card = ?;

-- name: CreateAssociated :exec
INSERT INTO
  associated (number_card, name, group_id)
VALUES
  (?, ?, ?);

-- name: UpdateAssociated :execresult
UPDATE associated
SET
  name = ?,
  group_id = ?
WHERE
  number_card = ?;

-- name: DeleteAssociatedByNumberCard :execresult
DELETE FROM associated
WHERE
  number_card = ?;

-- name: GetAssociatedByGroup :many
SELECT
  *
FROM
  associated
WHERE
  group_id = ?;

-- GROUPS
-- name: GetGroups :many
SELECT
  *
FROM
  groups;

-- name: GetGroupByID :one
SELECT
  *
FROM
  groups
WHERE
  id = ?;

-- name: CreateGroup :exec
INSERT INTO
  groups (name, hours)
VALUES
  (?, ?);

-- name: UpdateGroup :execresult
UPDATE groups
SET
  name = ?,
  hours = ?
WHERE
  id = ?;

-- name: DeleteGroupById :execresult
DELETE FROM groups
WHERE
  id = ?;

-- MEETING
-- name: CreateMeeting :exec
INSERT INTO
  meeting (group_id, address, date)
VALUES
  (?, ?, ?);

-- name: GetMeetings :many
SELECT
  *
FROM
  meeting;

-- name: GetMeetingByID :one
SELECT
  *
FROM
  meeting
WHERE
  id = ?;

-- name: UpdateMeeting :execresult
UPDATE meeting
SET
  group_id = ?,
  address = ?,
  date = ?
WHERE
  id = ?;

-- name: DeleteMeetingById :execresult
DELETE FROM meeting
WHERE
  id = ?;

-- name: GetMeetingsByGroup :many
SELECT
  *
FROM
  meeting
WHERE
  group_id = ?;

-- PRESENCE
-- name: GetPresence :many
SELECT
  *
FROM
  presence;

-- name: GetPresenceByCompositeKey :one
SELECT
  *
FROM
  presence
WHERE
  number_card = ?
  AND meeting_id = ?;

-- name: GetPresenceByMeeting :many
SELECT
  *
FROM
  presence
WHERE
  meeting_id = ?;

-- name: GetPresenceByAssociated :many
SELECT
  *
FROM
  presence
WHERE
  number_card = ?;

-- name: CreatePresence :exec
INSERT INTO
  presence (number_card, meeting_id, is_presence)
VALUES
  (?, ?, ?);

-- name: UpdatePresence :execresult
UPDATE presence
SET
  is_presence = ?
WHERE
  number_card = ?
  AND meeting_id = ?;

-- name: DeletePresenceByCompositeKey :execresult
DELETE FROM presence
WHERE
  number_card = ?
  AND meeting_id = ?;

-- PAYMENT
-- name: GetPayment :many
SELECT
  *
FROM
  payment;

-- name: GetPaymentByID :one
SELECT
  *
FROM
  payment
WHERE
  id = ?;

-- name: GetPaymentByAssociated :many
SELECT
  *
FROM
  payment
WHERE
  number_card = ?;

-- name: GetPaymentByMonthYear :many
SELECT
  *
FROM
  payment
WHERE
  strftime('%m', ref_month) = ?
  AND strftime('%Y', ref_month) = ?;

-- name: CreatePayment :exec
INSERT INTO
  payment (number_card, ref_month, payment_date)
VALUES
  (?, ?, ?);

-- name: UpdatePayment :execresult
UPDATE payment
SET
  ref_month = ?,
  payment_date = ?
WHERE
  id = ?;

-- name: DeletePaymentById :execresult
DELETE FROM payment
WHERE
  id = ?;

-- name: CheckPaymentExists :one
SELECT
  1
FROM
  payment
WHERE
  number_card = ?
  AND ref_month = ?;
