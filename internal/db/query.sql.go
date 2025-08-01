// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const checkPaymentExists = `-- name: CheckPaymentExists :one
SELECT
  1
FROM
  payment
WHERE
  number_card = ?
  AND ref_month = ?
`

type CheckPaymentExistsParams struct {
	NumberCard int64
	RefMonth   string
}

func (q *Queries) CheckPaymentExists(ctx context.Context, arg CheckPaymentExistsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, checkPaymentExists, arg.NumberCard, arg.RefMonth)
	var column_1 int64
	err := row.Scan(&column_1)
	return column_1, err
}

const createAssociated = `-- name: CreateAssociated :exec
INSERT INTO
  associated (number_card, name, group_id)
VALUES
  (?, ?, ?)
`

type CreateAssociatedParams struct {
	NumberCard int64
	Name       string
	GroupID    int64
}

func (q *Queries) CreateAssociated(ctx context.Context, arg CreateAssociatedParams) error {
	_, err := q.db.ExecContext(ctx, createAssociated, arg.NumberCard, arg.Name, arg.GroupID)
	return err
}

const createGroup = `-- name: CreateGroup :exec
INSERT INTO
  groups (name, hours)
VALUES
  (?, ?)
`

type CreateGroupParams struct {
	Name  string
	Hours time.Time
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) error {
	_, err := q.db.ExecContext(ctx, createGroup, arg.Name, arg.Hours)
	return err
}

const createMeeting = `-- name: CreateMeeting :exec
INSERT INTO
  meeting (group_id, address, date)
VALUES
  (?, ?, ?)
`

type CreateMeetingParams struct {
	GroupID int64
	Address string
	Date    time.Time
}

// MEETING
func (q *Queries) CreateMeeting(ctx context.Context, arg CreateMeetingParams) error {
	_, err := q.db.ExecContext(ctx, createMeeting, arg.GroupID, arg.Address, arg.Date)
	return err
}

const createPayment = `-- name: CreatePayment :exec
INSERT INTO
  payment (number_card, ref_month, payment_date)
VALUES
  (?, ?, ?)
`

type CreatePaymentParams struct {
	NumberCard  int64
	RefMonth    string
	PaymentDate time.Time
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, createPayment, arg.NumberCard, arg.RefMonth, arg.PaymentDate)
	return err
}

const createPresence = `-- name: CreatePresence :exec
INSERT INTO
  presence (number_card, meeting_id, is_presence)
VALUES
  (?, ?, ?)
`

type CreatePresenceParams struct {
	NumberCard int64
	MeetingID  int64
	IsPresence bool
}

func (q *Queries) CreatePresence(ctx context.Context, arg CreatePresenceParams) error {
	_, err := q.db.ExecContext(ctx, createPresence, arg.NumberCard, arg.MeetingID, arg.IsPresence)
	return err
}

const deleteAssociatedByNumberCard = `-- name: DeleteAssociatedByNumberCard :execresult
DELETE FROM associated
WHERE
  number_card = ?
`

func (q *Queries) DeleteAssociatedByNumberCard(ctx context.Context, numberCard int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAssociatedByNumberCard, numberCard)
}

const deleteGroupById = `-- name: DeleteGroupById :execresult
DELETE FROM groups
WHERE
  id = ?
`

func (q *Queries) DeleteGroupById(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteGroupById, id)
}

const deleteMeetingById = `-- name: DeleteMeetingById :execresult
DELETE FROM meeting
WHERE
  id = ?
`

func (q *Queries) DeleteMeetingById(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteMeetingById, id)
}

const deletePaymentById = `-- name: DeletePaymentById :execresult
DELETE FROM payment
WHERE
  id = ?
`

func (q *Queries) DeletePaymentById(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePaymentById, id)
}

const deletePresenceByCompositeKey = `-- name: DeletePresenceByCompositeKey :execresult
DELETE FROM presence
WHERE
  number_card = ?
  AND meeting_id = ?
`

type DeletePresenceByCompositeKeyParams struct {
	NumberCard int64
	MeetingID  int64
}

func (q *Queries) DeletePresenceByCompositeKey(ctx context.Context, arg DeletePresenceByCompositeKeyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePresenceByCompositeKey, arg.NumberCard, arg.MeetingID)
}

const getAssociated = `-- name: GetAssociated :many
SELECT
  number_card, name, group_id
FROM
  associated
`

// ASSOCIATED
func (q *Queries) GetAssociated(ctx context.Context) ([]Associated, error) {
	rows, err := q.db.QueryContext(ctx, getAssociated)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Associated
	for rows.Next() {
		var i Associated
		if err := rows.Scan(&i.NumberCard, &i.Name, &i.GroupID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssociatedByGroup = `-- name: GetAssociatedByGroup :many
