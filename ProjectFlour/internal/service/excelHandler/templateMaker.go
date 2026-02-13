package excelHandler

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
)

// TODO: for headers and their width should make struct instead of double slices
// TODO: ROW with data example for user for each function

const (
	defaultSheet = "Sheet1"
)

type TemplateMaker struct {
}

func NewTemplateMaker() *TemplateMaker {
	return &TemplateMaker{}
}

func (t *TemplateMaker) CreateTemplateTypeProduct() ([]byte, error) {
	const op = "internal.service.excelHandler.templateMaker.go"
	const fileName = "Product_type_import"
	var (
		headers         = []string{"Тип продукции", "Коэффициент типа продукции"}
		widthForHeaders = []float64{15, 30}
	)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.SetSheetName(defaultSheet, fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+rune(i)))
		if err := f.SetCellValue(fileName, cell, header); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := f.SetCellStyle(fileName, "A1", fmt.Sprintf("%s1", string('A'+rune(len(headers)-1))), style); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i := 0; i < len(headers); i++ {
		col := string('A' + rune(i))
		if err := f.SetColWidth(fileName, col, col, widthForHeaders[i]); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	buf := new(bytes.Buffer)
	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}

func (t *TemplateMaker) CreateTemplateTypeMaterials() ([]byte, error) {
	const op = "internal.service.excelHandler.templateMaker.CreateTemplateTypeMaterials"
	const fileName = "Material_type_import"
	var (
		headers        = []string{"Тип материала", "Процент брака материала"}
		widthForHeader = []float64{15, 30}
	)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.SetSheetName(defaultSheet, fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+rune(i)))
		if err := f.SetCellValue(fileName, cell, header); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := f.SetCellStyle(fileName, "A1", fmt.Sprintf("%s1", string('A'+rune(len(headers)-1))), style); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i := 0; i < len(headers); i++ {
		col := string('A' + rune(i))
		if err := f.SetColWidth(fileName, col, col, widthForHeader[i]); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	buf := new(bytes.Buffer)
	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}

func (t *TemplateMaker) CreateTemplatePartners() ([]byte, error) {
	const op = "internal.service.excelHandler.templateMaker.CreateTemplatePartners"
	const fileName = "Partners_import"
	var (
		headers = []string{"Тип партнера", "Наименование партнера", "Диретор", "Электронная почта партнера",
			"Телефон партнера", "Юридический адрес партнера", "ИНН", "Рейтинг"}
		widthForHeaders = []float64{15, 25, 10, 30, 25, 30, 5, 10}
	)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.SetSheetName(defaultSheet, fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+rune(i)))
		if err := f.SetCellValue(fileName, cell, header); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := f.SetCellStyle(fileName, "A1", fmt.Sprintf("%s1", string('A'+rune(len(headers)-1))), style); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i := 0; i < len(headers); i++ {
		col := string('A' + rune(i))
		if err := f.SetColWidth(fileName, col, col, widthForHeaders[i]); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

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
	var (
		headers = []string{"Тип продукции", "Наименование продукции",
			"Артикул", "Минимальная стоимость для партнера"}
		widthForHeaders = []float64{15, 25, 10, 40}
	)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err := f.SetSheetName(defaultSheet, fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+rune(i)))
		if err := f.SetCellValue(fileName, cell, header); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := f.SetCellStyle(fileName, "A1", fmt.Sprintf("%s1", string('A'+rune(len(headers)-1))), style); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i := 0; i < len(headers); i++ {
		col := string('A' + rune(i))
		if err := f.SetColWidth(fileName, col, col, widthForHeaders[i]); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

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
	var (
		headers = []string{"Продукция", "Наименование партнера",
			"Количество продукции", "Дата продажи"}
		widthForHeaders = []float64{15, 25, 25, 15}
	)

	f := excelize.NewFile()
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	err := f.SetSheetName(defaultSheet, fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+rune(i)))
		if err := f.SetCellValue(fileName, cell, header); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := f.SetCellStyle(fileName, "A1", fmt.Sprintf("%s1", string('A'+rune(len(headers)-1))), style); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i := 0; i < len(headers); i++ {
		col := string('A' + rune(i))
		if err := f.SetColWidth(fileName, col, col, widthForHeaders[i]); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	buf := new(bytes.Buffer)

	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return buf.Bytes(), nil
}
