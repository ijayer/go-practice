/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-01 15:33
 * 更新：
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

var srv = "http://127.0.0.1:8081"

func main() {
	DirectUploadToQiNiu()
	UploadToQiNiuWithCallback()
	OverwrittenFileWithCallback()
}

func uploadToken(url string) (string, error) {
	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := Response{}
	err = json.Unmarshal(body, &data)

	return data.Data.(string), err
}

// DirectUploadToQiNiu 直接上传文件指定到七牛云，多个文件上传则多次调用该接口实现
func DirectUploadToQiNiu() {
	// 请求配置有回调业务的上传凭证
	url := fmt.Sprintf("%s%s", srv, "/api/qiniu/upload/token")
	upToken, err := uploadToken(url) // 请求上传令牌
	if err != nil {
		fmt.Printf("token error: %v\n", upToken)
		return
	}

	ctx := context.Background()               // 上下文
	localFile := `D:\others\Beauty\girl4.jpg` // 上传文件路径
	key := "girl4.jpg"                        // 存储在七牛的文件名
	putExtra := storage.PutExtra{             // 上传额外配置选项
		Params: map[string]string{
			"x:name":       "avatar",
			"x:id":         "5a94f6f3c7f41c3f3c3508bf",
			"x:collection": "users",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件, 文件名将由key指定
	var ret interface{}
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Printf("upload file error: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", ret)
}

// UploadToQiNiuWithCallback 上传文件到七牛云，并带有回调策略
func UploadToQiNiuWithCallback() {
	// 请求配置有回调业务的上传凭证
	url := fmt.Sprintf("%s%s", srv, "/api/qiniu/upload/token?ktw=o.jpg")
	upToken, err := uploadToken(url)
	if err != nil {
		fmt.Printf("token error: %v\n", upToken)
		return
	}

	ctx := context.Background()           // 上下文
	localFile := `D:\others\Beauty\o.jpg` // 上传文件路径
	key := "o.jpg"                        // 存储在七牛的文件名
	putExtra := storage.PutExtra{         // 上传额外配置选项
		Params: map[string]string{
			"x:name":       "avatar",
			"x:id":         "5a94f6f3c7f41c3f3c3508bf",
			"x:collection": "users",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件, 文件名将由key指定
	var ret interface{}
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Printf("upload file error: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", ret)
}

// OverwrittenFileWithCallback 覆盖七牛云上的文件，并带有回调业务
func OverwrittenFileWithCallback() {
	// 请求配置有回调业务的上传凭证
	url := fmt.Sprintf("%s%s", srv, "/api/qiniu/upload/token?ktw=o.jpg&tp=callback")
	upToken, err := uploadToken(url)
	if err != nil {
		fmt.Printf("token error: %v\n", upToken)
		return
	}

	ctx := context.Background()                           // 上下文
	localFile := `D:\others\Beauty\19217264230798708.jpg` // 上传文件路径
	key := "o.jpg"                                        // 存储在七牛的文件名
	putExtra := storage.PutExtra{                         // 上传额外配置选项
		Params: map[string]string{
			"x:name":       "avatar",
			"x:id":         "5a94f6f3c7f41c3f3c3508bf",
			"x:collection": "users",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件, 文件名将由key指定
	var ret interface{}
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Printf("upload file error: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", ret)
}
