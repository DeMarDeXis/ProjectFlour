package excelStorage

import (
	"ProjectFlour/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PartsExcel struct {
	db *sqlx.DB
}

func NewExcelPartsStorage(db *sqlx.DB) *PartsExcel {
	return &PartsExcel{
		db: db,
	}
}

func (p *PartsExcel) AddTypeProductFromExcel(production []model.TypeProduct) error {
	const op = "storage.postgres.excelStorage.AddTypeProductFromExcel"

	tx, err := p.db.Begin()
	if err != nil {
		return fmt.Errorf("tx.Begin - %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `INSERT INTO type_products (name_type_product, coefficient) VALUES ($1, $2) ON CONFLICT DO NOTHING`

	for i, prod := range production {
		_, err = tx.Exec(query, prod.NameType, prod.Coefficient)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		fmt.Printf("Запись %d из %d завершена\n", i+1, len(production))
	}

	return tx.Commit()
}

func (p *PartsExcel) AddTypeMaterialFromExcel(material []model.TypeMaterial) error {
	const op = "storage.postgres.excelStorage.AddTypeMaterialFromExcel"

	tx, err := p.db.Begin()
	if err != nil {
		return fmt.Errorf("tx.Begin - %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `INSERT INTO type_materials (name_type_material, percent_of_marriage) VALUES ($1, $2) ON CONFLICT DO NOTHING`

	for i, mat := range material {
		_, err = tx.Exec(query, mat.NameType, mat.PercentOfMarriage)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		fmt.Printf("Запись %d из %d завершена\n", i+1, len(material))
	}

	return tx.Commit()
}

func (p *PartsExcel) AddProductsFromExcel(production []model.Product) error {
	const op = "storage.postgres.excelStorage.AddProductsFromExcel"

	tx, err := p.db.Beginx()
	if err != nil {
		return fmt.Errorf("tx.Begin - %w in %s", err, op)
	}

	mainQ := `INSERT INTO products (name_product, type_product_id, article, min_price_for_partner) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	helperJoinQ := `SELECT id FROM type_products WHERE name_type_product = $1`

	var typeID int

	for _, prod := range production {
		if prod.TypeProduct == "" {
			prod.TypeProduct = "Не определено"
		}
		err = tx.QueryRow(helperJoinQ, prod.TypeProduct).Scan(&typeID)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		_, err := tx.Exec(mainQ, prod.Name, typeID, prod.Article, prod.MinPriceForPartner)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return tx.Commit()
}

func (p *PartsExcel) AddPartnersFromExcel(partners []model.Partners) error {
	const op = "storage.postgres.excelStorage.AddPartnersFromExcel"

	tx, err := p.db.Beginx()
	if err != nil {
		return fmt.Errorf("tx.Begin - %w in %s", err, op)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Запросы для работы с organization_type
	isExistOrgTypeQ := `SELECT id FROM organizations_type WHERE name_of_type_organization = $1`
	insertOrgTypeQ := `INSERT INTO organizations_type (name_of_type_organization) VALUES ($1) RETURNING id`

	// Запрос для вставки партнера
	insertPartnerQ := `INSERT INTO partners (type_organization, name_partner, full_name_boss, email, phone_number, legal_address, inn, rate) 
					   VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING`

	for i, partner := range partners {
		var orgTypeID int

		if partner.TypeOrg == "" {
			partner.TypeOrg = "Не определено"
		}

		// Проверяем существует ли организационно-правовая форма
		err = tx.QueryRow(isExistOrgTypeQ, partner.TypeOrg).Scan(&orgTypeID)
		if err != nil {
			// Если не найдена, создаем новую
			if err.Error() == "sql: no rows in result set" {
				err = tx.QueryRow(insertOrgTypeQ, partner.TypeOrg).Scan(&orgTypeID)
				if err != nil {
					return fmt.Errorf("%s: failed to insert organization type '%s' - %w", op, partner.TypeOrg, err)
				}
				fmt.Printf("Создана новая организационно-правовая форма: %s (ID: %d)\n", partner.TypeOrg, orgTypeID)
			} else {
				return fmt.Errorf("%s: failed to check organization type '%s' - %w", op, partner.TypeOrg, err)
			}
		}

		// Вставляем партнера с найденным или созданным ID организационно-правовой формы
		_, err = tx.Exec(insertPartnerQ, orgTypeID, partner.Name, partner.FullNameOfBoss,
			partner.Email, partner.PhoneNumber, partner.LegalAddress, partner.INN, partner.Rate)
		if err != nil {
			return fmt.Errorf("%s: failed to insert partner '%s' - %w", op, partner.Name, err)
		}

		fmt.Printf("Запись %d из %d завершена (партнер: %s)\n", i+1, len(partners), partner.Name)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: failed to commit transaction - %w", op, err)
	}

	return nil
}

func (p *PartsExcel) AddProductPartnersFromExcel(productsPartners []model.ProductsPartners) error {
	const op = "storage.postgres.excelStorage.AddProductPartnersFromExcel"

	tx, err := p.db.Beginx()
	if err != nil {
		return fmt.Errorf("tx.Begin - %w in %s", err, op)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	getPartnerIDQ := `SELECT id FROM partners WHERE name_partner=$1`
	getProductIDQ := `SELECT id FROM products WHERE name_product=$1`

	insertPartnerQ := `INSERT INTO partner_products (partner_id, product_id, quantity_of_product, date_of_sale) 
						VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`

	for _, productPartner := range productsPartners {
		var PartnerID int
		var ProductID int

		err = tx.QueryRow(getPartnerIDQ, productPartner.Partner).Scan(&PartnerID)
		if err != nil {
			return fmt.Errorf("%s: failed to get partner ID for '%s' - %w", op, productPartner.Partner, err)
		}
		err = tx.QueryRow(getProductIDQ, productPartner.Product).Scan(&ProductID)
		if err != nil {
			return fmt.Errorf("%s: failed to get product ID for '%s' - %w", op, productPartner.Product, err)
		}

		_, err = tx.Exec(insertPartnerQ, PartnerID, ProductID, productPartner.Quantity, productPartner.DateOfSale)
		if err != nil {
			return fmt.Errorf("%s: failed to insert partner-product relation - %w", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: failed to commit transaction - %w", op, err)
	}

	return nil
}
