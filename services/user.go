package services

import (
	"context"
	"jira-for-peasents/common"
	datastore "jira-for-peasents/db"
	db "jira-for-peasents/db/sqlc"
	service_utils "jira-for-peasents/services/utils"
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
	Name     string
	Email    string
	Password string
}

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (db.User, error) {
	hashedPassword, e := service_utils.GenerateFromPassword(params.Password)
	if e != nil {
		return db.User{}, e
	}
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return db.User{}, common.NewDBError(err.Error())
	}

	defer func() {
		err = s.db.RollbackTx(ctx, tx)
	}()

	newUser, e := db.New(tx).CreateUser(ctx, db.CreateUserParams{
		Name:     params.Name,
		Email:    params.Email,
		Password: hashedPassword,
		ID:       datastore.GenerateId(),
	})

	if e != nil {
		return db.User{}, e
	}

	e = s.db.CommitTx(ctx, tx)

	return newUser, e
}

func (s *UserService) GetUserFromId(ctx context.Context, id string) (db.User, error) {
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return db.User{}, common.NewDBError(err.Error())
	}

	existingUser, e := db.New(tx).GetUser(ctx, id)

	if e != nil {
		return db.User{}, e
	}

	return existingUser, e
}

func (s *UserService) GetUserFromEmail(ctx context.Context, email string) (db.User, error) {
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return db.User{}, common.NewDBError(err.Error())
	}

	existingUser, e := db.New(tx).GetUserFromEmail(ctx, email)

	if e != nil {
		return db.User{}, e
	}

	return existingUser, e
}

type UpdateUserParams struct {
	Name  string
	Email string
}

func (s *UserService) UpdateUser(ctx context.Context, id string, params UpdateUserParams) (db.User, error) {
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return db.User{}, err
	}

	defer func() {
		err = s.db.RollbackTx(ctx, tx)
	}()

	updatedUser, e := db.New(tx).UpdateUser(ctx, db.UpdateUserParams{
		ID:    id,
		Name:  params.Name,
		Email: params.Email,
	})

	if e != nil {
		return db.User{}, e
	}

	e = s.db.CommitTx(ctx, tx)

	return updatedUser, e
}
