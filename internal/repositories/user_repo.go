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

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) Create(user *models.User) error {

	query := `INSERT INTO users (name, email, password, phone, role, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`

	_, err := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Password,
		user.Phone,
		user.Role,
		time.Now(),
		time.Now(),
	)
	return err
}
