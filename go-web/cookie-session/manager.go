/*
 * 说明：实现 Session 管理器
 * 作者：zhe
 * 时间：2018-05-25 4:46 PM
 * 更新：
 */

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Provider 定义Session管理器底层的存储操作
type Provider interface {
	// SessionInit 实现Session初始化，操作成功后返回新建的Session对象
	SessionInit(sid string) (Session, error)

	// SessionRead 返回sid所代表的Session对象，如果不存在，则以sid为参数
	// 调用SessionInit函数创建并返回一个新的Session对象
	SessionRead(sid string) (Session, error)

	// SessionDestroy 函数用来销毁sid对应的Session对象
	SessionDestroy(sid string) error

	// SessionGC 根据maxLifeTime来删除过期的Session对象
	SessionGC(maxLifeTime int64)
}

// Session 定义用户对Session存储数据项的操作
type Session interface {
	// Set 设置Session的键值对
	Set(key, value interface{}) error

	// Get 根据key获取Session值
	Get(key interface{}) interface{}

	// Delete 删除key对应的Session值
	Delete(key interface{}) error

	// SessionId 返回当前的SessionId
	SessionId() string
}

// providers 所有Sessions的集合
//      map["memory"]    = Provider
//      map["file"]      = Provider
//      map["mongodb"]   = Provider
//      ...
// Provider 存储实现了Session管理器的对象
var providers = make(map[string]Provider)

// Register 在providers集合中注册Session
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: register called twice for provider " + name)
	}
	providers[name] = provider
}

// Manager 定义Session管理器
type Manager struct {
	cookieName  string     // private cookie name
	mu          sync.Mutex // protects session
	maxLifeTime int64
	provider    Provider
}

// NewManager 初始化Session管理器Manager
//   providerName：在 providers 中获取实现了Session管理器的对象，完成 Manager.Provider的初始化
//   cookieName:
//   maxLifeTime:  Session 的最大生命周期(转化为时间戳存储，便于GC处理)
func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknow provide %s", providerName)
	}
	return &Manager{
		provider:    provider,
		cookieName:  cookieName,
		maxLifeTime: maxLifeTime,
		mu:          sync.Mutex{},
	}, nil
}

// NewSessionId 返回全局唯一的Session ID
func (m *Manager) NewSessionId() string {
	var sid string
	defer func() {
		logrus.Debugf("sid: %v", sid)
	}()
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	sid = base64.URLEncoding.EncodeToString(b)
	return sid
}

// SessionStart 为每个来访用户分配或获取与他相关连的Session，以便后面根据
// Session信息来验证操作。SessionStart这个函数就是用来检测是否已经有某个
// Session与当前来访用户发生了关联，如果没有则创建之。
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) Session {
	var session Session
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := m.NewSessionId()
		session, _ = m.provider.SessionInit(sid)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}
	return session
}

// SessionDestroy
func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		// m.mu.Lock()
		// defer m.mu.Unlock()

		m.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{
			Name:     m.cookieName,
			Path:     "/",
			HttpOnly: true,
			Expires:  expiration,
			MaxAge:   -1,
		}
		http.SetCookie(w, &cookie)
	}
}

// GC
func (m *Manager) GC() {
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime), func() {
		m.GC()
	})
}
