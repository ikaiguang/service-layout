package events

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	bizrepos "github.com/ikaiguang/service-layout/app/testing-service/internal/biz/repo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type testingEvent struct {
	log *log.Helper
}

func NewTestingEvent(
	logger log.Logger,
) bizrepos.TestingEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/event"))

	return &testingEvent{
		log: logHelper,
	}
}

func (t testingEvent) Send(ctx context.Context, msg emptypb.Empty) error {
	return errorpkg.WithStack(errorpkg.ErrorUnimplemented("implement me"))
}

func (t testingEvent) Receive(ctx context.Context, handler bizrepos.Handler) error {
	return errorpkg.WithStack(errorpkg.ErrorUnimplemented("implement me"))
}

func (t testingEvent) Close() error {
	return nil
}
