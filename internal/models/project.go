package models

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"go.zenithar.org/todo/internal/helpers"
)

// Project represents project entity information holder
type Project struct {
	ID          string    `codec:"id" bson:"id" db:"id"`
	Label       string    `codec:"label" bson:"label" db:"label"`
	Description string    `codec:"description" bson:"description" db:"description"`
	CreatedOn   time.Time `codec:"created_on" bson:"created_on" db:"created_on"`
}

// NewProject returns a project entity initialized
func NewProject(label string) *Project {
	return &Project{
		ID:        helpers.IDGeneratorFunc(),
		Label:     label,
		CreatedOn: helpers.TimeFunc(),
	}
}

// Validate entity constraints
func (p *Project) Validate() error {
	return validation.ValidateStruct(p,
		// ID is mandatory, must be 32 characters long and alphanumeric only.
		validation.Field(&p.ID, validation.Required, validation.Length(32, 32), is.Alphanumeric),
		// Label is mandatory, must be between 3 and 50 characters long and printable ascii only
		validation.Field(&p.Label, validation.Required, validation.Length(3, 50), is.PrintableASCII),
	)
}
