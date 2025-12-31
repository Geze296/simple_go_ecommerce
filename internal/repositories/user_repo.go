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
	GetAll() ([]*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
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



func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User

	query := `SELECT id,
			name,
			email,
			created_at,
			updated_at FROM users`

	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		var user models.User

		err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}
	return users, nil
}


func (r *UserRepository) FindById(id uint) (*models.User, error){
	var user models.User

	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`

	row := r.db.QueryRow(query,id)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}