package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"order" // This is the package generated from order.proto

	"google.golang.org/grpc"
)

// OrderServer implements the OrderServiceServer interface
type OrderServer struct {
	order.UnimplementedOrderServiceServer
}

// PlaceOrder handles placing an order
func (s *OrderServer) PlaceOrder(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	// Simulate order processing
	log.Printf("Order received: Customer: %s, Product: %s, Quantity: %d", req.GetCustomerName(), req.GetProduct(), req.GetQuantity())

	// Here, we could check stock, payment status, etc.
	// For simplicity, assume the order is always successfully placed.
	return &order.OrderResponse{
		Status:  "Success",
		Message: fmt.Sprintf("Order for %d %s(s) placed successfully!", req.GetQuantity(), req.GetProduct()),
	}, nil
}

func main() {
	// Set up the server to listen on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register our service with the gRPC server
	order.RegisterOrderServiceServer(grpcServer, &OrderServer{})

	// Start the server
	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
