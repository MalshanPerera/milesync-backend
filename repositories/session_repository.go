package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type SessionModel db.Session

type SessionRepository struct {
	*datastore.Trx
	db *datastore.DB
}

type CreateSessionParams struct {
	UserID       string
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

func NewSessionRepository(db *datastore.DB) *SessionRepository {
	return &SessionRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
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

func (repo *SessionRepository) GetSessionByUserId(ctx context.Context, userId string) (SessionModel, error) {
	session, err := repo.db.GetQuery().GetSessionByUserId(ctx, userId)
	if err != nil {
		return SessionModel{}, err
	}

	return SessionModel(session), nil
}

func (repo *SessionRepository) DeleteSession(ctx context.Context, tx pgx.Tx, userId string) error {
	return repo.db.GetQuery().WithTx(tx).DeleteSession(ctx, userId)
}
