package store

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ID          int            `json:"id"`
	ImageUrl    *string        `json:"image_url"`
	DisplayName string         `json:"display_name"`
	Rating      *int           `json:"rating"`
	Description *string        `json:"description"`
	Category    string         `json:"category"`
	Activation  *bool          `json:"activation"`
	Entries     []ProductEntry `json:"entries"`
}

type ProductEntry struct {
	ID             int     `json:"id"`
	Quantity       int     `json:"quantity"`
	Price          int     `json:"price"`
	Review         *string `json:"review"`
	WarrantyPeriod *string `json:"warranty"`
	Rating         *int    `json:"rating"`
}

type PostgresProductStore struct {
	db *sql.DB
}

func NewPostgresProductStore(db *sql.DB) *PostgresProductStore {
	return &PostgresProductStore{db: db}
}

type ProductStore interface {
	CreateProduct(*Product) (*Product, error)
	GetProductById(id int64) (*Product, error)
}

func (pg *PostgresProductStore) CreateProduct(product *Product) (*Product, error) {

	tx, err := pg.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query :=
		`
	INSERT INTO products (display_name, rating, description, category, activation, image_url)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	fmt.Println(product)
	err = tx.QueryRow(query, product.DisplayName, product.Rating, product.Description, product.Category, product.Activation, product.ImageUrl).Scan(&product.ID)

	if err != nil {
		return nil, err
	}

	for _, entry := range product.Entries {

		query := `
		INSERT INTO product_entries (product_id, quantity, price, review, warranty_period, rating)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`
		err := tx.QueryRow(query, product.ID, entry.Quantity, entry.Price, entry.Review, entry.WarrantyPeriod, entry.Rating).Scan(&entry.ID)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pg *PostgresProductStore) GetProductById(id int64) (*Product, error) {
	product := &Product{}
	return product, nil
}
