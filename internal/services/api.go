package services

import (
	"context"

	dto "go.zenithar.org/todo/pkg/protocol/todo"
)

// Project services contract
type Project interface {
	Get(context.Context, *dto.GetProjectReq) *dto.SingleProjectRes
}
