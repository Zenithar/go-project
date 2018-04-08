package project

import (
	"context"

	"go.uber.org/zap"

	"go.zenithar.org/todo/internal/services"
	"go.zenithar.org/todo/pkg/log"
	dto "go.zenithar.org/todo/pkg/protocol/todo"
)

type loggerMiddleware struct {
	next   services.Project
	logger log.Factory
}

// LoggerDecorator adds logging feature to project service
func LoggerDecorator(logger log.Factory, next services.Project) services.Project {
	return &loggerMiddleware{
		logger: logger,
		next:   next,
	}
}

// -----------------------------------------------------------------------------

func (s *loggerMiddleware) Get(ctx context.Context, req *dto.GetProjectReq) *dto.SingleProjectRes {
	s.logger.For(ctx).Debug("BEGIN project.Get", zap.Stringer("request", req))

	res := s.next.Get(ctx, req)

	s.logger.For(ctx).Debug("END project.Get", zap.Stringer("request", req), zap.Stringer("response", res), zap.Error(res.Error))
	return res
}
