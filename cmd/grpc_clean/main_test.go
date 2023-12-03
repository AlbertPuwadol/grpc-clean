package main_test

import (
	"context"
	"testing"

	pb "github.com/AlbertPuwadol/grpc-clean/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestHello(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.Empty{}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		t.Fatalf("HELLO FAILED: %v", err)
	}

	if res.Response != "OK" {
		t.Errorf("HELLO returned incorrect status, expected \"OK\" got %s", res.Response)
	}
}
