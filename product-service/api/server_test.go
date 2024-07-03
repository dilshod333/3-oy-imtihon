package api

import (
	"context"
	gen "product-service/internal/proto/gen3"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &Server{Db: db}

	createdAt := time.Now().Format(time.RFC3339)
	expectedQuery := `SELECT * FROM products WHERE product_id =$1`
	mock.ExpectQuery(expectedQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"product_id", "name", "type", "quantity", "description", "price", "created_at"}).
			AddRow(1, "ovqat", "oziq-ovqat", 10, "Dilshodd", 233.35, "2022-06-17"))

	req := &gen.ProductId{ProductId: 1}

	resp, err := s.GetProducts(context.Background(), req)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were no expectations: %s", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.ProductId)
	assert.Equal(t, "ovqat", resp.Name)
	assert.Equal(t, "oziq-ovqat", resp.Type)
	assert.Equal(t, int32(10), resp.Quantity)
	assert.Equal(t, "Dilshodd", resp.Description)
	assert.Equal(t, 233.35, resp.Price)
	assert.Equal(t, createdAt, resp.CreatedAt)
}
