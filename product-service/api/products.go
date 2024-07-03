package api

import (
	"context"
	"database/sql"
	"log"
	storage "product-service/internal"
	gen "product-service/internal/proto/gen3"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// var db *sql.DB

type Server struct {
	gen.UnimplementedProductServiceServer
	Db *sql.DB
}

func DbConn() *Server {
	db := storage.Connect()
	return &Server{Db: db}
}

func (u *Server) GetProducts(ctx context.Context, req *gen.ProductId) (*gen.ProductResponse, error) {
	query := `select * from products where product_id=$1`
	var product gen.ProductResponse

	err := u.Db.QueryRow(query, req.ProductId).Scan(&product.ProductId, &product.Name, &product.Type, &product.Quantity, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		log.Fatal(">>Error<<<", err)
	}
	return &product, nil
}

func (u *Server) CreateProducts(ctx context.Context, req *gen.ProductRequest) (*gen.ProductResponse, error) {
	var product gen.ProductResponse
	query := `insert into products (name, type, quantity, description, price, created_at) values($1, $2, $3, $4, $5, $6) returning *`
	err := u.Db.QueryRow(query, req.Name, req.Type, req.Quantity, req.Description, req.Price, time.Now().Format(time.RFC3339)).Scan(&product.ProductId, &product.Name, &product.Type, &product.Quantity, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		log.Fatal("Error maybe not found that ID...", err)

	}
	return &gen.ProductResponse{
		ProductId:   product.ProductId,
		Name:        product.Name,
		Type:        product.Type,
		Description: product.Description,
		Quantity:    product.Quantity,
		Price:       product.Price,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}, nil

}

func (u *Server) UpdateProduct(ctx context.Context, req *gen.UpdateProductReq) (*gen.ProductResponse, error) {
	var product gen.ProductResponse
	query := `UPDATE products SET name=$1, type=$2, quantity=$3, description=$4, price=$5, created_at=$6 WHERE product_id=$7 RETURNING *`
	err := u.Db.QueryRow(query, req.Name, req.Type, req.Quantity, req.Description, req.Price, time.Now().Format(time.RFC3339), req.ProductId).Scan(&product.ProductId, &product.Name, &product.Type, &product.Quantity, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return &product, nil

}

func (u *Server) ListAllProduct(req *gen.EmptyList, stream gen.ProductService_ListAllProductServer) error {
	query := `SELECT product_id, name, type, quantity, description, price, created_at FROM products`
	rows, err := u.Db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var product gen.ProductResponse
		if err := rows.Scan(&product.ProductId, &product.Name, &product.Type, &product.Quantity, &product.Description, &product.Price, &product.CreatedAt); err != nil {
			return err
		}
		if err := stream.Send(&product); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (u *Server) DeleteProduct(ctx context.Context, req *gen.ProductId) (*gen.DeleteProductResponse, error) {

	query := `DELETE FROM products WHERE product_id = $1`

	result, err := u.Db.ExecContext(ctx, query, req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, " not delete product: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "wronnngg: %v", err)
	}

	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "product with ID %d not found", req.ProductId)
	}

	return &gen.DeleteProductResponse{	
		Status: "User deleted successfully!",
	}, nil

}




