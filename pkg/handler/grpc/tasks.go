package grpc

import (
	"context"

	"github.com/AlbertPuwadol/grpc-clean/pkg/usecase"
	pb "github.com/AlbertPuwadol/grpc-clean/proto"
	"google.golang.org/grpc"
)

type IHanlder interface {
	Hello(context.Context, *pb.Empty) (*pb.SuccessResponse, error)
	Tasks(context.Context, *pb.TasksRequest) (*pb.TasksResponse, error)
	Task1(context.Context, *pb.TaskRequest) (*pb.Task1Response, error)
	Task2(context.Context, *pb.TaskRequest) (*pb.Task2Response, error)
	Task3(context.Context, *pb.TaskRequest) (*pb.Task3Response, error)
}

type handler struct {
	usecase usecase.IUsecase
	pb.UnimplementedGRPCCleanServiceServer
}

func NewHandler(usecase usecase.IUsecase, grpcServer *grpc.Server) {
	handlerGrpc := &handler{usecase: usecase}
	pb.RegisterGRPCCleanServiceServer(grpcServer, handlerGrpc)
}

func (h handler) Hello(ctx context.Context, req *pb.Empty) (*pb.SuccessResponse, error) {
	return &pb.SuccessResponse{Response: h.usecase.Hello()}, nil
}

func (h handler) Tasks(ctx context.Context, req *pb.TasksRequest) (*pb.TasksResponse, error) {
	text := req.GetText()
	tasks := req.GetTasks()

	var response *pb.TasksResponse

	for _, task := range tasks {
		if task == "task1" {
			task1result := h.usecase.Task1(text)
			response.Task1.ProcessedText = text
			for _, entity := range task1result.Entities {
				protoentity := &pb.Entity{Start: int32(entity.Start), End: int32(entity.End), Label: entity.Label, Entity: entity.Entity}
				response.Task1.Entities = append(response.Task1.Entities, protoentity)
			}
		} else if task == "task2" {
			response.Task2 = h.usecase.Task2(text)
		} else if task == "task3" {
			response.Task3 = int32(h.usecase.Task3(text))
		}
	}
	return response, nil
}

func (h handler) Task1(ctx context.Context, req *pb.TaskRequest) (*pb.Task1Response, error) {
	text := req.GetText()

	var response *pb.Task1Response

	response.ProcessedText = text
	task1result := h.usecase.Task1(text)
	for _, entity := range task1result.Entities {
		protoentity := &pb.Entity{Start: int32(entity.Start), End: int32(entity.End), Label: entity.Label, Entity: entity.Entity}
		response.Entities = append(response.Entities, protoentity)
	}

	return response, nil
}

func (h handler) Task2(ctx context.Context, req *pb.TaskRequest) (*pb.Task2Response, error) {
	text := req.GetText()

	return &pb.Task2Response{Task2: h.usecase.Task2(text)}, nil
}

func (h handler) Task3(ctx context.Context, req *pb.TaskRequest) (*pb.Task3Response, error) {
	text := req.GetText()

	return &pb.Task3Response{Task3: int32(h.usecase.Task3(text))}, nil
}
