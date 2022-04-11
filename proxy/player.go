package proxy

import "context"

type PlayerMsg interface {
	Send(ctx context.Context, userKey []interface{}, data ...interface{}) error
	ReConn(ctx context.Context, userKey []interface{}) error
}
