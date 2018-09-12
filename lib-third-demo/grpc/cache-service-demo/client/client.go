/*
 * 说明：
 * 作者：zhe
 * 时间：2018-09-10 5:22 PM
 * 更新：
 */

package client

import (
	"fmt"
	"io"
	"os"
	"time"

	"cache-service/interceptor"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Run() {
	if err := runClient(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run client: %v\n", err)
		os.Exit(1)
	}
}

// Server streaming RPCs
//
func runClient() error {
	// 建立连接
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure(), interceptor.WithClientInterceptor())
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}

	// 初始化 Cache 客户端对象
	cache := rpc.NewCacheClient(conn)

	// 调用 grpc 的 store() 方法存储键值对 {"gopher": "go go go"}
	ctx, _ := context.WithTimeout(context.Background(), 22*time.Second)    // 超时控制
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("dry-run", "0")) // 设置 Metadata(即类似：HTTP HEADER)

	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "aaa",
		Key:          "gopher",
		Val:          []byte("go go go"),
	})
	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "bbb",
		Key:          "grpc",
		Val:          []byte("rpc rpc rpc"),
	})
	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "ccc",
		Key:          "proto",
		Val:          []byte("pb pb pb"),
	})

	if err != nil {
		return fmt.Errorf("failed to store: %v", err)
	}

	// 调用 grpc 的 get() 方法取回键为 `gopher` 的值
	ctx, _ = context.WithTimeout(context.Background(), 22*time.Second)
	resp, err := cache.Get(ctx, &rpc.GetReq{Key: "gopher"})
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}
	fmt.Printf("Got cached value: %v\n", string(resp.Val))

	// 调用 grpc 的 Dump() 方法取回服务端存储的所有的键值对
	ctx, _ = context.WithTimeout(context.Background(), 22*time.Second)
	stream, err := cache.Dump(ctx, &rpc.DumpReq{})
	if err != nil {
		return fmt.Errorf("failed to dump: %v", err)
	}
	for {
		// 取回数据(流方式)
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to stream item: %v", err)
		}
		// 输出
		fmt.Printf("Dump result: %v\n", item.String())
	}
	return nil
}

/*
// Unary RPC
//
func runClient() error {
	// 建立连接
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure(), interceptor.WithClientInterceptor())
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}
	cache := rpc.NewCacheClient(conn)

	// 调用 grpc 的 store() 方法存储键值对 {"gopher": "con"}
	ctx, _ := context.WithTimeout(context.Background(), 22*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("dry-run", "0"))

	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "aaa",
		Key:          "gopher",
		Val:          []byte("go go go"),
	})
	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "bbb",
		Key:          "grpc",
		Val:          []byte("rpc rpc rpc"),
	})
	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "ccc",
		Key:          "proto",
		Val:          []byte("pb pb pb"),
	})

	if err != nil {
		return fmt.Errorf("failed to store: %v", err)
	}

	// 调用 grpc 的 get() 方法取回键为 `gopher` 的值
	ctx, _ = context.WithTimeout(context.Background(), 22*time.Second)
	resp, err := cache.Get(ctx, &rpc.GetReq{Key: "gopher"})
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}
	// 输出
	fmt.Printf("Got cached value: %v\n", string(resp.Val))

	// 调用 grpc 的 Dump() 方法取回服务端存储的所有的键值对
	ctx, _ = context.WithTimeout(context.Background(), 22*time.Second)
	dumpResp, err := cache.Dump(ctx, &rpc.DumpReq{})
	if err != nil {
		return fmt.Errorf("failed to dump: %v", err)
	}
	// 输出
	fmt.Printf("Dump result: %v\n", dumpResp.String())

	return nil
}
*/
