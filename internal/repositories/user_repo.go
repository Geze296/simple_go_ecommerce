package repositories

import (
	"database/sql"
	"time"

	"github.com/geze296/simple_go_ecommerce/internal/models"
)

type UserRepo interface {
	Create(user *models.User) error
	Update(id uint, user *models.User) (*models.User, error)
	FindById(id uint) (*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository{
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {

	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),
	)
	return err
}
