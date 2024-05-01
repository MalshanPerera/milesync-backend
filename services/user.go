package services

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"
	"jira-for-peasants/errors"
	"jira-for-peasants/utils"
)

type UserService struct {
	db *datastore.DB
}

func NewUserService(db *datastore.DB) *UserService {
	return &UserService{
		db: db,
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

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (db.User, db.Session, error) {
	hashedPassword, e := utils.GenerateFromPassword(params.Password)
	if e != nil {
		return db.User{}, db.Session{}, e
	}
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return db.User{}, db.Session{}, errors.NewDBError(err.Error())
	}

	defer func() {
		err = s.db.RollbackTx(ctx, tx)
	}()

	newUser, e := s.db.GetQuery().WithTx(tx).CreateUser(ctx, db.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  hashedPassword,
	})

	if e != nil {
		return db.User{}, db.Session{}, e
	}

	accessToken, expiredAt, err := utils.CreateToken(newUser.ID, utils.Type.AccessToken)
	if err != nil {
		return db.User{}, db.Session{}, err
	}

	refreshToken, _, err := utils.CreateToken(newUser.ID, utils.Type.RefreshToken)
	if err != nil {
		return db.User{}, db.Session{}, err
	}

	session, e := s.db.GetQuery().WithTx(tx).CreateSession(ctx, db.CreateSessionParams{
		UserID:       newUser.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	})

	if e != nil {
		return db.User{}, db.Session{}, e
	}

	e = s.db.CommitTx(ctx, tx)

	if e != nil {
		return db.User{}, db.Session{}, e
	}

	return newUser, session, nil
}

func (s *UserService) LoginUser(ctx context.Context, params LoginUserParams) (db.Session, error) {
	existingUser, e := s.db.GetQuery().GetUserFromEmail(ctx, params.Email)

	if e != nil {
		return db.Session{}, e
	}

	match, e := utils.ComparePasswordAndHash(params.Password, existingUser.Password)

	if !match || e != nil {
		return db.Session{}, e
	}

	accessToken, expiredAt, err := utils.CreateToken(existingUser.ID, utils.Type.AccessToken)
	if err != nil {
		return db.Session{}, err
	}

	refreshToken, _, err := utils.CreateToken(existingUser.ID, utils.Type.RefreshToken)
	if err != nil {
		return db.Session{}, err
	}

	session, e := s.db.GetQuery().UpdateSession(ctx, db.UpdateSessionParams{
		UserID:       existingUser.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	})

	if e == errors.NoResults {
		session, e = s.db.GetQuery().CreateSession(ctx, db.CreateSessionParams{
			UserID:       existingUser.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    expiredAt,
		})
	}

	if e != nil {
		return db.Session{}, e
	}

	return session, nil
}

func (s *UserService) GetUserFromId(ctx context.Context, id string) (db.User, error) {
	existingUser, e := s.db.GetQuery().GetUser(ctx, id)

	if e != nil {
		return db.User{}, e
	}

	return existingUser, e
}

func (s *UserService) GetUserFromEmail(ctx context.Context, email string) (db.User, error) {
	existingUser, e := s.db.GetQuery().GetUserFromEmail(ctx, email)

	if e != nil {
		return db.User{}, e
	}

	return existingUser, e
}

func (s *UserService) UpdateUser(ctx context.Context, id string, params UpdateUserParams) (db.User, error) {

	updatedUser, e := s.db.GetQuery().UpdateUser(ctx, db.UpdateUserParams{
		ID:        id,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
	})

	if e != nil {
		return db.User{}, e
	}

	return updatedUser, e
}
