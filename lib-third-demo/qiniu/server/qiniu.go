package server

import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// 七牛云服务相关参数
const (
	bucket    = "qcloud"
	accessKey = "access key from qiniu"
	secretKey = "secret key from qiniu"
)

// 服务器直接上传文件到七牛云(表单方式：不覆盖)
func ServerUpload() {
	// 文件参数
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s", bucket),
		// Scope:fmt.Sprintf("%s:%s", Bucket, "key"), //最终生成的upToken将包含key(文件名)
		Expires: 3600, // 1 hour
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)       // 上传凭证
	ctx := context.Background()                 // 上下文
	localFile := `D:\others\Beauty\1365962.jpg` // 上传文件路径
	key := "avatar.jpg"                         // 存储在七牛的文件名
	ret := storage.PutRet{}                     // 标准的上传回复内容
	putExtra := storage.PutExtra{               // 上传额外配置选项
		Params: map[string]string{
			"x:name": "avatar",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件, 文件名将由key指定
	err := formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ret)

	// 上传文件，upToken未指定key, 文件名将由文件hash作为key
	err = formUploader.PutFileWithoutKey(ctx, &ret, upToken, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ret)
}

// 自定义上传回复结构体
type MyPutRet struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Name   string `json:"name"`
	Fsize  int    `json:"fsize"`
	Bucket string `json:"bucket"`
}

// 服务器直接上传文件到七牛云(表单方式：自定义上传回复内容)
func ServerUploadWithDefinePutRet() {
	// 文件参数
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s", bucket),
		Expires:    3600,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)       // 上传凭证
	ctx := context.Background()                 // 上下文
	localFile := `D:\others\Beauty\1365962.jpg` // 上传文件路径
	key := "my_avatar.jpg"                      // 存储在七牛的文件名
	ret := MyPutRet{}                           // 自定义上传回复内容
	putExtra := storage.PutExtra{               // 上传额外配置选项
		Params: map[string]string{
			"x:name": "avatar",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件, 文件名将由key指定
	err := formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ret)
}

// 服务器直接上传文件到七牛云(表单方式：覆盖原始文件)
func ServerOverwriteUpload() {
	// 文件参数
	keyToOverwrite := "avatar_1.jpg"
	putPolicy := storage.PutPolicy{
		Scope:   fmt.Sprintf("%s:%s", bucket, keyToOverwrite), //最终生成的upToken将包含key(覆盖文件名)
		Expires: 3600,                                         // 1 hour
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)    // 上传凭证
	ctx := context.Background()              // 上下文
	localFile := `D:\others\Beauty\girl.jpg` // 上传文件路径
	key := "avatar_1.jpg"                    // 存储在七牛的文件名
	ret := storage.PutRet{}                  // 标准的上传回复内容
	putExtra := storage.PutExtra{            // 上传额外配置选项
		Params: map[string]string{
			"x:name": "avatar",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 覆盖上传：取本地另一个文件上传，key(文件名)依旧保持前面设定的值，然后在生成upToken时将keyToOverwrite值设置为key的值
	err := formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", ret)
}

// 服务器直接上传文件到七牛云(表单方式：启用回调)
// 回调地址通过 Ngrok(https://ngrok.com/) 将公网地址(随机生成的)映射到内网测试
var (
	callbackURL  = "http://f81b42d6.ngrok.io/api/qiniu/callback"                                                                    // 向业务服务器发起的POST回调请求
	callbackBody = `{"key":$(key),"name":$(fname),"hash":$(etag),"file_size":$(fsize), "id":$(x:id), "collection":$(x:collection)}` // 发送POST回调请求的Body内容
)

func ServerUploadWithCallback() {
	// 文件参数
	putPolicy := storage.PutPolicy{
		Scope:            bucket,             //
		Expires:          3600,               // 1 hour
		CallbackURL:      callbackURL,        //
		CallbackBody:     callbackBody,       //
		CallbackBodyType: "application/json", //
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)     // 上传凭证
	ctx := context.Background()               // 上下文
	localFile := `D:\others\Beauty\girl3.jpg` // 上传文件路径
	key := "girl3.jpg"                        // 存储在七牛的文件名
	ret := storage.PutRet{}                   // 标准的上传回复内容
	putExtra := storage.PutExtra{             // 上传额外配置选项
		Params: map[string]string{
			"x:name": "avatar",
		},
	}

	// 资源管理配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件, 文件名将由key指定
	err := formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%v\n", ret)
}
