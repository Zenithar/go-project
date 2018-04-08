package todo

import (
	"go.zenithar.org/todo/internal/models"
)

// FromProject returns a DTO from Entity
func FromProject(entity *models.Project) *Domain_Project {
	return &Domain_Project{
		Id:          entity.ID,
		Label:       entity.Label,
		Description: entity.Description,
		CreatedOn:   &entity.CreatedOn,
	}
}
