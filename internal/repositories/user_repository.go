package repositories

import (
	"context"

	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"
	"jira-for-peasants/pkg/utils"

	"github.com/jackc/pgx/v5"
)

type UserModel db.User

type CreateUserParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	ExpiresAt int64
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

type UserRepository struct {
	*datastore.Trx
	db *datastore.DB
}

func NewUserRepository(db *datastore.DB) *UserRepository {
	return &UserRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, tx pgx.Tx, params CreateUserParams) (UserModel, error) {
	hashedPassword, e := utils.GenerateFromPassword(params.Password)
	if e != nil {
		return UserModel{}, e
	}

	newUser, e := repo.db.GetQuery().WithTx(tx).CreateUser(ctx, db.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  hashedPassword,
	})

	if e != nil {
		return UserModel{}, e
	}

	return UserModel(newUser), nil
}

func (repo *UserRepository) GetUserFromId(ctx context.Context, id string) (UserModel, error) {
	existingUser, err := repo.db.GetQuery().GetUser(ctx, id)

	if err != nil {
		return UserModel{}, err
	}

	return UserModel(existingUser), err
}

func (repo *UserRepository) GetUserFromEmail(ctx context.Context, email string) (UserModel, error) {
	existingUser, err := repo.db.GetQuery().GetUserFromEmail(ctx, email)

	if err != nil {
		return UserModel{}, err
	}

	return UserModel(existingUser), err
}

func (repo *UserRepository) UpdateUser(ctx context.Context, tx pgx.Tx, id string, params UpdateUserParams) (UserModel, error) {

	updatedUser, err := repo.db.GetQuery().WithTx(tx).UpdateUser(ctx, db.UpdateUserParams{
		ID:        id,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
	})

	if err != nil {
		return UserModel{}, err
	}

	return UserModel(updatedUser), err
}
