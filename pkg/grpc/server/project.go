package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.zenithar.org/todo/internal/services"
	pb "go.zenithar.org/todo/pkg/grpc/internal/pb"
	dto "go.zenithar.org/todo/pkg/protocol/todo"
)

type projectServer struct {
	projects services.Project
}

// NewProjectServer initialize a gRPC server for project services
func NewProjectServer(projects services.Project) pb.ProjectServer {
	return &projectServer{
		projects: projects,
	}
}

// -----------------------------------------------------------------------------

func (s *projectServer) Create(context.Context, *dto.CreateProjectReq) (*dto.SingleProjectRes, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (s *projectServer) Get(ctx context.Context, req *dto.GetProjectReq) (*dto.SingleProjectRes, error) {
	return s.projects.Get(ctx, req), nil
}

func (s *projectServer) Update(context.Context, *dto.CreateProjectReq) (*dto.SingleProjectRes, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (s *projectServer) Delete(context.Context, *dto.GetProjectReq) (*dto.SingleProjectRes, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}