SELECT
  number_card, name, group_id
FROM
  associated
WHERE
  group_id = ?
`

func (q *Queries) GetAssociatedByGroup(ctx context.Context, groupID int64) ([]Associated, error) {
	rows, err := q.db.QueryContext(ctx, getAssociatedByGroup, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Associated
	for rows.Next() {
		var i Associated
		if err := rows.Scan(&i.NumberCard, &i.Name, &i.GroupID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssociatedByNumberCard = `-- name: GetAssociatedByNumberCard :one
SELECT
  number_card, name, group_id
FROM
  associated
WHERE
  number_card = ?
`

func (q *Queries) GetAssociatedByNumberCard(ctx context.Context, numberCard int64) (Associated, error) {
	row := q.db.QueryRowContext(ctx, getAssociatedByNumberCard, numberCard)
	var i Associated
	err := row.Scan(&i.NumberCard, &i.Name, &i.GroupID)
	return i, err
}

const getGroupById = `-- name: GetGroupById :one
SELECT
  id, name, hours
FROM
  groups
WHERE
  id = ?
`

func (q *Queries) GetGroupById(ctx context.Context, id int64) (Group, error) {
	row := q.db.QueryRowContext(ctx, getGroupById, id)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Hours)
	return i, err
}

const getGroups = `-- name: GetGroups :many
SELECT
  id, name, hours
FROM
  groups
`

// GROUPS
func (q *Queries) GetGroups(ctx context.Context) ([]Group, error) {
	rows, err := q.db.QueryContext(ctx, getGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Group
	for rows.Next() {
		var i Group
		if err := rows.Scan(&i.ID, &i.Name, &i.Hours); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMeetingById = `-- name: GetMeetingById :one
SELECT
  id, group_id, address, date
FROM
  meeting
WHERE
  id = ?
`

func (q *Queries) GetMeetingById(ctx context.Context, id int64) (Meeting, error) {
	row := q.db.QueryRowContext(ctx, getMeetingById, id)
	var i Meeting
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.Address,
		&i.Date,
	)
	return i, err
}

const getMeetings = `-- name: GetMeetings :many
SELECT
  id, group_id, address, date
FROM
  meeting
`

func (q *Queries) GetMeetings(ctx context.Context) ([]Meeting, error) {
	rows, err := q.db.QueryContext(ctx, getMeetings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Meeting
	for rows.Next() {
		var i Meeting
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.Address,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMeetingsByGroup = `-- name: GetMeetingsByGroup :many
SELECT
  id, group_id, address, date
FROM
  meeting
WHERE
  group_id = ?
`

func (q *Queries) GetMeetingsByGroup(ctx context.Context, groupID int64) ([]Meeting, error) {
	rows, err := q.db.QueryContext(ctx, getMeetingsByGroup, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Meeting
	for rows.Next() {
		var i Meeting
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.Address,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPayment = `-- name: GetPayment :many
SELECT
  id, number_card, ref_month, payment_date
FROM
  payment
`

// PAYMENT
func (q *Queries) GetPayment(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPayment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.NumberCard,
			&i.RefMonth,
			&i.PaymentDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaymentByAssociated = `-- name: GetPaymentByAssociated :many
SELECT
  id, number_card, ref_month, payment_date
FROM
  payment
WHERE
  number_card = ?
`

func (q *Queries) GetPaymentByAssociated(ctx context.Context, numberCard int64) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentByAssociated, numberCard)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.NumberCard,
			&i.RefMonth,
			&i.PaymentDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaymentById = `-- name: GetPaymentById :one
SELECT
  id, number_card, ref_month, payment_date
FROM
  payment
WHERE
  id = ?
`

func (q *Queries) GetPaymentById(ctx context.Context, id int64) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPaymentById, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.NumberCard,
		&i.RefMonth,
		&i.PaymentDate,
	)
	return i, err
}

const getPaymentByMonthYear = `-- name: GetPaymentByMonthYear :many
SELECT
  id, number_card, ref_month, payment_date
FROM
  payment
WHERE
  strftime('%m', ref_month) = ?
  AND strftime('%Y', ref_month) = ?
`

type GetPaymentByMonthYearParams struct {
	RefMonth   string
	RefMonth_2 string
}

func (q *Queries) GetPaymentByMonthYear(ctx context.Context, arg GetPaymentByMonthYearParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentByMonthYear, arg.RefMonth, arg.RefMonth_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.NumberCard,
			&i.RefMonth,
			&i.PaymentDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPresence = `-- name: GetPresence :many
