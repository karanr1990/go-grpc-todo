package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/karanr1990/go-grpc-todo/protos/todo"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":50051"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.TodoInputData) (*pb.TodoResponseData, error) {
	log.Printf("Received: %v", in.GetName())

	todo := &pb.TodoResponseData{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}
	return todo, nil
}

func main() {

	fmt.Println("welcome to grpc server")
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
