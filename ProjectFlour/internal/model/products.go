package model

type TypeMaterial struct {
	ID                int     `db:"id"`
	NameType          string  `db:"name_type_material"`
	PercentOfMarriage float32 `db:"percent_of_marriage"`
}

type TypeProduct struct {
	ID          int     `db:"id"`
	NameType    string  `db:"name_type_product"`
	Coefficient float32 `db:"coefficient"`
}

type Product struct {
	ID                 int     `db:"id"`
	TypeProduct        string  `db:"name_type_product"`
	Name               string  `db:"name_product"`
	Article            string  `db:"article"`
	MinPriceForPartner float32 `db:"min_price_for_partner"`
}
