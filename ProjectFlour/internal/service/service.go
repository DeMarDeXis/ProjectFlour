package service

import (
	"ProjectFlour/internal/events"
	"ProjectFlour/internal/model"
	"ProjectFlour/internal/service/excelHandler"
	"ProjectFlour/internal/storage"
)

type AuthService interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ExcelImportService interface {
	AddTypeProductFromExcel(filepath string) error
	AddTypeMaterialsFromExcel(filepath string) error
	AddProductsFromExcel(filepath string) error
	AddPartnersFromExcel(filepath string) error
	AddProductPartnersFromExcel(filepath string) error
}

type ProductsService interface {
	//AddTypeMaterial
	GetAllTypesMaterial() (*[]model.TypeMaterial, error)
	//AddTypeProduct
	GetAllTypeProduct() (*[]model.TypeProduct, error)
	//AddProduct
	GetAllProducts() (*[]model.Product, error)
	//AllPartners
	GetAllPartners() (*[]model.Partners, error)
	//AllProductsPartner
	GetAllProductsPartner() (*[]model.ProductsPartners, error)
}

type Service struct {
	AuthService
	ExcelImportService
	ProductsService
	eventBus *events.EventBus
}

func New(storage *storage.Storage, eb *events.EventBus) *Service {
	return &Service{
		AuthService:        NewAuthService(storage),
		ExcelImportService: excelHandler.New(storage, eb),
		ProductsService:    NewProductsService(storage),
		eventBus:           eb,
	}
}
