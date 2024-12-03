package main

import (
	"context"
	"fmt"
	"log"
	"order" // Import the generated order package

	"google.golang.org/grpc"
)

func main() {
	// Connect to the server at localhost:50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// Create a new OrderService client
	client := order.NewOrderServiceClient(conn)

	// Prepare an order request
	orderRequest := &order.OrderRequest{
		CustomerName: "John Doe",
		Product:      "Laptop",
		Quantity:     1,
	}

	// Send the order request to the server
	resp, err := client.PlaceOrder(context.Background(), orderRequest)
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}

	// Print the response
	fmt.Printf("Response: %s - %s\n", resp.GetStatus(), resp.GetMessage())
}
