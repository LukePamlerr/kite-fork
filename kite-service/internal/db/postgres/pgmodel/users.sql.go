// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package pgmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getUser = `-- name: GetUser :one
SELECT id, email, display_name, discord_id, discord_username, discord_avatar, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.DisplayName,
		&i.DiscordID,
		&i.DiscordUsername,
		&i.DiscordAvatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByDiscordID = `-- name: GetUserByDiscordID :one
SELECT id, email, display_name, discord_id, discord_username, discord_avatar, created_at, updated_at FROM users WHERE discord_id = $1
`

func (q *Queries) GetUserByDiscordID(ctx context.Context, discordID string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByDiscordID, discordID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.DisplayName,
		&i.DiscordID,
		&i.DiscordUsername,
		&i.DiscordAvatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, display_name, discord_id, discord_username, discord_avatar, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.DisplayName,
		&i.DiscordID,
		&i.DiscordUsername,
		&i.DiscordAvatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertUser = `-- name: UpsertUser :one
INSERT INTO users (
    id,
    email,
    display_name, 
    discord_id,
    discord_username,
    discord_avatar,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) ON CONFLICT (discord_id) 
DO UPDATE SET 
    email = EXCLUDED.email,
    display_name = EXCLUDED.display_name,
    discord_id = EXCLUDED.discord_id,
    discord_username = EXCLUDED.discord_username,
    discord_avatar = EXCLUDED.discord_avatar,
    updated_at = EXCLUDED.updated_at
RETURNING id, email, display_name, discord_id, discord_username, discord_avatar, created_at, updated_at
`

type UpsertUserParams struct {
	ID              string
	Email           string
	DisplayName     string
	DiscordID       string
	DiscordUsername string
	DiscordAvatar   pgtype.Text
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) (User, error) {
	row := q.db.QueryRow(ctx, upsertUser,
		arg.ID,
		arg.Email,
		arg.DisplayName,
		arg.DiscordID,
		arg.DiscordUsername,
		arg.DiscordAvatar,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.DisplayName,
		&i.DiscordID,
		&i.DiscordUsername,
		&i.DiscordAvatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
