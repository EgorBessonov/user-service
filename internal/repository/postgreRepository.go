//Package repository applies for connection with databases
package repository

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/user-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

//PostgresRepository type
type PostgresRepository struct {
	DBconn *pgxpool.Pool
}

//NewPostgresRepository returns new repository instance
func NewPostgresRepository(conn *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{DBconn: conn}
}

//Save method create new user instance in database
func (rps *PostgresRepository) Save(ctx context.Context, user *model.User) error {
	_, err := rps.DBconn.Exec(ctx, `insert into users (name, email, password)
	values ($1, $2, $3)`, user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("poostgres repository: can't save user - %e", err)
	}
	return nil
}

//Get method returns user information from database
func (rps *PostgresRepository) Get(ctx context.Context, userEmail string) (*model.User, error) {
	var user model.User
	err := rps.DBconn.QueryRow(ctx, `select id, name, email, password from users
	where email=$1`, userEmail).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("postgres: can't get user")
	}
	return &user, nil
}
