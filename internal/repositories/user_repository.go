package repositories

import (
	"context"
	"errors"
	"go-web/internal/database"
	"go-web/internal/models"
)

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByUsernameWithRoles(ctx context.Context, username string) (*models.User, error)
}

type UserRepoImpl struct {
	DB *database.DB
}

func NewUserRepo(db *database.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}

func (r *UserRepoImpl) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT * FROM users WHERE username = $1`
	row := r.DB.Pool.QueryRow(ctx, query, username)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
		&user.DeletedAt,
		&user.DeletedBy,
	)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepoImpl) FindByUsernameWithRoles(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT
		 users.*,
		 roles.*
		FROM users
		LEFT JOIN user_roles ON users.user_id = user_roles.user_id
		LEFT JOIN roles ON user_roles.role_id = roles.role_id
		WHERE users.username = $1
	`
	
	rows, err := r.DB.Pool.Query(ctx, query, username)
	if err != nil {
		return nil, errors.New("failed to query user with roles")
	}
	defer rows.Close()

	var user *models.User
	var roles []models.Role
	var role models.Role

	for rows.Next() {
		user = &models.User{}
		err := rows.Scan(
			&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy, &user.DeletedAt, &user.DeletedBy,
			&role.ID, &role.Name, &role.CreatedAt, &role.CreatedBy, &role.UpdatedAt, &role.UpdatedBy, &role.DeletedAt, &role.DeletedBy,
		)
		if err != nil {
			return nil, errors.New("failed to scan user and role")
		}

		roles = append(roles, role)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Roles = roles

	return user, nil
}