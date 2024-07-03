package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"order-service/api/models"
	"order-service/api/models/handleproduct"
	"order-service/api/models/handleuser"
	"order-service/proto/gen1"
	"time"
)

type Server struct {
	gen1.UnimplementedOrderServiceServer
	db *sql.DB
	Lg *log.Logger
}

func NewUserServer(DB *sql.DB, lg *log.Logger) *Server {
	return &Server{db: DB, Lg: lg}
}

func (s *Server) CreateOrder(ctx context.Context, req *gen1.CreateOrderReq) (*gen1.CreateOrderResp, error) {
	name, err := handleuser.UserClinet(int(req.UserId))
	product := handleproduct.ProductHandle(int(req.ProductId))

	if err != nil {
		log.Fatal("error on getting name...", err)
	}
	var order gen1.CreateOrderResp
	order.Name = name
	total_price := req.Quantity * int32(product.Price)
	query := `insert into orders(user_id, product_id, name, price, total_price, order_time) values($1, $2, $3, $4, $5, $6) returning *`
	if err := s.db.QueryRow(query, req.UserId, req.ProductId, product.NameProduct, product.Price, total_price, time.Now().Format(time.RFC3339)).Scan(&order.OrderId, &order.UserId, &order.ProductId, &order.Name, &order.Price, &order.TotalPrice, &order.OrderTime); err != nil {
		log.Fatalf("scanning error on the nserting %v", err)
	}
	s.Lg.Println("created orderrrrr")
	return &gen1.CreateOrderResp{
		OrderId:    order.OrderId,
		UserId:     order.UserId,
		ProductId:  req.ProductId,
		Name:       order.Name,
		Price:      order.Price,
		TotalPrice: float32(req.Quantity) * product.Price,
		OrderTime:  order.OrderTime,
	}, nil

}

func (s *Server) CreateOrders(stream gen1.OrderService_CreateOrdersServer) error {
	log.Println("Create orderga kirdiiiiiii")
	s.Lg.Println("CreateOrders method started")
	for {
		req, err := stream.Recv()
		fmt.Println("Nimadirrrrrr")
		if err == io.EOF {
			s.Lg.Println("End of stream")
			return stream.SendAndClose(&gen1.Response{Message: "all data saved"})
		}
		if err != nil {
			s.Lg.Println("error receiving stream data:", err)
			return err
		}

		s.Lg.Printf("Received order: %+v", req)

		product := handleproduct.ProductHandle(int(req.ProductId))
		s.AddtoDatabase(req, product)
		s.Lg.Println("created new orderssss")

	}
}

func (s *Server) AddtoDatabase(req *gen1.CreateOrderReq, product *models.Handle) error {

	orderTime := time.Now().Format(time.RFC3339)
	totalPrice := req.Quantity * int32(product.Price)
	query := `INSERT INTO orders (user_id, product_id, name, price, total_price, order_time) 
                  VALUES ($1, $2, $3, $4, $5, $6)`
	var orderId int64
	_, err := s.db.Exec(query, req.UserId, req.ProductId, product.NameProduct, product.Price, totalPrice, orderTime)
	if err != nil {
		s.Lg.Printf("Error inserting order: %v", err)
		return err
	}

	s.Lg.Printf("Order saved to database with OrderID: %d", orderId)
	return nil
}
