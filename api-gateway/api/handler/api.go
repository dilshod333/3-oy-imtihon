package handler

import (
	user1 "api-gateway/methods/user"
	"io"

	"api-gateway/proto/gen"
	"api-gateway/proto/gen1"
	"api-gateway/proto/gen3"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"
	// "google.golang.org/protobuf/internal/order"
)

type Server struct {
	userr gen.UserServiceClient
	pro   gen3.ProductServiceClient
	ord   gen1.OrderServiceClient
}

func Conn() *Server {
	user := user1.ConnectUser()
	product := user1.ConnectProduct()
	order := user1.OrderProductt()
	return &Server{userr: user, pro: product, ord: order}
}

func (u *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req gen.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatalf("Adduser request error %v", err)
	}

	resp, err := u.userr.CreateUser(context.Background(), &req)
	if err != nil {
		log.Fatalf("Sending request to the user service to add user %v", err)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal(err)
	}

}

func (u *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	var req = gen.GetUserReq{UserId: int32(id)}
	resp, err := u.userr.GetUser(context.Background(), &req)
	if err != nil {
		log.Fatalf("getting with id error... %v", err)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatalf("Errror on encoding the result...%v", err)
	}

}

func (u *Server) UpdateUserr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal("Error on getting id on updateuser...", err)
	}
	var req = gen.UpdateReq{UserId: int32(id)}
	if err:=json.NewDecoder(r.Body).Decode(&req);err!=nil{
		log.Fatal(err)
	}
	resp, err := u.userr.UpdateUser(context.Background(), &req)
	if err != nil {
		log.Fatal("error...", err)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal("error on encoding updateuser....", err)
	}

}

func (u *Server) DeleteUserr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	var req = gen.GetUserReq{UserId: int32(id)}
	resp, err := u.userr.DeleteUser(context.Background(), &req)
	if err != nil {
		log.Fatalf("getting with id error... %v", err)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatalf("Errror on encoding the result...%v", err)
	}
}

// PRODUCT //

func (u *Server) GetProductt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	var reqq = gen3.ProductId{ProductId: int32(id)}
	resp, err := u.pro.GetProducts(context.Background(), &reqq)
	if err != nil {
		log.Fatalf("getting with id error... %v", err)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatalf("Errror on encoding the result...%v", err)
	}
}

func (u *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req gen3.ProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatalf("Adduser request error %v", err)
	}

	resp, err := u.pro.CreateProducts(context.Background(), &req)
	if err != nil {
		log.Fatalf("Sending request to the user service to add user %v", err)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal(err)
	}

}

func (u *Server) UpdateProductt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req gen3.UpdateProductReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatal("Errorr while decoddingg...", err)
	}

	resp, err := u.pro.UpdateProduct(context.Background(), &req)

	if err != nil {
		log.Fatal("Erroor just...", err)
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal("error on encoding responsee....", err)
	}

}

func (u *Server) DeleteProductt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	var req = gen3.ProductId{ProductId: int32(id)}

	resp, err := u.pro.DeleteProduct(context.Background(), &req)

	if err != nil {
		log.Fatal("Error on...", err)
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal("error on encoding response....", err)
	}

}

// ORDER-SERVICE

func (u *Server) CreateOrderr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order gen1.CreateOrderReq

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Fatalf("Adduser request error %v", err)
	}

	resp, err := u.ord.CreateOrder(context.Background(), &order)
	if err != nil {
		log.Fatalf("Sending request to the user service to add user %v", err)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal(err)
	}
}

func (u *Server) CreateOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var list gen1.ListOrders
	byte, _ := io.ReadAll(r.Body)
	err := protojson.Unmarshal(byte, &list)

	if err != nil {
		http.Error(w,"xatolik unmarshal", http.StatusBadRequest)
		return 
	}


	stream, err := u.ord.CreateOrders(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Send each request in the slice
	for _, req := range list.Listord {
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err !=  nil {
		http.Error(w, "Xatolik cloaseandstreaming", http.StatusInternalServerError)
		return 
	}


	// Encode and send the response back to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal(err)
	}
}
