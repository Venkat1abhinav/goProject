package store

import (
	"context"
	"database/sql"
	"fmt"
)

type Product struct {
	ID          int            `json:"id"`
	ImageUrl    *string        `json:"image_url"`
	DisplayName string         `json:"display_name"`
	Rating      *int           `json:"rating"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Activation  *bool          `json:"activation"`
	Entries     []ProductEntry `json:"entries"`
}

type ProductEntry struct {
	ID             int     `json:"id"`
	Quantity       int     `json:"quantity"`
	Price          float32 `json:"price"`
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
	UpdateProduct(*Product) error
	DeleteProduct(id int64) error
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

	tx, err := pg.db.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly: true,
	})

	if err != nil {
		return product, err
	}

	query := `
	SELECT id, display_name, rating, description, category, activation, image_url FROM products WHERE id = $1 
	`

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	err = tx.QueryRow(query, id).Scan(&product.ID, &product.DisplayName, &product.Rating, &product.Description, &product.Category, &product.Activation, &product.ImageUrl)

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	entryQuery := `
	SELECT id, quantity, price, review, warranty_period, rating FROM product_entries 
	WHERE product_id = $1
	ORDER BY rating DESC
	`

	rows, err := tx.Query(entryQuery, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry ProductEntry

		err = rows.Scan(
			&entry.ID,
			&entry.Quantity,
			&entry.Price,
			&entry.Review,
			&entry.WarrantyPeriod,
			&entry.Rating,
		)

		if err != nil {
			return nil, err
		}
		product.Entries = append(product.Entries, entry)
	}

	err = tx.Commit()

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pg *PostgresProductStore) UpdateProduct(product *Product) error {

	tx, err := pg.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
	UPDATE products
	SET display_name = $1, rating = $2, description = $3, category = $4, activation = $5, image_url = $6
	WHERE id = $7
	`

	result, err := tx.Exec(query, product.DisplayName, product.Rating, product.Description, product.Category, product.Activation, product.ImageUrl, product.ID)

	if err != nil {
		return err
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsEffected == 0 {
		return sql.ErrNoRows
	}

	deleteQuery := `
	DELETE FROM product_entries WHERE product_id = $1
	`
	_, err = tx.Exec(deleteQuery, product.ID)

	if err != nil {
		return err
	}

	for _, entry := range product.Entries {
		query = `
		INSERT INTO product_entries (product_id, quantity, price, review, warranty_period, rating)
		VALUES ($1, $2, $3, $4, $5, $6)
		`
		_, err = tx.Exec(query,
			product.ID,
			entry.Quantity,
			entry.Price,
			entry.Review,
			entry.WarrantyPeriod,
			entry.Rating,
		)

		if err != nil {
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (pg *PostgresProductStore) DeleteProduct(id int64) error {

	tx, err := pg.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
	DELETE from products WHERE id = $1
	`

	result, err := tx.Exec(query, id)

	if err != nil {
		return err
	}

	res, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if res == 0 {
		return sql.ErrNoRows
	}

	err = tx.Commit()

	if err != nil {
		return err
	}
	return nil
}
