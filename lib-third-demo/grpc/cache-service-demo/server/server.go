/*
 * 说明：gRPC Server 业务实现
 * 作者：zhe
 * 时间：2018-09-10 4:36 PM
 * 更新：
 */

package server

import (
	"net"
	"os"

	interceptor "cache-service/interceptor"
	rpc "cache-service/proto"

	"golang.org/x/net/context"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	addr = "localhost:5051"
	max  = 64
)

// Run run server
func Run() {
	if err := runServer(); err != nil {
		Fprintf(os.Stderr, "failed to run cache server: %s\n", err)
		os.Exit(1)
	}
}

// runServer 初始化一个 grpc Server, 并在 addr 上启动运行(阻塞方式)
//
// netutil.LimitListener(l, 1024) & grpc.MaxConcurrentStreams(64) 两个结合起来基本控制了并发的总数
func runServer() error {
	// 初始化 grpc server
	srv := grpc.NewServer(
		interceptor.ServerInterceptor(),    // 注册拦截器
		grpc.MaxConcurrentStreams(max),     // 指定每个 grpc 的连接可以有 max 个并发流(stream)
		grpc.InTapHandle(NewTap().Handler), // 限制访问频率
	)

	// 注册服务
	rpc.RegisterCacheServer(srv, &CacheService{})
	rpc.RegisterAccountsServer(srv, &AccountService{})

	// 运行服务
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return Errorf("listen error: %v", err)
	}
	l = netutil.LimitListener(l, 1014) // 限制总共可以有多少个连接

	// blocks until complete
	return srv.Serve(l)
}

// CacheService 服务端缓存对象，须实现 .proto 定义的服务接口：service Cache{}
type CacheService struct {
	accounts      rpc.AccountsClient
	store         map[string][]byte
	keysByAccount map[string]int64
}

func (s *CacheService) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	val, ok := s.store[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "key not found %s", req.Key)
	}
	return &rpc.GetResp{Val: val}, nil
}

/*
// Unary RPC
//
// Dump 取回所有的键值对
func (s *CacheService) Dump(ctx context.Context, req *rpc.DumpReq) (*rpc.DumpResp, error) {
	var pbDumpResp rpc.DumpResp
	var pbItem rpc.DumpItem

	for k, v := range s.store {
		pbItem.Key = k
		pbItem.Val = v
		pbDumpResp.Items = append(pbDumpResp.Items, &pbItem)
	}

	return &pbDumpResp, nil
}
*/

// Server streaming RPCs
//
// Dump 以流的方式取回所有的键值对
func (s *CacheService) Dump(req *rpc.DumpReq, stream rpc.Cache_DumpServer) error {
	for k, v := range s.store {
		stream.Send(&rpc.DumpItem{Key: k, Val: v})
	}
	return nil
}

// Store
func (s *CacheService) Store(ctx context.Context, req *rpc.StoreReq) (*rpc.StoreResp, error) {
	s.initMap()

	// 创建 Account 的客户端连接
	if err := s.initAccounts(); err != nil {
		return nil, Errorf("failed to dial server in s.initAccounts: %v", err)
	}

	// 调用另一个服务获取账户的信息，包含其键值限制
	resp, err := s.accounts.GetByToken(ctx, &rpc.GetByTokenReq{Token: req.AccountToken})
	if err != nil {
		return nil, Errorf("failed s.accounts.GetByToken: %v", err)
	}

	// 检查是否超量使用
	if s.keysByAccount[req.AccountToken] >= resp.Account.MaxCacheKeys {
		return nil, status.Errorf(codes.FailedPrecondition, "Account %s exceeds max key limit %d", req.AccountToken, resp.Account.MaxCacheKeys)
	}

	// 如果 key 不存在，需要新加键值对，那么我们就对计数器 +1
	if !dryRun(ctx) {
		if _, ok := s.store[req.Key]; !ok {
			s.keysByAccount[req.AccountToken] += 1
		}
		s.store[req.Key] = req.Val
	}
	return &rpc.StoreResp{}, nil
}

// initMap 初始化 Map
func (s *CacheService) initMap() {
	if s.store == nil {
		s.store = make(map[string][]byte)
	}
	if s.keysByAccount == nil {
		s.keysByAccount = make(map[string]int64)
	}
}

// initAccounts 初始化 account 的客户端对象
func (s *CacheService) initAccounts() error {
	// 建立连接
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure(), interceptor.WithClientInterceptor())
	if err != nil {
		return Errorf("failed to dial server: %v", err)
	}
	s.accounts = rpc.NewAccountsClient(conn)
	return nil
}

// AccountService 账号服务对象，须实现 .proto 定义的服务接口：service Accounts{}
type AccountService struct {
	cacheLimitByToken map[string]int64
}

// GetByToken 根据 token 获取该用户被允许存放的最大 k-v 数量，如果该 token 的值还未创建，则初始化后返回
func (s *AccountService) GetByToken(ctx context.Context, req *rpc.GetByTokenReq) (*rpc.GetByTokenResp, error) {
	if s.cacheLimitByToken == nil {
		s.cacheLimitByToken = make(map[string]int64)
	}

	v, ok := s.cacheLimitByToken[req.Token]
	if !ok {
		v = 256
		s.cacheLimitByToken[req.Token] = v // init max limits: 256
	}
	return &rpc.GetByTokenResp{Account: &rpc.Account{MaxCacheKeys: v}}, nil
}

// dryRun 根据客户端的设置，判断是否执行存放操作
func dryRun(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}
	val, ok := md["dry-run"]
	if !ok {
		return false
	}
	if len(val) < 1 {
		return false
	}
	return val[0] == "1"
}
