package excelHandler

import (
	"ProjectFlour/internal/events"
	"ProjectFlour/internal/model"
	"ProjectFlour/internal/storage"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ExcelHandService struct {
	storage storage.ExcelImportStorage
	event   *events.EventBus
}

func New(storage storage.ExcelImportStorage, eventBus *events.EventBus) *ExcelHandService {
	return &ExcelHandService{
		storage: storage,
		event:   eventBus,
	}
}

func (l *ExcelHandService) AddTypeProductFromExcel(filepath string) error {
	const op = "service.excelHandler.AddTypeProduct"
	const sheet = "Product_type_import"

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("%s: %w", op, err)
	}

	var typeProduct []model.TypeProduct

	for _, row := range rows[1:] {
		coef, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		typeProduct = append(typeProduct, model.TypeProduct{
			NameType:    row[0],
			Coefficient: float32(coef),
		})
	}

	if err := l.storage.AddTypeProductFromExcel(typeProduct); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	l.event.Publish(events.EventFileImported, events.FileImportedEvent{
		Type:    "product_types",
		Message: "Типы продукции успешно обновлены",
		Count:   len(typeProduct),
	})

	return nil
}

func (l *ExcelHandService) AddTypeMaterialsFromExcel(filepath string) error {
	const op = "service.excelHandler.AddTypeMaterialsFromExcel"
	const sheet = "Material_type_import"

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("%s: %w", op, err)
	}

	var allTypeMaterial []model.TypeMaterial

	for _, row := range rows[1:] {
		trimPercent := row[1][:len(row[1])-1]

		percent, err := strconv.ParseFloat(trimPercent, 32)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		allTypeMaterial = append(allTypeMaterial, model.TypeMaterial{
			NameType:          row[0],
			PercentOfMarriage: float32(percent),
		})
	}

	if err := l.storage.AddTypeMaterialFromExcel(allTypeMaterial); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	l.event.Publish(events.EventFileImported, events.FileImportedEvent{
		Type:    "materials_type",
		Message: "Типы материалов успешно добавлены",
		Count:   len(allTypeMaterial),
	})

	return nil
}

func (l *ExcelHandService) AddProductsFromExcel(filepath string) error {
	const op = "service.excelHandler.AddProductsFromExcel"
	const sheet = "Products_import"

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("%s: %w", op, err)
	}

	var allProducts []model.Product

	for _, row := range rows[1:] {
		priceToFloat, err := strconv.ParseFloat(row[3], 32)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		allProducts = append(allProducts, model.Product{
			TypeProduct:        row[0],
			Name:               row[1],
			Article:            row[2],
			MinPriceForPartner: float32(priceToFloat),
		})
	}

	if err := l.storage.AddProductsFromExcel(allProducts); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	l.event.Publish(events.EventFileImported, events.FileImportedEvent{
		Type:    "products",
		Message: "Продукция успешно добавлена",
		Count:   len(allProducts),
	})

	return nil

}

func (l *ExcelHandService) AddPartnersFromExcel(filepath string) error {
	const op = "service.excelHandler.AddPartnersFromExcel"
	const sheet = "Partners_import"

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("%s: %w", op, err)
	}

	var allPartners []model.Partners

	for _, row := range rows[1:] {

		rateToInt, err := strconv.ParseInt(row[7], 10, 32)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		allPartners = append(allPartners, model.Partners{
			TypeOrg:        row[0],
			Name:           row[1],
			FullNameOfBoss: row[2],
			Email:          row[3],
			PhoneNumber:    row[4],
			LegalAddress:   row[5],
			INN:            row[6],
			Rate:           int(rateToInt),
		})
	}

	if err := l.storage.AddPartnersFromExcel(allPartners); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	l.event.Publish(events.EventFileImported, events.FileImportedEvent{
		Type:    "partners",
		Message: "Партнеры успешно добавлены",
		Count:   len(allPartners),
	})

	return nil
}

func (l *ExcelHandService) AddProductPartnersFromExcel(filepath string) error {
	const op = "storage.postgres.excelStorage.AddProductPartnersFromExcel"
	const sheet = "Partner_products_import"

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("%s: %w", op, err)
	}

	var allProductionFromFile []model.ProductsPartners

	for _, row := range rows[1:] {
		quantityToInt, err := strconv.ParseInt(row[2], 10, 32)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		allProductionFromFile = append(allProductionFromFile, model.ProductsPartners{
			Partner:    row[1],
			Product:    row[0],
			Quantity:   int(quantityToInt),
			DateOfSale: row[3],
		})
	}

	if err := l.storage.AddProductPartnersFromExcel(allProductionFromFile); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	l.event.Publish(events.EventFileImported, events.FileImportedEvent{
		Type:    "prod_partners",
		Message: "Продукции партнеров успешно добавлена",
		Count:   len(allProductionFromFile),
	})

	return nil
}
