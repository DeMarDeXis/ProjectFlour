package httpHandler

import (
	"ProjectFlour/internal/handlers/httpHandler/dataHandler"
	"ProjectFlour/internal/handlers/httpHandler/excelHand"
	"ProjectFlour/internal/handlers/httpHandler/mw/logger"
	"ProjectFlour/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

type ExcelHandlerInterface interface {
	ImportTypeProduct(w http.ResponseWriter, r *http.Request)
	ImportTypeMaterial(w http.ResponseWriter, r *http.Request)
	ImportProducts(w http.ResponseWriter, r *http.Request)
	ImportPartners(w http.ResponseWriter, r *http.Request)
	ImportProductPartners(w http.ResponseWriter, r *http.Request)
}

type TemplateExcelInterface interface {
	TemplateTypeProduct(w http.ResponseWriter, r *http.Request)
	TemplateTypesMaterial(w http.ResponseWriter, r *http.Request)
	TemplatePartners(w http.ResponseWriter, r *http.Request)
	TemplateProducts(w http.ResponseWriter, r *http.Request)
	TemplateProductsPartners(w http.ResponseWriter, r *http.Request)
}

type ProductsHTTPMethods interface {
	GetAllTypesMaterial(w http.ResponseWriter, r *http.Request)
	GetAllTypeProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetAllPartners(w http.ResponseWriter, r *http.Request)
	GetAllProductsPartner(w http.ResponseWriter, r *http.Request)
}

type HTTPHandler struct {
	ExcelHandlerInterface
	TemplateExcelInterface
	ProductsHTTPMethods
	service *service.Service
	logg    *slog.Logger
}

func NewHTTTPHandler(service *service.Service, logg *slog.Logger) *HTTPHandler {
	return &HTTPHandler{
		ExcelHandlerInterface:  excelHand.NewExcelHTTPHandler(service.ExcelImportService, logg),
		TemplateExcelInterface: excelHand.NewExcelTemplateHandler(service.TemplateMakerForExcel, logg),
		ProductsHTTPMethods:    dataHandler.NewProductsHTTPHandler(service.ProductsService, logg),
		service:                service,
		logg:                   logg,
	}
}

func (h *HTTPHandler) InitRoutes(logg *slog.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(logger.New(logg))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.RealIP)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Auth - It's done
	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	// Excel - It's done
	router.Route("/excel", func(rex chi.Router) {
		rex.Use(h.userIdentity)
		rex.Route("/import", func(reximp chi.Router) {
			reximp.Post("/types_of_materials", h.ExcelHandlerInterface.ImportTypeMaterial)
			reximp.Post("/types_of_products", h.ExcelHandlerInterface.ImportTypeProduct)
			reximp.Post("/products", h.ExcelHandlerInterface.ImportProducts)
			reximp.Post("/partners", h.ExcelHandlerInterface.ImportPartners)
			reximp.Post("/product_partners", h.ExcelHandlerInterface.ImportProductPartners)
		})
		rex.Route("/template", func(rextemp chi.Router) {
			rextemp.Get("/types_of_products", h.TemplateExcelInterface.TemplateTypeProduct)
			rextemp.Get("/types_of_materials", h.TemplateExcelInterface.TemplateTypesMaterial)
			rextemp.Get("/products", h.TemplateExcelInterface.TemplateProducts)
			rextemp.Get("/partners", h.TemplateExcelInterface.TemplatePartners)
			rextemp.Get("/product_partners", h.TemplateExcelInterface.TemplateProductsPartners)
		})
	})

	// Data getters - It's done
	router.Route("/data/company", func(routData chi.Router) {
		routData.Use(h.userIdentity)
		routData.Get("/types_of_materials", h.ProductsHTTPMethods.GetAllTypesMaterial)
		routData.Get("/types_of_products", h.ProductsHTTPMethods.GetAllTypeProduct)
		routData.Get("/products", h.ProductsHTTPMethods.GetAllProducts)
		routData.Get("/partners", h.ProductsHTTPMethods.GetAllPartners)
		routData.Get("/product_partners", h.ProductsHTTPMethods.GetAllProductsPartner)
	})

	return router
}
