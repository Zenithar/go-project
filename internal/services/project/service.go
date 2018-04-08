package project

import (
	"context"
	"fmt"
	"net/http"

	"go.zenithar.org/todo/internal/repositories"
	"go.zenithar.org/todo/internal/services"
	dto "go.zenithar.org/todo/pkg/protocol/todo"
)

type projectService struct {
	projects repositories.ProjectRepository
}

// NewService returns a project service instance
func NewService(projects repositories.ProjectRepository) services.Project {
	return &projectService{
		projects: projects,
	}
}

// -----------------------------------------------------------------------------

func (s *projectService) Get(ctx context.Context, req *dto.GetProjectReq) *dto.SingleProjectRes {

	res := &dto.SingleProjectRes{}

	// Validate request
	if err := req.Validate(); err != nil {
		res.Error = &dto.Error{
			Code:    http.StatusPreconditionFailed,
			Message: fmt.Sprintf("Validation error: %s", err.Error()),
		}
		return res
	}

	// Do the database call
	entity, err := s.projects.Read(ctx, req.Id)
	if err != nil && err != repositories.ErrNoResult {
		res.Error = &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to query database",
		}
		return res
	}
	if entity == nil || err == repositories.ErrNoResult {
		res.Error = &dto.Error{
			Code:    http.StatusNotFound,
			Message: "Project not found",
		}
		return res
	}

	// Map result
	res.Entity = dto.FromProject(entity)

	// Return result
	return res
}
