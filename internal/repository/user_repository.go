package repository

import (
	"nongki/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	GetMe(userID string) (*domain.User, error)
	UpdateUser(userID string, userUpdate domain.User) (*domain.User, error)
	DeleteUser(userID string, deletedBy string) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetMe(userID string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&user, query, userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	err := user.BeforeCreate()
	if err != nil {
		return err
	}

	query := `INSERT INTO users (id, name, email, password, address, gender, marital_status, created_at, created_by, updated_at, updated_by)
              VALUES (:id, :name, :email, :password, :address, :gender, :marital_status, :created_at, :created_by, :updated_at, :updated_by)`

	_, err = r.db.NamedExec(query, user)
	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(userID string, userUpdate domain.User) (*domain.User, error) {
	updatedUser := domain.User{}
	query := `
			UPDATE users
			SET name = $1, email = $2, address = $3, gender = $4, marital_status = $5, updated_at = NOW()
			WHERE id = $6 AND deleted_at IS NULL
			RETURNING id, name, email, address, gender, marital_status, updated_at
	`
	err := r.db.QueryRow(query, userUpdate.Name, userUpdate.Email, userUpdate.Address, userUpdate.Gender, userUpdate.MaritalStatus, userID).
		Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.Address, &updatedUser.Gender, &updatedUser.MaritalStatus, &updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *userRepository) DeleteUser(userID string, deletedBy string) error {
	query := `
			UPDATE users
			SET deleted_at = $1, deleted_by = $2
			WHERE id = $3 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, time.Now(), deletedBy, userID)
	if err != nil {
		return err
	}

	return nil
}
