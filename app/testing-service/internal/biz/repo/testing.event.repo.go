package bizrepos

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler func(context.Context, *emptypb.Empty) error

type TestingEventRepo interface {
	Send(ctx context.Context, msg *emptypb.Empty) error
	Receive(ctx context.Context, handler Handler) error
	Close(ctx context.Context) error
}
