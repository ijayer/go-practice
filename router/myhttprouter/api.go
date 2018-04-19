package myhttprouter

import (
	"fmt"
	"net/http"

	"context"

	"errors"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Ctx  context.Context // 请求上下文
	Data interface{}     // 响应数据
	Code int             // 状态码
}

// Index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (Response, error) {
	fmt.Fprint(w, "Not protected, Welcome!\n")

	return Response{Ctx: r.Context(), Data: "index", Code: 200}, nil
}

// Hello
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (Response, error) {
	var ctx = r.Context()
	name := ps.ByName("name")

	fmt.Fprint(w, name)

	ctx = context.WithValue(ctx, "key", "123")

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
