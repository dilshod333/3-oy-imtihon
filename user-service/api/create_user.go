package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"user-service/config"
	"user-service/internal/proto/gen"
	"user-service/pkg"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	User_id  int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Server struct {
	gen.UnimplementedUserServiceServer
	Db *sql.DB
	Lg *log.Logger
}

func NewServer(lg *log.Logger) (*Server, error) {
	cfg := config.Load(".env")
	db, err := pkg.ConnectToDB(cfg)
	if err != nil {
		return nil, err
	}
	return &Server{Db: db, Lg: lg}, nil
}
func (u *Server) CreateUser(ctx context.Context, req *gen.UserRequest) (*gen.UserResponse, error) {
	var user gen.UserResponse

	u.Lg.Printf("Received CreateUser request: %+v\n", req)

	query := `INSERT INTO users (name, age, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING user_id, name, age, email, password, created_at`

	err := u.Db.QueryRow(query, req.Name, req.Age, req.Email, req.Password, time.Now().Format(time.RFC3339)).Scan(&user.UserId, &user.Name, &user.Age, &user.Email, &user.Password, &user.CreatedAt)
	fmt.Println(">>>>>>>>>>>", err)
	if err != nil {

		u.Lg.Printf("Error executing query: %v\n", err)
		return nil, status.Errorf(codes.Internal, "could not create user: %v", err)
	}

	u.Lg.Printf("User created successfully: %+v\n", &user)

	return &user, nil
}

func (s *Server) GetUser(ctx context.Context, req *gen.GetUserReq) (*gen.UserResponse, error) {
	query := `SELECT user_id, name, age, email, password, created_at FROM users WHERE user_id = $1`
	var user gen.UserResponse

	err := s.Db.QueryRow(query, req.UserId).Scan(&user.UserId, &user.Name, &user.Age, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Fatal("Error on queryrow...", err)
	}

	return &gen.UserResponse{
		UserId:    user.UserId,
		Name:      user.Name,
		Age:       user.Age,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}, nil
}

// func (d *Server) UpdateUser(ctx context.Context, req *gen.UpdateReq) (*gen.UserResponse, error) {
// 	query := `UPDATE users SET name=$1, age=$2, email=$3, password=$4, created_at=$5 WHERE user_id=$6 RETURNING user_id, name, age, email, password, created_at`
// 	var user gen.UserResponse

// 	err := d.Db.QueryRowContext(ctx, query, req.Name, req.Age, req.Email, req.Password, time.Now().Format(time.RFC3339), req.UserId).Scan(&user.UserId, &user.Name, &user.Age, &user.Email, &user.Password, &user.CreatedAt)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
// 		}
// 		return nil, status.Errorf(codes.Internal, "not update user: %v", err)
// 	}

// 	return &user, nil
// }

func (d *Server) UpdateUser(ctx context.Context, req *gen.UpdateReq) (*gen.UserResponse, error) {
    d.Lg.Printf("Received UpdaaateUser request: %+v\n", req)

    query := `UPDATE users SET name=$1, age=$2, email=$3, password=$4, created_at=$5 WHERE user_id=$6 RETURNING user_id, name, age, email, password, created_at`
    var user gen.UserResponse

    err := d.Db.QueryRow(query, req.Name, req.Age, req.Email, req.Password, time.Now().Format(time.RFC3339), req.UserId).Scan(
        &user.UserId, &user.Name, &user.Age, &user.Email, &user.Password, &user.CreatedAt,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            d.Lg.Printf("No user found with ID: %d\n", req.UserId)
            return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
        }
        d.Lg.Printf("Error executing query: %v\n", err)
        return nil, status.Errorf(codes.Internal, "could not update user: %v", err)
    }

    d.Lg.Printf("User updated successfully: %+v\n", &user)

    return &user, nil
}


func (u *Server) DeleteUser(ctx context.Context, req *gen.GetUserReq) (*gen.DeleteResponse, error) {

	query := `DELETE FROM users WHERE user_id = $1`

	result, err := u.Db.ExecContext(ctx, query, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "wronnngg: %v", err)
	}

	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user with ID %d not found", req.UserId)
	}

	return &gen.DeleteResponse{
		Status: "User deleted successfully!",
	}, nil
}
