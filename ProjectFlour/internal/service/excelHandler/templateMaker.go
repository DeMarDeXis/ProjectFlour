package excelHandler

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	defaultSheet = "Sheet1"
)

type TemplateMaker struct {
}

func NewTemplateMaker() *TemplateMaker {
	return &TemplateMaker{}
}

// TODO:[example row] set example separeted row in this function
func (t *TemplateMaker) CreateTemplateTypeProduct() ([]byte, error) {
	const op = "internal.service.excelHandler.templateMaker.go"
	const fileName = "Product_type_import"

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.DeleteSheet(defaultSheet)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = f.NewSheet(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	f.SetCellValue(fileName, "A1", "Тип продукции")
	f.SetCellValue(fileName, "B1", "Коэффициент типа продукции")

	// TODO: for what exist method 'f.WriteToBuf'. Explore it
	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}

// TODO: to test it
func (t *TemplateMaker) CreateTemplateTypeMaterials() ([]byte, error) {
	const op = "internal.service.excelHandler.templateMaker.CreateTemplateTypeMaterials"
	const fileName = "Material_type_import"

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.DeleteSheet(defaultSheet)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = f.NewSheet(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	f.SetCellValue(fileName, "A1", "Тип материала")
	f.SetCellValue(fileName, "B1", "Процент брака материала")

	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}

func (t *TemplateMaker) CreateTemplatePartners() ([]byte, error) {
	const op = "internal.service.excelHandler.templateMaker.CreateTemplatePartners"
	const fileName = "Partners_import"

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.DeleteSheet(defaultSheet)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = f.NewSheet(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	f.SetCellValue(fileName, "A1", "Тип партнера")
	f.SetCellValue(fileName, "B1", "Наименование партнера")
	f.SetCellValue(fileName, "C1", "Директор")
	f.SetCellValue(fileName, "D1", "Электронная почта партнера")
	f.SetCellValue(fileName, "E1", "Телефон партнера")
	f.SetCellValue(fileName, "F1", "Юридический адрес партнера")
	f.SetCellValue(fileName, "G1", "ИНН")
	f.SetCellValue(fileName, "H1", "Рейтинг")

	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}

func (t *TemplateMaker) CreateTemplateProducts() ([]byte, error) {
	const (
		op       = "internal.service.excelHandler.templateMaker.go"
		fileName = "Products_import"
	)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.DeleteSheet(defaultSheet)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = f.NewSheet(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	f.SetCellValue(fileName, "A1", "Тип продукции")
	f.SetCellValue(fileName, "B1", "Наименование продукции")
	f.SetCellValue(fileName, "C1", "Артикул")
	f.SetCellValue(fileName, "D1", "Минимальная стоимость для партнера")

	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}

func (t *TemplateMaker) CreateTemplateProductsPartners() ([]byte, error) {
	const (
		op       = "internal.service.excelHandler.templateMaker.go.Partner_products_import"
		fileName = "Partner_products_import"
	)

	f := excelize.NewFile()
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	err := f.DeleteSheet(defaultSheet)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = f.NewSheet(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	f.SetCellValue(fileName, "A1", "Продукция")
	f.SetCellValue(fileName, "B1", "Наименование партнера")
	f.SetCellValue(fileName, "C1", "Количество продукции")
	f.SetCellValue(fileName, "D1", "Дата продажи")

	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}
