package dataHandler

import (
	"ProjectFlour/internal/handlers/httpHandler/httplib"
	"ProjectFlour/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ProductsHTTP struct {
	service service.ProductsService
	logg    *slog.Logger
}

func NewProductsHTTPHandler(srvs service.ProductsService, logg *slog.Logger) *ProductsHTTP {
	return &ProductsHTTP{
		service: srvs,
		logg:    logg,
	}
}

// @Summary Получить все типы материалов
// @Description Возвращает все типы материалов
// @Tags materials
// @Produce json
// @Success 200 {array} model.TypeMaterial
// @Router /data/company/types_of_materials [get]
func (h *ProductsHTTP) GetAllTypesMaterial(w http.ResponseWriter, r *http.Request) {
	typesMaterial, err := h.service.GetAllTypesMaterial()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(typesMaterial)
}

// @Summary Получить все типы продуктов
// @Description Возвращает все типы продуктов
// @Tags products
// @Produce json
// @Success 200 {array} model.TypeProduct
// @Router /data/company/types_of_products [get]
func (h *ProductsHTTP) GetAllTypeProduct(w http.ResponseWriter, r *http.Request) {
	const op = "httpHandler.products.GetAllTypeProduct"

	typeProducts, err := h.service.GetAllTypeProduct()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(typeProducts)
}

// @Summary Получить все продукты
// @Description Возвращает все продукты
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Router /data/company/products [get]
func (h *ProductsHTTP) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	allProducts, err := h.service.GetAllProducts()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allProducts)
}

// @Summary Получить всех партнеров
// @Description Возвращает всех партнеров
// @Tags partners
// @Produce json
// @Success 200 {array} model.Partners
// @Router /data/company/partners [get]
func (h *ProductsHTTP) GetAllPartners(w http.ResponseWriter, r *http.Request) {
	allPartners, err := h.service.GetAllPartners()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allPartners)
}

// @Summary Получить все продукты партнера
// @Description Возвращает все продукты партнера
// @Tags products
// @Produce json
// @Success 200 {array} model.ProductsPartners
// @Router /data/company/product_partners [get]
func (h *ProductsHTTP) GetAllProductsPartner(w http.ResponseWriter, r *http.Request) {
	allProdsPartner, err := h.service.GetAllProductsPartner()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allProdsPartner)
}
