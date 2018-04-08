package todo

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Validate protobuf request
func (m *GetProjectReq) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Id, validation.Required, validation.Length(32, 32), is.Alphanumeric),
	)
}
