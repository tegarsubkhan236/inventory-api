package repository

import (
	"encoding/json"
	"inventory-api/pkg/api/request"
	"log"
)

func (s *storage) CreateProduct(r request.Product, categories []request.ProductCategory) error {

	statement := `INSERT INTO inv_product (name, product_desc,product_photo, product_category_id) VALUES ($1, $2, $3, $4);`

	log.Print("REPO :", r)
	uye, _ := json.Marshal(categories)
	err := s.db.QueryRow(statement, r.Name, r.ProductDesc, r.ProductPhotoUrl, uye).Err() //pq.Array()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetProductByID(ProductID int) (request.Product, error) {
	var jsonData []byte
	var item request.Product

	err := s.db.QueryRow(`SELECT id,name, product_desc,product_photo, product_category_id FROM inv_product WHERE id=$1`, ProductID).Scan(
		&item.ID, &item.Name, &item.ProductDesc, &item.ProductPhotoUrl, &jsonData)
	_ = json.Unmarshal(jsonData, &item.ProductCategoryID)
	if err != nil {
		return request.Product{}, err
	}
	return item, nil
}

func (s *storage) GetProductByCategory(Category string) ([]request.Product, error) {
	statement := `SELECT * FROM inv_product WHERE product_category_id @> $1;`

	rows, err := s.db.Query(statement, Category)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var data []request.Product
	var jsonData []byte

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item request.Product
		if err := rows.Scan(&item.ID, &item.Name, &item.ProductDesc, &item.ProductPhotoUrl, &jsonData); err != nil {
			return data, err
		}
		_ = json.Unmarshal(jsonData, &item.ProductCategoryID)
		data = append(data, item)
	}

	return data, nil
}

func (s *storage) ListProduct() ([]request.Product, error) {
	statement := `SELECT * FROM inv_product`

	rows, err := s.db.Query(statement)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var data []request.Product
	var jsonData []byte

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item request.Product
		if err := rows.Scan(&item.ID, &item.Name, &item.ProductDesc, &item.ProductPhotoUrl, &jsonData); err != nil {
			return data, err
		}
		_ = json.Unmarshal(jsonData, &item.ProductCategoryID)
		data = append(data, item)
	}

	return data, nil
}

func (s *storage) UpdateProduct(ProductID int, request request.UpdateProductRequest) error {
	statement := `UPDATE inv_product SET name = $1, product_desc = $2 WHERE id = $3`

	err := s.db.QueryRow(statement, request.Name, request.ProductDesc, ProductID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return err
}

func (s *storage) UpdateCategoryByProductID(ProductID int, CategoryID []request.ProductCategory) error {
	statement := `UPDATE inv_product SET product_category_id = $1 WHERE id = $2`

	uye, _ := json.Marshal(CategoryID)
	err := s.db.QueryRow(statement, uye, ProductID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return err
}

func (s *storage) DeleteProduct(ProductID int) error {
	statement := `DELETE FROM inv_product WHERE id = $1`

	err := s.db.QueryRow(statement, ProductID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
