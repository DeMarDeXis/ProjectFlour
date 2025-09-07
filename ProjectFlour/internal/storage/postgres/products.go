package postgres

import (
	"ProjectFlour/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProductsPartners struct {
	db *sqlx.DB
}

func NewStorageProduct(db *sqlx.DB) *ProductsPartners {
	return &ProductsPartners{
		db: db,
	}
}

func (t *ProductsPartners) GetAllTypesMaterial() (*[]model.TypeMaterial, error) {
	const op = "storage.postgres.products.GetAllTypesMaterial"

	var typeMaterials []model.TypeMaterial

	q := "SELECT id, name_type_material, percent_of_marriage FROM type_materials"
	err := t.db.Select(&typeMaterials, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &typeMaterials, nil
}

func (t *ProductsPartners) GetAllTypeProduct() (*[]model.TypeProduct, error) {
	const op = "storage.postgres.products.GetAllTypeProduct"

	var typeProducts []model.TypeProduct
	q := "SELECT id, name_type_product, coefficient FROM type_products"

	err := t.db.Select(&typeProducts, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &typeProducts, nil
}

func (t *ProductsPartners) GetAllProducts() (*[]model.Product, error) {
	const op = "storage.postgres.products.GetAllProduct"

	var typeProducts []model.Product

	q := `SELECT p.id, tp.name_type_product, p.name_product, p.article, p.min_price_for_partner FROM type_products tp
			JOIN products p ON tp.id = p.type_product_id`
	err := t.db.Select(&typeProducts, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &typeProducts, nil
}
