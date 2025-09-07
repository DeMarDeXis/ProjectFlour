package postgres

import (
	"ProjectFlour/internal/model"
	"fmt"
)

//type PartnersStorage struct {
//	db *sqlx.DB
//}
//
//func NewStoragePartners(db *sqlx.DB) *PartnersStorage {
//	return &PartnersStorage{
//		db: db,
//	}
//}

func (p *ProductsPartners) GetAllPartners() (*[]model.Partners, error) {
	const op = "storage.postgres.products.GetAllPartners"

	var partners []model.Partners

	q := `SELECT p.id, ot.name_of_type_organization, p.name_partner, p.full_name_boss, 
       		p.email, p.phone_number, p.legal_address, p.inn, p.rate 
			FROM partners p
			JOIN organizations_type ot
			ON p.type_organization = ot.id`
	err := p.db.Select(&partners, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &partners, nil
}

func (p *ProductsPartners) GetAllProductsPartner() (*[]model.ProductsPartners, error) {
	const op = "storage.postgres.products.GetAllProductsPartner"

	var prodPartners []model.ProductsPartners

	q := `SELECT pp.id, p.name_partner, pr.name_product, pp.quantity_of_product, pp.date_of_sale   
       	FROM partner_products pp 
    	JOIN partners p 
    	    ON pp.partner_id = p.id
		JOIN products pr
			ON pp.product_id = pr.id`

	err := p.db.Select(&prodPartners, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &prodPartners, nil
}
