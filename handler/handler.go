package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		fmt.Printf("from: %s\n", request.Host)
		// 返回上传 html 页面
		file, err := ioutil.ReadFile("static/view/index.html")
		if err != nil {
			return
		}
		if _, err = io.WriteString(writer, string(file)); err != nil {
			return
		}
	} else if request.Method == "POST" {
		// 接受文件流
		file, head, err := request.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err:%v\n", err.Error())
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)
		newFile, err := os.Create("./" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file by : %v\n", err.Error())
			return
		}
		defer func(newFile *os.File) {
			err := newFile.Close()
			if err != nil {
				panic(err)
			}
		}(newFile)
		_, err = io.Copy(newFile, file)
		if err != nil {
			log.Fatalf("copy failed %v\n", err.Error())
			return
		}
		http.Redirect(writer, request, "/file/upload/suc", http.StatusFound)
	}

}

func UploadSucceedHandler(write http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(write, "Uploading finish!")
	if err != nil {
		panic(err)
	}
}