SELECT
  number_card, meeting_id, is_presence
FROM
  presence
`

// PRESENCE
func (q *Queries) GetPresence(ctx context.Context) ([]Presence, error) {
	rows, err := q.db.QueryContext(ctx, getPresence)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Presence
	for rows.Next() {
		var i Presence
		if err := rows.Scan(&i.NumberCard, &i.MeetingID, &i.IsPresence); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPresenceByAssociated = `-- name: GetPresenceByAssociated :many
SELECT
  number_card, meeting_id, is_presence
FROM
  presence
WHERE
  number_card = ?
`

func (q *Queries) GetPresenceByAssociated(ctx context.Context, numberCard int64) ([]Presence, error) {
	rows, err := q.db.QueryContext(ctx, getPresenceByAssociated, numberCard)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Presence
	for rows.Next() {
		var i Presence
		if err := rows.Scan(&i.NumberCard, &i.MeetingID, &i.IsPresence); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPresenceByCompositeKey = `-- name: GetPresenceByCompositeKey :one
SELECT
  number_card, meeting_id, is_presence
FROM
  presence
WHERE
  number_card = ?
  AND meeting_id = ?
`

type GetPresenceByCompositeKeyParams struct {
	NumberCard int64
	MeetingID  int64
}

func (q *Queries) GetPresenceByCompositeKey(ctx context.Context, arg GetPresenceByCompositeKeyParams) (Presence, error) {
	row := q.db.QueryRowContext(ctx, getPresenceByCompositeKey, arg.NumberCard, arg.MeetingID)
	var i Presence
	err := row.Scan(&i.NumberCard, &i.MeetingID, &i.IsPresence)
	return i, err
}

const getPresenceByMeeting = `-- name: GetPresenceByMeeting :many
SELECT
  number_card, meeting_id, is_presence
FROM
  presence
WHERE
  meeting_id = ?
`

func (q *Queries) GetPresenceByMeeting(ctx context.Context, meetingID int64) ([]Presence, error) {
	rows, err := q.db.QueryContext(ctx, getPresenceByMeeting, meetingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Presence
	for rows.Next() {
		var i Presence
		if err := rows.Scan(&i.NumberCard, &i.MeetingID, &i.IsPresence); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAssociated = `-- name: UpdateAssociated :execresult
UPDATE associated
SET
  name = ?,
  group_id = ?
WHERE
  number_card = ?
`

type UpdateAssociatedParams struct {
	Name       string
	GroupID    int64
	NumberCard int64
}

func (q *Queries) UpdateAssociated(ctx context.Context, arg UpdateAssociatedParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateAssociated, arg.Name, arg.GroupID, arg.NumberCard)
}

const updateGroup = `-- name: UpdateGroup :execresult
UPDATE groups
SET
  name = ?,
  hours = ?
WHERE
  id = ?
`

type UpdateGroupParams struct {
	Name  string
	Hours time.Time
	ID    int64
}

func (q *Queries) UpdateGroup(ctx context.Context, arg UpdateGroupParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateGroup, arg.Name, arg.Hours, arg.ID)
}

const updateMeeting = `-- name: UpdateMeeting :execresult
UPDATE meeting
SET
  group_id = ?,
  address = ?,
  date = ?
WHERE
  id = ?
`

type UpdateMeetingParams struct {
	GroupID int64
	Address string
	Date    time.Time
	ID      int64
}

func (q *Queries) UpdateMeeting(ctx context.Context, arg UpdateMeetingParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateMeeting,
		arg.GroupID,
		arg.Address,
		arg.Date,
		arg.ID,
	)
}

const updatePayment = `-- name: UpdatePayment :execresult
UPDATE payment
SET
  ref_month = ?,
  payment_date = ?
WHERE
  id = ?
`

type UpdatePaymentParams struct {
	RefMonth    string
	PaymentDate time.Time
	ID          int64
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updatePayment, arg.RefMonth, arg.PaymentDate, arg.ID)
}

const updatePresence = `-- name: UpdatePresence :execresult
UPDATE presence
SET
  is_presence = ?
WHERE
  number_card = ?
  AND meeting_id = ?
`

type UpdatePresenceParams struct {
	IsPresence bool
	NumberCard int64
	MeetingID  int64
}

func (q *Queries) UpdatePresence(ctx context.Context, arg UpdatePresenceParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updatePresence, arg.IsPresence, arg.NumberCard, arg.MeetingID)
}
