/*
 * 说明：Server 端限流
 * 作者：zhe
 * 时间：2018-09-11 4:14 PM
 * 更新：
 */

package server

import (
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

// Tap
type Tap struct {
	lim *rate.Limiter
}

func NewTap() *Tap {
	return &Tap{lim: rate.NewLimiter(150, 5)}
}

func (t *Tap) Handler(ctx context.Context, info *tap.Info) (context.Context, error) {
	if !t.lim.Allow() {
		return nil, status.Errorf(codes.ResourceExhausted, "service is over rate limit")
	}
	return ctx, nil
}
