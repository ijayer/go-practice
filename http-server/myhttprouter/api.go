package myhttprouter

import (
	"fmt"
	"net/http"

	"context"

	"errors"

	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Ctx  context.Context // 请求上下文
	Data interface{}     // 响应数据
	Code int             // 状态码
}

// Index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (Response, error) {
	cookie := r.Header.Get("Cookie")
	logrus.Infof("Cookie: %v", cookie)

	// Write Response
	fmt.Fprint(w, "Not protected, Welcome!\n")
	fmt.Fprintf(w, "Cookie: %v\n", cookie)

	return Response{Ctx: r.Context(), Data: cookie, Code: 200}, nil
}

// Hello
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (Response, error) {
	ctx := r.Context()
	name := ps.ByName("name")
	ctx = context.WithValue(ctx, "key", "123")

	fmt.Fprint(w, name)
	return Response{Ctx: ctx, Data: name, Code: 404}, errors.New("404 Forbidden")
}

// AllParaHello
func AllParaHello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s\n", ps.ByName("name"))
}

// Protected
func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}

// BasicAuth
func BasicAuth(h httprouter.Handle, requiredUsername, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username, password, hasAuth := r.BasicAuth()

		if hasAuth && username == requiredUsername && password == requiredPassword {
			h(w, r, ps)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// Cookie
func Cookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set Cookie
	cookie := new(http.Cookie)
	cookie.Name = "foo.1"
	cookie.Value = "bar.1"
	http.SetCookie(w, cookie)

	cookie2 := new(http.Cookie)
	cookie2.Name = "foo.2"
	cookie2.Value = "bar.2"
	cookie2.Path = "/index"
	cookie2.Secure = true
	cookie2.Domain = "test.robot-qixing.com"
	cookie2.Expires = time.Now().UTC().Add(30 * time.Second)
	http.SetCookie(w, cookie2)

	cookie3 := new(http.Cookie)
	cookie3.Name = "foo.3"
	cookie3.Value = "bar.3"
	http.SetCookie(w, cookie3)

	fmt.Fprintf(w, "%s\n%s\n%s", cookie.String(), cookie2.String(), cookie3.String())
}

func Redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "http://test.robot-qixing.com:1011/#/login", http.StatusTemporaryRedirect)
}
