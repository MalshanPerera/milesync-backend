package services

import (
	"context"
	"errors"
	errpkg "jira-for-peasants/errors"
	repo "jira-for-peasants/repositories"
	"jira-for-peasants/utils"
)

type UserService struct {
	userRepository    repo.UserRepository
	sessionRepository repo.SessionRepository
}

func NewUserService(
	userRepository *repo.UserRepository,
	sessionRepository *repo.SessionRepository,
) *UserService {
	return &UserService{
		userRepository:    *userRepository,
		sessionRepository: *sessionRepository,
	}
}

type CreateUserParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type LoginUserParams struct {
	Email    string
	Password string
}

type UpdateUserParams struct {
	FirstName string
	LastName  string
	Email     string
}

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (repo.UserModel, repo.SessionModel, error) {
	tx, err := s.userRepository.BeginTx(ctx)

	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	defer func() {
		err = s.userRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	newUser, err := s.userRepository.CreateUser(ctx, tx, repo.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  params.Password,
	})

	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	accessToken, expiredAt, err := utils.CreateToken(newUser.ID, utils.Type.AccessToken)
	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	refreshToken, _, err := utils.CreateToken(newUser.ID, utils.Type.RefreshToken)
	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	session, err := s.sessionRepository.CreateSession(ctx, tx, repo.CreateSessionParams{
		UserID:       newUser.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	})

	tx.Commit(ctx)

	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	return newUser, session, nil
}

func (s *UserService) LoginUser(ctx context.Context, params LoginUserParams) (repo.UserModel, repo.SessionModel, error) {
	tx, err := s.userRepository.BeginTx(ctx)

	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	defer func() {
		err = s.userRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	existingUser, err := s.userRepository.GetUserFromEmail(ctx, params.Email)

	if err != nil {
		if errors.Is(err, errpkg.NoResults) {
			return repo.UserModel{}, repo.SessionModel{}, errpkg.NoResults
		}
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	match, err := utils.ComparePasswordAndHash(params.Password, existingUser.Password)

	if !match || err != nil {
		return repo.UserModel{}, repo.SessionModel{}, errpkg.NoResults
	}

	accessToken, expiredAt, err := utils.CreateToken(existingUser.ID, utils.Type.AccessToken)
	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	refreshToken, _, err := utils.CreateToken(existingUser.ID, utils.Type.RefreshToken)
	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	session, err := s.sessionRepository.UpdateSession(ctx, tx, repo.CreateSessionParams{
		UserID:       existingUser.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	})

	if errors.Is(err, errpkg.NoResults) {
		session, err = s.sessionRepository.CreateSession(ctx, tx, repo.CreateSessionParams{
			UserID:       existingUser.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    expiredAt,
		})
	}

	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.UserModel{}, repo.SessionModel{}, err
	}

	return existingUser, session, nil
}
