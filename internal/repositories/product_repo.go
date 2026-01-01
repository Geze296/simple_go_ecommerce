package repositories

import (
	"database/sql"
	"time"

	"github.com/geze296/simple_go_ecommerce/internal/models"
)

type ProductRepo interface{
	Create(product *models.Product) error
	Update(id uint, product *models.Product) error
	GetAll() ([]*models.Product, error) 
	GetById(id uint) (*models.Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepository {
	return &ProductRepository{db:db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	query := `INSERT INTO products (name, price, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	_,err := r.db.Exec(query, 
		product.Name,
		product.Price,
		time.Now(),
		time.Now(),
	)
	if err != nil{
		return err
	}

	return nil
}


func (r *ProductRepository) GetAll() ([]models.Product, error){
	var products []models.Product
	
	query := `SELECT id, name, price, created_at, updated_at FROM products`
	row, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	for row.Next(){
		var product models.Product

		err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Created_At,
			&product.Updated_At,
		)

		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}