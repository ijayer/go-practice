/*
 * 说明：
 * 作者：zhe
 * 时间：2018-05-29 21:44
 * 更新：
 */

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// SessionStore 定义Session数据项存储结构
type SessionStore struct {
	Id           string                      // session id
	Items        map[interface{}]interface{} // session 存储的值
	TimeAccessed time.Time                   // session 最后的访问时间
}

// NewSessionStore 初始化并返回SessionStore的指针对象
func NewSessionStore(sid string) *SessionStore {
	return &SessionStore{
		Id:           sid,
		Items:        make(map[interface{}]interface{}),
		TimeAccessed: time.Now(),
	}
}

// SessionId
func (st *SessionStore) SessionId() string {
	return st.Id
}

// Set
func (st *SessionStore) Set(key, value interface{}) error {
	globalMemProvider.SessionUpdate(st.Id)
	st.Items[key] = value
	return nil
}

// Get
func (st *SessionStore) Get(key interface{}) interface{} {
	globalMemProvider.SessionUpdate(st.Id)
	return st.Items[key]
}

// Delete
func (st *SessionStore) Delete(key interface{}) error {
	globalMemProvider.SessionUpdate(st.Id)
	delete(st.Items, key)
	return nil
}

// MemoryProvider 定义基于内存的Session存储对象
type MemoryProvider struct {
	mu       sync.Mutex             // 互斥锁
	sessions map[string]interface{} // 存储在内存中的Session
}

// SessionInit
func (mp *MemoryProvider) SessionInit(sid string) (Session, error) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	newSession := NewSessionStore(sid)
	mp.sessions[sid] = newSession

	return newSession, nil
}

// SessionRead
func (mp *MemoryProvider) SessionRead(sid string) (Session, error) {
	session, ok := mp.sessions[sid]
	if ok {
		return session.(*SessionStore), nil
	}
	return mp.SessionInit(sid)
}

// SessionDestroy
func (mp *MemoryProvider) SessionDestroy(sid string) error {
	delete(mp.sessions, sid)
	return nil
}

// SessionGC
func (mp *MemoryProvider) SessionGC(maxLifeTime int64) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	now := time.Now().Unix()
	for _, session := range mp.sessions {
		sessionExpiredTime := session.(*SessionStore).TimeAccessed.Unix() + maxLifeTime
		if now > sessionExpiredTime {
			delete(mp.sessions, session.(*SessionStore).Id)
			logrus.Debugf("[DELETE]: session=%s", session.(*SessionStore).Id)
			break
		}
	}
}

// SessionUpdate
func (mp *MemoryProvider) SessionUpdate(sid string) error {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	if session, ok := mp.sessions[sid]; ok {
		session.(*SessionStore).TimeAccessed = time.Now()
		return nil
	}
	return fmt.Errorf("the session[%s] not exist", sid)
}
