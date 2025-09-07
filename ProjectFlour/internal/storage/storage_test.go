package storage

import (
	"github.com/jmoiron/sqlx"
	"testing"

	_ "github.com/lib/pq"
)

func TestNewStorage(t *testing.T) {
	db := &sqlx.DB{}
	strg := NewStorage(db)

	if strg.AuthorizationStorage == nil {
		t.Errorf("AuthorizationStorage is nil")
	}

	if strg.ProductStorage == nil {
		t.Errorf("ProductStorage is nil")
	}

	if strg.ExcelImportStorage == nil {
		t.Errorf("ExcelImportStorage is nil")
	}
}
