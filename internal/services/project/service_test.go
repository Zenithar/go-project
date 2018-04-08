package project_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"

	"go.zenithar.org/todo/internal/models"
	"go.zenithar.org/todo/internal/repositories"
	"go.zenithar.org/todo/internal/repositories/mocks"
	"go.zenithar.org/todo/internal/services"
	"go.zenithar.org/todo/internal/services/project"
	"go.zenithar.org/todo/pkg/log"
	dto "go.zenithar.org/todo/pkg/protocol/todo"
	"go.zenithar.org/todo/pkg/tracer"
)

func TestProject_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	projectRepo := mocks.NewMockProjectRepository(mockCtrl)

	logFactory := log.With(zap.String("service", "project"))
	tr := tracer.Init("service.Project", logFactory)

	var srv services.Project
	{
		srv = project.NewService(projectRepo)
		srv = project.LoggerDecorator(logFactory, srv)
		srv = project.TracingDecorator(tr, srv)
	}

	Convey("Given a valid project creation request", t, func() {

		req := &dto.GetProjectReq{
			Id: "JKAMWXCVBNAZERTYUIOPQSDFGHJKLMWX",
		}

		So(req.Validate(), ShouldBeNil)

		Convey("When trying to retrieve from database", func() {
			// Arm mocks
			projectRepo.EXPECT().Read(ctx, gomock.Any()).Return(models.NewProject("test"), nil).Times(1)

			// Call
			res := srv.Get(ctx, req)

			Convey("Then entity should be found", func() {
				// Checks
				So(res, ShouldNotBeNil)
				So(res.Error, ShouldBeNil)
				So(res.Entity, ShouldNotBeNil)
			})
		})
	})
}

func TestProject_Get_InvalidRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	projectRepo := mocks.NewMockProjectRepository(mockCtrl)

	srv := project.NewService(projectRepo)

	Convey("Given an invalid project creation request", t, func() {

		req := &dto.GetProjectReq{
			Id: "JK-éWXCVBNAZERTYUIOPQSDFGHJKLMWX",
		}

		So(req.Validate(), ShouldNotBeNil)

		Convey("When trying to retrieve from database", func() {

			// Call
			res := srv.Get(ctx, req)

			Convey("Then entity should not be retrievable", func() {
				// Checks
				So(res, ShouldNotBeNil)
				So(res.Error, ShouldNotBeNil)
				So(res.Entity, ShouldBeNil)
				So(res.Error.Code, ShouldEqual, http.StatusPreconditionFailed)
			})
		})
	})
}

func TestProject_Get_DatabaseError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	projectRepo := mocks.NewMockProjectRepository(mockCtrl)

	srv := project.NewService(projectRepo)

	Convey("Given a valid project creation request", t, func() {

		req := &dto.GetProjectReq{
			Id: "JK11WXCVBNAZERTYUIOPQSDFGHJKLMWX",
		}

		So(req.Validate(), ShouldBeNil)

		Convey("When trying to retrieve from database", func() {
			// Arm mocks
			projectRepo.EXPECT().Read(ctx, gomock.Any()).Return(nil, fmt.Errorf("Test Error")).Times(1)

			// Call
			res := srv.Get(ctx, req)

			Convey("Then entity should not be retrievable", func() {
				// Checks
				So(res, ShouldNotBeNil)
				So(res.Error, ShouldNotBeNil)
				So(res.Entity, ShouldBeNil)
				So(res.Error.Code, ShouldEqual, http.StatusInternalServerError)
			})
		})
	})
}

func TestProject_Get_EntityNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	projectRepo := mocks.NewMockProjectRepository(mockCtrl)

	srv := project.NewService(projectRepo)

	Convey("Given a valid project creation request", t, func() {

		req := &dto.GetProjectReq{
			Id: "JK11WXCVBNAZERTYUIOPQSDFGHJKLMWX",
		}

		So(req.Validate(), ShouldBeNil)

		Convey("When trying to retrieve from database", func() {
			// Arm mocks
			projectRepo.EXPECT().Read(ctx, gomock.Any()).Return(nil, repositories.ErrNoResult).Times(1)

			// Call
			res := srv.Get(ctx, req)

			Convey("Then entity should not be retrievable", func() {
				// Checks
				So(res, ShouldNotBeNil)
				So(res.Error, ShouldNotBeNil)
				So(res.Entity, ShouldBeNil)
				So(res.Error.Code, ShouldEqual, http.StatusNotFound)
			})
		})
	})
}
