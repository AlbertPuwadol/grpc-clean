package main_test

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/AlbertPuwadol/grpc-clean/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"gotest.tools/assert"
)

func Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(metadata.AppendToOutgoingContext(ctx, "authorization", ""), method, req, reply, cc, opts...) // Put JWT Token here
	}
}

func TestHello(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(Unary()))
	if err != nil {
		t.Log(err)
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()
	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.Empty{}
	expected := &pb.SuccessResponse{Status: "OK"}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		t.Fatalf("Hello FAILED: %v\n", err)
	}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTasks(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(Unary()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()
	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TasksRequest{Text: "this is a test tesk", Tasks: []string{"task1", "task2", "task3"}}
	expected := &pb.TasksResponse{Task1: &pb.Task1Response{ProcessedText: request.Text, Entities: []*pb.Entity{{Start: 1, End: 2, Entity: "Hospital", Label: "PLACE"}, {Start: 3, End: 10, Entity: "Doctor", Label: "PERSON"}}}, Task2: "POSITIVE", Task3: 100}

	res, err := client.Tasks(context.Background(), request)
	if err != nil {
		t.Fatalf("Tasks FAILED: %v\n", err)
	}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTask1(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(Unary()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()
	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TaskRequest{Text: "this is a test tesk"}
	expected := &pb.Task1Response{ProcessedText: request.Text, Entities: []*pb.Entity{{Start: 1, End: 2, Entity: "Hospital", Label: "PLACE"}, {Start: 3, End: 10, Entity: "Doctor", Label: "PERSON"}}}

	res, err := client.Task1(context.Background(), request)
	if err != nil {
		t.Fatalf("Task1 FAILED: %v\n", err)
	}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTask2(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(Unary()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()
	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TaskRequest{Text: "this is a test tesk"}
	expected := &pb.Task2Response{Task2: "POSITIVE"}

	res, err := client.Task2(context.Background(), request)
	if err != nil {
		t.Fatalf("Task2 FAILED: %v\n", err)
	}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}

func TestTask3(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(Unary()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()
	client := pb.NewGRPCCleanServiceClient(conn)

	request := &pb.TaskRequest{Text: "this is a test tesk"}
	expected := &pb.Task3Response{Task3: 100}

	res, err := client.Task3(context.Background(), request)
	if err != nil {
		t.Fatalf("Task3 FAILED: %v\n", err)
	}

	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", res))
}
