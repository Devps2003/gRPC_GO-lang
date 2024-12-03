package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"order"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type OrderHandler struct {
	client order.OrderServiceClient
}

func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	// Read order details from HTTP request body
	var req order.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	// Call gRPC server to place the order
	resp, err := h.client.PlaceOrder(context.Background(), &req)
	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}

	// Send response back to HTTP client
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderServiceClient(conn)

	// Set up HTTP server with routes
	r := mux.NewRouter()
	handler := &OrderHandler{client: client}
	r.HandleFunc("/order", handler.PlaceOrder).Methods("POST")

	// Start HTTP server
	log.Println("HTTP server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
