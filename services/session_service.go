package services

import (
	"context"
	errpkg "jira-for-peasants/errors"
	repo "jira-for-peasants/repositories"
	"jira-for-peasants/utils"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type SessionService struct {
	sessionRepository *repo.SessionRepository
}

type CreateSessionParams struct {
	UserID       string
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

func NewSessionService(sessionRepo *repo.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepo,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, params CreateSessionParams) (repo.SessionModel, error) {
	tx, err := s.sessionRepository.BeginTx(ctx)
	if err != nil {
		return repo.SessionModel{}, err
	}

	defer func() {
		err = s.sessionRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	session, err := s.sessionRepository.CreateSession(ctx, tx, repo.CreateSessionParams{
		UserID:       params.UserID,
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
		ExpiresAt:    params.ExpiresAt,
	})

	if err != nil {
		return repo.SessionModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.SessionModel{}, err
	}

	return session, nil
}

func (s *SessionService) UpdateSession(ctx context.Context, params CreateSessionParams) (repo.SessionModel, error) {
	tx, err := s.sessionRepository.BeginTx(ctx)
	if err != nil {
		return repo.SessionModel{}, err
	}

	defer func() {
		err = s.sessionRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	session, err := s.sessionRepository.UpdateSession(ctx, tx, repo.CreateSessionParams{
		UserID:       params.UserID,
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
		ExpiresAt:    params.ExpiresAt,
	})

	if err != nil {
		return repo.SessionModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.SessionModel{}, err
	}

	return session, nil
}

func (s *SessionService) GetSessionByUserId(ctx context.Context, userId string) (repo.SessionModel, error) {
	session, err := s.sessionRepository.GetSessionByUserId(ctx, userId)
	if err != nil {
		return repo.SessionModel{}, err
	}

	return session, nil
}

func (s *SessionService) DeleteSession(ctx context.Context, tx pgx.Tx, userId string) error {
	return s.sessionRepository.DeleteSession(ctx, tx, userId)
}

func (s *SessionService) ValidateUserSession(ctx context.Context, userId string) (repo.SessionModel, error) {
	session, err := s.sessionRepository.GetSessionByUserId(ctx, userId)
	if err != nil {
		return repo.SessionModel{}, errpkg.NewApiError(http.StatusUnauthorized, "Unauthorized")
	}

	accessToken, expiredAt, err := utils.CreateToken(userId, utils.Type.AccessToken)
	if err != nil {
		return repo.SessionModel{}, err
	}

	refreshToken, _, err := utils.CreateToken(userId, utils.Type.RefreshToken)
	if err != nil {
		return repo.SessionModel{}, err
	}

	newSession, err := s.sessionRepository.UpdateSession(ctx, nil, repo.CreateSessionParams{
		UserID:       session.UserID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	})

	if err != nil {
		return repo.SessionModel{}, err
	}

	return newSession, nil
}
