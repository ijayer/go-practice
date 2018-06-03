/*
 * 说明：
 * 作者：zhe
 * 时间：2018-05-28 4:50 PM
 * 更新：https://www.sohamkamani.com/blog/2018/03/25/golang-session-authentication/
 */

package main

import (
	"fmt"
	"net/http"
	"net/url"

	"time"

	"github.com/julienschmidt/httprouter"
	"instance.golang.com/utils"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser() *User {
	return &User{
		Username: "test",
		Password: "test",
	}
}

// SignIn
func (u *User) Signin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := User{}

	// 验证登陆凭证
	err := utils.ReadJSON(r, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if u.Username != user.Username || u.Password != user.Password {
		http.Error(w, "incorrect username|password", http.StatusUnauthorized)
		return
	}

	// 处理session：如果 session 已完成过初始化，则取到该值并返回；
	// 如果不存在，则创建并返回
	session := globalSessionManager.SessionStart(w, r)
	session.Set("username", user.Username)

	// Set Cookie
	cookie := &http.Cookie{
		Name:     globalSessionManager.cookieName,
		Value:    url.QueryEscape(session.SessionId()), // 查询转义
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(globalSessionManager.maxLifeTime) * time.Second),
		Secure:   false,
		HttpOnly: false,
	}
	http.SetCookie(w, cookie)

	utils.WriteJson(w, struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{
		Code: http.StatusOK,
		Msg:  "sign in success",
	})
}

// SignOut
func (u *User) Signout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 根据客户端请求中的Cookie，读取保存在服务端的Session信息
	session, code := obtainSession(r)
	if code != 0 {
		w.WriteHeader(code)
		return
	}
	sid := session.SessionId()

	// 清理该账户的Session会话
	if err := globalMemProvider.SessionDestroy(sid); err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Welcome
func (u *User) Welcome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 根据客户端请求中的Cookie，读取保存在服务端的Session信息
	session, code := obtainSession(r)
	if code != 0 {
		w.WriteHeader(code)
		return
	}

	// Finally, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome：%+v", session)))
}

// Refresh Refresh HTTP handler to refresh the users session token every time
// they hit the /refresh route in our application
func Refresh(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 根据客户端请求中的Cookie，读取保存在服务端的Session信息
	session, code := obtainSession(r)
	if code != 0 {
		w.WriteHeader(code)
		return
	}
	sid := session.SessionId()

	// 删除旧的 Session 会话
	if err := globalMemProvider.SessionDestroy(sid); err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 创建新的 Session 会话
	sid = globalSessionManager.NewSessionId()
	session, _ = globalMemProvider.SessionInit(sid)
	cookie := &http.Cookie{
		Name:     globalSessionManager.cookieName,
		Value:    url.QueryEscape(session.SessionId()), // 查询转义
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(globalSessionManager.maxLifeTime) * time.Second),
		Secure:   false,
		HttpOnly: false,
	}
	http.SetCookie(w, cookie)
}

// obtainSession 获取服务端存储的Session会话
func obtainSession(r *http.Request) (Session, int) {
	// 读取客户端的 Cookies，Cookie 会随着每个请求发送过来
	cookie, err := r.Cookie(globalSessionManager.cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return nil, http.StatusUnauthorized
		}
		// For any other type of error, return a bad request status
		return nil, http.StatusBadRequest
	}
	clientSid, _ := url.QueryUnescape(cookie.Value)

	// 根据 sid 读取 Session 会话信息
	session, err := globalMemProvider.SessionRead(clientSid)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		return nil, http.StatusInternalServerError
	}
	if session == nil {
		// If the session token is not present in cache, return an unauthorized error
		return nil, http.StatusUnauthorized
	}

	return session, 0
}
