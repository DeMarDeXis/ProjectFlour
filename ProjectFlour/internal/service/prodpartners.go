package service

import (
	"ProjectFlour/internal/model"
	"ProjectFlour/internal/storage"
)

type Products struct {
	storage storage.ProductStorage
}

func NewProductsService(strg storage.ProductStorage) *Products {
	return &Products{
		storage: strg,
	}
}

func (p *Products) GetAllTypesMaterial() (*[]model.TypeMaterial, error) {
	return p.storage.GetAllTypesMaterial()
}

func (p *Products) GetAllTypeProduct() (*[]model.TypeProduct, error) {
	return p.storage.GetAllTypeProduct()
}

func (p *Products) GetAllProducts() (*[]model.Product, error) {
	return p.storage.GetAllProducts()
}

func (p *Products) GetAllPartners() (*[]model.Partners, error) {
	return p.storage.GetAllPartners()
}

func (p *Products) GetAllProductsPartner() (*[]model.ProductsPartners, error) {
	return p.storage.GetAllProductsPartner()
}
