package excelHand

import (
	"ProjectFlour/internal/handlers/httpHandler/httplib"
	"ProjectFlour/internal/service"
	"fmt"
	"log/slog"
	"net/http"
)

type ExcelTemplateHandler struct {
	service service.TemplateMakerForExcel
	logg    *slog.Logger
}

func NewExcelTemplateHandler(srvc service.TemplateMakerForExcel, logg *slog.Logger) *ExcelTemplateHandler {
	return &ExcelTemplateHandler{
		service: srvc,
		logg:    logg,
	}
}

const extensionExcel = ".xlsx"

// @Summary TemplateTypeProduct maker
// @Description Method for create template Excel file for import types products
// @Tags excel
// @Produce file
// @Success 200 {file} file
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/template/types_of_products [get]
func (h *ExcelTemplateHandler) TemplateTypeProduct(w http.ResponseWriter, r *http.Request) {
	const fileName = "Product_type_import"

	fileData, err := h.service.CreateTemplateTypeProduct()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("ContentType", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName+extensionExcel))
	httplib.NewStatusResponse(w, h.logg, http.StatusOK, "OK")

	_, err = w.Write(fileData)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary TemplateTypeMaterial maker
// @Description Method for create template Excel file for import types materials
// @Tags excel
// @Produce file
// @Success 200 {file} file
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/template/types_of_materials [get]
func (h *ExcelTemplateHandler) TemplateTypesMaterial(w http.ResponseWriter, r *http.Request) {
	const fileName = "Material_type_import"

	fileData, err := h.service.CreateTemplateTypeMaterials()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName+extensionExcel))

	_, err = w.Write(fileData)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary TemplatePartners maker
// @Description Method for create template Excel file for import partners
// @Tags excel
// @Produce file
// @Success 200 {file} file
// @Failure 500 {object} map[string]string "Internal server error"
// @Route /excel/template/partners [get]
func (h *ExcelTemplateHandler) TemplatePartners(w http.ResponseWriter, r *http.Request) {
	const fileName = "Partners_import"

	fileData, err := h.service.CreateTemplatePartners()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName+extensionExcel))

	_, err = w.Write(fileData)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary TemplateProducts maker
// @Description Method for create template Excel file for import products
// @Tags excel
// @Produce file
// @Success 200 {file} file
// @Failure 500 {object} map[string]string "Internal server error"
// @Route /excel/template/products [get]
func (h *ExcelTemplateHandler) TemplateProducts(w http.ResponseWriter, r *http.Request) {
	const fileName = "Products_import"

	fileData, err := h.service.CreateTemplateProducts()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName+extensionExcel))

	_, err = w.Write(fileData)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary TemplateProductsPartners maker
// @Description Method for create template Excel file for import partner's productions
// @Tags excel
// @Produce file
// @Success 200 {file} file
// @Failure 500 {object} map[string]string "Internal server error"
// @Route /excel/template/product_partners [get]
func (h *ExcelTemplateHandler) TemplateProductsPartners(w http.ResponseWriter, r *http.Request) {
	const fileName = "Partner_products_import"

	fileData, err := h.service.CreateTemplateProductsPartners()
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName+extensionExcel))

	_, err = w.Write(fileData)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
}
