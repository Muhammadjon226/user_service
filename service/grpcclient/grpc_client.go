package grpcclient

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/Muhammadjon226/user_service/config"
	pb "github.com/Muhammadjon226/user_service/genproto/task_service"
)

//IServiceManager ...
type IServiceManager interface {
	TodoService() pb.ToDoServiceClient
}

type serviceManager struct {
	cfg         config.Config
	todoService pb.ToDoServiceClient
}

//New ...
func New(cfg config.Config) (IServiceManager, error) {
	connTodo, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.TodoServiceHost, cfg.TodoServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("todo service dial host: %s port: %d",
			cfg.TodoServiceHost, cfg.TodoServicePort)
	}
	serviceManager := &serviceManager{
		cfg:         cfg,
		todoService: pb.NewToDoServiceClient(connTodo),
	}

	return serviceManager, nil
}

// TodoService ...
func (s *serviceManager) TodoService() pb.ToDoServiceClient {
	return s.todoService
}
