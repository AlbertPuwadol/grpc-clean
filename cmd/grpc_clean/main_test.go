package main_test

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/AlbertPuwadol/grpc-clean/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gotest.tools/assert"
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

	if res.Status != "OK" {
		t.Errorf("HELLO returned incorrect status, expected \"OK\" got %s", res.Status)
	}
}

func TestTasks(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TasksRequest{Text: "this is a test tesk", Tasks: []string{"task1", "task2", "task3"}}

	res, err := client.Tasks(context.Background(), request)
	if err != nil {
		t.Fatalf("Tasks FAILED: %v", err)
	}

	expected := &pb.TasksResponse{Task1: &pb.Task1Response{ProcessedText: request.Text, Entities: []*pb.Entity{{Start: 1, End: 2, Entity: "Hospital", Label: "PLACE"}, {Start: 3, End: 10, Entity: "Doctor", Label: "PERSON"}}}, Task2: "POSITIVE", Task3: 100}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTask1(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TaskRequest{Text: "this is a test tesk"}

	res, err := client.Task1(context.Background(), request)
	if err != nil {
		t.Fatalf("Tasks FAILED: %v", err)
	}

	expected := &pb.Task1Response{ProcessedText: request.Text, Entities: []*pb.Entity{{Start: 1, End: 2, Entity: "Hospital", Label: "PLACE"}, {Start: 3, End: 10, Entity: "Doctor", Label: "PERSON"}}}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTask2(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TaskRequest{Text: "this is a test tesk"}

	res, err := client.Task2(context.Background(), request)
	if err != nil {
		t.Fatalf("Tasks FAILED: %v", err)
	}

	expected := &pb.Task2Response{Task2: "POSITIVE"}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTask3(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TaskRequest{Text: "this is a test tesk"}

	res, err := client.Task3(context.Background(), request)
	if err != nil {
		t.Fatalf("Tasks FAILED: %v", err)
	}

	expected := &pb.Task3Response{Task3: 100}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}
