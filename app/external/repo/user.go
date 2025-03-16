package repo

import (
	"auth-service/app/domain/user"
	"auth-service/types"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.Service
}

type userRepo struct {
	tableName string
	readDB    *sqlx.DB
	writeDB   *sqlx.DB
	psql      sq.StatementBuilderType
}

func NewUserRepo(db *DB) UserRepo {
	return &userRepo{
		tableName: `users`,
		readDB:    db.ReadDB,
		writeDB:   db.WriteDB,
		psql:      db.psql,
	}
}

func (u userRepo) Create(ctx context.Context, user types.SignUpUserPayload) error {
	query, arg, err := u.psql.Insert(u.tableName).Columns("first_name", "last_name", "email", "password").Values(user.FirstName, user.LastName, user.Email, user.Password).ToSql()

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	_, err = u.writeDB.ExecContext(ctx, query, arg...)

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (u userRepo) GetUserByEmail(ctx context.Context, email string) (*types.SignUpUser, error) {
	query, args, err := u.psql.Select("id", "first_name", "last_name", "email", "password", "created_at", "updated_at").From(u.tableName).Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	var user types.SignUpUser

	if err := u.readDB.GetContext(ctx, &user, query, args...); err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	return &user, nil
}
