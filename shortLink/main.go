package main

import (
	"net/http"
	"os"
	"syscall"
	"os/signal"
	"github.com/mr_litt/go-case/shortLink/app"
)

func main() {

	// 获取短链接服务
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", app.GetShortUrl)
		http.ListenAndServe(":8081", mux)
	}()

	// 跳转长链接服务
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", app.RedirectLongUrl)
		http.ListenAndServe(":8082", mux)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
