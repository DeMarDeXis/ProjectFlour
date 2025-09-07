package excelHand

import (
	"ProjectFlour/internal/handlers/httpHandler/httplib"
	"ProjectFlour/internal/service"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ExcelHTTP struct {
	service service.ExcelImportService
	logg    *slog.Logger
}

func NewExcelHTTPHandler(srvs service.ExcelImportService, logg *slog.Logger) *ExcelHTTP {
	return &ExcelHTTP{
		service: srvs,
		logg:    logg,
	}
}

const typeFormFile = "excel_file"

// @Summary Import type product from Excel
// @Description Upload and process Excel file to import type products
// @Tags excel
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file to upload (.xlsx, .xls)"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/import/types_of_products [post]
func (h *ExcelHTTP) ImportTypeProduct(w http.ResponseWriter, r *http.Request) {
	const op = "httpHandler.importTypeProduct"

	err := r.ParseMultipartForm(10 << 22) // 10 MB
	if err != nil {
		h.logg.Error("failed to parse multipart form", slog.Any("error", err.Error()))
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile(typeFormFile)
	if err != nil {
		h.logg.Error("failed to get file from form", slog.Any("error", err.Error()))
		http.Error(w, "failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if filepath.Ext(header.Filename) != ".xlsx" && filepath.Ext(header.Filename) != ".xls" {
		h.logg.Error("invalid file extension", slog.String("filename", header.Filename))
		http.Error(w, "only Excel files (.xlsx, .xls) are allowed", http.StatusBadRequest)
		return
	}

	tempFile, err := h.createTempExcelFile(file, header.Filename)
	if err != nil {
		h.logg.Error("failed to create temp file", slog.Any("error", err.Error()))
		http.Error(w, "failed to create temp file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile)

	err = h.service.AddTypeProductFromExcel(tempFile)
	if err != nil {
		h.logg.Error("failed to add type product from excel", slog.Any("error", err.Error()))
		http.Error(w, "failed to add type product from excel", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Type product added successfully"}`))
	w.Header().Set("Content-Type", "application/json")

}

// @Summary Import type material from Excel
// @Description Upload and process Excel file to import type materials
// @Tags excel
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file to upload (.xlsx, .xls)"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/import/types_of_materials [post]
func (h *ExcelHTTP) ImportTypeMaterial(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 22)
	if err != nil {
		h.logg.Error("failed to parse multipart form", slog.Any("error", err.Error()))
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile(typeFormFile)
	if err != nil {
		h.logg.Error("failed to get file from form", slog.Any("error", err.Error()))
		http.Error(w, "failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if filepath.Ext(header.Filename) != ".xlsx" && filepath.Ext(header.Filename) != ".xls" {
		h.logg.Error("invalid file extension", slog.String("filename", header.Filename))
		http.Error(w, "only Excel files (.xlsx, .xls) are allowed", http.StatusBadRequest)
		return
	}

	tempFile, err := h.createTempExcelFile(file, header.Filename)
	if err != nil {
		h.logg.Error("failed to create temp file", slog.Any("error", err.Error()))
		http.Error(w, "failed to create temp file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile)

	err = h.service.AddTypeMaterialsFromExcel(tempFile)
	if err != nil {
		h.logg.Error("failed to add type material from excel", slog.Any("error", err.Error()))
		http.Error(w, "failed to add type material from excel", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Type material added successfully"}`))
	w.Header().Set("Content-Type", "application/json")
}

// @Summary Import products from Excel
// @Description Upload and process Excel file to import products
// @Tags excel
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file to upload (.xlsx, .xls)"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/import/products [post]
func (h *ExcelHTTP) ImportProducts(w http.ResponseWriter, r *http.Request) {
	const op = "httpHandler.importProducts"

	err := r.ParseMultipartForm(10 << 22)
	if err != nil {
		h.logg.Error("failed to parse multipart form", slog.Any("error", err.Error()))
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile(typeFormFile)
	if err != nil {
		h.logg.Error("failed to get file from form", slog.Any("error", err.Error()))
		http.Error(w, "failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if filepath.Ext(header.Filename) != ".xlsx" && filepath.Ext(header.Filename) != ".xls" {
		h.logg.Error("invalid file extension", slog.String("filename", header.Filename))
		http.Error(w, "only Excel files (.xlsx, .xls) are allowed", http.StatusBadRequest)
		return
	}

	tempFile, err := h.createTempExcelFile(file, header.Filename)
	if err != nil {
		h.logg.Error("failed to create temp file", slog.Any("error", err.Error()))
		http.Error(w, "failed to create temp file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile)

	err = h.service.AddProductsFromExcel(tempFile)
	if err != nil {
		h.logg.Error("failed to add products from excel", slog.Any("error", err.Error()))
		http.Error(w, "failed to add products from excel", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Products added successfully"}`))
	w.Header().Set("Content-Type", "application/json")
}

// @Summary Import partners from Excel
// @Description Upload and process Excel file to import partners
// @Tags excel
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file to upload (.xlsx, .xls)"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/import/partners [post]
func (h *ExcelHTTP) ImportPartners(w http.ResponseWriter, r *http.Request) {
	const op = "httpHandler.ImportPartners"

	err := r.ParseMultipartForm(10 << 22)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	file, header, err := r.FormFile(typeFormFile)
	if err != nil {
		h.logg.Error("failed to get file from form", slog.Any("error", err.Error()))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	if filepath.Ext(header.Filename) != ".xlsx" && filepath.Ext(header.Filename) != ".xls" {
		h.logg.Error("invalid file extension", slog.String("filename", header.Filename))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	tempFile, err := h.createTempExcelFile(file, header.Filename)
	if err != nil {
		h.logg.Error("failed to create temp file", slog.Any("error", err.Error()))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
	defer os.Remove(tempFile)

	err = h.service.AddPartnersFromExcel(tempFile)
	if err != nil {
		h.logg.Error("failed to add partners", slog.Any("error", err.Error()), slog.String("op", op))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Partners added successfully"}`))
	w.Header().Set("Content-Type", "application/json")
}

// @Summary Import product and partner relations from Excel
// @Description Upload and process Excel file to import product and partner relations
// @Tags excel
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file to upload (.xlsx, .xls)"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /excel/import/product_partners [post]
func (h *ExcelHTTP) ImportProductPartners(w http.ResponseWriter, r *http.Request) {
	const op = "httpHandler.excelHand.ImportProductPartners"

	err := r.ParseMultipartForm(10 << 22)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	file, header, err := r.FormFile(typeFormFile)
	if err != nil {
		h.logg.Error("failed to get file from form", slog.Any("error", err.Error()))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	if filepath.Ext(header.Filename) != ".xlsx" && filepath.Ext(header.Filename) != ".xls" {
		h.logg.Error("invalid file extension", slog.String("filename", header.Filename))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	tempFile, err := h.createTempExcelFile(file, header.Filename)
	if err != nil {
		h.logg.Error("failed to create temp file", slog.Any("error", err.Error()))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}
	defer os.Remove(tempFile)

	err = h.service.AddProductPartnersFromExcel(tempFile)
	if err != nil {
		h.logg.Error("failed to add partners", slog.Any("error", err.Error()), slog.String("op", op))
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ProdPartners added successfully"}`))
	w.Header().Set("Content-Type", "application/json")
}

// createTempExcelFile creates a temporary file from an Excel file.
func (h *ExcelHTTP) createTempExcelFile(file io.Reader, originalFilename string) (string, error) {
	const op = "httpHandler.createTempExcelFile"
	const tempDir = "temp"

	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	timeStamp := time.Now().Unix()
	ext := filepath.Ext(originalFilename)
	tempFileName := fmt.Sprintf("excel_import_%d%s", timeStamp, ext)
	tempFilePath := filepath.Join(tempDir, tempFileName)

	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		os.Remove(tempFilePath)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tempFilePath, nil
}
