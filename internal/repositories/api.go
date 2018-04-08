package repositories

import (
	"context"

	"go.zenithar.org/todo/internal/models"
)

// ProjectSearchFilter represents project available search criterias
type ProjectSearchFilter struct {
	Label string
}

//go:generate mockgen -source api.go -destination=./mocks/mock_project.go -package=mocks ProjectRepository

// ProjectRepository defines project repository contract for CRUD operations
type ProjectRepository interface {
	Search(context.Context, *ProjectSearchFilter)
	Create(context.Context, models.Project) (*models.Project, error)
	Read(context.Context, string) (*models.Project, error)
	Update(context.Context, models.Project) error
	Delete(context.Context, string) error
}
