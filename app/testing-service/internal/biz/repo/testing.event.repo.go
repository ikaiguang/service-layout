package bizrepos

import (
	"context"
	"gitlab.realibox.cn/hubv3tools/protobuf/types/known/emptypb"
)

type Handler func(context.Context, emptypb.Empty) error

type TestingEventRepo interface {
	Send(ctx context.Context, msg emptypb.Empty) error
	Receive(ctx context.Context, handler Handler) error
	Close() error
}
