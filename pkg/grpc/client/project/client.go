package project

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"

	"go.zenithar.org/todo/internal/services"
	"go.zenithar.org/todo/pkg/grpc/internal/pb"
	dto "go.zenithar.org/todo/pkg/protocol/todo"
)

var (
	// PoolInitial defines the initial pool size
	PoolInitial = 2
	// PoolMax defines the maximum pool size
	PoolMax = 10
	// PoolTimeout defines the connection timeout
	PoolTimeout = 30 * time.Second
)

type client struct {
	pool *grpcpool.Pool
}

// NewClient initializes a remote project service client
func NewClient(uri string, opts ...grpc.DialOption) (services.Project, error) {
	// Initialize connection pool
	pool, err := grpcpool.New(func() (*grpc.ClientConn, error) {
		return grpc.Dial(uri, opts...)
	}, PoolInitial, PoolMax, PoolTimeout)

	// Close on destruction
	runtime.SetFinalizer(pool, func(p *grpcpool.Pool) {
		if !p.IsClosed() {
			p.Close()
		}
	})

	// Return wrapper
	return &client{
		pool: pool,
	}, err
}

// -----------------------------------------------------------------------------

func (c *client) withClient(ctx context.Context, request func(pb.ProjectClient) error) *dto.Error {
	// Acquire a connection
	conn, err := c.pool.Get(ctx)
	if err != nil {
		return &dto.Error{
			Code:    http.StatusServiceUnavailable,
			Message: fmt.Sprintf("Unable to acquire a connection, %v", err),
		}
	}
	defer conn.Close()

	// Initialize a client
	remote := pb.NewProjectClient(conn.ClientConn)

	// Do the call
	if errCall := request(remote); errCall != nil {
		return &dto.Error{
			Code:    http.StatusServiceUnavailable,
			Message: fmt.Sprintf("Error occurs during remote call, %v", errCall),
		}
	}

	// Return no error
	return nil
}

// -----------------------------------------------------------------------------

func (c *client) Get(ctx context.Context, in *dto.GetProjectReq) *dto.SingleProjectRes {
	var err error
	res := &dto.SingleProjectRes{}

	// Do the call
	if errCall := c.withClient(ctx, func(remote pb.ProjectClient) error {
		res, err = remote.Get(ctx, in)
		return err
	}); errCall != nil {
		res.Error = errCall
	}

	// Return result
	return res
}
