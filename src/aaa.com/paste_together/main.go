package main

import (
	"aaa.com/paste_together/common"
	"aaa.com/paste_together/controller"
	"aaa.com/paste_together/middleware"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 执行文件所在目录
	basePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	// 模板目录
	templateDirPath := filepath.Join(basePath, "template")
	if !common.CheckPathExist(templateDirPath) {
		templateDirPath2, _ := filepath.Abs("template")
		if !common.CheckPathExist(templateDirPath2) {
			panic("Template path is not exist")
		} else {
			templateDirPath = templateDirPath2
		}
	}
	log.Println("Template path: " + templateDirPath)

	// 消息存储目录
	storageDirPath := filepath.Join(basePath, "data")
	log.Println("Data path: " + storageDirPath)

	// 获取监听地址
	var listenAddr string
	flag.StringVar(&listenAddr, "listen", ":46699", "--listen=:46699")
	flag.Parse()
	log.Println("Listen: " + listenAddr)

	// 初始化控制器
	messageController := controller.MessageController{
		TemplateDirPath: templateDirPath,
		StorageDirPath:  storageDirPath,
	}
	messageController.Init()

	// 默认调用的中间件
	defaultMiddleware := func(h func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
		return middleware.LogHandler(middleware.RecoverHandler(h))
	}

	// 首页
	http.HandleFunc("/", defaultMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		messageController.GetContents(w, r)
	}))

	// 消息添加接口
	http.HandleFunc("/api/v1/messages", defaultMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			messageController.Create(w, r)
		case http.MethodDelete:
			messageController.DeleteAll(w, r)
		default:
			common.ResponseJsonError(
				w,
				common.AppError{Code: http.StatusMethodNotAllowed, Err: errors.New(http.StatusText(http.StatusMethodNotAllowed))},
			)
		}
	}))

	// 启动服务
	s := &http.Server{
		Addr:           listenAddr,
		Handler:        nil,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
