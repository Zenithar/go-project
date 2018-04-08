package project

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tlog "github.com/opentracing/opentracing-go/log"

	"go.zenithar.org/todo/internal/services"
	dto "go.zenithar.org/todo/pkg/protocol/todo"
)

type tracingMiddleware struct {
	next   services.Project
	tracer opentracing.Tracer
}

// TracingDecorator decorates a Project service to add OpenTracing features
func TracingDecorator(tracer opentracing.Tracer, next services.Project) services.Project {
	return &tracingMiddleware{
		next:   next,
		tracer: tracer,
	}
}

// -----------------------------------------------------------------------------

func (s *tracingMiddleware) withSpan(ctx context.Context, methodName string, request func(context.Context, opentracing.Span) *dto.Error) {
	var err *dto.Error

	// OpenTracing
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span = s.tracer.StartSpan(methodName, opentracing.ChildOf(span.Context()))
		defer func(span opentracing.Span) {
			if err != nil {
				ext.Error.Set(span, true)
				span.LogFields(tlog.Error(err))
			}
			span.Finish()
		}(span)
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	// Do the call
	err = request(ctx, span)
}

// -----------------------------------------------------------------------------

func (s *tracingMiddleware) Get(ctx context.Context, req *dto.GetProjectReq) *dto.SingleProjectRes {
	var res *dto.SingleProjectRes

	s.withSpan(ctx, "Get", func(spanCtx context.Context, span opentracing.Span) *dto.Error {
		// Instrument span
		if span != nil {
			span.SetTag("params.project_id", req.Id)
		}

		// Real call
		res = s.next.Get(spanCtx, req)

		// No error
		return res.Error
	})

	// Return result
	return res
}
