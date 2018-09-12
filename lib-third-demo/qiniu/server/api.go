/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-01 14:41
 * 更新：
 */

package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// Response
type Response struct {
	Data interface{} `json:"data"`
}

// HTTP Error
type HTTPError struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}

// 七牛云测试域名
var domain = "http://p4wdbkaxc.bkt.clouddn.com"

// UploadToken 给客户端颁发七牛云的文件上传凭证
// URL Params: ktw 覆盖文件名：有值时，生成的token启用覆盖上传
//             tp  token类型(tokenType): 值为 `callback`时，生成的token启用回掉策略
// 备注：callbackXXX 和 returnXXX 不可混用
func UploadToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 72000 * 12, // 默认token有效期12小时
	}

	// 覆盖上传(替换)
	keyToWritten := r.URL.Query().Get("ktw")
	if keyToWritten != "" {
		putPolicy.Scope = fmt.Sprintf("%s:%s", bucket, keyToWritten)
		putPolicy.ReturnBody = `{"key":$(key), "name":$(fname), "hash":$(etag),"file_size":$(fsize)}`
	}

	// 回调策略: 由七牛云发起回调请求到业务服务器
	tokenType := r.URL.Query().Get("tp")
	if tokenType == "callback" {
		putPolicy.CallbackURL = callbackURL
		putPolicy.CallbackBody = callbackBody
		putPolicy.CallbackBodyType = "application/json"
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	data := Response{Data: upToken}
	WriteJson(w, http.StatusOK, data)
}

// DownloadURL 颁发给客户端七牛云文件的私有下载地址
func DownloadURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	key := r.URL.Query().Get("key")
	mac := qbox.NewMac(accessKey, secretKey)
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)

	data := Response{Data: privateAccessURL}
	WriteJson(w, http.StatusOK, data)
}

// QiNiuCallback 七牛云回调请求接口
func QiNiuCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Printf("callback_addr: %v\n", r.Host)

	mac := qbox.NewMac(accessKey, secretKey)
	if ok, err := qbox.VerifyCallback(mac, r); !ok || err != nil {
		e := HTTPError{Code: http.StatusNetworkAuthenticationRequired, Error: err.Error()}
		WriteJson(w, http.StatusNetworkAuthenticationRequired, &e)
		return
	}

	//data, _ := ioutil.ReadAll(r.Body)
	//fmt.Printf("body: %v\n", string(data))

	type callbackBody struct {
		Key        string `json:"key"`
		Name       string `json:"name"`
		Hash       string `json:"hash"`
		FileSize   int    `json:"file_size"`
		Id         string `json:"id"`
		Collection string `json:"collection"`
	}
	cb := callbackBody{}
	if err := ReadJson(r, &cb); err != nil {
		fmt.Printf("read callback body faild: %v\n", err)
		return
	}

	deadline := time.Now().Add(time.Second * 3600 * 24 * 365).Unix()
	privateAccessURL := storage.MakePrivateURL(mac, domain, cb.Key, deadline)

	type Resp struct {
		Name     string `json:"name"`
		FileSize int    `json:"file_size"`
		LinkURL  string `json:"link_url"`
	}
	resp := Resp{Name: cb.Key, FileSize: cb.FileSize, LinkURL: privateAccessURL}
	WriteJson(w, http.StatusOK, &Response{Data: &resp})
}

func ReadJson(r *http.Request, v interface{}) (err error) {
	b, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(b, &v)
	return err
}

func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	b, _ := json.Marshal(v)
	WriteResult(w, b, code)
}

func WriteResult(w http.ResponseWriter, data []byte, custom int) {
	w.Header().Set("Content-Type", "application/json")
	Write(w, custom, data)
}

func Write(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	w.Write(data)
}
