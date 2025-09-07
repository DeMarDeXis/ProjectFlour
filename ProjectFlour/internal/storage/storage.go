package storage

import (
	"ProjectFlour/internal/model"
	"ProjectFlour/internal/storage/postgres"
	"ProjectFlour/internal/storage/postgres/excelStorage"
	"github.com/jmoiron/sqlx"
)

type AuthorizationStorage interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string, password string) (model.User, error)
}

type ExcelImportStorage interface {
	AddTypeProductFromExcel(production []model.TypeProduct) error
	AddTypeMaterialFromExcel(material []model.TypeMaterial) error
	AddProductsFromExcel(production []model.Product) error
	AddPartnersFromExcel(partners []model.Partners) error
	AddProductPartnersFromExcel(productsPartners []model.ProductsPartners) error
}

type MaterialStorage interface {
	//AddTypeMaterial
}

type ProductStorage interface {
	GetAllTypesMaterial() (*[]model.TypeMaterial, error)
	GetAllTypeProduct() (*[]model.TypeProduct, error)
	GetAllProducts() (*[]model.Product, error)
	GetAllPartners() (*[]model.Partners, error)
	GetAllProductsPartner() (*[]model.ProductsPartners, error)
}

type Storage struct {
	AuthorizationStorage
	ExcelImportStorage
	ProductStorage
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		AuthorizationStorage: postgres.NewAuthStorage(db),
		ProductStorage:       postgres.NewStorageProduct(db),
		ExcelImportStorage:   excelStorage.NewExcelPartsStorage(db),
	}
}
