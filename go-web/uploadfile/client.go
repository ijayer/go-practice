package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("获取页面失败")
		}
	}()

	// 上传页面
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	html :=
		`<html>
	    <head>
	        <title>Golang Upload Files</title>
	    </head>
	    <body>
	        <form id="uploadForm"  enctype="multipart/form-data" action="/upload" method="POST">
	            <p>Golang Upload</p>
	            <br/>
	                <input type="file" id="file1" name="userfile" multiple="multiple" />
	            <br/>
	            <br/>
			<input type="text" name="tips" value="tips tips tips">
	            <br/>
	            <input type="submit" value="Upload">
	        </form>
	    </body>
	</html>`
	io.WriteString(w, html)
}

func uploadServer(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("文件上传异常")
		}
	}()

	if "POST" == r.Method {
		//在使用r.MultipartForm前必须先调用ParseMultipartForm方法，参数为最大缓存
		err := r.ParseMultipartForm(32 << 20) //
		if err != nil {
			fmt.Printf("#error: %v\n", err.Error())
		}
		//fmt.Println(r.MultipartForm)
		//fmt.Println(r.MultipartReader())

		// file
		if r.MultipartForm != nil && r.MultipartForm.File != nil {
			fileHeaderS := r.MultipartForm.File["userfile"] //获取所有上传文件信息
			num := len(fileHeaderS)
			fmt.Printf("总文件数：%d 个文件\n", num)

			// 循环对每个文件进行处理
			for n, fileHeader := range fileHeaderS {
				// 获取文件名
				filename := fileHeader.Filename

				// 结束文件
				file, err := fileHeader.Open()
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				// read form file
				src := make([]byte, 5000000)
				num, err := file.Read(src)
				if err != nil {
					fmt.Printf("#error: %v\n", err.Error())
				}

				// encode to string and save into file
				srcString := base64.StdEncoding.EncodeToString(src[0:num])
				err = ioutil.WriteFile("./uploadfile/"+filename+".txt", []byte(srcString), 0666)
				if err != nil {
					fmt.Printf("#error: %v\n", err.Error())
				}

				// read from file and decode
				tb, _ := ioutil.ReadFile("./uploadfile/" + filename + ".txt")
				dist, _ := base64.StdEncoding.DecodeString(string(tb))

				// create a new image file
				f, _ := os.OpenFile("./uploadfile/new"+filename+".png", os.O_RDWR|os.O_CREATE, os.ModePerm)
				defer f.Close()
				f.Write(dist)

				fmt.Printf(
					"%s  NO.: %2d, Size: %4d KB, Name：%s,\n",
					time.Now().Format("2006-01-02 15:04:05"),
					n,
					num/1024,
					filename,
				)

				//f, err := os.Create("./uploadfile/"+filename)
				//defer f.Close()
				//io.Copy(f, file)
				//
				//// 获取文件状态信息
				//fstat, _ := f.Stat()
				//
				//// 打印接收信息
				//fmt.Fprintf(
				//	w, "%s  NO.: %2d, Size: %4d KB, Name：%s\n",
				//	time.Now().Format("2006-01-02 15:04:05"),
				//	n,
				//	fstat.Size()/1024,
				//	filename,
				//)
				//fmt.Printf(
				//	"%s  NO.: %2d, Size: %4d KB, Name：%s,\n",
				//	time.Now().Format("2006-01-02 15:04:05"),
				//	n,
				//	fstat.Size()/1024,
				//	filename,
				//)
			}
		}

		// text
		tips := r.Form.Get("tips")
		fmt.Printf("%s  Key: %v  Value: %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			"tips",
			tips,
		)
		return

	} else {
		indexHandle(w, r)
	}
}
