package model

type Partners struct {
	ID             int    `db:"id"`
	TypeOrg        string `db:"name_of_type_organization"`
	Name           string `db:"name_partner"`
	FullNameOfBoss string `db:"full_name_boss"`
	Email          string `db:"email"`
	PhoneNumber    string `db:"phone_number"`
	LegalAddress   string `db:"legal_address"`
	INN            string `db:"inn"`
	Rate           int    `db:"rate"`
}

type OrgType struct {
	ID            int    `db:"id"`
	NameOfTypeOrg string `db:"name_of_type_organization"`
}

type ProductsPartners struct {
	ID              int     `db:"id"`
	Partner         string  `db:"name_partner"`
	Product         string  `db:"name_product"`
	Quantity        int     `db:"quantity_of_product"`
	DateOfSale      string  `db:"date_of_sale"`
	PriceForPartner float32 `db:"price_for_partner"`
}
