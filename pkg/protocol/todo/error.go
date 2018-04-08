package todo

import "fmt"

func (m *Error) Error() string {
	return fmt.Sprintf("%d - %s", m.Code, m.Message)
}
