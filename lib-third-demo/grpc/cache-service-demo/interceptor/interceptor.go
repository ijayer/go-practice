/*
 * 说明: 中间件(拦截器): 只要在客户端和服务端分别注册了 Interceptor,
 * 那么进行 RPC 调用的时候，这些中间件会先被调用，因此这个中间件可以对
 * 调用进行一层包装，然后再进行调用。
 * 作者：zhe
 * 时间：2018-09-11 1:33 PM
 * 更新：
 */

package interceptor

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// WithClientInterceptor 客户端 RPC 调用拦截器
func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(ctx context.Context, method string, req interface{},
	resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	start := time.Now()
	err := invoker(ctx, method, req, resp, cc, opts...)
	log.Printf("invoke remote method=%s duration=%s error=%v", method, time.Since(start), err)
	return err
}

// NotIdempotent 置为 '非幂等'
func NotIdempotent(ctx context.Context) context.Context {
	return context.WithValue(ctx, "idempotent", "false")
}

// IsIdempotent 判断是否幂等
func IsIdempotent(ctx context.Context) bool {
	val, ok := ctx.Value("idempotent").(bool)
	if !ok {
		return true
	}
	return val
}

// ServerInterceptor 服务端 RPC 调用拦截器
func ServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("invoke server method=%s duration=%s error=%v", info.FullMethod, time.Since(start), err)
	return resp, err
}
