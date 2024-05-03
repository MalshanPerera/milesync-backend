package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type SessionModel db.Session

type SessionRepository struct {
	db *datastore.DB
}

func NewSessionRepository(db *datastore.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

type CreateSessionParams struct {
	UserID       string
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

func (repo *SessionRepository) CreateSession(ctx context.Context, tx pgx.Tx, params CreateSessionParams) (SessionModel, error) {
	session, e := repo.db.GetQuery().WithTx(tx).CreateSession(ctx, db.CreateSessionParams{
		UserID:       params.UserID,
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
		ExpiresAt:    params.ExpiresAt,
	})

	if e != nil {
		return SessionModel{}, e
	}

	return SessionModel(session), nil
}

func (repo *SessionRepository) UpdateSession(ctx context.Context, tx pgx.Tx, params CreateSessionParams) (SessionModel, error) {
	session, e := repo.db.GetQuery().WithTx(tx).UpdateSession(ctx, db.UpdateSessionParams{
		UserID:       params.UserID,
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
		ExpiresAt:    params.ExpiresAt,
	})
	if e != nil {
		return SessionModel{}, e
	}

	return SessionModel(session), nil
}
