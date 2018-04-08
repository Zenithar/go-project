package models

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"go.zenithar.org/todo/internal/helpers"
)

// Task represents task entity information holder
type Task struct {
	ProjectID   string     `codec:"project_id" bson:"project_id" db:"project_id"`
	ID          string     `codec:"id" bson:"id" db:"id"`
	Label       string     `codec:"label" bson:"label" db:"label"`
	Description string     `codec:"description" bson:"description" db:"description"`
	Done        bool       `codec:"done" bson:"done" db:"done"`
	CreatedOn   time.Time  `codec:"created_on" bson:"created_on" db:"created_on"`
	CompletedOn *time.Time `codec:"completed_on" bson:"completed_on" db:"completed_on"`
}

// NewTask returns a task entity initialized and associated to the given project
func NewTask(projectID, label string) *Task {
	return &Task{
		ProjectID: projectID,
		ID:        helpers.IDGeneratorFunc(),
		Label:     label,
		Done:      false,
		CreatedOn: helpers.TimeFunc(),
	}
}

// Validate entity constraints
func (p *Task) Validate() error {
	return validation.ValidateStruct(p,
		// ProjectID is mandatory, must be 32 characters long and alphanumeric only.
		validation.Field(&p.ProjectID, validation.Required, validation.Length(32, 32), is.Alphanumeric),
		// ID is mandatory, must be 32 characters long and alphanumeric only.
		validation.Field(&p.ID, validation.Required, validation.Length(32, 32), is.Alphanumeric),
		// Label is mandatory, must be between 3 and 50 characters long and printable ascii only
		validation.Field(&p.Label, validation.Required, validation.Length(3, 50), is.PrintableASCII),
	)
}

// Complete the given task
func (p *Task) Complete() {
	now := helpers.TimeFunc()
	p.Done = true
	p.CompletedOn = &now
}
