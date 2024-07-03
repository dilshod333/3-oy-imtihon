package api

import (
	"context"
	"log"
	"testing"
	"time"
	"user-service/internal/proto/gen"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' Not openned database", err)
	}
	defer db.Close()

	s := &Server{Db: db}

	mock.ExpectExec(`DELETE FROM users WHERE user_id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &gen.GetUserReq{UserId: 1}

	resp, err := s.DeleteUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "User deleted successfully!", resp.Status)
}

func TestUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' N opened database", err)
	}
	defer db.Close()

	s := &Server{Db: db}

	mock.ExpectQuery(`SELECT user_id, name, age, email, password, created_at FROM users WHERE user_id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "age", "email", "password", "created_at"}).
			AddRow(1, "Dilshod", 30, "Dilshod@example.com", "Dilshodd", "2022-07-01"))

	req := &gen.GetUserReq{UserId: 1}

	resp, err := s.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.UserId)
	assert.Equal(t, "Dilshod", resp.Name)
	assert.Equal(t, int32(30), resp.Age)
	assert.Equal(t, "Dilshod@example.com", resp.Email)
	assert.Equal(t, "Dilshodd", resp.Password)
	assert.Equal(t, "2022-07-01", resp.CreatedAt)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("notooppened database", err)
	}
	defer db.Close()

	s := &Server{Db: db}

	mock.ExpectQuery(`SELECT user_id, name, age, email, password, created_at FROM users WHERE user_id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "age", "email", "password", "created_at"}).
			AddRow(1, "Dilshod", 30, "Dilshod@example.com", "Dilshodd", "2022-07-01"))

	req := &gen.GetUserReq{UserId: 1}

	resp, err := s.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.UserId)
	assert.Equal(t, "Dilshod", resp.Name)
	assert.Equal(t, int32(30), resp.Age)
	assert.Equal(t, "Dilshod@example.com", resp.Email)
	assert.Equal(t, "Dilshodd", resp.Password)
	assert.Equal(t, "2022-07-01", resp.CreatedAt)
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Error creating mock database...", err)
	}
	defer db.Close()

	lg := log.New(log.Writer(), "test: ", log.LstdFlags)

	s := &Server{Db: db, Lg: lg}

	mock.ExpectQuery(`INSERT INTO users \(name, age, email, password, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING user_id, name, age, email, password, created_at`).
		WithArgs("Dilshod", 30, "Dilshod@example.com", "Dilshodd", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "age", "email", "password", "created_at"}).
			AddRow(1, "Dilshod", 30, "Dilshod@example.com", "Dilshodd", "2022-07-01"))

	req := &gen.UserRequest{
		Name:     "Dilshod",
		Age:      30,
		Email:    "Dilshod@example.com",
		Password: "Dilshodd",
	}

	resp, err := s.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.UserId)
	assert.Equal(t, "Dilshod", resp.Name)
	assert.Equal(t, int32(30), resp.Age)
	assert.Equal(t, "Dilshod@example.com", resp.Email)
	assert.Equal(t, "Dilshodd", resp.Password)
	assert.Equal(t, "2022-07-01", resp.CreatedAt)
}

func TestUpdateUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	server := &Server{Db: db}

	mock.ExpectQuery(`UPDATE users SET name=\$1, age=\$2, email=\$3, password=\$4, created_at=\$5 WHERE user_id=\$6 RETURNING user_id, name, age, email, password, created_at`).
		WithArgs("Dilshod Updated", 31, "DilshodUpdated@example.com", "newpassword123", sqlmock.AnyArg(), 1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "age", "email", "password", "created_at"}).
			AddRow(1, "Dilshod Updated", 31, "DilshodUpdated@example.com", "newpassword123", time.Now().Format(time.RFC3339)))

	req := &gen.UpdateReq{
		UserId:   1,
		Name:     "Dilshod Updated",
		Age:      31,
		Email:    "DilshodUpdated@example.com",
		Password: "newpassword123",
	}
	resp, err := server.UpdateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.UserId)
	assert.Equal(t, "Dilshod Updated", resp.Name)
	assert.Equal(t, int32(31), resp.Age)
	assert.Equal(t, "DilshodUpdated@example.com", resp.Email)
	assert.Equal(t, "newpassword123", resp.Password)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("wroogoggggm: %s", err)
	}
}
