package excelHandler_test

import (
	"ProjectFlour/internal/service/excelHandler"
	"bytes"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestTemplateMakerCreateTemplateTypeProduct(t *testing.T) {
	maker := excelHandler.NewTemplateMaker()

	data, err := maker.CreateTemplateTypeProduct()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer f.Close()

	val, err := f.GetCellValue("Producttype_import", "A1")
	if val == "" {
		t.Errorf("[ERROR]Val is empty. Test has been failed: %v", err)
	}
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if val != "Тип продукции" {
		t.Errorf("Expected 'Тип продукции', got %s", val)
	}
}
