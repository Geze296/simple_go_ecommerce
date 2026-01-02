package repositories

import (
	"database/sql"
	"time"

	"github.com/geze296/simple_go_ecommerce/internal/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	query := `INSERT INTO orders (price, user_id, product_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query,
		order.Price,
		order.UserID,
		order.ProductID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, user_id, product_id, price FROM orders`

	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var order models.Order
		err := row.Scan(
			&order.ID,
			&order.UserID,
			&order.ProductID,
			&order.Price,
		)

		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
